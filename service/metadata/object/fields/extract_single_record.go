// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

import "github.com/byted-apaas/server-sdk-go/common/structs"

type ExtractSingleRecord struct {
	FieldBase
	CompositeTypeAPIName string                 `json:"compositeTypeAPIName"`
	SubFields            map[string]interface{} `json:"subFields"`
	Filter               *structs.Criterion     `json:"filter"`
	SortConditions       *structs.Sorts         `json:"sortConditions"`
	RecordPosition       int64                  `json:"recordPosition"`
}

type ExtractSingleRecordV3 struct {
	FieldBaseV3
	CompositeTypeAPIName string                 `json:"compositeTypeAPIName"`
	SubFields            map[string]interface{} `json:"subFields"`
	Filter               *structs.Criterion     `json:"filter"`
	SortConditions       *structs.Sorts         `json:"sortConditions"`
	RecordPosition       int64                  `json:"recordPosition"`
}
