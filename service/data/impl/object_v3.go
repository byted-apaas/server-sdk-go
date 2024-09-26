package data

import (
	"context"
	"reflect"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/common/utils"
	"github.com/byted-apaas/server-sdk-go/request"
	"github.com/byted-apaas/server-sdk-go/service/data"
)

type ObjectV3 struct {
	appCtx        *structs.AppCtx
	objectAPIName string
	err           error
	authType      *string
}

func NewObjectV3(s *structs.AppCtx, objectAPIName string) *ObjectV3 {
	o := &ObjectV3{
		appCtx:        s,
		objectAPIName: objectAPIName,
	}

	if objectAPIName == "" {
		o.err = cExceptions.InvalidParamError("[Object] objectAPIName is empty")
	}
	return o
}

func (o *ObjectV3) Create(ctx context.Context, record interface{}) (*structs.RecordIDV3, error) {
	ctx = cUtils.SetUserAndAuthTypeToCtx(ctx, o.authType)
	if err := o.check(); err != nil {
		return nil, err
	}

	return request.GetInstance(ctx).CreateRecordV3(ctx, o.appCtx, o.objectAPIName, record)
}

func (o *ObjectV3) BatchCreate(ctx context.Context, records interface{}) ([]string, error) {
	ctx = cUtils.SetUserAndAuthTypeToCtx(ctx, o.authType)
	if err := o.check(); err != nil {
		return nil, err
	}

	return request.GetInstance(ctx).BatchCreateRecordV3(ctx, o.appCtx, o.objectAPIName, records)
}

func (o *ObjectV3) Update(ctx context.Context, _id string, record interface{}) error {
	ctx = cUtils.SetUserAndAuthTypeToCtx(ctx, o.authType)
	if err := o.check(); err != nil {
		return err
	}

	return request.GetInstance(ctx).UpdateRecordV3(ctx, o.appCtx, o.objectAPIName, _id, record)
}

func (o *ObjectV3) BatchUpdate(ctx context.Context, records map[string]interface{}, result ...interface{}) error {
	ctx = cUtils.SetUserAndAuthTypeToCtx(ctx, o.authType)
	if err := o.check(); err != nil {
		return err
	}

	resp, err := request.GetInstance(ctx).BatchUpdateRecordV3(ctx, o.appCtx, o.objectAPIName, records)

	if err != nil {
		return err
	}

	if len(result) > 0 {
		return utils.ParseBatchResultV3(resp, result[0])
	}
	return nil
}

func (o *ObjectV3) Delete(ctx context.Context, _id string) error {
	ctx = cUtils.SetUserAndAuthTypeToCtx(ctx, o.authType)
	if err := o.check(); err != nil {
		return err
	}
	return request.GetInstance(ctx).DeleteRecordV3(ctx, o.appCtx, o.objectAPIName, _id)
}

func (o *ObjectV3) BatchDelete(ctx context.Context, _ids []string, result ...interface{}) error {
	ctx = cUtils.SetUserAndAuthTypeToCtx(ctx, o.authType)
	if err := o.check(); err != nil {
		return err
	}

	resp, err := request.GetInstance(ctx).BatchDeleteRecordV3(ctx, o.appCtx, o.objectAPIName, _ids)
	if err != nil {
		return err
	}

	if len(result) > 0 {
		return utils.ParseBatchResultV3(resp, result[0])
	}
	return nil
}

func (o *ObjectV3) Count(ctx context.Context) (int64, error) {
	if err := o.check(); err != nil {
		return 0, err
	}
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).Count(ctx)
}

func (o *ObjectV3) FindStream(ctx context.Context, recordType reflect.Type, handler func(ctx context.Context, records interface{}) error, params ...structs.FindStreamParam) error {
	if err := o.check(); err != nil {
		return err
	}
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).FindStream(ctx, recordType, handler, params...)
}

// FindAll Deprecated
func (o *ObjectV3) FindAll(ctx context.Context, records interface{}) error {
	if err := o.check(); err != nil {
		return err
	}
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).FindAll(ctx, records)
}

func (o *ObjectV3) Find(ctx context.Context, records interface{}, unauthFields ...interface{}) error {
	if err := o.check(); err != nil {
		return err
	}
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).Find(ctx, records, unauthFields...)
}

func (o *ObjectV3) FindOne(ctx context.Context, record interface{}, unauthFields ...interface{}) error {
	if err := o.check(); err != nil {
		return err
	}
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).FindOne(ctx, record, unauthFields...)
}

func (o *ObjectV3) Where(condition interface{}) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).Where(condition)
}

func (o *ObjectV3) FuzzySearch(keyword string, fieldAPINames []string) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).FuzzySearch(keyword, fieldAPINames)
}

func (o *ObjectV3) Offset(offset int64) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).Offset(offset)
}

func (o *ObjectV3) Limit(limit int64) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).Limit(limit)
}

func (o *ObjectV3) OrderBy(fieldAPINames ...string) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).OrderBy(fieldAPINames...)
}

func (o *ObjectV3) OrderByDesc(fieldAPINames ...string) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).OrderByDesc(fieldAPINames...)
}

func (o *ObjectV3) Select(fieldAPINames ...string) data.IQuery {
	return newQuery(o.appCtx, o.objectAPIName, o.authType, o.err).Select(fieldAPINames...)
}

func (o *ObjectV3) UseUserAuth() data.IObjectV3 {
	o.authType = cUtils.StringPtr(cConstants.AuthTypeUser)
	return o
}

func (o *ObjectV3) UseSystemAuth() data.IObjectV3 {
	o.authType = cUtils.StringPtr(cConstants.AuthTypeSystem)
	return o
}

func (o *ObjectV3) check() error {
	return o.err
}
