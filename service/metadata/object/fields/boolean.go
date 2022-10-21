// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

import "github.com/byted-apaas/server-sdk-go/common/structs"

type Boolean struct {
	FieldBase
	DescriptionWhenTrue  structs.I18ns `json:"descriptionWhenTrue"`
	DescriptionWhenFalse structs.I18ns `json:"descriptionWhenFalse"`
	DefaultValue         bool          `json:"defaultValue"`
}
