package main

import (
	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
)

var (
	recordTransaction = map[string]interface{}{
		"text":         "recordTransaction",
		"text2":        "recordTransaction",
		"bigintType":   "2",
		"dateType":     "2024-08-20",
		"datetimeType": 1724688780000,
		"decimal":      "1",
		"formula":      "4",
		"option":       "option_7f97916560b",
		"richText": map[string]interface{}{
			"raw": "<div style=\"white-space: pre-wrap;\">&lt;p&gt;test&lt;/p&gt;</div>",
		},
		//"richText": "<div style=\"white-space: pre-wrap;\">&lt;p&gt;test&lt;/p&gt;</div>",
		"avatar": map[string]interface{}{
			"image": map[string]interface{}{
				"token": "e09c1f0c43ff4019bd0463fc3bbe8821_c",
				"uri":   "/img/553944/e09c1f0c43ff4019bd0463fc3bbe8821_l.jpg",
			},
			"source": "image",
		},
	}

	recordTransactionStruct = TestCreateObjectV2{
		Text:       "testV3Struct",
		BigintType: "3",
		DateType:   "2024-09-09",
		Datetime:   1724688780000,
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
		Option: &faassdk.OptionV3{ // 这里虽然可以传结构过去，并且也自己改了颜色和 label，但是创建的时候还是只会取 api_name，label 用的是之前设置的 label
			APIName: "option_7f97916560b",
			Color:   "red",
			Label: faassdk.MultilingualV3{
				Zh: "选项1",
			},
		},
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
			ID: "1808488223853568",
			Name: faassdk.MultilingualV3{
				Zh: "1808488223853568",
			},
		},
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
