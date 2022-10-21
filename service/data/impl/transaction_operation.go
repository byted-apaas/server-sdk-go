// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package data

import (
	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/structs"
)

func newTransactionOperation(operationType constants.OperationType, objectAPIName string, originInput interface{}) (*structs.TransactionOperation, error) {
	input, err := cUtils.JsonMarshalBytes(originInput)
	if err != nil {
		return nil, cExceptions.InvalidParamError("Marshal input failed, err: %v", err)
	}

	return &structs.TransactionOperation{
		OperationType: operationType,
		ObjectAPIName: objectAPIName,
		Input:         string(input),
	}, nil
}
