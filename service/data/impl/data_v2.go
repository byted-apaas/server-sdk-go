// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package data

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data"
)

type DataV2 struct {
	appCtx *structs.AppCtx
}

func NewDataV2(s *structs.AppCtx) data.IDataV2 {
	if s == nil {
		s = &structs.AppCtx{}
	}
	s.DataVersion = structs.DataVersionV2
	return &DataV2{appCtx: s}
}

func (d *DataV2) Object(objectAPIName string) data.IObjectV2 {
	return NewObject(d.appCtx, objectAPIName)
}

func (d *DataV2) NewTransaction() data.ITransaction {
	return NewTransaction(d.appCtx)
}

func (d *DataV2) Oql(oql string, args ...interface{}) data.IOql {
	return NewOql(d.appCtx, oql, args...)
}
