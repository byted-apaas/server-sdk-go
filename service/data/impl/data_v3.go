// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package data

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data"
)

type DataV3 struct {
	appCtx *structs.AppCtx
}

func NewDataV3(s *structs.AppCtx) data.IDataV3 {
	if s == nil {
		s = &structs.AppCtx{}
	}
	s.DataVersion = structs.DataVersionV3
	return &DataV3{appCtx: s}
}

func (d *DataV3) Object(objectAPIName string) data.IObjectV3 {
	return NewObjectV3(d.appCtx, objectAPIName)
}

func (d *DataV3) NewTransaction() data.ITransaction {
	return NewTransaction(d.appCtx)
}

func (d *DataV3) Oql(oql string, args ...interface{}) data.IOql {
	return NewOql(d.appCtx, oql, args...)
}
