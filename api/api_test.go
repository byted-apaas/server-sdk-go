// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package api

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data/cond"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
	"github.com/byted-apaas/server-sdk-go/service/tasks"
)

var (
	ctx = context.Background()
)

func init() {
	ctx = cUtils.SetDebugTypeToCtx(ctx, cConstants.DebugTypeLocal)
	ctx = context.WithValue(ctx, cConstants.CtxKeyEnvBoe, "boe")
	_ = os.Setenv("ENV", "development")
	_ = os.Setenv("KTenantName", "xxx")
	_ = os.Setenv("KNamespace", "xxx")
	_ = os.Setenv("KSvcID", "xxx")
	_ = os.Setenv("KClientID", "xxx")
	_ = os.Setenv("KClientSecret", "xxx")
	//_ = os.Setenv("KOpenApiDomain", "xxx")
	_ = os.Setenv("KOpenApiDomain", "xxx")
	_ = os.Setenv("KFaaSInfraDomain", "xxx")
	//_ = os.Setenv("KFaaSInfraDomain", "xxx")
	_ = os.Setenv("CONSUL_HTTP_HOST", "xxx")
	_ = os.Setenv("RUNTIME_IDC_NAME", "boe")
}

func TestDataFindOne(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	var record TestObject
	err := Data.Object("testObject").
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
	err := Data.Object("testObject").Where(nil).Select("_id", "text", "number", "lookup").Find(ctx, &records)
	assert.Nil(t, err)
	cUtils.PrintLog(records)
}

func TestDataCreate(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup,omitempty"`
	}
	result, err := Data.Object("test").Create(ctx,
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
	result, err := Data.Object("testObject").BatchCreate(ctx,
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

func TestDataBatchCreateAsync(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	result, err := Data.Object("testObject").BatchCreateAsync(ctx,
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
	err := Data.Object("testObject").Update(ctx, 1738115468760103,
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
	err := Data.Object("testObject").BatchUpdate(ctx,
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

func TestDataBatchUpdateAsync(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}
	taskID, err := Data.Object("testObject").BatchUpdateAsync(ctx,
		map[int64]interface{}{
			1738115468760103: &TestObject{
				Text:   fmt.Sprintf("update: %d", time.Now().Unix()),
				Number: -4,
			},
			1742103170465799: &TestObject{
				Text:   fmt.Sprintf("update: %d", time.Now().Unix()),
				Number: -4,
			},
		})
	assert.Nil(t, err)
	cUtils.PrintLog(taskID)
}

func TestDataDelete(t *testing.T) {
	err := Data.Object("testObject").Delete(ctx, 1742103159605255)
	assert.Nil(t, err)
}

func TestDataBatchDelete(t *testing.T) {
	err := Data.Object("testObject").BatchDelete(ctx, []int64{1742103159605255, 1742103160448023})
	assert.Nil(t, err)
}

func TestDataBatchDeleteAsync(t *testing.T) {
	taskID, err := Data.Object("testObject").BatchDeleteAsync(ctx, []int64{1742103159605255, 1742103160448023})
	assert.Nil(t, err)
	cUtils.PrintLog(taskID)
}

func TestDataCount(t *testing.T) {
	count, err := Data.Object("testObject").Count(ctx)
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
	err := Data.Object("testObject").FindAll(ctx, &records)
	assert.Nil(t, err)
	cUtils.PrintLog(records)
}

func TestOql(t *testing.T) {
	var rs interface{}
	err := Data.Oql("insert into testObject(text, number) values ('oql', 1)").Execute(ctx, &rs)
	assert.Nil(t, err)
	cUtils.PrintLog(rs)

	var records interface{}
	err = Data.Oql("select _id, text, number, lookup from testObject").Execute(ctx, &records)
	assert.Nil(t, err)
	cUtils.PrintLog(records)
}

func TestTransactionCreate(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}

	tx := Data.NewTransaction()
	result1, err := tx.Object("testObject").RegisterCreate(&TestObject{
		Text:   "事务创建1",
		Number: -1,
	})
	assert.Nil(t, err)
	cUtils.PrintLog(result1)

	result2, err := tx.Object("testObject").RegisterCreate(&TestObject{
		Text:   "事务创建2",
		Number: -1,
		Lookup: map[string]interface{}{"id": result1.ID},
	})
	assert.Nil(t, err)
	cUtils.PrintLog(result2)

	err = tx.Commit(ctx)
	assert.Nil(t, err)

	cUtils.PrintLog(result1)
	cUtils.PrintLog(result2)
}

func TestTransactionUpdateDelete(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}

	tx := Data.NewTransaction()
	tx.Object("testObject").RegisterUpdate(1742103171874903, &TestObject{
		Text:   "事务创建3",
		Number: -1,
	})
	tx.Object("testObject").RegisterDelete(1742103171874887)

	err := tx.Commit(ctx)
	assert.Nil(t, err)
}

func TestTransactionBatchCreate(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}

	tx := Data.NewTransaction()
	ids1, err := tx.Object("testObject").RegisterBatchCreate([]*TestObject{{
		Text:   "事务创建4",
		Number: -1,
	}, {
		Text:   "事务创建5",
		Number: -1,
	}})
	assert.Nil(t, err)
	cUtils.PrintLog(ids1)

	ids2, err := tx.Object("testObject").RegisterBatchCreate([]*TestObject{{
		Text:   "事务创建6",
		Number: -1,
		Lookup: map[string]interface{}{"id": ids1[0]},
	}, {
		Text:   "事务创建7",
		Number: -1,
		Lookup: map[string]interface{}{"id": ids1[1]},
	}})
	assert.Nil(t, err)
	cUtils.PrintLog(ids2)

	err = tx.Commit(ctx)
	assert.Nil(t, err)
	cUtils.PrintLog(ids1)
	cUtils.PrintLog(ids2)
}

func TestTransactionBatchUpdateDelete(t *testing.T) {
	type TestObject struct {
		ID     int64       `json:"_id"`
		Text   string      `json:"text"`
		Number int64       `json:"number"`
		Lookup interface{} `json:"lookup"`
	}

	tx := Data.NewTransaction()
	tx.Object("testObject").RegisterBatchUpdate([]*TestObject{{
		ID:     1742103165680734,
		Text:   "事务创建44",
		Number: -1,
	}, {
		ID:     1742103165680750,
		Text:   "事务创建55",
		Number: -1,
	}})

	tx.Object("testObject").RegisterBatchDelete([]int64{1742103165681742, 1742103165680766})

	err := tx.Commit(ctx)
	assert.Nil(t, err)
}

func TestFunctionInvoke(t *testing.T) {
	var result interface{}
	err := Function("syncFunc_v2").Invoke(ctx, map[string]interface{}{"a": "aa"}, &result)
	assert.Nil(t, err)
	cUtils.PrintLog(result)
}

func TestFunctionInvokeAsync(t *testing.T) {
	taskID, err := Tasks.CreateAsyncTask(ctx, "syncFunc_v2", map[string]interface{}{"a": "aa"})
	assert.Nil(t, err)
	cUtils.PrintLog(taskID)
}

func TestFunctionInvokeDistributed(t *testing.T) {
	taskID, err := Tasks.CreateDistributedTask(ctx, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "handlerFunc_v2",
		"progressCallbackFunc_v2", "completedCallbackFunc_v2", tasks.NewOptions(5, 5, 1))
	assert.Nil(t, err)
	cUtils.PrintLog(taskID)
}

func TestCreateMsg(t *testing.T) {
	taskID, err := Msg.NotifyCenter.Create(ctx, &structs.MessageBody{
		Icon:        "info",
		Percent:     10,
		TargetUsers: []int64{1737592670388231, 1736970625231931},
		Title:       &faassdk.Multilingual{{1033, "英文标题112"}, {2052, "中文标题112"}},
		Detail:      &faassdk.Multilingual{{2052, "中文详情112"}, {1033, "英文详情112"}},
	})
	assert.Nil(t, err)
	cUtils.PrintLog(taskID)
}

func TestUpdateMsg(t *testing.T) {
	err := Msg.NotifyCenter.Update(ctx, 1742127940874270, &structs.MessageBody{
		Icon:        "info",
		Percent:     10,
		TargetUsers: []int64{1737592670388231},
		Title:       &faassdk.Multilingual{{1033, "英文标题2"}, {2052, "中文标题2"}},
		Detail:      &faassdk.Multilingual{{2052, "中文详情2"}, {1033, "英文详情2"}},
	})
	assert.Nil(t, err)
}

func TestUploadAttachment(t *testing.T) {
	result, err := Resources.File.UploadByPath(ctx, "api_test.go", "api_test.go", 0)
	assert.Nil(t, err)
	cUtils.PrintLog(result)
}

func TestDownloadAttachment(t *testing.T) {
	result, err := Resources.File.Download(ctx, "2d57ddafd22d4c9a9f285dbef7b127d8")
	assert.Nil(t, err)
	cUtils.PrintLog(string(result))
}

func TestGlobalVar(t *testing.T) {
	result, err := GetVar(ctx, "key1")
	assert.Nil(t, err)
	cUtils.PrintLog(result)
}

func TestGetField(t *testing.T) {
	var field interface{}
	err := Metadata.Object("testObject").GetField(ctx, "lookup", &field)
	assert.Nil(t, err)
	cUtils.PrintLog(field)
}

func TestGetFields(t *testing.T) {
	var fields map[string]interface{}
	err := Metadata.Object("testObject").GetFields(ctx, &fields)
	assert.Nil(t, err)
	cUtils.PrintLog(fields)
}
