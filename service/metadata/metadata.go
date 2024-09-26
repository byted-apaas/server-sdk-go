// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package metadata

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/metadata/object"
)

// IMetadata 元数据读写接口
type IMetadata interface {
	// Object 操作指定对象的元数据信息
	Object(objectAPIName string) IObject
}

// IObject 对象读写接口
type IObject interface {
	// GetFields 读取对象的字段元数据信息列表
	GetFields(ctx context.Context, fields interface{}) (err error)
	// GetField 读取指定字段的元数据信息
	GetField(ctx context.Context, fieldAPIName string, field interface{}) (err error)
}

type Metadata struct {
	appCtx *structs.AppCtx
}

func (m *Metadata) Object(objectAPIName string) IObject {
	return object.NewObject(m.appCtx, objectAPIName)
}

func NewMetadata(appCtx *structs.AppCtx) IMetadata {
	if appCtx == nil {
		appCtx = &structs.AppCtx{}
	}
	appCtx.DataVersion = structs.DataVersionV1
	return &Metadata{
		appCtx: appCtx,
	}
}
