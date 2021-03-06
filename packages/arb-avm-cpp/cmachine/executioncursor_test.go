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

package cmachine

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
	"os"
	"testing"
)

func TestExecutionCursor(t *testing.T) {
	dePath := "dbPath"

	if err := os.RemoveAll(dePath); err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}

	defer func() {
		if err := os.RemoveAll(dePath); err != nil {
			logger.Error().Stack().Err(err).Send()
			t.Fatal(err)
		}
	}()

	arbStorage, err := NewArbStorage(dePath)
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}

	if err := arbStorage.Initialize(codeFile); err != nil {
		t.Fatal(err)
	}
	defer arbStorage.CloseArbStorage()

	lookup := arbStorage.GetArbCore()
	cursor, err := lookup.GetExecutionCursor(big.NewInt(0))
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}
	if !cursor.InboxAcc().Equals(common.Hash{}) {
		logger.Error().Msg("inbox acc isn't zero at beginning")
	}
	if !cursor.SendAcc().Equals(common.Hash{}) {
		logger.Error().Msg("send acc isn't zero at beginning")
	}
	if !cursor.LogAcc().Equals(common.Hash{}) {
		logger.Error().Msg("log acc isn't zero at beginning")
	}

	err = lookup.AdvanceExecutionCursor(cursor, big.NewInt(10000), true)
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}
	if cursor.InboxAcc().Equals(common.Hash{}) {
		logger.Error().Msg("inbox acc is zero after execution")
	}
	if cursor.SendAcc().Equals(common.Hash{}) {
		logger.Error().Msg("send acc is zero after execution")
	}
	if cursor.LogAcc().Equals(common.Hash{}) {
		logger.Error().Msg("log acc is zero after execution")
	}
}
