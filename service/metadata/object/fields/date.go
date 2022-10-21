// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type Date struct {
	FieldBase
	Required bool `json:"required"`
}
