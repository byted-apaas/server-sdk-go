// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
)

type Boolean struct {
	FieldBase
	DescriptionWhenTrue  structs.I18ns `json:"descriptionWhenTrue"`
	DescriptionWhenFalse structs.I18ns `json:"descriptionWhenFalse"`
	DefaultValue         bool          `json:"defaultValue"`
}

type BooleanV3 struct {
	FieldBaseV3
	DescriptionWhenTrue  *faassdk.MultilingualV3 `json:"descriptionWhenTrue"`
	DescriptionWhenFalse *faassdk.MultilingualV3 `json:"descriptionWhenFalse"`
	DefaultValue         bool                    `json:"defaultValue"`
}
