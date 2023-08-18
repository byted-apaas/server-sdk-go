package app

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/structs"
)

type IApp interface {
	// GetAppInfo 获取应用信息
	GetAppInfo(ctx context.Context) (*structs.AppInfo, error)
}
