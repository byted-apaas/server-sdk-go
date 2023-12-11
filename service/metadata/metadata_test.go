// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package metadata

import (
	"context"
	"os"
	"testing"

	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/service/metadata/object/fields"
)

var (
	m   = &Metadata{}
	ctx = context.Background()
)

func Init() {
	os.Setenv("KClientID", "01d8a2f7a5fbeb687e17f9d118cd6e68db941689e224ab8e158b930880b0d3ae")
	os.Setenv("KClientSecret", "b912ef8f3dc309ae4568369168f8b5b60e274ddfb5a8706d082277beba699b8ebdbfeafe091fe9efd3f152b519f369d7")
	os.Setenv("KInnerAPIDomain", "https://apaas-innerapi-boe.bytedance.net")
	os.Setenv("KOpenApiDomain", "http://oapi-kunlun-staging-boe.byted.org")
	os.Setenv("KFaaSInfraDomain", "http://apaas-faasinfra-staging-boe.bytedance.net")
	os.Setenv("KTenantName", "zhouweixin02-dev16")
	os.Setenv("KNamespace", "sdk_sample__c")
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
	err := m.Object("allFieldObject").GetFields(ctx, &res)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(res)

	// 2. 通过 map 获取所有, key 即为 apiName
	mapRes := make(map[string]interface{})
	err = m.Object("allFieldObject").GetFields(ctx, &mapRes)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(mapRes)
}

func TestGetField(t *testing.T) {
	// 1. 使用 fields 定义好的结构体，获取
	res := fields.Multilingual{}
	err := m.Object("allFieldObject").GetField(ctx, "_createdBy", &res)
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

func TestGetFieldBigint(t *testing.T) {
	var bigint interface{}
	err := m.Object("allFieldObject").GetField(ctx, "bigint", &bigint)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(bigint)
}

func TestGetFieldDecimal(t *testing.T) {
	var decimal interface{}
	err := m.Object("allFieldObject").GetField(ctx, "decimal", &decimal)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(decimal)
}

func TestGetFieldRegion(t *testing.T) {
	var region interface{}
	err := m.Object("allFieldObject").GetField(ctx, "region", &region)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(region)
}

func TestGetFieldRollupCount(t *testing.T) {
	var field interface{}
	err := m.Object("allFieldObject").GetField(ctx, "rollup_count", &field)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(field)
}

func TestGetFieldRollupAvg(t *testing.T) {
	var field interface{}
	err := m.Object("allFieldObject").GetField(ctx, "rollup_avg", &field)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(field)
}

func TestGetFieldRollupSum(t *testing.T) {
	var field interface{}
	err := m.Object("allFieldObject").GetField(ctx, "rollup_sum", &field)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(field)
}

func TestGetFieldRollupMax(t *testing.T) {
	var field interface{}
	err := m.Object("allFieldObject").GetField(ctx, "rollup_max", &field)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(field)
}

func TestGetFieldRollupMin(t *testing.T) {
	var field interface{}
	err := m.Object("allFieldObject").GetField(ctx, "rollup_min", &field)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(field)
}

func TestGetFieldText(t *testing.T) {
	var region interface{}
	err := m.Object("allFieldObject").GetField(ctx, "text", &region)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(region)
}
