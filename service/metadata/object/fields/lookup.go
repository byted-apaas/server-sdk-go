// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type SortCondition struct {
	FieldAPIName string `json:"fieldAPIName"`
	Sort         string `json:"sort"` //"asc"|"desc"
}

type Lookup struct {
	FieldBase
	Required       bool            `json:"required"`
	Multiple       bool            `json:"multiple"`
	ObjectAPIName  string          `json:"objectAPIName"`
	Hierarchy      bool            `json:"hierarchy"`
	DisplayStyle   string          `json:"displayStyle"`
	SortConditions []SortCondition `json:"sortConditions"`
	Filter         []interface{}   `json:"filter"` // TODO 未返回
}

type LookupV3 struct {
	FieldBaseV3
	Required       bool            `json:"required"`
	Multiple       bool            `json:"multiple"`
	ObjectAPIName  string          `json:"objectAPIName"`
	Hierarchy      bool            `json:"hierarchy"`
	DisplayStyle   string          `json:"displayStyle"`
	SortConditions []SortCondition `json:"sortConditions"`
	Filter         []interface{}   `json:"filter"` // TODO 未返回
}
