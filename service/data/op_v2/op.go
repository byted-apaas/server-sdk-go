// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package op_v2

import (
	"github.com/byted-apaas/server-sdk-go/service/data/op"
)

const (
	And        = "and"
	Or         = "or"
	Eq         = "eq"
	Neq        = "neq"
	Gt         = "gt"
	Gte        = "gte"
	Lt         = "lt"
	Lte        = "lte"
	Contain    = "contain"
	NotContain = "notContain"
	In         = "in"
	NotIn      = "notIn"
	Empty      = "empty"
	NotEmpty   = "notEmpty"
	HasAnyOf   = "hasAnyOf"
	HasNoneOf  = "hasNoneOf"
)

var (
	opToOpV2 = map[string]string{
		op.And:          And,
		op.Or:           Or,
		op.Eq:           Eq,
		op.Neq:          Neq,
		op.Gt:           Gt,
		op.Gte:          Gte,
		op.Lt:           Lt,
		op.Lte:          Lte,
		op.Contain:      Contain,
		op.NotContain:   NotContain,
		op.In:           In,
		op.NotIn:        NotIn,
		op.Empty:        Empty,
		op.NotEmpty:     NotEmpty,
		op.HasAnyOf:     HasAnyOf,
		op.HasNoneOf:    HasNoneOf,
		op.IsOnOrBefore: Lt,
		op.IsOnOrAfter:  Gt,
	}
)

func ConvertOpToOpV2(op string) string {
	opV2, _ := opToOpV2[op]
	return opV2
}
