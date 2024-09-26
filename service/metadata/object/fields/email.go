// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type Email struct {
	FieldBase
	Required bool `json:"required"`
	Unique   bool `json:"unique"`
}

type EmailV3 struct {
	FieldBaseV3
	Required bool `json:"required"`
	Unique   bool `json:"unique"`
}
