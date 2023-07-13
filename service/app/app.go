package app

import (
	"context"

	cStructs "github.com/byted-apaas/server-common-go/structs"
)

type IApp interface {
	// GetAppInfo 获取应用信息
	GetAppInfo(ctx context.Context) (*cStructs.AppInfo, error)
}
