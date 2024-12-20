package main

import (
	"time"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
)

var (
	recordInterface = map[string]interface{}{
		"text":         "testv3Interface",
		"bigintType":   "2",
		"dateType":     "2024-08-20",
		"datetimeType": time.Now().UnixMilli(),
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

func createRecord() (string, string) {
	var id1, id2 string
	// creat1
	application.GetLogger(ctx).Infof("=========== create by interface ==============")
	r1, err := application.DataV3.Object("objectForAll").Create(ctx, recordInterface)
	if err != nil {
		application.GetLogger(ctx).Errorf("create record error: %+v", err)
		return "", ""
	}
	id1 = r1.ID
	application.GetLogger(ctx).Infof("create record success, id: %+v", id1)

	// create2
	application.GetLogger(ctx).Infof("=========== create by struct ==============")
	r2, err := application.DataV3.Object("objectForAll").Create(ctx, recordStruct)
	if err != nil {
		application.GetLogger(ctx).Errorf("create record error: %+v", err)
		return "", ""
	}
	id2 = r2.ID
	application.GetLogger(ctx).Infof("create record success, id: %+v", id2)
	return id1, id2
}

func batchCreate() {
	application.GetLogger(ctx).Infof("=========== batch create ==============")
	r, err := application.DataV3.Object("objectForAll").BatchCreate(ctx, []interface{}{recordInterface, recordStruct, recordStruct})
	if err != nil {
		application.GetLogger(ctx).Errorf("create record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("create record success, id: %+v", r)
}
