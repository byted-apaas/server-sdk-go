// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package cond

import (
	"encoding/json"
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

func (c *Criterion) ToCriterionV3() (*CriterionV3, error) {
	if c == nil {
		return nil, nil
	}
	expressionV3 := make([]*ArithmeticExpressionV3, 0)
	for _, exp := range c.Conditions {
		tmp := &ArithmeticExpressionV3{
			Index:    exp.Index,
			Operator: exp.Operator,
		}

		if exp.Left.Settings.FieldPath != nil && len(exp.Left.Settings.FieldPath) > 0 {
			settings := exp.Left.Settings.ToLeftSettingTypeV3()
			left, err := json.Marshal(settings)
			if err != nil {
				return nil, err
			}
			tmp.Left = ExpressionV3{
				Type:     exp.Left.Type,
				Settings: string(left),
			}
		}

		if exp.Right != nil {
			settings := exp.Right.Settings.ToRightSettingTypeV3()
			right, err := json.Marshal(settings)
			if err != nil {
				return nil, err
			}
			tmp.Right = &ExpressionV3{
				Type:     exp.Right.Type,
				Settings: string(right),
			}
		}

		expressionV3 = append(expressionV3, tmp)
	}
	criterionV3 := &CriterionV3{
		Conditions: expressionV3,
		Expression: c.Logic,
	}
	return criterionV3, nil
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

func (s *SettingType) ToLeftSettingTypeV3() *SettingLeftTypeV3 {
	settingV3 := &SettingLeftTypeV3{}
	for _, fieldPath := range s.FieldPath {
		fieldPathV3 := &FieldPathV3{
			ObjectAPIName: fieldPath.ObjectAPIName,
			FieldAPIName:  fieldPath.FieldAPIName,
		}
		settingV3.FieldPath = append(settingV3.FieldPath, fieldPathV3)
	}

	return settingV3
}

func (s *SettingType) ToRightSettingTypeV3() *SettingRightTypeV3 {
	settingV3 := &SettingRightTypeV3{
		Data: s.Data,
	}

	return settingV3
}

type Expression struct {
	Type     string      `json:"type"`
	Settings SettingType `json:"settings"`
}

type FieldPathV3 struct {
	ObjectAPIName string `json:"objectApiName"`
	FieldAPIName  string `json:"fieldApiName"`
}

type SettingLeftTypeV3 struct {
	FieldPath []*FieldPathV3 `json:"fieldPath,omitempty"`
}

type SettingRightTypeV3 struct {
	Data interface{} `json:"data,omitempty"`
}

type ExpressionV3 struct {
	Type     string `json:"type"`
	Settings string `json:"settings"`
}

type CriterionV3 struct {
	Conditions []*ArithmeticExpressionV3 `json:"conditions"`
	Expression string                    `json:"expression"`
}

type ArithmeticExpressionV3 struct {
	Index    int64         `json:"index"`
	Left     ExpressionV3  `json:"left"`
	Operator string        `json:"operator"`
	Right    *ExpressionV3 `json:"right,omitempty"`
}
