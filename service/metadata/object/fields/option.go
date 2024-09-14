// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
)

type Option struct {
	FieldBase
	Required            bool          `json:"required"`
	Multiple            bool          `json:"multiple"`
	OptionSource        string        `json:"optionSource"` // 暴露给开发者：custom, global; 底层存储：local,  global
	GlobalOptionAPIName string        `json:"globalOptionAPIName"`
	OptionList          []*OptionItem `json:"optionList"`
}

func (o *Option) ToOptionV3() *OptionV3 {
	optionV3 := &OptionV3{
		FieldBaseV3:         o.FieldBase.ToFieldBaseV3(),
		Required:            o.Required,
		Multiple:            o.Multiple,
		OptionSource:        o.OptionSource,
		GlobalOptionAPIName: o.GlobalOptionAPIName,
	}

	for _, optionItem := range o.OptionList {
		optionV3.OptionList = append(optionV3.OptionList, optionItem.ToOptionItemV3())
	}
	return optionV3
}

type OptionItem struct {
	Label       structs.I18ns `json:"label"`
	APIName     string        `json:"apiName"`
	Description structs.I18ns `json:"description"`
	Color       string        `json:"color"`
	Active      bool          `json:"active"`
}

func (o *OptionItem) ToOptionItemV3() *OptionItemV3 {
	return &OptionItemV3{
		Label:       o.Label.TransToMultilingualV3(),
		APIName:     o.APIName,
		Description: o.Description.TransToMultilingualV3(),
		Color:       o.Color,
		Active:      o.Active,
	}
}

type OptionV3 struct {
	FieldBaseV3
	Required            bool            `json:"required"`
	Multiple            bool            `json:"multiple"`
	OptionSource        string          `json:"optionSource"` // 暴露给开发者：custom, global; 底层存储：local,  global
	GlobalOptionAPIName string          `json:"globalOptionAPIName"`
	OptionList          []*OptionItemV3 `json:"optionList"`
}

type OptionItemV3 struct {
	Label       *faassdk.MultilingualV3 `json:"label"`
	APIName     string                  `json:"apiName"`
	Description *faassdk.MultilingualV3 `json:"description"`
	Color       string                  `json:"color"`
	Active      bool                    `json:"active"`
}
