package main

import (
	"github.com/byted-apaas/server-sdk-go/application"
)

func createRecord() {
	record := map[string]interface{}{
		"text":         "testv3",
		"bigintType":   "2",
		"dateType":     "2024-08-20",
		"datetimeType": 1724688780000,
		"decimal":      "1",
		"formula":      "2",
	}
	id, err := application.DataV2.Object("objectForAll").Create(ctx, record)
	if err != nil {
		application.GetLogger(ctx).Errorf("create record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("create record success, id: %+v", id)
}
