package user

import (
	"context"

	cStructs "github.com/byted-apaas/server-common-go/structs"
	cUtils "github.com/byted-apaas/server-common-go/utils"
)

type UserContext = cStructs.UserContext

// GetContext 获得一些上游传入的系统上下文参数, 指定特定结构
func GetContext(ctx context.Context) UserContext {
	return cUtils.GetUserContext(ctx)
}

// GetContextMap 获取上游传入的系统上下文参数, Map 通用结构
func GetContextMap(ctx context.Context) map[string]interface{} {
	return cUtils.GetUserContextMap(ctx)
}

//type IUserContext interface {
//	// GetFlow 获取流程的一些系统参数上下文
//	GetFlow(ctx context.Context) Flow
//}
//
//type UserContext struct {
//	appCtx *structs.AppCtx
//}
//
//func (u *UserContext) GetFlow(ctx context.Context) cStructs.Flow {
//	return cUtils.GetUserContext(ctx).Flow
//}
//
//func NewUserContext(appCtx *structs.AppCtx) IUserContext {
//	return &UserContext{appCtx: appCtx}
//}
