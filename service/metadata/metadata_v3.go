package metadata

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/metadata/object"
)

// IMetadataV3 元数据读写接口
type IMetadataV3 interface {
	// Object 操作指定对象的元数据信息
	Object(objectAPIName string) IObject
}

type MetadataV3 struct {
	appCtx *structs.AppCtx
}

func (m *MetadataV3) Object(objectAPIName string) IObject {
	return object.NewObject(m.appCtx, objectAPIName)
}

func NewMetadataV3(appCtx *structs.AppCtx) IMetadataV3 {
	if appCtx == nil {
		appCtx = &structs.AppCtx{}
	}
	appCtx.DataVersion = structs.DataVersionV3
	return &MetadataV3{
		appCtx: appCtx,
	}
}
