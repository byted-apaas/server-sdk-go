package main

import (
	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
)

var (
	updateRecordInterface = map[string]interface{}{
		"text":         "testv3Interface",
		"bigintType":   "2",
		"dateType":     "2024-08-20",
		"datetimeType": 1724688780000,
		"decimal":      "1",
		"formula":      "2",
		"option":       "option_7f97916560b",
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
	}

	updateRecordStruct = TestObjectV2{
		Text:       "updateStruct",
		BigintType: "212",
		DateType:   "2024-09-10",
		Datetime:   1724688780000,
		Number:     12.3,
		Decimal:    "12",
		Formula:    "4",
		Boolean:    true,
		Email:      "lf@bytedance.com",
		Phone: &faassdk.PhoneNumberV3{
			RegionCode:  "CN",
			DialingCode: "+86",
			Number:      "18902845222",
		},
		Multilingual: &structs.Multilingual{
			Zh: "多语",
			En: "multilingual",
		},
		RichText: &faassdk.RichTextV3{
			Raw: "<div style=\"white-space: pre-wrap;\">&lt;p&gt;testUpdate&lt;/p&gt;</div>",
		},
		Option: &structs.OptionV3{ // 这里虽然可以传结构过去，并且也自己改了颜色和 label，但是创建的时候还是只会取 api_name，label 用的是之前设置的 label
			APIName: "option_7f97916560b",
			Color:   "red",
			Label: structs.Multilingual{
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
				Name:     "update.xlsx",
				Size:     "12237",
				Token:    "BIZ_06aa3803148f4d7791ab3699e84a5ac3",
				URI:      "/ae/api/v1/assets/attachment/download?token=BIZ_06aa3803148f4d7791ab3699e84a5ac3",
			},
		},
	}
)

func update(id string) {
	application.GetLogger(ctx).Infof("=========== update %d ==============", id)
	//updateRecord := TestObjectV2{Text: "update", ID: id} // 这里不用这样用，如果要用结构体需要定义更新的结构，不能用完整的结构
	updateRecord := map[string]interface{}{
		"text": "update",
	}
	err := application.DataV2.Object("objectForAll").Update(ctx, id, updateRecord)
	if err != nil {
		application.GetLogger(ctx).Errorf("update record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("update record success, id: %+v", id)
}

func update2(id string) {
	application.GetLogger(ctx).Infof("=========== update2 %d ==============", id)
	err := application.DataV2.Object("objectForAll").Update(ctx, id, updateRecordInterface)
	if err != nil {
		application.GetLogger(ctx).Errorf("update2 record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("update2 record success, id: %+v", id)
}

func update3(id string) {
	application.GetLogger(ctx).Infof("=========== update3 %d ==============", id)

	err := application.DataV2.Object("objectForAll").Update(ctx, id, updateRecordStruct)
	if err != nil {
		application.GetLogger(ctx).Errorf("update2 record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("update2 record success, id: %+v", id)
}
