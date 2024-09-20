package main

import (
	"time"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
)

var (
	recordTransaction = map[string]interface{}{
		"text":         "recordTransaction",
		"text2":        "recordTransaction",
		"bigintType":   "2",
		"number":       "1.1",
		"dateType":     "2024-08-20",
		"datetimeType": time.Now().UnixMilli(),
		"decimal":      "1",
		"formula":      "4",
		"email":        "transaction@interface.com",
		"option":       "option_7f97916560b",
		"option2":      []string{"option_8a6e42e40af", "option_79ff1b98f80"},
		//"richText": map[string]interface{}{
		//	"raw": "<div style=\"white-space: pre-wrap;\">&lt;p&gt;test&lt;/p&gt;</div>",
		//},
		"richText": "<div style=\"white-space: pre-wrap;\">&lt;p&gt;test&lt;/p&gt;</div>",
		"avatar": map[string]interface{}{
			"image": map[string]interface{}{
				"token": "e09c1f0c43ff4019bd0463fc3bbe8821_c",
				"uri":   "/img/553944/e09c1f0c43ff4019bd0463fc3bbe8821_l.jpg",
			},
			"source": "image",
		},
		"attachment": []map[string]interface{}{
			{
				"mime_type": "xlsx",
				"name":      "待升级 app.xlsx",
				"size":      "12237",
				"token":     "BIZ_06aa3803148f4d7791ab3699e84a5ac3",
				"uri":       "/ae/api/v1/assets/attachment/download?token=BIZ_06aa3803148f4d7791ab3699e84a5ac3",
			},
		},
		"phone": map[string]interface{}{
			"dialing_code": "+86",
			"number":       "18888888888",
			"region_code":  "CN",
		},
		"multilingual": map[string]interface{}{
			"zh-CN": "多语",
			"en-US": "multilingual",
		},
		"lookup": "1810084405164122",
		"region": map[string]interface{}{
			"_id":         "1736322142527502",
			"region_code": "FSM",
			"fullPath": map[string]interface{}{
				"zh-CN": "密克罗尼西亚联邦",
				"en-US": "Micronesia (Federated States of)",
			},
		},
	}

	recordTransactionStruct = TestTransactionCreateObjectV2{
		Text:       "recordTransactionStruct",
		Text2:      "xx",
		BigintType: "3",
		DateType:   "2024-09-09",
		Datetime:   time.Now().UnixMilli(),
		Decimal:    "2",
		Formula:    "3",
		Number:     111.1,
		Boolean:    false,
		RichText: &faassdk.RichTextV3{
			Raw: "<div style=\"white-space: pre-wrap;\">&lt;p&gt;test111&lt;/p&gt;</div>",
		},
		Phone: &faassdk.PhoneNumberV3{
			DialingCode: "+86",
			Number:      "18888888888",
			RegionCode:  "CN",
		},
		Option:  "option_7f97916560b",
		Option2: []string{"option_8a6e42e40af", "option_79ff1b98f80"},
		Avatar: &faassdk.AvatarV3{
			Image: &faassdk.AttachmentModelV3{
				Token: "e09c1f0c43ff4019bd0463fc3bbe8821_c",
				URI:   "/img/553944/e09c1f0c43ff4019bd0463fc3bbe8821_l.jpg",
			},
			Source: "image",
		},
		Attachment: []*faassdk.AttachmentModelV3{
			{
				MimeType: "xlsx",
				Name:     "待升级 app.xlsx",
				Size:     "12237",
				Token:    "BIZ_06aa3803148f4d7791ab3699e84a5ac3",
				URI:      "/ae/api/v1/assets/attachment/download?token=BIZ_06aa3803148f4d7791ab3699e84a5ac3",
			},
		},
		Region: &faassdk.RegionV3{
			ID:         "1736322104113191",
			RegionCode: "MDCY00003394",
			FullPath: faassdk.MultilingualV3{
				Zh: " / 瓜达拉哈拉",
				En: "Guadalajara /  / Spain",
			},
		},
		Lookup: &faassdk.LookupV3{
			ID: "1810393764956212",
			Name: faassdk.MultilingualV3{
				Zh: "1810393764956212",
			},
		},
	}
)

func transactionV3() {
	// 创建一个新的空事务
	tx := application.DataV3.NewTransaction()

	var record []TestObjectV2
	err := application.DataV3.Object("objectForAll").Select(AllFieldAPINames...).Limit(2).Find(ctx, &record)
	if err != nil {
		application.GetLogger(ctx).Errorf("find record error: %+v", err)
		return
	}

	if len(record) < 2 {
		application.GetLogger(ctx).Errorf("find record error: %+v", err)
		return
	}

	//// 注册删除
	//tx.Object("objectForAll").RegisterDelete(record[0].ID)
	//application.GetLogger(ctx).Infof("delete record, id: %+v", record[0].ID)

	//// 注册更新
	updateRecord := recordTransactionStruct
	updateRecord.Text2 = "updateTransaction"
	tx.Object("objectForAll").RegisterUpdate(record[1].ID, updateRecord)
	application.GetLogger(ctx).Infof("update record, id: %+v", record[1].ID)

	//// 注册创建
	//id, err := tx.Object("objectForAll").RegisterCreate(recordTransactionStruct)
	//if err != nil {
	//	application.GetLogger(ctx).Errorf("create record error: %+v", err)
	//	return
	//}
	//application.GetLogger(ctx).Infof("create record, id: %+v", id)

	err = tx.Commit(ctx)
	if err != nil {
		application.GetLogger(ctx).Errorf("commit transaction error: %+v", err)
		return
	}
}

func batchTransactionV3() {
	tx := application.DataV3.NewTransaction()

	//id1, id2 := createRecord()
	//application.GetLogger(ctx).Infof("delete record, id: %+v, %+v", id1, id2)
	//
	////// 注册批量删除
	//tx.Object("objectForAll").RegisterBatchDelete([]interface{}{id1, id2})

	id3, id4 := createRecord()
	application.GetLogger(ctx).Infof("update record, id: %+v, %+v", id3, id4)

	//// 注册批量更新
	updateRecord := recordTransactionStruct
	updateRecord.Text2 = "batchUpdateTransaction"
	updateRecord.ID = id3
	//updateRecord2 := recordInterface
	//updateRecord2["text2"] = "batchUpdateTransaction"
	//updateRecord2["_id"] = id4
	updateRecord2 := recordTransactionStruct
	updateRecord2.Text2 = "batchUpdateTransaction"
	updateRecord2.ID = id4
	tx.Object("objectForAll").RegisterBatchUpdate([]interface{}{updateRecord, updateRecord2})

	//// 注册批量 db 操作，单个批量操作数据量限制 100 条，多个批量操作的总数据量限制 500 条
	////// 注册批量创建
	//ids, err := tx.Object("objectForAll").RegisterBatchCreate([]interface{}{recordTransaction, recordTransactionStruct})
	//if err != nil {
	//	application.GetLogger(ctx).Errorf("batch create record error: %+v", err)
	//	return
	//}
	//application.GetLogger(ctx).Infof("create record success, ids: %+v", ids)

	// 提交事务
	err := tx.Commit(ctx)
	if err != nil {
		application.GetLogger(ctx).Errorf("commit transaction error: %+v", err)
		return
	}
}

func batchTransactionV1() {
	tx := application.Data.NewTransaction()

	//id1, id2 := createRecord()
	//application.GetLogger(ctx).Infof("delete record, id: %+v, %+v", id1, id2)
	//
	////// 注册批量删除
	//tx.Object("objectForAll").RegisterBatchDelete([]interface{}{id1, id2})

	id3, id4 := createRecord()
	application.GetLogger(ctx).Infof("update record, id: %+v, %+v", id3, id4)

	//// 注册批量更新
	recordTransactionV1 := map[string]interface{}{
		"text":  "v1Transaction",
		"text2": "111",
	}
	updateRecord, _ := deepCopy(recordTransactionV1)
	updateRecord["text"] = "[datav1]batchUpdateTransaction"
	updateRecord["_id"] = id3
	//updateRecord2 := recordInterface
	//updateRecord2["text2"] = "batchUpdateTransaction"
	//updateRecord2["_id"] = id4
	updateRecord2, _ := deepCopy(recordTransactionV1)
	updateRecord2["text"] = "[datav1]batchUpdateTransaction"
	updateRecord2["_id"] = id4
	tx.Object("objectForAll").RegisterBatchUpdate([]interface{}{updateRecord, updateRecord2})

	//// 注册批量 db 操作，单个批量操作数据量限制 100 条，多个批量操作的总数据量限制 500 条
	////// 注册批量创建
	//ids, err := tx.Object("objectForAll").RegisterBatchCreate([]interface{}{recordTransaction, recordTransactionStruct})
	//if err != nil {
	//	application.GetLogger(ctx).Errorf("batch create record error: %+v", err)
	//	return
	//}
	//application.GetLogger(ctx).Infof("create record success, ids: %+v", ids)

	// 提交事务
	err := tx.Commit(ctx)
	if err != nil {
		application.GetLogger(ctx).Errorf("commit transaction error: %+v", err)
		return
	}
}
