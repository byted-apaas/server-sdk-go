// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package data

import (
	"context"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
	"github.com/byted-apaas/server-sdk-go/service/data"
)

type Oql struct {
	appCtx *structs.AppCtx
	oql    string
	err    error
}

func NewOql(s *structs.AppCtx, oql string) data.IOql {
	o := &Oql{appCtx: s, oql: oql}
	if oql == "" {
		o.err = cExceptions.InvalidParamError("[Object] objectAPIName is empty")
	}
	return o
}

func (o *Oql) Execute(ctx context.Context, resultSet interface{}) error {
	if err := o.check(); err != nil {
		return err
	}

	return request.GetInstance(ctx).Oql(ctx, o.appCtx, o.oql, nil, nil, resultSet)
}

func (o *Oql) check() error {
	return o.err
}
