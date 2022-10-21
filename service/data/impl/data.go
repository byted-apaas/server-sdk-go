// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package data

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data"
)

type Data struct {
	appCtx *structs.AppCtx
}

func NewData(s *structs.AppCtx) data.IData {
	return &Data{appCtx: s}
}

func (d *Data) Object(objectAPIName string) data.IObject {
	return NewObject(d.appCtx, objectAPIName)
}

func (d *Data) NewTransaction() data.ITransaction {
	return NewTransaction(d.appCtx)
}

func (d *Data) Oql(oql string) data.IOql {
	return NewOql(d.appCtx, oql)
}
