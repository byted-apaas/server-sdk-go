package main

import "github.com/byted-apaas/server-sdk-go/application"

var (
	recordTransaction = map[string]interface{}{
		"text": "recordTransaction",
		//"bigintType":   "2",
		//"dateType":     "2024-08-20",
		//"datetimeType": 1724688780000,
		//"decimal":      "1",
		//"formula":      "2",
		//"option":       "option_7f97916560b",
		////"richText": map[string]interface{}{
		////	"raw": "<div style=\"white-space: pre-wrap;\">&lt;p&gt;test&lt;/p&gt;</div>",
		////},
		//"richText": "<div style=\"white-space: pre-wrap;\">&lt;p&gt;test&lt;/p&gt;</div>",
		//"avatar": map[string]interface{}{
		//	"image": map[string]interface{}{
		//		"token": "e09c1f0c43ff4019bd0463fc3bbe8821_c",
		//		"uri":   "/img/553944/e09c1f0c43ff4019bd0463fc3bbe8821_l.jpg",
		//	},
		//	"source": "image",
		//},
	}
)

func transactionV3() {
	// 创建一个新的空事务
	tx := application.DataV3.NewTransaction()
	// 注册创建
	id, err := tx.Object("objectForAll").RegisterCreate(recordTransaction)
	if err != nil {
		application.GetLogger(ctx).Errorf("create record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("create record success, id: %+v", id)

	//// 注册更新
	//tx.Object("objectForAll").RegisterUpdate(id, updateRecordStruct)
	//// 注册删除
	//tx.Object("objectForAll").RegisterDelete(id)

	err = tx.Commit(ctx)
	if err != nil {
		application.GetLogger(ctx).Errorf("commit transaction error: %+v", err)
		return
	}
}

func batchTransactionV3() {
	tx := application.DataV3.NewTransaction()

	// 注册批量 db 操作，单个批量操作数据量限制 100 条，多个批量操作的总数据量限制 500 条
	// 注册批量创建
	ids, err := tx.Object("objectForAll").RegisterBatchCreate([]interface{}{recordTransaction})
	if err != nil {
		application.GetLogger(ctx).Errorf("batch create record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("create record success, ids: %+v", ids)

	//// 注册批量更新
	//tx.Object("objectForAll").RegisterBatchUpdate([]interface{}{updateRecordStruct})
	//// 注册批量删除
	//tx.Object("objectForAll").RegisterBatchDelete([]interface{}{recordInterface})

	// 提交事务
	tx.Commit(ctx)
}
