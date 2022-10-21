// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

import "github.com/byted-apaas/server-sdk-go/common/structs"

type FieldBase struct {
	Type    string        `json:"type"`
	APIName string        `json:"apiName"`
	Label   structs.I18ns `json:"label"`
}
