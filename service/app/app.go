package app

import (
	"context"
)

type IApp interface {
	// GetAppInfo 获取应用信息
	GetAppInfo(ctx context.Context) (*cStructs.AppInfo, error)
}
