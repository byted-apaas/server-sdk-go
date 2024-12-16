package main

import (
	"time"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
)

var (
	updateRecordInterface = map[string]interface{}{
		"text":         "update_testv3Interface",
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

	updateRecordStruct = TestObjectV2{
		Text:       "updateStruct",
		BigintType: "212",
		DateType:   "2024-09-10",
		Datetime:   time.Now().UnixMilli(),
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
				Name:     "update.xlsx",
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

func update(id string) {
	application.GetLogger(ctx).Infof("=========== update %d ==============", id)
	//updateRecord := TestObjectV2{Text: "update", ID: id} // 这里不用这样用，如果要用结构体需要定义更新的结构，不能用完整的结构
	updateRecord := map[string]interface{}{
		"text": "update",
	}
	err := application.DataV3.Object("objectForAll").Update(ctx, id, updateRecord)
	if err != nil {
		application.GetLogger(ctx).Errorf("update record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("update record success, id: %+v", id)
}

func update2(id string) {
	application.GetLogger(ctx).Infof("=========== update2 %d ==============", id)
	err := application.DataV3.Object("objectForAll").Update(ctx, id, updateRecordInterface)
	if err != nil {
		application.GetLogger(ctx).Errorf("update2 record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("update2 record success, id: %+v", id)
}

func update3(id string) {
	application.GetLogger(ctx).Infof("=========== update3 %d ==============", id)

	err := application.DataV3.Object("objectForAll").Update(ctx, id, updateRecordStruct)
	if err != nil {
		application.GetLogger(ctx).Errorf("update2 record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("update2 record success, id: %+v", id)
}

func batchUpdate() {
	application.GetLogger(ctx).Infof("=========== batchUpdate ==============")
	// 先查
	var records []TestObjectV2
	err := application.DataV3.Object("objectForAll").Select("_id").Offset(4).Limit(2).Find(ctx, &records)
	if err != nil {
		application.GetLogger(ctx).Errorf("Find record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("records: %+v", records)

	updateRecords := make(map[string]interface{})
	for i, r := range records {
		if i%2 == 0 {
			updateRecords[r.ID] = updateRecordInterface
		} else {
			updateRecords[r.ID] = updateRecordStruct
		}
	}
	application.GetLogger(ctx).Infof("updateRecords: %+v", updateRecords)

	err = application.DataV3.Object("objectForAll").BatchUpdate(ctx, updateRecords)
	if err != nil {
		application.GetLogger(ctx).Errorf("batchUpdate record error: %+v", err)
		return
	}
}

func batchUpdateError() {
	application.GetLogger(ctx).Infof("=========== batchUpdateError ==============")

	updateRecords := make(map[string]interface{})
	up := make(map[string]interface{})
	up = updateRecordInterface
	up["puhhh"] = "123"
	updateRecords["1810633284089911"] = up
	updateRecords["1810633284099900"] = updateRecordStruct

	err := application.DataV3.Object("objectForAll").BatchUpdate(ctx, updateRecords)
	if err != nil {
		application.GetLogger(ctx).Errorf("batchUpdate record error: %+v", err)
		// 请求参数不合法：[{"success":false,"_id":"1810633284089911","errors":[{"code":"k_ec_000015","message":"请求参数不合法：\"objectForAll\" 中的字段 \"puhhh\" 不存在，或是不支持的数ull}]},{"success":false,"_id":"1810633284099900","errors":[{"code":"k_mt_ec_400001859","message":"其他记录写入失败导致当前记录同时失败","sub_code":null,"fields":null}]}] [k_ec_000015]
		return
	}
}
