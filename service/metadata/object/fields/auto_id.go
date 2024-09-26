// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type AutoID struct {
	FieldBase
	GenerateMethod string `json:"generateMethod"` // "random"|"incremental"
	DigitsNumber   int64  `json:"digitsNumber"`
	Prefix         string `json:"prefix"`
	Suffix         string `json:"suffix"`
}

type AutoIDV3 struct {
	FieldBaseV3
	GenerateMethod string `json:"generateMethod"` // "random"|"incremental"
	DigitsNumber   int64  `json:"digitsNumber"`
	Prefix         string `json:"prefix"`
	Suffix         string `json:"suffix"`
}
