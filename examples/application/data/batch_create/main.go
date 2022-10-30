// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"fmt"

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

	// 批量创建
	records1, err := app.Data.Object("testObject").BatchCreate(context.Background(),
		[]*TestObject{{
			Text:   "代码创建文本1",
			Number: 1,
		}, {
			Text:   "代码创建文本2",
			Number: 2,
		}})
	if err != nil {
		panic(err)
	}
	fmt.Printf("records1: %+v\n", records1)

	// 批量关联创建
	records2, err := app.Data.Object("testObject").BatchCreate(context.Background(),
		[]*TestObject{{
			Text:   "代码创建文本1",
			Number: 11,
			Lookup: &Lookup{
				ID: records1[0],
			},
		}, {
			Text:   "代码创建文本2",
			Number: 22,
			Lookup: &Lookup{
				ID: records1[1],
			},
		}})
	if err != nil {
		panic(err)
	}
	fmt.Printf("records2: %+v\n", records2)
}
