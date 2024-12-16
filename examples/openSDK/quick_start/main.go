package main

import (
	"context"
	"fmt"

	utils2 "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/utils"
	"github.com/byted-apaas/server-sdk-go/opensdk"
)

// User 定义目标结构
type User struct {
	ID        int64  `json:"_id"`
	Email     string `json:"_email"`
	IsDeleted bool   `json:"_isDeleted"`
}

func main() {
	// 初始化空的 ctx
	ctx := context.Background()

	// 默认为云端运行模式，走的 RPC 协议，本地通过 RPC 访问不了 TCE。
	// 因此，在本地运行时需要设置为本地调试模式
	ctx = utils.LocalDebugMode(ctx)

	// 创建实例（注意：1.如何申请 API 凭证；2.Env 是什么）
	app := application.NewApplication("***", "***").Env(constants.PlatformEnvDEV)

	// 创建变量接收查询的结果
	var users []*User

	// 查询数据（注意：users 变量前的取地址符号不可省略）
	err := app.Data.Object("_user").Select("_id", "_email", "_isDeleted").Find(ctx, &users)
	if err != nil {
		// 示例代码中的 err 不做特殊处理，开发者自行按需处理
		panic(err)
	}

	// 序列化数据，方便打印
	usersBytes, err := utils2.JsonMarshalBytes(users)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(usersBytes))
}
