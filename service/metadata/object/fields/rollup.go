// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
)

type Rollup struct {
	FieldBase

	RollupType               string             `json:"rollupType"`
	RollupObjectApiName      string             `json:"rollupObjectApiName"`
	RollupFieldApiName       string             `json:"rollupFieldApiName"`
	RollupLookupFieldApiName string             `json:"rollupLookupFieldApiName"`
	Filter                   *structs.Criterion `json:"filter"`
}

type RollupV3 struct {
	FieldBaseV3

	RollupType               string             `json:"rollupType"`
	RollupObjectApiName      string             `json:"rollupObjectApiName"`
	RollupFieldApiName       string             `json:"rollupFieldApiName"`
	RollupLookupFieldApiName string             `json:"rollupLookupFieldApiName"`
	Filter                   *structs.Criterion `json:"filter"`
}
