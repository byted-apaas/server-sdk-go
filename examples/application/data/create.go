package main

import (
	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
)

var (
	recordInterface = map[string]interface{}{
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

	recordStruct = TestCreateObjectV2{
		Text:       "testV3Struct",
		BigintType: "3",
		DateType:   "2024-09-09",
		Datetime:   1724688780000,
		Decimal:    "2",
		Formula:    "3",
		Boolean:    false,
		RichText: &faassdk.RichTextV3{
			Raw: "<div style=\"white-space: pre-wrap;\">&lt;p&gt;test111&lt;/p&gt;</div>",
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
				Name:     "待升级 app.xlsx",
				Size:     "12237",
				Token:    "BIZ_06aa3803148f4d7791ab3699e84a5ac3",
				URI:      "/ae/api/v1/assets/attachment/download?token=BIZ_06aa3803148f4d7791ab3699e84a5ac3",
			},
		},
	}
)

func createRecord() (int64, int64) {
	var id1, id2 int64
	// creat1
	application.GetLogger(ctx).Infof("=========== 1 ==============")
	r1, err := application.DataV2.Object("objectForAll").Create(ctx, recordInterface)
	if err != nil {
		application.GetLogger(ctx).Errorf("create record error: %+v", err)
		return 0, 0
	}
	id1 = r1.ID
	application.GetLogger(ctx).Infof("create record success, id: %+v", id1)

	// create2
	application.GetLogger(ctx).Infof("=========== 2 ==============")
	r2, err := application.DataV2.Object("objectForAll").Create(ctx, recordStruct)
	if err != nil {
		application.GetLogger(ctx).Errorf("create record error: %+v", err)
		return 0, 0
	}
	id2 = r2.ID
	application.GetLogger(ctx).Infof("create record success, id: %+v", id2)
	return id1, id2
}
