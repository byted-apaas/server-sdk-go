// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package application

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/service/data/cond"
)

var (
	app     *Application
	faasApp *Application
	ctx     = context.Background()
)

func init() {
	app = NewApplication("xxx", "xxx").Env(constants.PlatformEnvPRE)
	ctx = cUtils.SetDebugTypeToCtx(context.Background(), cConstants.DebugTypeLocal)
}

func TestDataFindOne(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	var record TestObject
	err := app.Data.Object("testObject").
		Where(cond.And(cond.Eq("number", 456), cond.Eq("lookup.number", 123))).
		Select("_id", "text", "number", "lookup").FindOne(ctx, &record)
	assert.Nil(t, err)
	cUtils.PrintLog(record)
}

func TestDataFind(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	var records []*TestObject
	err := app.Data.Object("testObject").Where(nil).Select("_id", "text", "number", "lookup").Find(ctx, &records)
	assert.Nil(t, err)
	cUtils.PrintLog(records)
}

func TestDataWhereNotEmpty(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	var records []*TestObject
	err := app.Data.Object("testObject").Where(cond.NotEmpty("text")).Select("_id", "text", "number", "lookup").Find(ctx, &records)
	assert.Nil(t, err)
	cUtils.PrintLog(records)

	err = app.Data.Object("testObject").Where(cond.Empty("text")).Select("_id", "text", "number", "lookup").Find(ctx, &records)
	assert.Nil(t, err)
	cUtils.PrintLog(records)
}

func TestDataCreate(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	result, err := app.Data.Object("testObject").Create(ctx,
		&TestObject{
			Text:   "代码创建文本",
			Number: -1,
		})
	assert.Nil(t, err)
	cUtils.PrintLog(result)
}

func TestDataBatchCreate(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	result, err := app.Data.Object("testObject").BatchCreate(ctx,
		[]*TestObject{{
			Text:   "代码创建文本1",
			Number: -1,
		}, {
			Text:   "代码创建文本2",
			Number: -1,
		}})
	assert.Nil(t, err)
	cUtils.PrintLog(result)
}

func TestDataUpdate(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id,omitempty"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	err := app.Data.Object("testObject").Update(ctx, 1738115468760103,
		&TestObject{
			Text:   fmt.Sprintf("update: %d", time.Now().Unix()),
			Number: -2,
		})
	assert.Nil(t, err)
}

func TestDataBatchUpdate(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	err := app.Data.Object("testObject").BatchUpdate(ctx,
		map[int64]interface{}{
			1738115468760103: &TestObject{
				Text:   fmt.Sprintf("update: %d", time.Now().Unix()),
				Number: -3,
			},
			1742103170465799: &TestObject{
				Text:   fmt.Sprintf("update: %d", time.Now().Unix()),
				Number: -3,
			},
		})
	assert.Nil(t, err)
}

func TestDataDelete(t *testing.T) {
	err := app.Data.Object("testObject").Delete(ctx, 1742103159605255)
	assert.Nil(t, err)
}

func TestDataBatchDelete(t *testing.T) {
	err := app.Data.Object("testObject").BatchDelete(ctx, []int64{1742103159605255, 1742103160448023})
	assert.Nil(t, err)
}

func TestDataCount(t *testing.T) {
	count, err := app.Data.Object("testObject").Count(ctx)
	assert.Nil(t, err)
	cUtils.PrintLog(count)
}

func TestDataFindAll(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	var records []*TestObject
	err := app.Data.Object("testObject").FindAll(ctx, &records)
	assert.Nil(t, err)
	cUtils.PrintLog(records)
}

func TestFunctionInvoke(t *testing.T) {
	var result interface{}
	err := app.Function("syncFunc_v2").Invoke(ctx, map[string]interface{}{"a": "aa"}, &result)
	assert.Nil(t, err)
	cUtils.PrintLog(result)
}

func TestUploadAttachment(t *testing.T) {
	result, err := app.Resources.File.UploadByPath(ctx, "opensdk_test.go", "opensdk_test.go")
	assert.Nil(t, err)
	cUtils.PrintLog(result)
}

func TestDownloadAttachment(t *testing.T) {
	result, err := app.Resources.File.Download(ctx, "2d57ddafd22d4c9a9f285dbef7b127d8")
	assert.Nil(t, err)
	cUtils.PrintLog(string(result))
}

func TestUploadAvatar(t *testing.T) {
	result, err := app.Resources.File.UploadAvatarByPath(ctx, "avatar.png", "/Users/zhouweixin/Downloads/avatar.png")
	assert.Nil(t, err)
	cUtils.PrintLog(result)
}

func TestDownloadAvatar(t *testing.T) {
	result, err := app.Resources.File.DownloadAvatar(ctx, "7a474cd18ab74366aed5f058cc4c9c55_o")
	assert.Nil(t, err)

	err = os.WriteFile("/Users/zhouweixin/Downloads/avatar1.png", result, 0666)
	assert.Nil(t, err)
}
