// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package cond

import (
	"strings"

	"github.com/byted-apaas/server-sdk-go/service/data/op"
)

// LogicalExpression 逻辑表达式，树状结构
type LogicalExpression struct {
	Type                  string                  `json:"type"`
	ArithmeticExpressions []*ArithmeticExpression `json:"arithmeticExpressions"`
	LogicalExpressions    []*LogicalExpression    `json:"logicalExpressions"`
}

// ArithmeticExpression 算术表达式
type ArithmeticExpression struct {
	Index    int64         `json:"index"`
	Left     Expression    `json:"left"`
	Operator string        `json:"operator"`
	Right    *Expression   `json:"right,omitempty"`
	Expr     *ExpressionV2 `json:"-"` // v2 版本专用
}

func NewLogicalExpression(logicType string, expressions []*ArithmeticExpression, logics []*LogicalExpression) *LogicalExpression {
	return &LogicalExpression{
		Type:                  logicType,
		ArithmeticExpressions: expressions,
		LogicalExpressions:    logics,
	}
}

func NewArithmeticExpression(leftValue string, operator string, rightValue interface{}) *ArithmeticExpression {
	var fieldPaths []*FieldPath
	leftFields := strings.Split(leftValue, ".")
	for i := range leftFields {
		// 这里拿不到 ObjectAPIName，后面再补全
		fieldPaths = append(fieldPaths, &FieldPath{FieldAPIName: leftFields[i]})
	}
	left := Expression{
		Type: "metadataVariable",
		Settings: SettingType{
			FieldPath: fieldPaths,
		}}
	right := &Expression{
		Type: "constant",
		Settings: SettingType{
			Data: rightValue,
		},
	}
	if operator == op.Empty || operator == op.NotEmpty {
		right = nil
	}
	return &ArithmeticExpression{Operator: operator, Left: left, Right: right, Expr: NewExpressionV2(leftValue, operator, rightValue)}
}

func (l *LogicalExpression) AddLogicalExpression(exp *LogicalExpression) {
	if exp != nil {
		l.LogicalExpressions = append(l.LogicalExpressions, exp)
	}
}

func (l *LogicalExpression) AddArithmeticExpression(exp *ArithmeticExpression) {
	if exp != nil {
		l.ArithmeticExpressions = append(l.ArithmeticExpressions, exp)
	}
}

type Criterion struct {
	Conditions []*ArithmeticExpression `json:"conditions"`
	Logic      string                  `json:"logic"`
}

type FieldPath struct {
	ObjectAPIName   string   `json:"objectApiName"`
	FieldAPIName    string   `json:"fieldApiName"`
	ExtendLogicTags []string `json:"extendLogicTags"`
}

type SettingType struct {
	Data      interface{}  `json:"data,omitempty"`
	FieldPath []*FieldPath `json:"fieldPath,omitempty"`
}

type Expression struct {
	Type     string      `json:"type"`
	Settings SettingType `json:"settings"`
}
