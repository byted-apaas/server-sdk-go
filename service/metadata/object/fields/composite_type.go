// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type CompositeType struct {
	FieldBase
	CompositeTypeAPIName string                 `json:"compositeTypeAPIName"`
	Required             bool                   `json:"required"`
	Multiple             bool                   `json:"multiple"`
	SubFields            map[string]interface{} `json:"subFields"`
}

type CompositeTypeV3 struct {
	FieldBaseV3
	CompositeTypeAPIName string                 `json:"compositeTypeAPIName"`
	Required             bool                   `json:"required"`
	Multiple             bool                   `json:"multiple"`
	SubFields            map[string]interface{} `json:"subFields"`
}
