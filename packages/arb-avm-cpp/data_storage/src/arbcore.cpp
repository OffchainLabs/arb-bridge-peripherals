/*
 * Copyright 2020-2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <data_storage/arbcore.hpp>

#include "value/utils.hpp"

#include <avm/inboxmessage.hpp>
#include <avm/machinethread.hpp>
#include <data_storage/aggregator.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/readsnapshottransaction.hpp>
#include <data_storage/readwritetransaction.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/value.hpp>
#include <data_storage/value/valuecache.hpp>

#include <ethash/keccak.hpp>
#include <set>
#include <sstream>
#include <vector>

#ifdef __linux__
#include <sys/prctl.h>
#endif

namespace {
constexpr auto log_inserted_key = std::array<char, 1>{-60};
constexpr auto log_processed_key = std::array<char, 1>{-61};
constexpr auto send_inserted_key = std::array<char, 1>{-62};
constexpr auto send_processed_key = std::array<char, 1>{-63};
constexpr auto message_entry_inserted_key = std::array<char, 1>{-64};
constexpr auto logscursor_current_prefix = std::array<char, 1>{-66};

constexpr auto sideload_cache_size = 20;
}  // namespace

ArbCore::ArbCore(std::shared_ptr<DataStorage> data_storage_)
    : data_storage(std::move(data_storage_)),
      code(std::make_shared<Code>(getNextSegmentID(data_storage))) {
    if (logs_cursors.size() > 255) {
        throw std::runtime_error("Too many logscursors");
    }
    for (size_t i = 0; i < logs_cursors.size(); i++) {
        logs_cursors[i].current_total_key.insert(
            logs_cursors[i].current_total_key.end(),
            logscursor_current_prefix.begin(), logscursor_current_prefix.end());
        logs_cursors[i].current_total_key.emplace_back(i);
    }
}

bool ArbCore::machineIdle() {
    return machine_idle;
}

ArbCore::message_status_enum ArbCore::messagesStatus() {
    auto current_status = message_data_status.load();
    if (current_status != MESSAGES_ERROR && current_status != MESSAGES_READY) {
        message_data_status = MESSAGES_EMPTY;
    }
    return current_status;
}

std::string ArbCore::messagesClearError() {
    if (message_data_status != ArbCore::MESSAGES_ERROR &&
        message_data_status != ArbCore::MESSAGES_NEED_OLDER) {
        return "";
    }

    message_data_status = MESSAGES_EMPTY;
    auto str = core_error_string;
    core_error_string.clear();

    return str;
}

std::optional<std::string> ArbCore::machineClearError() {
    if (!machine_error) {
        return std::nullopt;
    }

    machine_error = false;
    auto str = machine_error_string;
    machine_error_string.clear();

    return str;
}

bool ArbCore::startThread() {
    abortThread();

    core_thread =
        std::make_unique<std::thread>(std::reference_wrapper<ArbCore>(*this));

    return true;
}

void ArbCore::abortThread() {
    if (core_thread) {
        arbcore_abort = true;
        core_thread->join();
        core_thread = nullptr;
    }
    arbcore_abort = false;
}

// deliverMessages sends messages to core thread
bool ArbCore::deliverMessages(
    std::vector<std::vector<unsigned char>> messages,
    const uint256_t& previous_inbox_acc,
    bool last_block_complete,
    const std::optional<uint256_t>& reorg_message_count) {
    if (message_data_status != MESSAGES_EMPTY) {
        return false;
    }

    message_data.messages = std::move(messages);
    message_data.previous_inbox_acc = previous_inbox_acc;
    message_data.last_block_complete = last_block_complete;
    message_data.reorg_message_count = reorg_message_count;

    message_data_status = MESSAGES_READY;

    return true;
}

rocksdb::Status ArbCore::initialize(const LoadedExecutable& executable) {
    // Use latest existing checkpoint
    ValueCache cache{1, 0};

    auto status = reorgToMessageOrBefore(0, true, cache);
    if (status.ok()) {
        return status;
    }
    if (!status.IsNotFound()) {
        std::cerr << "Error with initial reorg: " << status.ToString()
                  << std::endl;
        return status;
    }

    code->addSegment(executable.code);
    machine = std::make_unique<MachineThread>(
        MachineState{code, executable.static_val});

    ReadWriteTransaction tx(data_storage);
    // Need to initialize database from scratch

    auto s = saveCheckpoint(tx);
    if (!s.ok()) {
        std::cerr << "failed to save initial checkpoint into db: "
                  << s.ToString() << std::endl;
        return s;
    }

    status = updateLogInsertedCount(tx, 0);
    if (!status.ok()) {
        throw std::runtime_error("failed to initialize log inserted count");
    }
    status = updateSendInsertedCount(tx, 0);
    if (!status.ok()) {
        throw std::runtime_error("failed to initialize log inserted count");
    }
    status = updateMessageEntryInsertedCount(tx, 0);
    if (!status.ok()) {
        throw std::runtime_error("failed to initialize log inserted count");
    }

    for (size_t i = 0; i < logs_cursors.size(); i++) {
        status = logsCursorSaveCurrentTotalCount(tx, i, 0);
        if (!status.ok()) {
            throw std::runtime_error("failed to initialize logscursor counts");
        }
    }

    s = tx.commit();
    if (!s.ok()) {
        std::cerr << "failed to commit initial state into db: " << s.ToString()
                  << std::endl;
        return s;
    }

    return rocksdb::Status::OK();
}

bool ArbCore::initialized() const {
    ReadTransaction tx(data_storage);
    std::vector<unsigned char> key;
    marshal_uint256_t(0, key);
    return tx.checkpointGetVector(vecToSlice(key)).status.ok();
}

template <class T>
std::unique_ptr<T> ArbCore::getMachineImpl(ReadTransaction& tx,
                                           uint256_t machineHash,
                                           ValueCache& value_cache) {
    auto results = getMachineStateKeys(tx, machineHash);
    if (std::holds_alternative<rocksdb::Status>(results)) {
        throw std::runtime_error("failed to load machine state");
    }

    return getMachineUsingStateKeys<T>(
        tx, std::get<CountedData<MachineStateKeys>>(results).data, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getMachineImpl(
    ReadTransaction& tx,
    uint256_t machineHash,
    ValueCache& value_cache);
template std::unique_ptr<MachineThread> ArbCore::getMachineImpl(
    ReadTransaction& tx,
    uint256_t machineHash,
    ValueCache& value_cache);

template <class T>
std::unique_ptr<T> ArbCore::getMachine(uint256_t machineHash,
                                       ValueCache& value_cache) {
    ReadSnapshotTransaction tx(data_storage);
    return getMachineImpl<T>(tx, machineHash, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getMachine(uint256_t, ValueCache&);
template std::unique_ptr<MachineThread> ArbCore::getMachine(uint256_t,
                                                            ValueCache&);

// triggerSaveCheckpoint is meant for unit tests and should not be called from
// multiple threads at the same time.
rocksdb::Status ArbCore::triggerSaveCheckpoint() {
    save_checkpoint = true;
    std::cerr << "Triggering checkpoint save" << std::endl;
    while (save_checkpoint) {
        // Wait until snapshot has been saved
        std::this_thread::sleep_for(std::chrono::milliseconds(10));
    }
    std::cerr << "Checkpoint saved" << std::endl;

    return save_checkpoint_status;
}

rocksdb::Status ArbCore::saveCheckpoint(ReadWriteTransaction& tx) {
    auto& state = machine->machine_state;
    if (!isValid(tx, state.output.fully_processed_inbox)) {
        std::cerr << "Attempted to save invalid checkpoint at gas "
                  << state.output.arb_gas_used << std::endl;
        assert(false);
        return rocksdb::Status::OK();
    }

    auto status = saveMachineState(tx, *machine);
    if (!status.ok()) {
        return status;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(state.output.arb_gas_used, key);
    auto key_slice = vecToSlice(key);
    std::vector<unsigned char> value_vec;
    serializeMachineStateKeys(MachineStateKeys{state}, value_vec);
    auto put_status = tx.checkpointPut(key_slice, vecToSlice(value_vec));
    if (!put_status.ok()) {
        std::cerr << "ArbCore unable to save checkpoint : "
                  << put_status.ToString() << "\n";
        return put_status;
    }

    return rocksdb::Status::OK();
}

rocksdb::Status ArbCore::saveAssertion(ReadWriteTransaction& tx,
                                       const Assertion& assertion,
                                       const uint256_t arb_gas_used) {
    auto status = saveLogs(tx, assertion.logs);
    if (!status.ok()) {
        return status;
    }

    status = saveSends(tx, assertion.sends);
    if (!status.ok()) {
        return status;
    }

    if (assertion.sideloadBlockNumber) {
        status = saveSideloadPosition(tx, *assertion.sideloadBlockNumber,
                                      arb_gas_used);
        if (!status.ok()) {
            return status;
        }
    }

    return rocksdb::Status::OK();
}

// reorgToMessageOrBefore resets the checkpoint and database entries
// such that machine state is at or before the requested message. cleaning
// up old references as needed.
// If use_latest is true, message_sequence_number is ignored and the latest
// checkpoint is used.
rocksdb::Status ArbCore::reorgToMessageOrBefore(
    const uint256_t& message_sequence_number,
    bool use_latest,
    ValueCache& cache) {
    std::variant<MachineStateKeys, rocksdb::Status> setup =
        rocksdb::Status::OK();

    if (use_latest) {
        std::cerr << "Reloading latest checkpoint" << std::endl;
    } else {
        std::cerr << "Reorganizing to message " << message_sequence_number
                  << std::endl;
    }

    {
        ReadWriteTransaction tx(data_storage);

        auto it = tx.checkpointGetIterator();

        // Find first checkpoint to delete
        it->SeekToLast();
        if (!it->status().ok()) {
            return it->status();
        }

        if (!it->Valid()) {
            return rocksdb::Status::NotFound();
        }

        // Delete each checkpoint until at or below message_sequence_number
        setup = [&]() -> std::variant<MachineStateKeys, rocksdb::Status> {
            while (it->Valid()) {
                std::vector<unsigned char> checkpoint_vector(
                    it->value().data(),
                    it->value().data() + it->value().size());
                auto checkpoint = extractMachineStateKeys(
                    checkpoint_vector.begin(), checkpoint_vector.end());
                if (checkpoint.getTotalMessagesRead() == 0 || use_latest ||
                    (message_sequence_number >=
                     checkpoint.getTotalMessagesRead() - 1)) {
                    if (isValid(tx, checkpoint.output.fully_processed_inbox)) {
                        // Good checkpoint
                        return checkpoint;
                    }

                    std::cerr << "Error: Invalid checkpoint found at gas: "
                              << checkpoint.output.arb_gas_used << std::endl;
                    assert(false);
                }

                // Obsolete checkpoint, need to delete referenced machine
                deleteMachineState(tx, checkpoint);

                // Delete checkpoint to make sure it isn't used later
                tx.checkpointDelete(it->key());

                it->Prev();
                if (!it->status().ok()) {
                    return it->status();
                }
            }
            return it->status();
        }();

        it = nullptr;
        if (std::holds_alternative<rocksdb::Status>(setup)) {
            return std::get<rocksdb::Status>(setup);
        }

        auto status = tx.commit();
        if (!status.ok()) {
            return status;
        }
    }
    MachineStateKeys checkpoint = std::get<MachineStateKeys>(std::move(setup));

    auto log_inserted_count = logInsertedCount();
    if (!log_inserted_count.status.ok()) {
        std::cerr << "Error getting inserted count in Cursor Reorg: "
                  << log_inserted_count.status.ToString() << "\n";
        return log_inserted_count.status;
    }

    if (checkpoint.output.log_count < log_inserted_count.data) {
        // Update log cursors, must be called before logs are deleted
        for (size_t i = 0; i < logs_cursors.size(); i++) {
            auto status =
                handleLogsCursorReorg(i, checkpoint.output.log_count, cache);
            if (!status.ok()) {
                return status;
            }
        }
    }

    uint256_t next_sideload_block_number = 0;
    if (checkpoint.output.last_sideload) {
        next_sideload_block_number = *checkpoint.output.last_sideload + 1;
    }

    ReadWriteTransaction tx(data_storage);
    auto status = deleteSideloadsStartingAt(tx, next_sideload_block_number);
    if (!status.ok()) {
        return status;
    }

    // Delete logs individually to handle reference counts
    auto optional_status =
        deleteLogsStartingAt(tx, checkpoint.output.log_count);
    if (optional_status && !optional_status->ok()) {
        return *optional_status;
    }

    status = updateLogInsertedCount(tx, checkpoint.output.log_count);
    if (!status.ok()) {
        return status;
    }

    status = updateSendInsertedCount(tx, checkpoint.output.send_count);
    if (!status.ok()) {
        return status;
    }

    // Machine was executing obsolete messages so restore machine
    // from last checkpoint
    if (machine != nullptr) {
        machine->abortMachine();
    }

    machine = getMachineUsingStateKeys<MachineThread>(tx, checkpoint, cache);

    // Update last machine output
    {
        std::unique_lock<std::shared_mutex> guard(last_machine_output_mutex);
        last_machine_output = machine->machine_state.output;
    }

    return tx.commit();
}

std::variant<rocksdb::Status, MachineStateKeys> ArbCore::getCheckpoint(
    ReadTransaction& tx,
    const uint256_t& arb_gas_used) const {
    std::vector<unsigned char> key;
    marshal_uint256_t(arb_gas_used, key);

    auto result = tx.checkpointGetVector(vecToSlice(key));
    if (!result.status.ok()) {
        return result.status;
    }
    return extractMachineStateKeys(result.data.begin(), result.data.end());
}

bool ArbCore::isCheckpointsEmpty(ReadTransaction& tx) const {
    auto it = std::unique_ptr<rocksdb::Iterator>(tx.checkpointGetIterator());
    it->SeekToLast();
    return !it->Valid();
}

uint256_t ArbCore::maxCheckpointGas() {
    ReadTransaction tx(data_storage);
    auto it = tx.checkpointGetIterator();
    it->SeekToLast();
    if (it->Valid()) {
        auto keyBuf = it->key().data();
        return deserializeUint256t(keyBuf);
    } else {
        return 0;
    }
}

// getCheckpointUsingGas returns the checkpoint at or before the specified gas
// if `after_gas` is false. If `after_gas` is true, checkpoint after specified
// gas is returned.
std::variant<rocksdb::Status, MachineStateKeys> ArbCore::getCheckpointUsingGas(
    ReadTransaction& tx,
    const uint256_t& total_gas,
    bool after_gas) {
    auto it = tx.checkpointGetIterator();
    std::vector<unsigned char> key;
    marshal_uint256_t(total_gas, key);
    auto key_slice = vecToSlice(key);
    it->SeekForPrev(key_slice);
    if (!it->Valid()) {
        if (!it->status().ok()) {
            return it->status();
        }
        return rocksdb::Status::NotFound();
    }
    if (after_gas) {
        it->Next();
        if (!it->status().ok()) {
            return it->status();
        }
        if (!it->Valid()) {
            return rocksdb::Status::NotFound();
        }
    }
    if (!it->status().ok()) {
        return it->status();
    }

    std::vector<unsigned char> saved_value(
        it->value().data(), it->value().data() + it->value().size());
    return extractMachineStateKeys(saved_value.begin(), saved_value.end());
}

template <class T>
std::unique_ptr<T> ArbCore::getMachineUsingStateKeys(
    const ReadTransaction& transaction,
    const MachineStateKeys& state_data,
    ValueCache& value_cache) const {
    std::set<uint64_t> segment_ids;

    auto static_results = ::getValueImpl(transaction, state_data.static_hash,
                                         segment_ids, value_cache);

    if (std::holds_alternative<rocksdb::Status>(static_results)) {
        std::stringstream ss;
        ss << "failed loaded core machine static: "
           << std::get<rocksdb::Status>(static_results).ToString();
        throw std::runtime_error(ss.str());
    }

    auto register_results = ::getValueImpl(
        transaction, state_data.register_hash, segment_ids, value_cache);
    if (std::holds_alternative<rocksdb::Status>(register_results)) {
        std::stringstream ss;
        ss << "failed loaded core machine register: "
           << std::get<rocksdb::Status>(register_results).ToString();
        throw std::runtime_error(ss.str());
    }

    auto stack_results = ::getValueImpl(transaction, state_data.datastack_hash,
                                        segment_ids, value_cache);
    if (std::holds_alternative<rocksdb::Status>(stack_results) ||
        !std::holds_alternative<Tuple>(
            std::get<CountedData<value>>(stack_results).data)) {
        throw std::runtime_error("failed to load machine stack");
    }

    auto auxstack_results = ::getValueImpl(
        transaction, state_data.auxstack_hash, segment_ids, value_cache);
    if (std::holds_alternative<rocksdb::Status>(auxstack_results)) {
        throw std::runtime_error("failed to load machine auxstack");
    }
    if (!std::holds_alternative<Tuple>(
            std::get<CountedData<value>>(auxstack_results).data)) {
        throw std::runtime_error(
            "failed to load machine auxstack because of format error");
    }

    segment_ids.insert(state_data.pc.pc.segment);
    segment_ids.insert(state_data.err_pc.pc.segment);

    bool loaded_segment = true;
    while (loaded_segment) {
        loaded_segment = false;
        std::set<uint64_t> next_segment_ids;
        for (auto it = segment_ids.rbegin(); it != segment_ids.rend(); ++it) {
            if (code->containsSegment(*it)) {
                // If the segment is already loaded, no need to restore it
                continue;
            }
            auto segment =
                getCodeSegment(transaction, *it, next_segment_ids, value_cache);
            code->restoreExistingSegment(std::move(segment));
            loaded_segment = true;
        }
        segment_ids = std::move(next_segment_ids);
    };
    auto state = MachineState{
        code,
        std::move(std::get<CountedData<value>>(register_results).data),
        std::move(std::get<CountedData<value>>(static_results).data),
        Datastack(
            std::get<Tuple>(std::get<CountedData<value>>(stack_results).data)),
        Datastack(std::get<Tuple>(
            std::get<CountedData<value>>(auxstack_results).data)),
        state_data.arb_gas_remaining,
        state_data.status,
        state_data.pc.pc,
        state_data.err_pc,
        state_data.staged_message,
        state_data.output};

    return std::make_unique<T>(state);
}

template std::unique_ptr<Machine> ArbCore::getMachineUsingStateKeys(
    const ReadTransaction& transaction,
    const MachineStateKeys& state_data,
    ValueCache& value_cache) const;
template std::unique_ptr<MachineThread> ArbCore::getMachineUsingStateKeys(
    const ReadTransaction& transaction,
    const MachineStateKeys& state_data,
    ValueCache& value_cache) const;

// operator() runs the main thread for ArbCore.  It is responsible for adding
// messages to the queue, starting machine thread when needed and collecting
// results of machine thread.
// This thread will update `delivering_messages` if and only if
// `delivering_messages` is set to MESSAGES_READY
void ArbCore::operator()() {
#ifdef __linux__
    prctl(PR_SET_NAME, "ArbCore", 0, 0, 0);
#endif
    ValueCache cache{5, 0};
    MachineExecutionConfig execConfig;
    execConfig.stop_on_sideload = true;
    size_t max_message_batch_size = 10;

    while (!arbcore_abort) {
        bool isMachineValid;
        {
            ReadTransaction tx(data_storage);
            isMachineValid = isValid(tx, machine->getReorgData());
        }
        if (!isMachineValid) {
            std::cerr
                << "Core thread operating on invalid machine. Rolling back."
                << std::endl;
            assert(false);
            auto status = reorgToMessageOrBefore(0, true, cache);
            if (!status.ok()) {
                std::cerr
                    << "Error in core thread calling reorgToMessageOrBefore: "
                    << status.ToString() << std::endl;
            }
        }
        if (message_data_status == MESSAGES_READY) {
            // Reorg might occur while adding messages
            auto add_status = addMessages(
                message_data.messages, message_data.last_block_complete,
                message_data.previous_inbox_acc,
                message_data.reorg_message_count, cache);
            if (!add_status) {
                // Messages from previous block invalid because of reorg so
                // request older messages
                message_data_status = MESSAGES_NEED_OLDER;
            } else if (!add_status->ok()) {
                core_error_string = add_status->ToString();
                message_data_status = MESSAGES_ERROR;
                std::cerr << "ArbCore inbox processed stopped with error: "
                          << core_error_string << "\n";
                break;
            } else {
                machine_idle = false;
                message_data_status = MESSAGES_SUCCESS;
            }
        }

        // Check machine thread
        if (machine->status() == MachineThread::MACHINE_ERROR) {
            core_error_string = machine->getErrorString();
            std::cerr << "AVM machine stopped with error: " << core_error_string
                      << "\n";
            break;
        }

        if (machine->status() == MachineThread::MACHINE_SUCCESS) {
            ReadWriteTransaction tx(data_storage);

            auto last_assertion = machine->nextAssertion();

            // Save last machine output
            {
                std::unique_lock<std::shared_mutex> guard(
                    last_machine_output_mutex);
                last_machine_output = machine->machine_state.output;
            }

            // Save logs and sends
            auto status = saveAssertion(
                tx, last_assertion, machine->machine_state.output.arb_gas_used);
            if (!status.ok()) {
                core_error_string = status.ToString();
                std::cerr << "ArbCore assertion saving failed: "
                          << core_error_string << "\n";
                break;
            }

            // Cache pre-sideload machines
            if (last_assertion.sideloadBlockNumber) {
                {
                    auto block = *last_assertion.sideloadBlockNumber;
                    std::unique_lock<std::shared_mutex> lock(
                        sideload_cache_mutex);
                    sideload_cache[block] = std::make_unique<Machine>(*machine);
                    // Remove any sideload_cache entries that are either more
                    // than sideload_cache_size blocks old, or in the future
                    // (meaning they've been reorg'd out).
                    auto it = sideload_cache.begin();
                    while (it != sideload_cache.end()) {
                        // Note: we check if block > sideload_cache_size here
                        // to prevent an underflow in the following check.
                        if ((block > sideload_cache_size &&
                             it->first < block - sideload_cache_size) ||
                            it->first > block) {
                            it = sideload_cache.erase(it);
                        } else {
                            it++;
                        }
                    }
                }

                // Save checkpoint for every sideload
                status = saveCheckpoint(tx);
                if (!status.ok()) {
                    core_error_string = status.ToString();
                    std::cerr << "ArbCore checkpoint saving failed: "
                              << core_error_string << "\n";
                    break;
                }

                // Clear oldest cache and start populating next cache
                cache.nextCache();

                // Machine was stopped to save sideload, update execConfig
                // and start machine back up where it stopped
                auto machine_success = machine->continueRunningMachine();
                if (!machine_success) {
                    core_error_string = "Error starting machine thread";
                    machine_error = true;
                    std::cerr << "ArbCore error: " << core_error_string << "\n";
                    break;
                }
            }

            status = tx.commit();
            if (!status.ok()) {
                core_error_string = status.ToString();
                machine_error = true;
                std::cerr << "ArbCore database update failed: "
                          << core_error_string << "\n";
                break;
            }
        }

        if (machine->status() == MachineThread::MACHINE_ABORTED) {
            // Just reset status so machine can be restarted
            machine->clearError();
        }

        if (machine->status() == MachineThread::MACHINE_NONE) {
            // Start execution of machine if new message available
            ReadSnapshotTransaction tx(data_storage);
            auto messages_result = readNextMessages(
                tx, machine->machine_state.output.fully_processed_inbox,
                max_message_batch_size);
            if (!messages_result.status.ok()) {
                core_error_string = messages_result.status.ToString();
                machine_error = true;
                std::cerr << "ArbCore failed getting message entry: "
                          << core_error_string << "\n";
                break;
            }

            if (!messages_result.data.empty()) {
                execConfig.inbox_messages = messages_result.data;

                auto success = machine->runMachine(execConfig);
                if (!success) {
                    core_error_string = "Error starting machine thread";
                    machine_error = true;
                    std::cerr << "ArbCore error: " << core_error_string << "\n";
                    break;
                }

                if (delete_checkpoints_before_message != uint256_t(0)) {
                    /*
                    deleteOldCheckpoints(delete_checkpoints_before_message,
                                         save_checkpoint_message_interval,
                                         ignore_checkpoints_after_message);
                    */
                    ignore_checkpoints_after_message = 0;
                    save_checkpoint_message_interval = 0;
                    delete_checkpoints_before_message = 0;
                }
            } else {
                // Machine all caught up, no messages to process
                machine_idle = true;
            }
        }

        for (size_t i = 0; i < logs_cursors.size(); i++) {
            if (logs_cursors[i].status == DataCursor::REQUESTED) {
                ReadTransaction tx(data_storage);
                handleLogsCursorRequested(tx, i, cache);
            }
        }

        if (save_checkpoint) {
            ReadWriteTransaction tx(data_storage);
            save_checkpoint_status = saveCheckpoint(tx);
            tx.commit();
            save_checkpoint = false;
        }

        if (!machineIdle() || message_data_status != MESSAGES_READY) {
            // Machine is already running or new messages, so sleep for a short
            // while
            std::this_thread::sleep_for(std::chrono::milliseconds(5));
        }
    }

    // Error occurred, make sure machine stops cleanly
    machine->abortMachine();
}

rocksdb::Status ArbCore::saveLogs(ReadWriteTransaction& tx,
                                  const std::vector<value>& vals) {
    if (vals.empty()) {
        return rocksdb::Status::OK();
    }
    auto log_result = logInsertedCountImpl(tx);
    if (!log_result.status.ok()) {
        return log_result.status;
    }

    auto log_index = log_result.data;
    for (const auto& val : vals) {
        auto value_result = saveValue(tx, val);
        if (!value_result.status.ok()) {
            return value_result.status;
        }

        std::vector<unsigned char> key;
        marshal_uint256_t(log_index, key);
        auto key_slice = vecToSlice(key);

        std::vector<unsigned char> value_hash;
        marshal_uint256_t(hash_value(val), value_hash);
        rocksdb::Slice value_hash_slice(
            reinterpret_cast<const char*>(value_hash.data()),
            value_hash.size());

        auto status = tx.logPut(key_slice, value_hash_slice);
        if (!status.ok()) {
            return status;
        }
        log_index += 1;
    }

    return updateLogInsertedCount(tx, log_index);
}

ValueResult<std::vector<value>> ArbCore::getLogs(uint256_t index,
                                                 uint256_t count,
                                                 ValueCache& valueCache) {
    ReadSnapshotTransaction tx(data_storage);

    return getLogsNoLock(tx, index, count, valueCache);
}

ValueResult<std::vector<value>> ArbCore::getLogsNoLock(ReadTransaction& tx,
                                                       uint256_t index,
                                                       uint256_t count,
                                                       ValueCache& valueCache) {
    if (count == 0) {
        return {rocksdb::Status::OK(), {}};
    }

    // Check if attempting to get entries past current valid logs
    auto log_count = logInsertedCountImpl(tx);
    if (!log_count.status.ok()) {
        return {log_count.status, {}};
    }
    auto max_log_count = log_count.data;
    if (index >= max_log_count) {
        return {rocksdb::Status::OK(), {}};
    }
    if (index + count > max_log_count) {
        count = max_log_count - index;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(index, key);

    auto hash_result = tx.logGetUint256Vector(vecToSlice(key),
                                              intx::narrow_cast<size_t>(count));
    if (!hash_result.status.ok()) {
        return {hash_result.status, {}};
    }

    std::vector<value> logs;
    for (const auto& hash : hash_result.data) {
        auto val_result = getValue(tx, hash, valueCache);
        if (std::holds_alternative<rocksdb::Status>(val_result)) {
            return {std::get<rocksdb::Status>(val_result), {}};
        }
        logs.push_back(
            std::move(std::get<CountedData<value>>(val_result).data));
    }

    return {rocksdb::Status::OK(), std::move(logs)};
}

rocksdb::Status ArbCore::saveSends(
    ReadWriteTransaction& tx,
    const std::vector<std::vector<unsigned char>>& sends) {
    if (sends.empty()) {
        return rocksdb::Status::OK();
    }
    auto send_result = sendInsertedCountImpl(tx);
    if (!send_result.status.ok()) {
        return send_result.status;
    }

    auto send_count = send_result.data;
    for (const auto& send : sends) {
        std::vector<unsigned char> key;
        marshal_uint256_t(send_count, key);
        auto key_slice = vecToSlice(key);

        auto status = tx.sendPut(key_slice, vecToSlice(send));
        if (!status.ok()) {
            return status;
        }
        send_count += 1;
    }

    return updateSendInsertedCount(tx, send_count);
}

ValueResult<std::vector<std::vector<unsigned char>>> ArbCore::getMessages(
    uint256_t index,
    uint256_t count) const {
    ReadSnapshotTransaction tx(data_storage);

    auto result = getMessagesImpl(tx, index, count, std::nullopt);
    if (!result.status.ok()) {
        return {result.status, {}};
    }

    std::vector<std::vector<unsigned char>> bytes_vec(result.data.size());
    for (auto& message_and_acc : result.data) {
        bytes_vec.push_back(std::move(message_and_acc.message));
    }

    return {result.status, bytes_vec};
}

ValueResult<std::vector<RawMessageAndAccumulator>> ArbCore::getMessagesImpl(
    const ReadConsistentTransaction& tx,
    uint256_t index,
    uint256_t count,
    std::optional<uint256_t> start_acc) const {
    std::vector<RawMessageAndAccumulator> messages;

    uint256_t start = index;
    uint256_t end = start + count;
    bool needs_consistency_check = false;
    if (start > 0) {
        // Check the previous item to ensure the inbox state is valid
        start -= 1;
        needs_consistency_check = true;
    }

    std::vector<unsigned char> tmp(32 * 2);
    rocksdb::Slice seq_batch_lower_bound;
    {
        auto ptr = reinterpret_cast<const char*>(tmp.data() + tmp.size());
        marshal_uint256_t(start, tmp);
        seq_batch_lower_bound = {ptr, 32};
    }
    auto seq_batch_it =
        tx.sequencerBatchItemGetIterator(&seq_batch_lower_bound);

    uint256_t prev_delayed_count = 0;
    rocksdb::Slice delayed_msg_lower_bound;
    std::unique_ptr<rocksdb::Iterator> delayed_msg_it;
    for (seq_batch_it->Seek(seq_batch_lower_bound); seq_batch_it->Valid();
         seq_batch_it->Next()) {
        auto item_key_ptr =
            reinterpret_cast<const unsigned char*>(seq_batch_it->key().data());
        auto item_value_ptr = reinterpret_cast<const unsigned char*>(
            seq_batch_it->value().data());
        auto last_sequence_number = extractUint256(item_key_ptr);
        auto item =
            deserializeSequencerBatchItem(last_sequence_number, item_value_ptr);

        if (needs_consistency_check) {
            if (start_acc && item.accumulator != *start_acc) {
                return {rocksdb::Status::NotFound(), {}};
            }
            needs_consistency_check = false;
            if (count == 0) {
                // Skip some possible work attempting to read delayed messages
                break;
            }
            prev_delayed_count = item.total_delayed_count;
            if (item.last_sequence_number >= index) {
                // We are in the middle of a delayed batch
                assert(!item.sequencer_message);
                // Offset prev_delayed_count by the distance to the end of the
                // batch
                prev_delayed_count -= item.last_sequence_number + 1 - index;
            } else {
                // We are just after this batch item
                assert(item.last_sequence_number + 1 == index);
                continue;
            }
        }

        if (item.sequencer_message) {
            messages.emplace_back(*item.sequencer_message, item.accumulator);
            if (prev_delayed_count != item.total_delayed_count) {
                throw std::runtime_error(
                    "Sequencer batch item included both sequencer message and "
                    "delayed messages");
            }
        } else if (item.total_delayed_count > prev_delayed_count) {
            if (!delayed_msg_it) {
                {
                    auto ptr =
                        reinterpret_cast<const char*>(tmp.data() + tmp.size());
                    marshal_uint256_t(prev_delayed_count, tmp);
                    delayed_msg_lower_bound = {ptr, 32};
                }
                delayed_msg_it =
                    tx.delayedMessageGetIterator(&delayed_msg_lower_bound);
                delayed_msg_it->Seek(delayed_msg_lower_bound);
            }

            while (delayed_msg_it->Valid() &&
                   prev_delayed_count < item.total_delayed_count &&
                   messages.size() < count) {
                auto delayed_key_ptr = reinterpret_cast<const unsigned char*>(
                    delayed_msg_it->key().data());
                auto delayed_value_ptr = reinterpret_cast<const unsigned char*>(
                    delayed_msg_it->value().data());
                if (extractUint256(delayed_key_ptr) != prev_delayed_count) {
                    throw std::runtime_error(
                        "Got wrong delayed message from database");
                }
                auto delayed_message = deserializeDelayedMessage(
                    prev_delayed_count, delayed_value_ptr);
                messages.emplace_back(delayed_message.message,
                                      item.accumulator);
                prev_delayed_count += 1;
                delayed_msg_it->Next();
            }

            if (!delayed_msg_it->status().ok()) {
                return {delayed_msg_it->status(), {}};
            }
            if (messages.size() < count &&
                prev_delayed_count != item.total_delayed_count) {
                throw std::runtime_error(
                    "Sequencer batch item referenced nonexistent delayed "
                    "messages");
            }
        } else {
            // This batch item does nothing?
            assert(false);
        }
        if (messages.size() >= count) {
            break;
        }
        assert(item.last_sequence_number + 1 == index + messages.size());
    }

    if (!seq_batch_it->status().ok()) {
        return {seq_batch_it->status(), {}};
    }
    if (needs_consistency_check) {
        return {rocksdb::Status::NotFound(), {}};
    }

    return {rocksdb::Status::OK(), messages};
}

ValueResult<SequencerBatchItem> ArbCore::getNextSequencerBatchItem(
    const ReadTransaction& tx,
    uint256_t sequence_number) const {
    std::vector<unsigned char> tmp(32);
    rocksdb::Slice seq_batch_lower_bound;
    {
        auto ptr = reinterpret_cast<const char*>(tmp.data());
        marshal_uint256_t(sequence_number, tmp);
        seq_batch_lower_bound = {ptr, 32};
    }
    auto seq_batch_it =
        tx.sequencerBatchItemGetIterator(&seq_batch_lower_bound);
    seq_batch_it->Seek(seq_batch_lower_bound);
    if (!seq_batch_it->Valid()) {
        if (seq_batch_it->status().ok()) {
            return {rocksdb::Status::NotFound(), SequencerBatchItem{}};
        } else {
            return {seq_batch_it->status(), SequencerBatchItem{}};
        }
    }
    auto key_ptr =
        reinterpret_cast<const unsigned char*>(seq_batch_it->key().data());
    auto value_ptr =
        reinterpret_cast<const unsigned char*>(seq_batch_it->value().data());
    auto last_sequence_number = extractUint256(key_ptr);
    auto item = deserializeSequencerBatchItem(last_sequence_number, value_ptr);
    return {rocksdb::Status::OK(), item};
}

ValueResult<std::vector<std::vector<unsigned char>>> ArbCore::getSends(
    uint256_t index,
    uint256_t count) const {
    ReadSnapshotTransaction tx(data_storage);

    if (count == 0) {
        return {rocksdb::Status::OK(), {}};
    }

    // Check if attempting to get entries past current valid sends
    auto send_count = sendInsertedCountImpl(tx);
    if (!send_count.status.ok()) {
        return {send_count.status, {}};
    }
    auto max_send_count = send_count.data;
    if (index >= max_send_count) {
        return {rocksdb::Status::NotFound(), {}};
    }
    if (index + count > max_send_count) {
        count = max_send_count - index;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(index, key);
    auto key_slice = vecToSlice(key);

    return tx.sendGetVectorVector(key_slice, intx::narrow_cast<size_t>(count));
}

ValueResult<uint256_t> ArbCore::getInboxAcc(uint256_t index) {
    ReadTransaction tx(data_storage);

    auto result = getNextSequencerBatchItem(tx, index);
    if (!result.status.ok()) {
        return {result.status, 0};
    }

    return {rocksdb::Status::OK(), result.data.accumulator};
}

ValueResult<std::pair<uint256_t, uint256_t>> ArbCore::getInboxAccPair(
    uint256_t index1,
    uint256_t index2) {
    ReadSnapshotTransaction tx(data_storage);

    auto result1 = getNextSequencerBatchItem(tx, index1);
    if (!result1.status.ok()) {
        return {result1.status, {0, 0}};
    }

    auto result2 = getNextSequencerBatchItem(tx, index2);
    if (!result2.status.ok()) {
        return {result2.status, {0, 0}};
    }

    return {rocksdb::Status::OK(),
            {result1.data.accumulator, result2.data.accumulator}};
}

uint256_t ArbCore::machineMessagesRead() {
    std::shared_lock<std::shared_mutex> guard(last_machine_output_mutex);
    return last_machine_output.fully_processed_inbox.count;
}

ValueResult<std::unique_ptr<ExecutionCursor>> ArbCore::getExecutionCursor(
    uint256_t total_gas_used,
    ValueCache& cache) {
    std::unique_ptr<ExecutionCursor> execution_cursor;
    {
        ReadSnapshotTransaction tx(data_storage);

        auto closest_checkpoint =
            getClosestExecutionMachine(tx, total_gas_used);
        if (std::holds_alternative<rocksdb::Status>(closest_checkpoint)) {
            std::cerr << "No execution machine available" << std::endl;
            return {std::get<rocksdb::Status>(closest_checkpoint), nullptr};
        }

        execution_cursor = std::make_unique<ExecutionCursor>(
            std::get<MachineStateKeys>(closest_checkpoint));
    }

    auto status = advanceExecutionCursorImpl(*execution_cursor, total_gas_used,
                                             false, 10, cache);

    if (!status.ok()) {
        std::cerr << "Couldn't advance execution machine" << std::endl;
    }

    return {status, std::move(execution_cursor)};
}

constexpr uint256_t checkpoint_load_gas_cost = 100'000'000;

rocksdb::Status ArbCore::advanceExecutionCursor(
    ExecutionCursor& execution_cursor,
    uint256_t max_gas,
    bool go_over_gas,
    ValueCache& cache) {
    auto gas_target = execution_cursor.getOutput().arb_gas_used + max_gas;
    {
        ReadSnapshotTransaction tx(data_storage);

        auto closest_checkpoint = getClosestExecutionMachine(tx, gas_target);
        if (std::holds_alternative<rocksdb::Status>(closest_checkpoint)) {
            return std::get<rocksdb::Status>(closest_checkpoint);
        }

        auto machine_state_keys =
            std::get<MachineStateKeys>(closest_checkpoint);
        bool already_newer = false;
        if (execution_cursor.getOutput().arb_gas_used +
                checkpoint_load_gas_cost >
            machine_state_keys.output.arb_gas_used) {
            // The existing execution cursor is far enough ahead that running it
            // up to the target gas will be cheaper than loading the checkpoint
            // from disk and running it. We just need to check that the
            // execution cursor is still valid (a reorg hasn't occurred).
            auto result =
                executionCursorGetMessagesNoLock(tx, execution_cursor, 0);
            if (result.status.ok() && result.data.first) {
                // Execution cursor machine still valid, so use it
                already_newer = true;
            }
        }

        if (!already_newer) {
            execution_cursor.machine =
                std::get<MachineStateKeys>(closest_checkpoint);
        }
    }

    return advanceExecutionCursorImpl(execution_cursor, gas_target, go_over_gas,
                                      10, cache);
}

MachineState& resolveExecutionVariant(std::unique_ptr<Machine>& mach) {
    return mach->machine_state;
}

MachineStateKeys& resolveExecutionVariant(MachineStateKeys& mach) {
    return mach;
}

std::unique_ptr<Machine>& ArbCore::resolveExecutionCursorMachine(
    const ReadTransaction& tx,
    ExecutionCursor& execution_cursor,
    ValueCache& cache) const {
    if (std::holds_alternative<MachineStateKeys>(execution_cursor.machine)) {
        auto machine_state_keys =
            std::get<MachineStateKeys>(execution_cursor.machine);
        execution_cursor.machine =
            getMachineUsingStateKeys<Machine>(tx, machine_state_keys, cache);
    }
    return std::get<std::unique_ptr<Machine>>(execution_cursor.machine);
}

std::unique_ptr<Machine> ArbCore::takeExecutionCursorMachineImpl(
    const ReadTransaction& tx,
    ExecutionCursor& execution_cursor,
    ValueCache& cache) const {
    auto mach =
        std::move(resolveExecutionCursorMachine(tx, execution_cursor, cache));
    execution_cursor.machine = MachineStateKeys{mach->machine_state};
    return mach;
}

std::unique_ptr<Machine> ArbCore::takeExecutionCursorMachine(
    ExecutionCursor& execution_cursor,
    ValueCache& cache) const {
    ReadSnapshotTransaction tx(data_storage);
    return takeExecutionCursorMachineImpl(tx, execution_cursor, cache);
}

rocksdb::Status ArbCore::advanceExecutionCursorImpl(
    ExecutionCursor& execution_cursor,
    uint256_t total_gas_used,
    bool go_over_gas,
    size_t message_group_size,
    ValueCache& cache) {
    auto handle_reorg = true;
    size_t reorg_attempts = 0;
    while (handle_reorg) {
        handle_reorg = false;
        if (reorg_attempts > 0) {
            if (reorg_attempts % 4 == 0) {
                std::cerr
                    << "Execution cursor has attempted to handle "
                    << reorg_attempts
                    << " reorgs. Checkpoints may be inconsistent with messages."
                    << std::endl;
            }
            assert(false);
            std::this_thread::sleep_for(std::chrono::milliseconds(250));
            if (reorg_attempts >= 16) {
                return rocksdb::Status::Busy();
            }
        }
        reorg_attempts++;

        while (true) {
            // Run machine until specified gas is reached
            MachineExecutionConfig execConfig;
            execConfig.max_gas = total_gas_used;
            execConfig.go_over_gas = go_over_gas;

            {
                ReadSnapshotTransaction tx(data_storage);

                auto& mach =
                    resolveExecutionCursorMachine(tx, execution_cursor, cache);

                uint256_t gas_used = execution_cursor.getOutput().arb_gas_used;
                if (gas_used == total_gas_used) {
                    break;
                } else if (go_over_gas && gas_used > total_gas_used) {
                    break;
                } else if (!go_over_gas &&
                           gas_used + mach->machine_state.nextGasCost() >
                               total_gas_used) {
                    break;
                }

                auto get_messages_result = readNextMessages(
                    tx, execution_cursor.getOutput().fully_processed_inbox,
                    message_group_size);
                if (!get_messages_result.status.IsNotFound()) {
                    // Reorg occurred, need to recreate machine
                    handle_reorg = true;
                    break;
                }
                if (!get_messages_result.status.ok()) {
                    std::cout << "Error getting messages for execution cursor"
                              << std::endl;
                    return get_messages_result.status;
                }
                execConfig.inbox_messages = std::move(get_messages_result.data);
            }

            auto& mach =
                std::get<std::unique_ptr<Machine>>(execution_cursor.machine);
            mach->machine_state.context = AssertionContext(execConfig);
            auto assertion = mach->run();
            if (assertion.gasCount == 0) {
                break;
            }
        }

        if (handle_reorg) {
            ReadSnapshotTransaction tx(data_storage);

            auto closest_checkpoint =
                getClosestExecutionMachine(tx, total_gas_used);
            if (std::holds_alternative<rocksdb::Status>(closest_checkpoint)) {
                std::cerr << "No execution machine available" << std::endl;
                return std::get<rocksdb::Status>(closest_checkpoint);
            }
            execution_cursor.machine =
                std::get<MachineStateKeys>(closest_checkpoint);
        }
    }

    return rocksdb::Status::OK();
}

std::variant<rocksdb::Status, MachineStateKeys>
ArbCore::getClosestExecutionMachine(ReadTransaction& tx,
                                    const uint256_t& total_gas_used,
                                    bool is_for_sideload) {
    auto target_gas_used = total_gas_used;
    while (true) {
        const std::lock_guard<std::mutex> lock(core_reorg_mutex);
        auto checkpoint_result =
            getCheckpointUsingGas(tx, target_gas_used, false);

        if (std::holds_alternative<rocksdb::Status>(checkpoint_result)) {
            return std::get<rocksdb::Status>(checkpoint_result);
        }

        auto& machine_state_keys =
            std::get<MachineStateKeys>(checkpoint_result);

        return checkpoint_result;
    }
}

ValueResult<std::vector<MachineMessage>> ArbCore::readNextMessages(
    const ReadConsistentTransaction& tx,
    const InboxState& fully_processed_inbox,
    size_t count) const {
    std::vector<MachineMessage> messages(count);

    auto raw_result = getMessagesImpl(tx, fully_processed_inbox.count, count,
                                      fully_processed_inbox.accumulator);
    if (!raw_result.status.ok()) {
        return {raw_result.status, messages};
    }

    for (auto& raw_message : raw_result.data) {
        messages.emplace_back(extractInboxMessage(raw_message.message),
                              raw_message.accumulator);
    }

    return {rocksdb::Status::OK(), messages};
}

bool ArbCore::isValid(const ReadTransaction& tx,
                      const InboxState& fully_processed_inbox) const {
    if (fully_processed_inbox.count == 0) {
        return true;
    }
    auto result =
        getNextSequencerBatchItem(tx, fully_processed_inbox.count - 1);
    return result.status.ok() &&
           result.data.accumulator == fully_processed_inbox.accumulator;
}

ValueResult<uint256_t> ArbCore::logInsertedCount() const {
    ReadTransaction tx(data_storage);

    return logInsertedCountImpl(tx);
}

ValueResult<uint256_t> ArbCore::logInsertedCountImpl(
    const ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(log_inserted_key));
}
rocksdb::Status ArbCore::updateLogInsertedCount(ReadWriteTransaction& tx,
                                                const uint256_t& log_index) {
    std::vector<unsigned char> value;
    marshal_uint256_t(log_index, value);

    return tx.statePut(vecToSlice(log_inserted_key), vecToSlice(value));
}

ValueResult<uint256_t> ArbCore::logProcessedCount(ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(log_processed_key));
}
rocksdb::Status ArbCore::updateLogProcessedCount(ReadWriteTransaction& tx,
                                                 rocksdb::Slice value_slice) {
    return tx.statePut(vecToSlice(log_processed_key), value_slice);
}

ValueResult<uint256_t> ArbCore::sendInsertedCount() const {
    ReadTransaction tx(data_storage);

    return sendInsertedCountImpl(tx);
}

ValueResult<uint256_t> ArbCore::sendInsertedCountImpl(
    const ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(send_inserted_key));
}

rocksdb::Status ArbCore::updateSendInsertedCount(ReadWriteTransaction& tx,
                                                 const uint256_t& send_index) {
    std::vector<unsigned char> value;
    marshal_uint256_t(send_index, value);

    return tx.statePut(vecToSlice(send_inserted_key), vecToSlice(value));
}

ValueResult<uint256_t> ArbCore::sendProcessedCount(ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(send_processed_key));
}
rocksdb::Status ArbCore::updateSendProcessedCount(ReadWriteTransaction& tx,
                                                  rocksdb::Slice value_slice) {
    return tx.statePut(vecToSlice(send_processed_key), value_slice);
}

ValueResult<uint256_t> ArbCore::messageEntryInsertedCount() const {
    ReadTransaction tx(data_storage);

    return messageEntryInsertedCountImpl(tx);
}

ValueResult<uint256_t> ArbCore::messageEntryInsertedCountImpl(
    const ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(message_entry_inserted_key));
}

rocksdb::Status ArbCore::updateMessageEntryInsertedCount(
    ReadWriteTransaction& tx,
    const uint256_t& message_index) {
    std::vector<unsigned char> value;
    marshal_uint256_t(message_index, value);
    return tx.statePut(vecToSlice(message_entry_inserted_key),
                       vecToSlice(value));
}

// deleteLogsStartingAt deletes the given index along with any
// newer logs. Returns std::nullopt if nothing deleted.
std::optional<rocksdb::Status> deleteLogsStartingAt(ReadWriteTransaction& tx,
                                                    uint256_t log_index) {
    auto it = tx.logGetIterator();

    // Find first message to delete
    std::vector<unsigned char> key;
    marshal_uint256_t(log_index, key);
    it->Seek(vecToSlice(key));
    if (it->status().IsNotFound()) {
        // Nothing to delete
        return std::nullopt;
    }
    if (!it->status().ok()) {
        return it->status();
    }

    while (it->Valid()) {
        // Remove reference to value
        auto value_hash_ptr = reinterpret_cast<const char*>(it->value().data());
        deleteValue(tx, deserializeUint256t(value_hash_ptr));

        it->Next();
    }
    if (!it->status().ok()) {
        return it->status();
    }

    return rocksdb::Status::OK();
}

void ArbCore::handleLogsCursorRequested(ReadTransaction& tx,
                                        size_t cursor_index,
                                        ValueCache& cache) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    logs_cursors[cursor_index].data.clear();

    // Provide requested logs
    auto log_inserted_count = logInsertedCountImpl(tx);
    if (!log_inserted_count.status.ok()) {
        logs_cursors[cursor_index].error_string =
            log_inserted_count.status.ToString();
        logs_cursors[cursor_index].status = DataCursor::ERROR;

        std::cerr << "logscursor index " << cursor_index
                  << " error getting inserted count: "
                  << log_inserted_count.status.ToString() << std::endl;
        return;
    }

    auto current_count_result =
        logsCursorGetCurrentTotalCount(tx, cursor_index);
    if (!current_count_result.status.ok()) {
        logs_cursors[cursor_index].error_string =
            current_count_result.status.ToString();
        logs_cursors[cursor_index].status = DataCursor::ERROR;

        std::cerr << "logscursor index" << cursor_index
                  << " error getting cursor current total count: "
                  << current_count_result.status.ToString() << std::endl;
        return;
    }

    if (current_count_result.data == log_inserted_count.data) {
        // No new messages, so don't post any changes
        return;
    }
    if (current_count_result.data > log_inserted_count.data) {
        logs_cursors[cursor_index].error_string =
            "current_count_result greater than log_inserted_count";
        logs_cursors[cursor_index].status = DataCursor::ERROR;

        std::cerr << "handleLogsCursor current count: "
                  << current_count_result.data << " > "
                  << " log inserted count: " << log_inserted_count.data
                  << std::endl;
        return;
    }
    if (current_count_result.data +
            logs_cursors[cursor_index].number_requested >
        log_inserted_count.data) {
        // Too many entries requested
        logs_cursors[cursor_index].number_requested =
            log_inserted_count.data - current_count_result.data;
    }
    if (logs_cursors[cursor_index].number_requested == 0) {
        logs_cursors[cursor_index].status = DataCursor::READY;
        // No new logs to provide
        return;
    }
    auto requested_logs =
        getLogs(current_count_result.data,
                logs_cursors[cursor_index].number_requested, cache);
    if (!requested_logs.status.ok()) {
        logs_cursors[cursor_index].error_string =
            requested_logs.status.ToString();
        logs_cursors[cursor_index].status = DataCursor::ERROR;

        std::cerr << "logscursor index " << cursor_index
                  << " error getting logs: " << requested_logs.status.ToString()
                  << std::endl;
        return;
    }
    logs_cursors[cursor_index].data = std::move(requested_logs.data);
    logs_cursors[cursor_index].status = DataCursor::READY;
}

// handleLogsCursorReorg must be called before logs are deleted.
// Note that this function should not update logs_cursors[cursor_index].status
// because it is happening out of line.
// Note that cursor reorg never adds new messages, but might add deleted
// messages.
rocksdb::Status ArbCore::handleLogsCursorReorg(size_t cursor_index,
                                               uint256_t log_count,
                                               ValueCache& cache) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    ReadWriteTransaction tx(data_storage);

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    auto current_count_result =
        logsCursorGetCurrentTotalCount(tx, cursor_index);
    if (!current_count_result.status.ok()) {
        std::cerr << "Unable to get logs cursor current total count: "
                  << cursor_index << "\n";
        return current_count_result.status;
    }

    if (current_count_result.data >
        logs_cursors[cursor_index].pending_total_count) {
        logs_cursors[cursor_index].pending_total_count =
            current_count_result.data;
    }

    if (log_count < logs_cursors[cursor_index].pending_total_count) {
        // Need to save logs that will be deleted
        auto logs = getLogsNoLock(
            tx, log_count,
            logs_cursors[cursor_index].pending_total_count - log_count, cache);
        if (!logs.status.ok()) {
            std::cerr << "Error getting "
                      << logs_cursors[cursor_index].pending_total_count -
                             log_count
                      << " logs starting at " << log_count
                      << " in Cursor reorg : " << logs.status.ToString()
                      << "\n";
            return logs.status;
        }
        logs_cursors[cursor_index].deleted_data.insert(
            logs_cursors[cursor_index].deleted_data.end(), logs.data.rbegin(),
            logs.data.rend());

        logs_cursors[cursor_index].pending_total_count = log_count;

        if (current_count_result.data > log_count) {
            auto status =
                logsCursorSaveCurrentTotalCount(tx, cursor_index, log_count);
            if (!status.ok()) {
                std::cerr << "unable to save current total count during reorg"
                          << std::endl;
                return status;
            }
        }
    }

    if (!logs_cursors[cursor_index].data.empty()) {
        if (current_count_result.data >= log_count) {
            // Don't save anything
            logs_cursors[cursor_index].data.clear();
        } else if (current_count_result.data +
                       logs_cursors[cursor_index].data.size() >
                   log_count) {
            // Only part of the data needs to be removed
            auto logs_to_keep = intx::narrow_cast<size_t>(
                log_count - current_count_result.data);
            logs_cursors[cursor_index].data.erase(
                logs_cursors[cursor_index].data.begin() + logs_to_keep,
                logs_cursors[cursor_index].data.end());
        }
    }

    if (logs_cursors[cursor_index].status == DataCursor::READY &&
        logs_cursors[cursor_index].data.empty() &&
        logs_cursors[cursor_index].deleted_data.empty()) {
        logs_cursors[cursor_index].status = DataCursor::REQUESTED;
    }

    return tx.commit();
}

bool ArbCore::logsCursorRequest(size_t cursor_index, uint256_t count) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    if (logs_cursors[cursor_index].status != DataCursor::EMPTY) {
        return false;
    }

    logs_cursors[cursor_index].number_requested = count;
    logs_cursors[cursor_index].status = DataCursor::REQUESTED;

    return true;
}

ValueResult<ArbCore::logscursor_logs> ArbCore::logsCursorGetLogs(
    size_t cursor_index) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    if (logs_cursors[cursor_index].status != DataCursor::READY) {
        // No new logs yet
        return {rocksdb::Status::TryAgain(), {}};
    }

    ReadTransaction tx(data_storage);
    auto current_count_result =
        logsCursorGetCurrentTotalCount(tx, cursor_index);
    if (!current_count_result.status.ok()) {
        std::cerr << "logs cursor " << cursor_index
                  << " unable to get current total count: "
                  << current_count_result.status.ToString() << "\n";
        return {current_count_result.status, {}};
    }

    logs_cursors[cursor_index].pending_total_count =
        current_count_result.data + logs_cursors[cursor_index].data.size();

    ArbCore::logscursor_logs logs{};
    logs.first_log_index = current_count_result.data;
    logs.logs = std::move(logs_cursors[cursor_index].data);
    logs.deleted_logs = std::move(logs_cursors[cursor_index].deleted_data);
    logs_cursors[cursor_index].data.clear();
    logs_cursors[cursor_index].deleted_data.clear();

    return {rocksdb::Status::OK(), std::move(logs)};
}

bool ArbCore::logsCursorConfirmReceived(size_t cursor_index) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    if (logs_cursors[cursor_index].status != DataCursor::READY) {
        logs_cursors[cursor_index].error_string =
            "logsCursorConfirmReceived called at wrong state";
        std::cerr << "logsCursorConfirmReceived called at wrong state: "
                  << logs_cursors[cursor_index].status << "\n";
        logs_cursors[cursor_index].status = DataCursor::ERROR;
        return false;
    }

    if (!logs_cursors[cursor_index].data.empty()) {
        // Still have logs to get
        std::cerr << "logs cursor " << cursor_index
                  << " has messages left in cursor when trying to confirm"
                  << std::endl;
        return false;
    }

    if (!logs_cursors[cursor_index].data.empty() ||
        !logs_cursors[cursor_index].deleted_data.empty()) {
        // Still have logs to get
        return false;
    }

    ReadWriteTransaction tx(data_storage);
    auto status = logsCursorSaveCurrentTotalCount(
        tx, cursor_index, logs_cursors[cursor_index].pending_total_count);
    tx.commit();

    logs_cursors[cursor_index].status = DataCursor::EMPTY;

    return true;
}

bool ArbCore::logsCursorCheckError(size_t cursor_index) const {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    return logs_cursors[cursor_index].status == DataCursor::ERROR;
}

ValueResult<uint256_t> ArbCore::logsCursorPosition(size_t cursor_index) const {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    ReadTransaction tx(data_storage);
    return logsCursorGetCurrentTotalCount(tx, cursor_index);
}

std::string ArbCore::logsCursorClearError(size_t cursor_index) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return "Invalid logsCursor index";
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    if (logs_cursors[cursor_index].status != DataCursor::ERROR) {
        std::cerr << "logsCursorClearError called when status not ERROR"
                  << std::endl;
        return "logsCursorClearError called when sttaus not ERROR";
    }

    auto str = logs_cursors[cursor_index].error_string;
    logs_cursors[cursor_index].error_string.clear();
    logs_cursors[cursor_index].data.clear();
    logs_cursors[cursor_index].deleted_data.clear();
    logs_cursors[cursor_index].status = DataCursor::EMPTY;

    return str;
}

rocksdb::Status ArbCore::logsCursorSaveCurrentTotalCount(
    ReadWriteTransaction& tx,
    size_t cursor_index,
    uint256_t count) {
    std::vector<unsigned char> value_data;
    marshal_uint256_t(count, value_data);
    return tx.statePut(vecToSlice(logs_cursors[cursor_index].current_total_key),
                       vecToSlice(value_data));
}

ValueResult<uint256_t> ArbCore::logsCursorGetCurrentTotalCount(
    const ReadTransaction& tx,
    size_t cursor_index) const {
    return tx.stateGetUint256(
        vecToSlice(logs_cursors[cursor_index].current_total_key));
}

rocksdb::Status ArbCore::saveSideloadPosition(ReadWriteTransaction& tx,
                                              const uint256_t& block_number,
                                              const uint256_t& arb_gas_used) {
    std::vector<unsigned char> key;
    marshal_uint256_t(block_number, key);
    auto key_slice = vecToSlice(key);

    std::vector<unsigned char> value;
    marshal_uint256_t(arb_gas_used, value);
    auto value_slice = vecToSlice(value);

    return tx.sideloadPut(key_slice, value_slice);
}

ValueResult<uint256_t> ArbCore::getSideloadPosition(
    ReadTransaction& tx,
    const uint256_t& block_number) {
    std::vector<unsigned char> key;
    marshal_uint256_t(block_number, key);
    auto key_slice = vecToSlice(key);

    auto it = tx.sideloadGetIterator();

    it->SeekForPrev(key_slice);

    auto s = it->status();
    if (!s.ok()) {
        return {s, 0};
    }

    auto value_slice = it->value();

    return {s, intx::be::unsafe::load<uint256_t>(
                   reinterpret_cast<const unsigned char*>(value_slice.data()))};
}

rocksdb::Status ArbCore::deleteSideloadsStartingAt(
    ReadWriteTransaction& tx,
    const uint256_t& block_number) {
    // Clear the cache
    {
        std::unique_lock<std::shared_mutex> guard(sideload_cache_mutex);
        auto it = sideload_cache.lower_bound(block_number);
        while (it != sideload_cache.end()) {
            it = sideload_cache.erase(it);
        }
    }

    // Clear the DB
    std::vector<unsigned char> key;
    marshal_uint256_t(block_number, key);
    auto key_slice = vecToSlice(key);

    auto it = tx.sideloadGetIterator();

    it->Seek(key_slice);

    while (it->Valid()) {
        tx.sideloadDelete(it->key());
        it->Next();
    }
    auto s = it->status();
    if (s.IsNotFound()) {
        s = rocksdb::Status::OK();
    }
    return s;
}

ValueResult<std::unique_ptr<Machine>> ArbCore::getMachineForSideload(
    const uint256_t& block_number,
    ValueCache& cache) {
    // Check the cache
    {
        std::shared_lock<std::shared_mutex> lock(sideload_cache_mutex);
        // Look for the first value after the value we want
        auto it = sideload_cache.upper_bound(block_number);
        if (it != sideload_cache.begin()) {
            // Go back a value to find the one we want
            it--;
            return {rocksdb::Status::OK(),
                    std::make_unique<Machine>(*it->second)};
        }
    }

    uint256_t gas_target;
    std::unique_ptr<ExecutionCursor> execution_cursor;
    {
        // Not found in cache, try the DB
        ReadSnapshotTransaction tx(data_storage);
        auto position_res = getSideloadPosition(tx, block_number);
        if (!position_res.status.ok()) {
            return {position_res.status, std::unique_ptr<Machine>(nullptr)};
        }

        auto closest_checkpoint =
            getClosestExecutionMachine(tx, position_res.data, true);
        if (std::holds_alternative<rocksdb::Status>(closest_checkpoint)) {
            return {std::get<rocksdb::Status>(closest_checkpoint), nullptr};
        }

        gas_target = position_res.data;
        execution_cursor = std::make_unique<ExecutionCursor>(
            std::get<MachineStateKeys>(closest_checkpoint));
    }

    auto status = advanceExecutionCursorImpl(*execution_cursor, gas_target,
                                             false, 10, cache);

    ReadSnapshotTransaction tx(data_storage);
    return {status,
            takeExecutionCursorMachineImpl(tx, *execution_cursor, cache)};
}
