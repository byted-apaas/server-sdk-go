package global_var

import (
	"context"
)

type IGlobalVar interface {
	// GetVar 获取全局变量
	GetVar(ctx context.Context, key string) (value string, err error)

	// GetVariable 获取全局变量
	GetVariable(ctx context.Context, key string) (string, error)
}
