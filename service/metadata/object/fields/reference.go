// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type ReferenceField struct {
	FieldBase
	GuideFieldAPIName string `json:"guideFieldAPIName"`
	FieldAPIName      string `json:"fieldAPIName"`
}

type ReferenceFieldV3 struct {
	FieldBaseV3
	GuideFieldAPIName string `json:"guideFieldAPIName"`
	FieldAPIName      string `json:"fieldAPIName"`
}
