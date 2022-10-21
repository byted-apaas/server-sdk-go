// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type RichText struct {
	FieldBase
	Required  bool `json:"required"`  // 默认值：false
	MaxLength int  `json:"maxLength"` // 默认值：1000
}
