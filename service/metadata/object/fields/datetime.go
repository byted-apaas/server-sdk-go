// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type DateTime struct {
	FieldBase
	Required bool `json:"required"`
}
