// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package metadata

import (
	"context"
	"testing"

	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/service/metadata/object/fields"
)

var (
	m   = &Metadata{}
	ctx = context.Background()
)

func Init() {
}

func TestMain(m *testing.M) {
	Init()
	m.Run()
}

func TestGetFields(t *testing.T) {
	// 1. 定义对象元数据结构体，字段类型使用 fields 下的结构
	type Result struct { // 定义对应的结构，注意 tag 对应 apiName
		CreatedBy fields.Lookup `json:"_createdBy"`
	}
	// 调用接口，decode 拿到结果
	res := Result{}
	err := m.Object("all_object").GetFields(ctx, &res)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(res)

	// 2. 通过 map 获取所有, key 即为 apiName
	mapRes := make(map[string]interface{})
	err = m.Object("all_object").GetFields(ctx, &mapRes)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(mapRes)
}

func TestGetField(t *testing.T) {
	// 1. 使用 fields 定义好的结构体，获取
	res := fields.Multilingual{}
	err := m.Object("all_object").GetField(ctx, "_createdBy", &res)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(res)

	// 2. 使用 map 获取所有
	mapRes := make(map[string]interface{})
	err = m.Object("all_object").GetField(ctx, "option", &mapRes)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(mapRes)
}

func TestGetField_All(t *testing.T) {
	autoID := fields.AutoID{}
	err := m.Object("all_object").GetField(ctx, "autoid", &autoID)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(autoID)

	avatar := fields.AvatarOrLogo{}
	err = m.Object("all_object").GetField(ctx, "avatar", &avatar)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(avatar)

	boolean := fields.Boolean{}
	err = m.Object("all_object").GetField(ctx, "bool", &boolean)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(boolean)

	comType := fields.CompositeType{}
	err = m.Object("all_object").GetField(ctx, "subobj", &comType)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(comType)

	date := fields.Date{}
	err = m.Object("all_object").GetField(ctx, "date", &date)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(date)

	datetime := fields.DateTime{}
	err = m.Object("all_object").GetField(ctx, "datetime", &datetime)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(datetime)

	email := fields.Email{}
	err = m.Object("all_object").GetField(ctx, "email", &email)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(email)

	m2s := fields.ExtractSingleRecord{}
	err = m.Object("all_object").GetField(ctx, "m2s", &m2s)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(m2s)

	file := fields.File{}
	err = m.Object("all_object").GetField(ctx, "file", &file)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(file)

	formula := fields.Formula{}
	err = m.Object("all_object").GetField(ctx, "formula", &formula)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(formula)

	lookup := fields.Lookup{}
	err = m.Object("all_object").GetField(ctx, "lookup", &lookup)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(lookup)

	phone := fields.MobileNumber{}
	err = m.Object("all_object").GetField(ctx, "phone", &phone)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(phone)

	multil := fields.Multilingual{}
	err = m.Object("all_object").GetField(ctx, "multil", &multil)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(multil)

	number := fields.Number{}
	err = m.Object("all_object").GetField(ctx, "number", &number)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(number)

	option := fields.Option{}
	err = m.Object("all_object").GetField(ctx, "option", &option)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(option)

	refer := fields.ReferenceField{}
	err = m.Object("all_object").GetField(ctx, "refer", &refer)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(refer)

	richtext := fields.RichText{}
	err = m.Object("all_object").GetField(ctx, "richtext", &richtext)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(richtext)

	text := fields.Text{}
	err = m.Object("all_object").GetField(ctx, "text", &text)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(text)
}
