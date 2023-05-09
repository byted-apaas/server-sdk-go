// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package data

import (
	"context"
	"reflect"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
	"github.com/byted-apaas/server-sdk-go/service/data"
)

type Object struct {
	appCtx        *structs.AppCtx
	objectAPIName string
	err           error
}

func NewObject(s *structs.AppCtx, objectAPIName string) *Object {
	o := &Object{
		appCtx:        s,
		objectAPIName: objectAPIName,
	}

	if objectAPIName == "" {
		o.err = cExceptions.InvalidParamError("[Object] objectAPIName is empty")
	}
	return o
}

func (o *Object) Create(ctx context.Context, record interface{}) (*structs.RecordID, error) {
	if err := o.check(); err != nil {
		return nil, err
	}

	if o.appCtx.IsOpenSDK() {
		return request.GetInstance(ctx).CreateRecordV2(ctx, o.appCtx, o.objectAPIName, record)
	}
	return request.GetInstance(ctx).CreateRecord(ctx, o.appCtx, o.objectAPIName, record)
}

func (o *Object) BatchCreate(ctx context.Context, records interface{}) ([]int64, error) {
	if err := o.check(); err != nil {
		return nil, err
	}

	if o.appCtx.IsOpenSDK() {
		return request.GetInstance(ctx).BatchCreateRecordV2(ctx, o.appCtx, o.objectAPIName, records)
	}
	return request.GetInstance(ctx).BatchCreateRecord(ctx, o.appCtx, o.objectAPIName, records)
}

func (o *Object) BatchCreateAsync(ctx context.Context, records interface{}) (int64, error) {
	if err := o.check(); err != nil {
		return 0, err
	}

	return request.GetInstance(ctx).BatchCreateRecordAsync(ctx, o.appCtx, o.objectAPIName, records)
}

func (o *Object) Update(ctx context.Context, _id int64, record interface{}) error {
	if err := o.check(); err != nil {
		return err
	}

	if o.appCtx.IsOpenSDK() {
		return request.GetInstance(ctx).UpdateRecordV2(ctx, o.appCtx, o.objectAPIName, _id, record)
	}
	return request.GetInstance(ctx).UpdateRecord(ctx, o.appCtx, o.objectAPIName, _id, record)
}

func (o *Object) BatchUpdate(ctx context.Context, records map[int64]interface{}) error {
	if err := o.check(); err != nil {
		return err
	}

	if o.appCtx.IsOpenSDK() {
		return request.GetInstance(ctx).BatchUpdateRecordV2(ctx, o.appCtx, o.objectAPIName, records)
	}
	return request.GetInstance(ctx).BatchUpdateRecord(ctx, o.appCtx, o.objectAPIName, records)
}

func (o *Object) BatchUpdateAsync(ctx context.Context, records map[int64]interface{}) (int64, error) {
	if err := o.check(); err != nil {
		return 0, err
	}

	return request.GetInstance(ctx).BatchUpdateRecordAsync(ctx, o.appCtx, o.objectAPIName, records)
}

func (o *Object) Delete(ctx context.Context, _id int64) error {
	if err := o.check(); err != nil {
		return err
	}

	if o.appCtx.IsOpenSDK() {
		return request.GetInstance(ctx).DeleteRecordV2(ctx, o.appCtx, o.objectAPIName, _id)
	}
	return request.GetInstance(ctx).DeleteRecord(ctx, o.appCtx, o.objectAPIName, _id)
}

func (o *Object) BatchDelete(ctx context.Context, _ids []int64) error {
	if err := o.check(); err != nil {
		return err
	}

	if o.appCtx.IsOpenSDK() {
		return request.GetInstance(ctx).BatchDeleteRecordV2(ctx, o.appCtx, o.objectAPIName, _ids)
	}
	return request.GetInstance(ctx).BatchDeleteRecord(ctx, o.appCtx, o.objectAPIName, _ids)
}

func (o *Object) BatchDeleteAsync(ctx context.Context, _ids []int64) (int64, error) {
	if err := o.check(); err != nil {
		return 0, err
	}

	return request.GetInstance(ctx).BatchDeleteRecordAsync(ctx, o.appCtx, o.objectAPIName, _ids)
}

func (o *Object) Count(ctx context.Context) (int64, error) {
	if err := o.check(); err != nil {
		return 0, err
	}
	return newQuery(o.appCtx, o.objectAPIName, o.err).Count(ctx)
}

func (o *Object) FindStream(ctx context.Context, recordType reflect.Type, handler func(ctx context.Context, records interface{}) error) error {
	if err := o.check(); err != nil {
		return err
	}
	return newQuery(o.appCtx, o.objectAPIName, o.err).FindStream(ctx, recordType, handler)
}

func (o *Object) FindAll(ctx context.Context, records interface{}) error {
	if err := o.check(); err != nil {
		return err
	}
	return newQuery(o.appCtx, o.objectAPIName, o.err).FindAll(ctx, records)
}

func (o *Object) Find(ctx context.Context, records interface{}) error {
	if err := o.check(); err != nil {
		return err
	}
	return newQuery(o.appCtx, o.objectAPIName, o.err).Find(ctx, records)
}

func (o *Object) FindOne(ctx context.Context, record interface{}) error {
	if err := o.check(); err != nil {
		return err
	}
	return newQuery(o.appCtx, o.objectAPIName, o.err).FindOne(ctx, record)
}

func (o *Object) Where(condition interface{}) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.err).Where(condition)
}

func (o *Object) Offset(offset int64) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.err).Offset(offset)
}

func (o *Object) Limit(limit int64) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.err).Limit(limit)
}

func (o *Object) OrderBy(fieldAPINames ...string) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.err).OrderBy(fieldAPINames...)
}

func (o *Object) OrderByDesc(fieldAPINames ...string) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.err).OrderByDesc(fieldAPINames...)
}

func (o *Object) Select(fieldAPINames ...string) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.err).Select(fieldAPINames...)
}

func (o *Object) check() error {
	return o.err
}
