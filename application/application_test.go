// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package application

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/service/data/cond"
	"github.com/stretchr/testify/assert"
)

var (
	app     *Application
	faasApp *Application
	ctx     = context.Background()
)

func init() {
	app = NewApplication("xxx", "xxx").Env(constants.PlatformEnvLR)
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
		Where(cond.And(cond.Eq("number", 123), cond.Eq("lookup.number", 123))).
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
			Number: 123,
			Lookup: map[string]interface{}{"_id": 1747281751947339},
		})
	assert.Nil(t, err)
	cUtils.PrintLog(result)
}

var (
	delList []int64
	delID   int64
)

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
		}, {
			Text:   "代码创建文本3",
			Number: -1,
		}})
	assert.Nil(t, err)
	cUtils.PrintLog(result)
	for i, r := range result {
		if i > 1 {
			break
		}
		delList = append(delList, r)
	}
	delID = result[2]
}

func TestDataUpdate(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id,omitempty"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	err := app.Data.Object("testObject").Update(ctx, 1747281618334803,
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
			1747280793268232: &TestObject{
				Text:   fmt.Sprintf("update: %d", time.Now().Unix()),
				Number: -3,
			},
			1747280884473971: &TestObject{
				Text:   fmt.Sprintf("update: %d", time.Now().Unix()),
				Number: -3,
			},
		})
	assert.Nil(t, err)
}

func TestDataDelete(t *testing.T) {
	err := app.Data.Object("testObject").Delete(ctx, delID)
	assert.Nil(t, err)
}

func TestDataBatchDelete(t *testing.T) {
	err := app.Data.Object("testObject").BatchDelete(ctx, delList)
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
	err := app.Function("long").Invoke(ctx, map[string]interface{}{"sleep": 100}, &result)
	assert.Nil(t, err)
	cUtils.PrintLog(result)
}

var fileID string

func TestUploadAttachment(t *testing.T) {
	result, err := app.Resources.File.UploadByPath(ctx, "application_test.go", "application_test.go")
	assert.Nil(t, err)
	cUtils.PrintLog(result)
	fileID = result.FileID
}

func TestDownloadAttachment(t *testing.T) {
	result, err := app.Resources.File.Download(ctx, fileID)
	assert.Nil(t, err)
	cUtils.PrintLog(string(result))
}

var image string

func TestUploadAvatar(t *testing.T) {
	result, err := app.Resources.File.UploadAvatarByPath(ctx, "avatar.png", "/Users/xxx/test/pic/avatar.png")
	assert.Nil(t, err)
	cUtils.PrintLog(result)
	image = result.ImageID
}

func TestDownloadAvatar(t *testing.T) {
	result, err := app.Resources.File.DownloadAvatar(ctx, image)
	assert.Nil(t, err)

	err = os.WriteFile("/Users/xxx/test/pic/avatar1.png", result, 0666)
	assert.Nil(t, err)
}
