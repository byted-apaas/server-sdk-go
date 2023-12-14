// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

import "github.com/byted-apaas/server-sdk-go/common/structs"

type Region struct {
	FieldBase
	Required    bool                    `json:"required"`
	Multiple    bool                    `json:"multiple"`
	OptionLevel bool                    `json:"optionLevel"`
	StrictLevel int64                   `json:"strictLevel"`
	Filter      []*structs.RegionFilter `json:"filter"`
}
