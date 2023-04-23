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
	appCtx    *structs.AppCtx
	oql       string
	namedArgs map[string]interface{}
	err       error
}

func NewOql(s *structs.AppCtx, oql string, args ...interface{}) data.IOql {
	o := &Oql{appCtx: s, oql: oql}
	if oql == "" {
		o.err = cExceptions.InvalidParamError("[Object] objectAPIName is empty")
	}
	if len(args) > 0 {
		if argMap, ok := args[0].(map[string]interface{}); ok {
			o.namedArgs = argMap
		}
	}
	return o
}

func (o *Oql) Execute(ctx context.Context, resultSet interface{}) error {
	if err := o.check(); err != nil {
		return err
	}

	return request.GetInstance(ctx).Oql(ctx, o.appCtx, o.oql, nil, o.namedArgs, resultSet)
}

func (o *Oql) check() error {
	return o.err
}
