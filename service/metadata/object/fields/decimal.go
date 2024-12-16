// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type Decimal struct {
	FieldBase
	Required            bool `json:"required"`
	Unique              bool `json:"unique"`
	DisplayAsPercentage bool `json:"displayAsPercentage"`
	DecimalPlacesNumber int  `json:"decimalPlacesNumber"`
}

type DecimalV3 struct {
	FieldBaseV3
	Required            bool `json:"required"`
	Unique              bool `json:"unique"`
	DisplayAsPercentage bool `json:"displayAsPercentage"`
	DecimalPlacesNumber int  `json:"decimalPlacesNumber"`
}
