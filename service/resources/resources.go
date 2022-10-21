// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package resources

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/resources/file"
)

type Resources struct {
	File file.IFile
}

func NewResources(s *structs.AppCtx) *Resources {
	return &Resources{
		File: file.NewFile(s),
	}
}
