// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package object

import (
	"context"

	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
	"github.com/byted-apaas/server-sdk-go/service/metadata/parser"
)

type Object struct {
	appCtx        *structs.AppCtx
	ObjectAPIName string
}

func NewObject(s *structs.AppCtx, objectAPIName string) *Object {
	return &Object{
		appCtx:        s,
		ObjectAPIName: objectAPIName,
	}
}

func (o *Object) GetField(ctx context.Context, fieldAPIName string, field interface{}) error {
	data, err := request.GetInstance(ctx).GetField(ctx, o.appCtx, o.ObjectAPIName, fieldAPIName)
	if err != nil {
		return err
	}

	if o.appCtx.IsDataV3() {
		parsedField, err := parser.ParseFieldV3(data)
		if err != nil {
			return err
		}
		return cUtils.Decode(parsedField, &field)
	}

	parsedField, err := parser.ParseField(data)
	if err != nil {
		return err
	}
	return cUtils.Decode(parsedField, &field)
}

func (o *Object) GetFields(ctx context.Context, fields interface{}) error {
	data, err := request.GetInstance(ctx).GetFields(ctx, o.appCtx, o.ObjectAPIName)

	if err != nil {
		return err
	}
	if o.appCtx.IsDataV3() {
		parsedField, err := parser.ParseFieldsV3(data.Fields)
		if err != nil {
			return err
		}
		return cUtils.Decode(parsedField, &fields)
	}

	parsedField, err := parser.ParseFields(data.Fields)
	if err != nil {
		return err
	}
	return cUtils.Decode(parsedField, &fields)
}
