// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

import "github.com/byted-apaas/server-sdk-go/common/structs"

type Formula struct {
	FieldBase
	ReturnType string        `json:"returnType"`
	Formula    structs.I18ns `json:"formula"`
}
