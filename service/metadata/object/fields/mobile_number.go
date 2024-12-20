// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type MobileNumber struct {
	FieldBase
	Required bool `json:"required"`
	Unique   bool `json:"unique"`
}

type MobileNumberV3 struct {
	FieldBaseV3
	Required bool `json:"required"`
	Unique   bool `json:"unique"`
}
