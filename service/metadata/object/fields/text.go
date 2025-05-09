// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type Text struct {
	FieldBase
	Required        bool   `json:"required"`
	Unique          bool   `json:"unique"`
	CaseSensitive   bool   `json:"caseSensitive"`
	Multiline       bool   `json:"multiline"`
	MaxLength       int64  `json:"maxLength"`
	ValidationRegex string `json:"validationRegex"`
	ErrorMsg        string `json:"errorMsg"`
}

type TextV3 struct {
	FieldBaseV3
	Required        bool   `json:"required"`
	Unique          bool   `json:"unique"`
	CaseSensitive   bool   `json:"caseSensitive"`
	Multiline       bool   `json:"multiline"`
	MaxLength       int64  `json:"maxLength"`
	ValidationRegex string `json:"validationRegex"`
	ErrorMsg        string `json:"errorMsg"`
}
