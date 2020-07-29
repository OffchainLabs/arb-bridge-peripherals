/*
 * Copyright 2019, Offchain Labs, Inc.
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

#include <avm/machine.hpp>

#include <secp256k1_recovery.h>
#include <ethash/keccak.hpp>

#define CATCH_CONFIG_ENABLE_BENCHMARKING 1
#include <catch2/catch.hpp>

using namespace intx;

MachineState runUnaryOp(uint256_t arg1, OpCode op) {
    MachineState m;
    m.stack.push(arg1);
    m.runOp(op);
    return m;
}

void testUnaryOp(uint256_t arg1, uint256_t result, OpCode op) {
    MachineState m = runUnaryOp(arg1, op);
    value res = m.stack.pop();
    auto num = nonstd::get_if<uint256_t>(&res);
    REQUIRE(num);
    REQUIRE(*num == result);
    REQUIRE(m.stack.stacksize() == 0);
}

MachineState runBinaryOp(uint256_t arg1, uint256_t arg2, OpCode op) {
    MachineState m;
    m.stack.push(arg2);
    m.stack.push(arg1);
    m.runOp(op);
    return m;
}

void testBinaryOp(uint256_t arg1,
                  uint256_t arg2,
                  uint256_t expected,
                  OpCode op) {
    MachineState m = runBinaryOp(arg1, arg2, op);
    value res = m.stack.pop();
    auto actual = nonstd::get_if<uint256_t>(&res);
    REQUIRE(actual);
    REQUIRE(*actual == expected);
    REQUIRE(m.stack.stacksize() == 0);
}

MachineState runTertiaryOp(uint256_t arg1,
                           uint256_t arg2,
                           uint256_t arg3,
                           OpCode op) {
    MachineState m;
    m.stack.push(arg3);
    m.stack.push(arg2);
    m.stack.push(arg1);
    m.runOp(op);
    return m;
}

void testTertiaryOp(uint256_t arg1,
                    uint256_t arg2,
                    uint256_t arg3,
                    uint256_t result,
                    OpCode op) {
    MachineState m = runTertiaryOp(arg1, arg2, arg3, op);
    value res = m.stack.pop();
    auto num = nonstd::get_if<uint256_t>(&res);
    REQUIRE(num);
    REQUIRE(*num == result);
    REQUIRE(m.stack.stacksize() == 0);
}

TEST_CASE("ADD opcode is correct") {
    SECTION("Non overlow is correct") {
        testBinaryOp(4_u256, 3_u256, 7_u256, OpCode::ADD);
    }

    SECTION("0+0 is correct") {
        testBinaryOp(0_u256, 0_u256, 0_u256, OpCode::ADD);
    }

    SECTION("Overlow is correct") {
        testBinaryOp(-1_u256, 4_u256, 3_u256, OpCode::ADD);
    }

    SECTION("-2+1 is correct") {
        testBinaryOp(-2_u256, 1_u256, -1_u256, OpCode::ADD);
    }
}

TEST_CASE("MUL opcode is correct") {
    SECTION("Non overlow is correct") {
        testBinaryOp(4_u256, 3_u256, 12_u256, OpCode::MUL);
    }

    SECTION("3*0 is correct") {
        testBinaryOp(3_u256, 0_u256, 0_u256, OpCode::MUL);
    }

    SECTION("-1*1 is correct") {
        testBinaryOp(-1_u256, 1_u256, -1_u256, OpCode::MUL);
    }

    SECTION("-2+1 is correct") {
        testBinaryOp(-2_u256, 1_u256, -2_u256, OpCode::MUL);
    }
}

TEST_CASE("SUB opcode is correct") {
    SECTION("Non overlow is correct") {
        testBinaryOp(4_u256, 3_u256, 1_u256, OpCode::SUB);
    }

    SECTION("Overlow is correct") {
        testBinaryOp(3_u256, 4_u256, -1_u256, OpCode::SUB);
    }
}

TEST_CASE("DIV opcode is correct") {
    SECTION("Non overlow is correct") {
        testBinaryOp(12_u256, 3_u256, 4_u256, OpCode::DIV);
    }

    SECTION("unsigned division is correct") {
        MachineState m = runBinaryOp(-6_u256, 2_u256, OpCode::DIV);
        value res = m.stack.pop();
        auto num = nonstd::get_if<uint256_t>(&res);
        REQUIRE(num);
        REQUIRE(*num != 3);
        REQUIRE(m.stack.stacksize() == 0);
    }

    SECTION("Divide by zero") {
        MachineState m = runBinaryOp(3_u256, 0_u256, OpCode::DIV);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("SDIV opcode is correct") {
    SECTION("Positive divided by positive") {
        testBinaryOp(12_u256, 3_u256, 4_u256, OpCode::SDIV);
    }
    SECTION("Positive divided by negative") {
        testBinaryOp(12_u256, -3_u256, -4_u256, OpCode::SDIV);
    }
    SECTION("Negative divided by positive") {
        testBinaryOp(-12_u256, 3_u256, -4_u256, OpCode::SDIV);
    }
    SECTION("Negative divided by negative") {
        testBinaryOp(-12_u256, -3_u256, 4_u256, OpCode::SDIV);
    }
    SECTION("Divide by zero") {
        MachineState m = runBinaryOp(3_u256, 0_u256, OpCode::SDIV);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("MOD opcode is correct") {
    SECTION("8 mod 3") { testBinaryOp(8_u256, 3_u256, 2_u256, OpCode::MOD); }

    SECTION("8 mod very large") {
        testBinaryOp(8_u256, -3_u256, 8_u256, OpCode::MOD);
    }

    SECTION("0 mod 3") { testBinaryOp(0_u256, 3_u256, 0_u256, OpCode::MOD); }

    SECTION("Mod by zero") {
        MachineState m = runBinaryOp(3_u256, 0_u256, OpCode::MOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("SMOD opcode is correct") {
    SECTION("Positive mod positive") {
        testBinaryOp(8_u256, 3_u256, 2_u256, OpCode::SMOD);
    }

    SECTION("Positive mod negative") {
        testBinaryOp(8_u256, -3_u256, 2_u256, OpCode::SMOD);
    }

    SECTION("Negative mod positive") {
        testBinaryOp(-8_u256, 3_u256, -2_u256, OpCode::SMOD);
    }

    SECTION("Negative mod negative") {
        testBinaryOp(-8_u256, -3_u256, -2_u256, OpCode::SMOD);
    }
    SECTION("Mod by zero") {
        MachineState m = runBinaryOp(3, 0, OpCode::SMOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("ADDMOD opcode is correct") {
    SECTION("(8+5)%3") {
        testTertiaryOp(8_u256, 5_u256, 3_u256, 1_u256, OpCode::ADDMOD);
    }

    SECTION("(-1+1)%7") {
        testTertiaryOp(-1_u256, 1_u256, 7_u256, 2_u256, OpCode::ADDMOD);
    }

    SECTION("(0+0)%7") {
        testTertiaryOp(0_u256, 0_u256, 7_u256, 0_u256, OpCode::ADDMOD);
    }

    SECTION("(3+3)%-4") {
        testTertiaryOp(3_u256, 3_u256, -4_u256, 6_u256, OpCode::ADDMOD);
    }

    SECTION("Mod by zero") {
        MachineState m = runTertiaryOp(8_u256, 3_u256, 0_u256, OpCode::ADDMOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("MULMOD opcode is correct") {
    SECTION("(8*2)%3") {
        testTertiaryOp(8_u256, 2_u256, 3_u256, 1_u256, OpCode::MULMOD);
    }

    SECTION("(-1*2)%7") {
        testTertiaryOp(-1_u256, 2_u256, 7_u256, 2_u256, OpCode::MULMOD);
    }

    SECTION("(0*0)%7") {
        testTertiaryOp(0_u256, 0_u256, 7_u256, 0_u256, OpCode::MULMOD);
    }

    SECTION("Mod by zero") {
        MachineState m = runTertiaryOp(8_u256, 3_u256, 0_u256, OpCode::MULMOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("EXP opcode is correct") {
    SECTION("All positive") {
        testBinaryOp(3_u256, 2_u256, 9_u256, OpCode::EXP);
    }
    SECTION("wrap") { testBinaryOp(2_u256, 256_u256, 0_u256, OpCode::EXP); }
}

TEST_CASE("LT opcode is correct") {
    SECTION("less") { testBinaryOp(3, 9, 1, OpCode::LT); }
    SECTION("greater") { testBinaryOp(9, 3, 0, OpCode::LT); }
    SECTION("equal") { testBinaryOp(3, 3, 0, OpCode::LT); }
    SECTION("First neg, second pos") { testBinaryOp(-3, 9, 0, OpCode::LT); }
}

TEST_CASE("GT opcode is correct") {
    testBinaryOp(3, 9, 0, OpCode::GT);
    testBinaryOp(9, 3, 1, OpCode::GT);
    testBinaryOp(3, 3, 0, OpCode::GT);
    testBinaryOp(-3, 9, 1, OpCode::GT);

    BENCHMARK_ADVANCED("gt 100x")(Catch::Benchmark::Chronometer meter) {
        MachineState sample_machine;
        for (int i = 0; i < 101; i++) {
            sample_machine.stack.push(uint256_t{100});
        }
        std::vector<MachineState> machines(meter.runs());
        std::fill(machines.begin(), machines.end(), sample_machine);
        meter.measure([&machines](int i) {
            auto& mach = machines[i];
            for (int i = 0; i < 100; i++) {
                mach.runOp(OpCode::GT);
            }
            return mach;
        });
    };
}

TEST_CASE("SLT opcode is correct") {
    SECTION("All positive true") {
        testBinaryOp(7_u256, 3_u256, 0_u256, OpCode::SLT);
    }
    SECTION("All positive false") {
        testBinaryOp(3_u256, 7_u256, 1_u256, OpCode::SLT);
    }
    SECTION("All negative true") {
        testBinaryOp(-3_u256, -7_u256, 0_u256, OpCode::SLT);
    }
    SECTION("All negative false") {
        testBinaryOp(-7_u256, -3_u256, 1_u256, OpCode::SLT);
    }
    SECTION("First pos, second neg true") {
        testBinaryOp(3_u256, -7_u256, 0_u256, OpCode::SLT);
    }
    SECTION("First neg, second pos false") {
        testBinaryOp(-3_u256, 7_u256, 1_u256, OpCode::SLT);
    }
    SECTION("equal") { testBinaryOp(3_u256, 3_u256, 0_u256, OpCode::SLT); }
}

TEST_CASE("SGT opcode is correct") {
    SECTION("All positive true") {
        testBinaryOp(7_u256, 3_u256, 1_u256, OpCode::SGT);
    }
    SECTION("All positive false") {
        testBinaryOp(3_u256, 7_u256, 0_u256, OpCode::SGT);
    }
    SECTION("All negative true") {
        testBinaryOp(-3_u256, -7_u256, 1_u256, OpCode::SGT);
    }
    SECTION("All negative false") {
        testBinaryOp(-7_u256, -3_u256, 0_u256, OpCode::SGT);
    }
    SECTION("First pos, second neg true") {
        testBinaryOp(3_u256, -7_u256, 1_u256, OpCode::SGT);
    }
    SECTION("First neg, second pos false") {
        testBinaryOp(-7_u256, 3_u256, 0_u256, OpCode::SGT);
    }
    SECTION("equal") { testBinaryOp(3_u256, 3_u256, 0_u256, OpCode::SGT); }
}

TEST_CASE("EQ opcode is correct") {
    SECTION("Not equal") { testBinaryOp(7_u256, 3_u256, 0_u256, OpCode::EQ); }
    SECTION("equal") { testBinaryOp(3_u256, 3_u256, 1_u256, OpCode::EQ); }
    SECTION("matching tuples") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}, m.pool.get()});
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}, m.pool.get()});
        m.runOp(OpCode::EQ);
        value res = m.stack.pop();
        auto actual = nonstd::get_if<uint256_t>(&res);
        REQUIRE(actual);
        REQUIRE(*actual == 1);
        REQUIRE(m.stack.stacksize() == 0);
    }
    SECTION("different tuples") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}, m.pool.get()});
        m.stack.push(Tuple{uint256_t{1}, uint256_t{3}, m.pool.get()});
        m.runOp(OpCode::EQ);
        value res = m.stack.pop();
        auto actual = nonstd::get_if<uint256_t>(&res);
        REQUIRE(actual);
        REQUIRE(*actual == 0);
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("ISZERO opcode is correct") {
    SECTION("true") { testUnaryOp(0, 1, OpCode::ISZERO); }
    SECTION("false") { testUnaryOp(2, 0, OpCode::ISZERO); }
}

TEST_CASE("AND opcode is correct") {
    SECTION("3 and 9 = 1") { testBinaryOp(3, 9, 1, OpCode::BITWISE_AND); }
    SECTION("3 and 3 = 3") { testBinaryOp(3, 3, 3, OpCode::BITWISE_AND); }
}

TEST_CASE("OR opcode is correct") {
    SECTION("3 or 9 = 11") { testBinaryOp(3, 9, 11, OpCode::BITWISE_OR); }
    SECTION("3 or 3 = 3") { testBinaryOp(3, 3, 3, OpCode::BITWISE_OR); }
}

TEST_CASE("XOR opcode is correct") {
    SECTION("3 or 9 = 11") { testBinaryOp(3, 9, 10, OpCode::BITWISE_XOR); }
    SECTION("3 or 3 = 3") { testBinaryOp(3, 3, 0, OpCode::BITWISE_XOR); }
}

TEST_CASE("NOT opcode is correct") {
    SECTION("0") { testUnaryOp(0_u256, -1_u256, OpCode::BITWISE_NOT); }
    SECTION("3") { testUnaryOp(3_u256, -4_u256, OpCode::BITWISE_NOT); }
    SECTION("-4") { testUnaryOp(-4_u256, 3_u256, OpCode::BITWISE_NOT); }
}

TEST_CASE("BYTE opcode is correct") {
    SECTION("31st byte of 16 = 16") { testBinaryOp(16, 31, 16, OpCode::BYTE); }
    SECTION("3rd byte of 3 = 0") { testBinaryOp(3, 3, 0, OpCode::BYTE); }
}

TEST_CASE("SIGNEXTEND opcode is correct") {
    SECTION("test1") {
        testBinaryOp(-1_u256, 0_u256, -1_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test2") {
        testBinaryOp(1_u256, 0_u256, 1_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test3") {
        testBinaryOp(127_u256, 0_u256, 127_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test4") {
        testBinaryOp(128_u256, 0_u256, -128_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test5") {
        testBinaryOp(254_u256, 0_u256, -2_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test6") {
        testBinaryOp(257_u256, 0_u256, 1_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test7") {
        testBinaryOp(65534_u256, 1_u256, -2_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test8") {
        testBinaryOp(65537_u256, 1_u256, 1_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test8") {
        testBinaryOp(65537_u256, 2_u256, 65537_u256, OpCode::SIGNEXTEND);
    }
}

TEST_CASE("HASH opcode is correct") {
    SECTION("10") {
        testUnaryOp(
            10,
            intx::from_string<uint256_t>("0xc65a7bb8d6351c1cf70c95a316cc6a92839"
                                         "c986682d98bc35f958f4883f9d2a8"),
            OpCode::HASH);
    }
}

TEST_CASE("TYPE opcode is correct") {
    SECTION("type int") {
        MachineState m;
        m.stack.push(value{uint256_t(3)});
        REQUIRE(m.stack.stacksize() == 1);
        m.runOp(OpCode::TYPE);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(0)});
        REQUIRE(m.stack.stacksize() == 0);
    }
    SECTION("type codepoint stub") {
        MachineState m;
        m.stack.push(value{CodePointStub({0, 0}, 0)});
        REQUIRE(m.stack.stacksize() == 1);
        m.runOp(OpCode::TYPE);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(1)});
        REQUIRE(m.stack.stacksize() == 0);
    }
    SECTION("type tuple") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}, m.pool.get()});
        REQUIRE(m.stack.stacksize() == 1);
        m.runOp(OpCode::TYPE);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("POP opcode is correct") {
    SECTION("pop") {
        MachineState m;
        m.stack.push(uint256_t{3});
        REQUIRE(m.stack.stacksize() == 1);
        m.runOp(OpCode::POP);
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("SPUSH opcode is correct") {
    SECTION("spush") {
        auto pool = std::make_shared<TuplePool>();
        auto code = std::make_shared<Code>();
        code->addSegment();
        MachineState m{std::move(code), uint256_t{5}, pool};
        m.runOp(OpCode::SPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("RPUSH opcode is correct") {
    SECTION("rpush") {
        MachineState m;
        m.registerVal = uint256_t(5);
        m.runOp(OpCode::RPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("RSET opcode is correct") {
    SECTION("rset") {
        MachineState m;
        m.stack.push(value{uint256_t(5)});
        m.runOp(OpCode::RSET);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.registerVal == value{uint256_t(5)});
    }
}

TEST_CASE("JUMP opcode is correct") {
    SECTION("jump") {
        MachineState m;
        CodePointRef cpr{0, 2};
        m.stack.push(value{CodePointStub(cpr, 73665)});
        m.runOp(OpCode::JUMP);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == cpr);
    }
}

TEST_CASE("CJUMP opcode is correct") {
    SECTION("cjump true") {
        MachineState m;
        CodePointRef cpr{0, 2};
        m.pc = {0, 3};
        m.stack.push(uint256_t{1});
        m.stack.push(value{CodePointStub(cpr, 73665)});
        m.runOp(OpCode::CJUMP);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == cpr);
    }
    SECTION("cjump false") {
        MachineState m;
        CodePointRef initial_pc{0, 3};
        m.pc = initial_pc;
        m.stack.push(uint256_t{0});
        m.stack.push(value{CodePointStub({0, 10}, 73665)});
        m.runOp(OpCode::CJUMP);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == initial_pc + 1);
    }
}

TEST_CASE("STACKEMPTY opcode is correct") {
    SECTION("empty") {
        MachineState m;
        m.runOp(OpCode::STACKEMPTY);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(1)});
        REQUIRE(m.stack.stacksize() == 0);
    }
    SECTION("not empty") {
        MachineState m;
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::STACKEMPTY);
        REQUIRE(m.stack.stacksize() == 2);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(0)});
        REQUIRE(m.stack.stacksize() == 1);
    }
}

TEST_CASE("PCPUSH opcode is correct") {
    SECTION("pcpush") {
        auto pool = std::make_shared<TuplePool>();
        auto code = std::make_shared<Code>();
        auto stub = code->addSegment();
        code->addOperation(stub.pc, Operation(OpCode::ADD));
        MachineState m{std::move(code), uint256_t(5), pool};
        auto initial_stub = CodePointStub(m.pc, m.loadCurrentInstruction());
        m.runOp(OpCode::PCPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        REQUIRE(m.pc == stub.pc);
        value res = m.stack.pop();
        REQUIRE(res == value{initial_stub});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("AUXPUSH opcode is correct") {
    SECTION("auxpush") {
        MachineState m;
        m.stack.push(value{uint256_t(5)});
        m.runOp(OpCode::AUXPUSH);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.auxstack.stacksize() == 1);
        value res = m.auxstack.pop();
        REQUIRE(res == value{uint256_t(5)});
    }
}

TEST_CASE("AUXPOP opcode is correct") {
    SECTION("auxpop") {
        MachineState m;
        m.auxstack.push(value{uint256_t(5)});
        m.runOp(OpCode::AUXPOP);
        REQUIRE(m.auxstack.stacksize() == 0);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
    }
}

TEST_CASE("AUXSTACKEMPTY opcode is correct") {
    SECTION("empty") {
        MachineState m;
        m.runOp(OpCode::AUXSTACKEMPTY);
        REQUIRE(m.auxstack.stacksize() == 0);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(1)});
    }
    SECTION("not empty") {
        MachineState m;
        m.auxstack.push(value{uint256_t(5)});
        m.runOp(OpCode::AUXSTACKEMPTY);
        REQUIRE(m.auxstack.stacksize() == 1);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(0)});
    }
}

MachineState createTestMachineState(OpCode op) {
    auto code = std::make_shared<Code>();
    auto stub = code->addSegment();
    stub = code->addOperation(stub.pc, {OpCode::HALT});
    code->addOperation(stub.pc, {op});
    auto pool = std::make_shared<TuplePool>();
    return {std::move(code), Tuple(), std::move(pool)};
}

TEST_CASE("NOP opcode is correct") {
    SECTION("nop") {
        MachineState m = createTestMachineState(OpCode::NOP);
        auto start_pc = m.pc;
        m.runOne();
        REQUIRE(m.auxstack.stacksize() == 0);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == start_pc + 1);
    }
}

TEST_CASE("ERRPUSH opcode is correct") {
    SECTION("errpush") {
        auto pool = std::make_shared<TuplePool>();
        auto code = std::make_shared<Code>();
        auto stub = code->addSegment();
        stub = code->addOperation(stub.pc, Operation(OpCode::ADD));
        MachineState m{std::move(code), uint256_t(5), pool};
        m.errpc = stub;
        m.runOp(OpCode::ERRPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        REQUIRE(m.pc == CodePointRef{0, 0});
        value res = m.stack.pop();
        REQUIRE(res == value{stub});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("ERRSET opcode is correct") {
    SECTION("errset") {
        MachineState m = createTestMachineState(OpCode::ERRSET);
        auto start_pc = m.pc;
        auto new_err_stub = CodePointStub({0, 54}, 968967);
        m.stack.push(value{new_err_stub});
        m.runOne();
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == start_pc + 1);
        REQUIRE(m.errpc == new_err_stub);
    }
}

TEST_CASE("DUP0 opcode is correct") {
    SECTION("dup") {
        MachineState m = createTestMachineState(OpCode::DUP0);
        auto start_pc = m.pc;
        m.stack.push(uint256_t{5});
        m.runOne();
        REQUIRE(m.stack.stacksize() == 2);
        REQUIRE(m.pc == start_pc + 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
    }
}

TEST_CASE("DUP1 opcode is correct") {
    SECTION("dup") {
        MachineState m = createTestMachineState(OpCode::DUP1);
        auto start_pc = m.pc;
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOne();
        REQUIRE(m.stack.stacksize() == 3);
        REQUIRE(m.pc == start_pc + 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
    }
}

TEST_CASE("DUP2 opcode is correct") {
    SECTION("dup") {
        MachineState m = createTestMachineState(OpCode::DUP2);
        auto start_pc = m.pc;
        m.stack.push(uint256_t{7});
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::DUP2);
        REQUIRE(m.stack.stacksize() == 4);
        REQUIRE(m.pc == start_pc + 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(7)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(7)});
    }
}

TEST_CASE("SWAP1 opcode is correct") {
    SECTION("swap") {
        MachineState m = createTestMachineState(OpCode::SWAP1);
        auto start_pc = m.pc;
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::SWAP1);
        REQUIRE(m.stack.stacksize() == 2);
        REQUIRE(m.pc == start_pc + 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
    }
}

TEST_CASE("SWAP2 opcode is correct") {
    SECTION("swap") {
        MachineState m = createTestMachineState(OpCode::SWAP2);
        auto start_pc = m.pc;
        m.stack.push(uint256_t{7});
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::SWAP2);
        REQUIRE(m.stack.stacksize() == 3);
        REQUIRE(m.pc == start_pc + 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(7)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
    }
}

TEST_CASE("TGET opcode is correct") {
    SECTION("tget") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7},
                           uint256_t{6}, m.pool.get()});
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::TGET);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(8)});
        REQUIRE(m.stack.stacksize() == 0);
    }

    SECTION("index out range") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7},
                           uint256_t{6}, m.pool.get()});
        m.stack.push(uint256_t{5});
        try {
            m.runOp(OpCode::TGET);
        } catch (const bad_tuple_index& e) {
            m.state = Status::Error;
        }
        // should throw bad_tuple_index and leave two items on stack
        REQUIRE(m.state == Status::Error);
        REQUIRE(m.stack.stacksize() == 2);
    }
}

TEST_CASE("TSET opcode is correct") {
    SECTION("2 tup") {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}, m.pool.get()});
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::TSET);
        value res = m.stack.pop();
        REQUIRE(res == value{Tuple{uint256_t{1}, uint256_t{3}, m.pool.get()}});
        REQUIRE(m.stack.stacksize() == 0);
    }

    SECTION("8 tup") {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.stack.push(Tuple{uint256_t{9}, uint256_t{9}, uint256_t{9},
                           uint256_t{9}, uint256_t{9}, uint256_t{9},
                           uint256_t{9}, uint256_t{9}, m.pool.get()});
        m.stack.push(uint256_t{7});
        m.runOp(OpCode::TSET);
        value res = m.stack.pop();
        REQUIRE(res == value{Tuple{uint256_t{9}, uint256_t{9}, uint256_t{9},
                                   uint256_t{9}, uint256_t{9}, uint256_t{9},
                                   uint256_t{9}, uint256_t{3}, m.pool.get()}});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("TLEN opcode is correct") {
    SECTION("tlen") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7},
                           uint256_t{6}, m.pool.get()});
        m.runOp(OpCode::TLEN);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(4)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("XGET opcode is correct") {
    SECTION("correct") {
        MachineState m;
        m.auxstack.push(Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7},
                              uint256_t{6}, m.pool.get()});
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::XGET);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(8)});
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.auxstack.stacksize() == 1);
    }

    SECTION("index out range") {
        MachineState m;
        m.auxstack.push(Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7},
                              uint256_t{6}, m.pool.get()});
        m.stack.push(uint256_t{5});

        CHECK_THROWS(m.runOp(OpCode::XGET));
        // should throw bad_tuple_index and leave two items on stack
        REQUIRE(m.stack.stacksize() == 1);
    }
}

TEST_CASE("XSET opcode is correct") {
    SECTION("2 tup") {
        MachineState m;
        m.auxstack.push(Tuple{uint256_t{1}, uint256_t{2}, m.pool.get()});
        m.stack.push(uint256_t{3});
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::XSET);
        value res = m.auxstack.pop();
        REQUIRE(res == value{Tuple{uint256_t{1}, uint256_t{3}, m.pool.get()}});
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.auxstack.stacksize() == 0);
    }
}

TEST_CASE("BREAKPOINT opcode is correct") {
    SECTION("break") {
        MachineState m;
        auto blockReason = m.runOp(OpCode::BREAKPOINT);
        REQUIRE(m.state == Status::Extensive);
        REQUIRE(nonstd::get_if<BreakpointBlocked>(&blockReason));
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("LOG opcode is correct") {
    SECTION("log") {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::LOG);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.context.logs[0] == value{uint256_t(3)});
    }
}

TEST_CASE("SEND opcode is correct") {
    SECTION("send") {
        // TODO: fill in send test
        MachineState m;
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2345}, uint256_t{1},
                           uint256_t{4}, m.pool.get()});

        m.runOp(OpCode::SEND);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.state == Status::Extensive);
    }
}

TEST_CASE("PUSHGAS opcode is correct") {
    MachineState m;
    m.arb_gas_remaining = 250;
    m.runOp(OpCode::PUSH_GAS);
    value res = m.stack.pop();
    REQUIRE(res == value{uint256_t(250)});
    REQUIRE(m.stack.stacksize() == 0);
    REQUIRE(m.auxstack.stacksize() == 0);
}

TEST_CASE("SET_GAS opcode is correct") {
    MachineState m;
    m.stack.push(uint256_t{100});
    m.runOp(OpCode::SET_GAS);
    REQUIRE(m.arb_gas_remaining == 100);
    REQUIRE(m.stack.stacksize() == 0);
    REQUIRE(m.auxstack.stacksize() == 0);
}

uint256_t& assumeInt(value& val) {
    auto aNum = nonstd::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

TEST_CASE("ecrecover opcode is correct") {
    std::array<unsigned char, 32> msg;
    std::generate(msg.begin(), msg.end(), std::rand);
    std::array<unsigned char, 32> seckey;
    std::generate(seckey.begin(), seckey.end(), std::rand);

    auto ctx = secp256k1_context_create(SECP256K1_CONTEXT_SIGN |
                                        SECP256K1_CONTEXT_VERIFY);
    secp256k1_ecdsa_recoverable_signature sig;
    secp256k1_pubkey pubkey;
    REQUIRE(secp256k1_ec_pubkey_create(ctx, &pubkey, seckey.data()) == 1);
    std::array<unsigned char, 65> pubkey_raw;
    size_t output_length = 65;
    REQUIRE(secp256k1_ec_pubkey_serialize(ctx, pubkey_raw.data(),
                                          &output_length, &pubkey,
                                          SECP256K1_EC_UNCOMPRESSED));
    REQUIRE(output_length == 65);

    REQUIRE(secp256k1_ecdsa_sign_recoverable(
                ctx, &sig, msg.data(), seckey.data(), nullptr, nullptr) == 1);

    std::array<unsigned char, 64> sig_raw;
    int recovery_id;
    REQUIRE(secp256k1_ecdsa_recoverable_signature_serialize_compact(
                ctx, sig_raw.data(), &recovery_id, &sig) == 1);

    MachineState s;
    s.stack.push(intx::be::unsafe::load<uint256_t>(msg.begin()));
    s.stack.push(uint256_t{recovery_id});
    s.stack.push(intx::be::unsafe::load<uint256_t>(sig_raw.begin() + 32));
    s.stack.push(intx::be::unsafe::load<uint256_t>(sig_raw.begin()));
    s.runOp(OpCode::ECRECOVER);
    REQUIRE(s.stack[0] != value(0));
    auto hash_val = ethash::keccak256(pubkey_raw.begin() + 1, 64);
    std::fill(&hash_val.bytes[0], &hash_val.bytes[12], 0);
    auto correct_address = intx::be::load<uint256_t>(hash_val);
    auto calculated_address = assumeInt(s.stack[0]);
    REQUIRE(correct_address == calculated_address);

    BENCHMARK_ADVANCED("ecrecover")(Catch::Benchmark::Chronometer meter) {
        MachineState sample_machine;
        sample_machine.stack.push(
            intx::be::unsafe::load<uint256_t>(msg.begin()));
        sample_machine.stack.push(uint256_t{recovery_id});
        sample_machine.stack.push(
            intx::be::unsafe::load<uint256_t>(sig_raw.begin() + 32));
        sample_machine.stack.push(
            intx::be::unsafe::load<uint256_t>(sig_raw.begin()));

        std::vector<MachineState> machines(meter.runs());
        std::fill(machines.begin(), machines.end(), sample_machine);
        meter.measure([&machines](int i) {
            return machines[i].runOp(OpCode::ECRECOVER);
        });
    };
}

TEST_CASE("INBOX opcode is correct") {
    SECTION("inbox") {
        // TODO: fill in inbox test
    }
}

TEST_CASE("ERROR opcode is correct") {
    SECTION("error") {
        // TODO: fill in error test
    }
}

TEST_CASE("HALT opcode is correct") {
    SECTION("halt") {
        // TODO: fill in halt test
    }
}
