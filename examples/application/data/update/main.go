package main

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/opensdk"
	fieldType "github.com/byted-apaas/server-sdk-go/service/data/field_type/opensdk"
)

type Lookup struct {
	ID        int64                  `json:"_id"`
	Name      fieldType.Multilingual `json:"_name"`
	IsDeleted bool                   `json:"_isDeleted"`
}

type TestObject struct {
	ID     int64   `json:"_id,omitempty"`
	Text   string  `json:"text"`
	Number int64   `json:"number"`
	Lookup *Lookup `json:"lookup"`
}

func main() {
	app := opensdk.NewApplication("***", "***").Env(constants.PlatformEnvDEV)

	err := app.Data.Object("testObject").Update(context.Background(), 1742602772528183,
		&TestObject{
			Text:   "代码创建文本：更新",
			Number: -1,
		})
	if err != nil {
		panic(err)
	}
}
