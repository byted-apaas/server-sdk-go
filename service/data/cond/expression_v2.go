// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package cond

import (
	"github.com/byted-apaas/server-sdk-go/service/data/op_v2"
)

type ExpressionV2 struct {
	LeftValue  string      `json:"leftValue"`
	Operator   string      `json:"operator"`
	RightValue interface{} `json:"rightValue"`
}

func NewExpressionV2(leftValue string, operator string, rightValue interface{}) *ExpressionV2 {
	return &ExpressionV2{Operator: op_v2.ConvertOpToOpV2(operator), LeftValue: leftValue, RightValue: rightValue}
}
