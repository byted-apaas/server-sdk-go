// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
)

type FieldBase struct {
	Type    string        `json:"type"`
	APIName string        `json:"apiName"`
	Label   structs.I18ns `json:"label"`
}

type FieldBaseV3 struct {
	Type    string                  `json:"type"`
	APIName string                  `json:"apiName"`
	Label   *faassdk.MultilingualV3 `json:"label"`
}

func (f *FieldBase) ToFieldBaseV3() FieldBaseV3 {
	return FieldBaseV3{
		Type:    f.Type,
		APIName: f.APIName,
		Label:   f.Label.TransToMultilingualV3(),
	}
}
