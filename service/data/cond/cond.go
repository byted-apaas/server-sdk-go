// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package cond

import (
	"github.com/byted-apaas/server-sdk-go/service/data/op"
)

// And 逻辑与
// @params exps：表达式列表，表达式的类型为 *LogicalExpression 或 *ArithmeticExpression，不合法的类型直接忽略
// @example exps：
//     cond.And(...)
//     cond.Or(...)
//     cond.Eq(...)
//     cond.Gt(...)
// @return 返回逻辑表达式
func And(exps ...interface{}) *LogicalExpression {
	logicalExps, arighmeticExps := logical(exps...)
	if len(logicalExps) == 0 && len(arighmeticExps) == 0 {
		return nil
	}

	return NewLogicalExpression(op.And, arighmeticExps, logicalExps)
}

// Or 逻辑或
// @params exps：表达式列表，表达式的类型为 *LogicalExpression 或 *ArithmeticExpression，不合法的类型直接忽略
// @example exps：
//     cond.And(...)
//     cond.Or(...)
//     cond.Eq(...)
//     cond.Gt(...)
// @return 返回逻辑表达式
func Or(exps ...interface{}) *LogicalExpression {
	logicalExps, arighmeticExps := logical(exps...)
	if len(logicalExps) == 0 && len(arighmeticExps) == 0 {
		return nil
	}

	return NewLogicalExpression(op.Or, arighmeticExps, logicalExps)
}

func logical(exps ...interface{}) ([]*LogicalExpression, []*ArithmeticExpression) {
	if len(exps) == 0 {
		return nil, nil
	}

	var logicalExps []*LogicalExpression
	var arighmeticExps []*ArithmeticExpression
	for i := range exps {
		if exps[i] == nil {
			continue
		}

		switch exps[i].(type) {
		case LogicalExpression:
			v, _ := exps[i].(LogicalExpression)
			logicalExps = append(logicalExps, &v)
		case *LogicalExpression:
			v, _ := exps[i].(*LogicalExpression)
			logicalExps = append(logicalExps, v)
		case ArithmeticExpression:
			v, _ := exps[i].(ArithmeticExpression)
			arighmeticExps = append(arighmeticExps, &v)
		case *ArithmeticExpression:
			v, _ := exps[i].(*ArithmeticExpression)
			arighmeticExps = append(arighmeticExps, v)
		}
	}
	return logicalExps, arighmeticExps
}

func Eq(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.Eq, rightValue)
}

func Neq(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.Neq, rightValue)
}

func Gt(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.Gt, rightValue)
}

func Gte(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.Gte, rightValue)
}

func Lt(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.Lt, rightValue)
}

func Lte(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.Lte, rightValue)
}

// IsOnOrBefore 日期比较，早于
func IsOnOrBefore(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.IsOnOrBefore, rightValue)
}

// IsOnOrAfter 日期比较，晚于
func IsOnOrAfter(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.IsOnOrAfter, rightValue)
}

func Contain(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.Contain, rightValue)
}

func NotContain(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.NotContain, rightValue)
}

func In(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.In, rightValue)
}

func NotIn(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.NotIn, rightValue)
}

func Empty(leftValue string) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.Empty, nil)
}

func NotEmpty(leftValue string) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.NotEmpty, nil)
}

func HasAnyOf(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.HasAnyOf, rightValue)
}

// HasAnyOfHierarchy 层级查询，需要设置 ExtendLogicTags
func HasAnyOfHierarchy(leftValue string, rightValue interface{}) *ArithmeticExpression {
	exp := NewArithmeticExpression(leftValue, op.HasAnyOf, rightValue)
	fieldPaths := exp.Left.Settings.FieldPath
	fieldPaths[len(fieldPaths)-1].ExtendLogicTags = []string{"all"}
	return exp
}

func HasNoneOf(leftValue string, rightValue interface{}) *ArithmeticExpression {
	return NewArithmeticExpression(leftValue, op.HasNoneOf, rightValue)
}

// HasNoneOfHierarchy 层级查询，需要设置 ExtendLogicTags
func HasNoneOfHierarchy(leftValue string, rightValue interface{}) *ArithmeticExpression {
	exp := NewArithmeticExpression(leftValue, op.HasNoneOf, rightValue)
	fieldPaths := exp.Left.Settings.FieldPath
	fieldPaths[len(fieldPaths)-1].ExtendLogicTags = []string{"all"}
	return exp
}
