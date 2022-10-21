// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

import "github.com/byted-apaas/server-sdk-go/common/structs"

type Option struct {
	FieldBase
	Required            bool          `json:"required"`
	Multiple            bool          `json:"multiple"`
	OptionSource        string        `json:"optionSource"` // 暴露给开发者：custom, global; 底层存储：local,  global
	GlobalOptionAPIName string        `json:"globalOptionAPIName"`
	OptionList          []*OptionItem `json:"optionList"`
}

type OptionItem struct {
	Label       structs.I18ns `json:"label"`
	APIName     string        `json:"apiName"`
	Description structs.I18ns `json:"description"`
	Color       string        `json:"color"`
	Active      bool          `json:"active"`
}
