// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package main

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/opensdk"
)

type Lookup struct {
	ID        int64                `json:"_id"`
	Name      opensdk.Multilingual `json:"_name"`
	IsDeleted bool                 `json:"_isDeleted"`
}

type TestObject struct {
	ID     int64       `json:"_id,omitempty"`
	Text   string      `json:"text"`
	Number int64       `json:"number"`
	Lookup interface{} `json:"lookup"`
}

func main() {
	app := application.NewApplication("xxx", "xxx").Env(constants.PlatformEnvPRE)

	updateRecords := map[int64]interface{}{
		1742643063468062: &TestObject{
			Text:   "代码创建文本：更新1",
			Number: 1,
		},
		1742643063468078: &TestObject{
			Text:   "代码创建文本：更新2",
			Number: 2,
		},
	}
	err := app.Data.Object("testObject").BatchUpdate(context.Background(), updateRecords)
	if err != nil {
		panic(err)
	}
}
