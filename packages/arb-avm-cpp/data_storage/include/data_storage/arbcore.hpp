/*
 * Copyright 2020, Offchain Labs, Inc.
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

#ifndef arbcore_hpp
#define arbcore_hpp

#include <avm/machine.hpp>
#include <avm_values/bigint.hpp>
#include <data_storage/checkpoint.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/messageentry.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/code.hpp>
#include <data_storage/value/valuecache.hpp>
#include <utility>

#include <avm/machinethread.hpp>
#include <map>
#include <memory>
#include <queue>
#include <shared_mutex>
#include <thread>
#include <utility>
#include <vector>
#include "datacursor.hpp"
#include "executioncursor.hpp"

namespace rocksdb {
class TransactionDB;
class Status;
struct Slice;
class ColumnFamilyHandle;
}  // namespace rocksdb

class ArbCore {
   public:
    typedef enum {
        MESSAGES_EMPTY,       // Out: Ready to receive messages
        MESSAGES_READY,       // In:  Messages in vector
        MESSAGES_SUCCESS,     // Out:  Messages processed successfully
        MESSAGES_NEED_OLDER,  // Out: Last message invalid, need older messages
        MESSAGES_ERROR        // Out: Error processing messages
    } message_status_enum;

   private:
    struct message_data_struct {
        std::vector<std::vector<unsigned char>> messages;
        uint256_t previous_inbox_acc;
        bool last_block_complete{false};
        std::optional<uint256_t> reorg_message_count;
    };

   private:
    std::unique_ptr<std::thread> core_thread;

    // Core thread input
    std::atomic<bool> arbcore_abort{false};

    // Core thread holds mutex only during reorg.
    // Routines accessing database for log entries will need to acquire mutex
    // because obsolete log entries have `Value` references removed causing
    // reference counts to be decremented and possibly deleted.
    // No mutex required to access Sends or Messages because obsolete entries
    // are not deleted.
    std::mutex core_reorg_mutex;
    std::shared_ptr<DataStorage> data_storage;

    std::unique_ptr<MachineThread> machine;
    std::shared_ptr<Code> code{};
    Checkpoint pending_checkpoint;

    // Cache a machine ready to sideload view transactions just after recent
    // blocks
    std::shared_mutex sideload_cache_mutex;
    std::map<uint256_t, std::unique_ptr<Machine>> sideload_cache;

    // Core thread inbox status input/output. Core thread will update if and
    // only if set to MESSAGES_READY
    std::atomic<message_status_enum> message_data_status{MESSAGES_EMPTY};

    // Core thread inbox input
    message_data_struct message_data;

    // Core thread inbox output
    std::string core_error_string;

    // Core thread logs output
    std::vector<DataCursor> logs_cursors{1};

    // Core thread machine state output
    std::atomic<bool> machine_idle{false};
    std::atomic<bool> machine_error{false};
    std::string machine_error_string;

   public:
    ArbCore() = delete;
    explicit ArbCore(std::shared_ptr<DataStorage> data_storage_);

    ~ArbCore() { abortThread(); }
    rocksdb::Status initialize(const LoadedExecutable& executable);
    bool initialized() const;
    void operator()();

   public:
    // Public Thread interaction
    bool startThread();
    void abortThread();

   private:
    // Private database interaction
    ValueResult<uint256_t> getInitialMachineHash(Transaction& tx);
    rocksdb::Status saveAssertion(Transaction& tx, const Assertion& assertion);
    ValueResult<Checkpoint> getCheckpoint(Transaction& tx,
                                          const uint256_t& arb_gas_used) const;
    rocksdb::Status resolveStagedMessage(Transaction& tx,
                                         value& message,
                                         uint256_t& inbox_acc) const;
    ValueResult<Checkpoint> getCheckpointUsingGas(Transaction& tx,
                                                  const uint256_t& total_gas,
                                                  bool after_gas);
    rocksdb::Status reorgToMessageOrBefore(
        Transaction& tx,
        const uint256_t& message_sequence_number,
        bool use_latest,
        ValueCache& cache);
    template <class T>
    std::unique_ptr<T> getMachineUsingStateKeys(Transaction& transaction,
                                                MachineStateKeys state_data,
                                                ValueCache& value_cache);

   public:
    // To be deprecated, use checkpoints instead
    template <class T>
    std::unique_ptr<T> getInitialMachine(ValueCache& value_cache);
    template <class T>
    std::unique_ptr<T> getInitialMachineImpl(Transaction& tx,
                                             ValueCache& value_cache);
    template <class T>
    std::unique_ptr<T> getMachine(uint256_t machineHash,
                                  ValueCache& value_cache);
    template <class T>
    std::unique_ptr<T> getMachineImpl(Transaction& tx,
                                      uint256_t machineHash,
                                      ValueCache& value_cache);

   public:
    // Useful for unit tests
    rocksdb::Status saveCheckpoint(Transaction& tx);
    bool isCheckpointsEmpty(Transaction& tx) const;
    uint256_t maxCheckpointGas();

   public:
    // Managing machine state
    bool machineIdle();
    std::optional<std::string> machineClearError();

   public:
    // Sending messages to core thread
    bool deliverMessages(std::vector<std::vector<unsigned char>>& messages,
                         const uint256_t& previous_inbox_acc,
                         bool last_block_complete,
                         const std::optional<uint256_t>& reorg_height);
    message_status_enum messagesStatus();
    std::string messagesClearError();

   public:
    // Logs Cursor interaction
    bool logsCursorRequest(size_t cursor_index, uint256_t count);
    std::optional<std::pair<uint256_t, std::vector<value>>> logsCursorGetLogs(
        size_t cursor_index);
    std::optional<std::pair<uint256_t, std::vector<value>>>
    logsCursorGetDeletedLogs(size_t cursor_index);
    bool logsCursorCheckError(size_t cursor_index) const;
    std::string logsCursorClearError(size_t cursor_index);
    bool logsCursorConfirmReceived(size_t cursor_index);

   private:
    // Logs cursor internal functions
    bool handleLogsCursorRequested(Transaction& tx,
                                   size_t cursor_index,
                                   ValueCache& cache);
    rocksdb::Status handleLogsCursorReorg(Transaction& tx,
                                          size_t cursor_index,
                                          uint256_t log_count,
                                          ValueCache& cache);

   public:
    // Execution Cursor interaction
    ValueResult<std::unique_ptr<ExecutionCursor>> getExecutionCursor(
        uint256_t total_gas_used,
        ValueCache& cache);
    rocksdb::Status advanceExecutionCursor(ExecutionCursor& execution_cursor,
                                           uint256_t max_gas,
                                           bool go_over_gas,
                                           ValueCache& cache);

   private:
    // Execution cursor internal functions
    rocksdb::Status getExecutionCursorImpl(Transaction& tx,
                                           ExecutionCursor& execution_cursor,
                                           uint256_t total_gas_used,
                                           bool go_over_gas,
                                           uint256_t message_group_size,
                                           ValueCache& cache);

   public:
    // Public database interaction
    std::unique_ptr<Transaction> makeTransaction();
    std::unique_ptr<const Transaction> makeConstTransaction() const;

    ValueResult<uint256_t> logInsertedCount() const;
    ValueResult<uint256_t> sendInsertedCount() const;
    ValueResult<uint256_t> messageEntryInsertedCount() const;
    ValueResult<uint256_t> messageEntryProcessedCount() const;
    ValueResult<std::vector<value>> getLogs(uint256_t index,
                                            uint256_t count,
                                            ValueCache& valueCache);
    ValueResult<std::vector<std::vector<unsigned char>>> getSends(
        uint256_t index,
        uint256_t count) const;

    ValueResult<std::vector<std::vector<unsigned char>>> getMessages(
        uint256_t index,
        uint256_t count) const;
    ValueResult<std::pair<std::vector<std::vector<unsigned char>>,
                          std::optional<uint256_t>>>
    getMessagesImpl(Transaction& tx, uint256_t index, uint256_t count) const;
    ValueResult<uint256_t> getInboxAcc(uint256_t index);
    ValueResult<std::pair<uint256_t, uint256_t>> getInboxAccPair(
        uint256_t index1,
        uint256_t index2);
    ValueResult<uint256_t> getSendAcc(uint256_t start_acc_hash,
                                      uint256_t start_index,
                                      uint256_t count) const;
    ValueResult<uint256_t> getLogAcc(uint256_t start_acc_hash,
                                     uint256_t start_index,
                                     uint256_t count,
                                     ValueCache& cache);

   private:
    // Private database interaction
    ValueResult<MessageEntry> getMessageEntry(
        Transaction& tx,
        uint256_t message_sequence_number) const;
    ValueResult<uint256_t> logInsertedCountImpl(Transaction& tx) const;

    ValueResult<uint256_t> logProcessedCount(Transaction& tx) const;
    rocksdb::Status updateLogProcessedCount(Transaction& tx,
                                            rocksdb::Slice value_slice);
    ValueResult<uint256_t> sendInsertedCountImpl(Transaction& tx) const;

    ValueResult<uint256_t> sendProcessedCount(Transaction& tx) const;
    rocksdb::Status updateSendProcessedCount(Transaction& tx,
                                             rocksdb::Slice value_slice);
    ValueResult<uint256_t> messageEntryInsertedCountImpl(Transaction& tx) const;

    ValueResult<uint256_t> messageEntryProcessedCountImpl(
        Transaction& tx) const;

    rocksdb::Status saveLogs(Transaction& tx, const std::vector<value>& val);
    rocksdb::Status saveSends(
        Transaction& tx,
        const std::vector<std::vector<unsigned char>>& send);

   private:
    std::optional<rocksdb::Status> addMessages(
        const std::vector<std::vector<unsigned char>>& new_messages,
        bool last_block_complete,
        const uint256_t& prev_inbox_acc,
        const uint256_t& message_count_in_machine,
        const std::optional<uint256_t>& reorg_message_count,
        ValueCache& cache);
    ValueResult<std::vector<value>> getLogsNoLock(Transaction& tx,
                                                  uint256_t index,
                                                  uint256_t count,
                                                  ValueCache& valueCache);

    ValueResult<bool> executionCursorAddMessages(
        Transaction& tx,
        ExecutionCursor& execution_cursor,
        const uint256_t& orig_message_group_size);
    ValueResult<bool> executionCursorAddMessagesNoLock(
        Transaction& tx,
        ExecutionCursor& execution_cursor,
        const uint256_t& orig_message_group_size);
    rocksdb::Status executionCursorSetup(Transaction& tx,
                                         ExecutionCursor& execution_cursor,
                                         const uint256_t& total_gas_used,
                                         ValueCache& cache,
                                         bool is_for_sideload = false);

    rocksdb::Status updateLogInsertedCount(Transaction& tx,
                                           const uint256_t& log_index);
    rocksdb::Status updateSendInsertedCount(Transaction& tx,
                                            const uint256_t& send_index);
    rocksdb::Status updateMessageEntryProcessedCount(
        Transaction& tx,
        const uint256_t& message_index);
    rocksdb::Status updateMessageEntryInsertedCount(
        Transaction& tx,
        const uint256_t& message_index);

   public:
    // Public sideload interaction
    ValueResult<std::unique_ptr<Machine>> getMachineForSideload(
        const uint256_t& block_number,
        ValueCache& cache);

   private:
    // Private sideload interaction
    rocksdb::Status saveSideloadPosition(Transaction& tx,
                                         const uint256_t& block_number);
    ValueResult<uint256_t> getSideloadPosition(Transaction& tx,
                                               const uint256_t& block_number);
    rocksdb::Status deleteSideloadsStartingAt(Transaction& tx,
                                              const uint256_t& block_number);
    rocksdb::Status logsCursorSaveCurrentTotalCount(Transaction& tx,
                                                    size_t cursor_index,
                                                    uint256_t count);
    ValueResult<uint256_t> logsCursorGetCurrentTotalCount(Transaction& tx,
                                                          size_t cursor_index);
};

std::optional<rocksdb::Status> deleteLogsStartingAt(Transaction& tx,
                                                    uint256_t log_index);

#endif /* arbcore_hpp */
