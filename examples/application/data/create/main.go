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
	ID     int64   `json:"_id,omitempty"`
	Text   string  `json:"text"`
	Number int64   `json:"number"`
	Lookup *Lookup `json:"lookup"`
}

func main() {
	app := application.NewApplication("xxx", "xxx").Env(constants.PlatformEnvPRE)



	// 创建记录1
	record1, err := app.Data.Object("testObject").Create(context.Background(),
		&TestObject{
			Text:   "代码创建文本",
			Number: -1,
		})
	if err != nil {
		panic(err)
	}
	fmt.Printf("record1: %+v\n", record1)

	// 创建记录2，关联记录1
	record2, err := app.Data.Object("testObject").Create(context.Background(),
		&TestObject{
			Text:   "代码创建文本",
			Number: -1,
			Lookup: &Lookup{
				ID: record1.ID,
			},
		})
	if err != nil {
		panic(err)
	}
	fmt.Printf("record2: %+v\n", record2)
}
