// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"fmt"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/service/data/cond"
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
	app := application.NewApplication("c_c9c02f05c7c14131a6b4", "b85740891cbd419c86887c47082e9e13").Env(constants.PlatformEnvDEV)

	var records []*TestObject
	err := app.Data.Object("testObject").Where(
		cond.And(
			cond.Lt("number", 0),
			cond.NotEmpty("text"),
		)).Select("_id", "text", "number").Find(context.Background(), &records)
	if err != nil {
		panic(err)
	}
	fmt.Printf("records: %d\n", len(records))
}
