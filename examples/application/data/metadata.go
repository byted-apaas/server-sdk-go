package main

import (
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/application"
)

func getFields() {
	var fields interface{}
	// extractSubObject / objectForAll
	err := application.MetadataV3.Object("extractSubObject").GetFields(ctx, &fields)
	if err != nil {
		application.GetLogger(ctx).Errorf("err: %v", err)
		return
	}
	application.GetLogger(ctx).Infof("fields: %s", cUtils.ToString(fields))
}

func getField() {
	var field interface{}
	err := application.MetadataV3.Object("objectForAll").GetField(ctx, "option2", &field)
	if err != nil {
		application.GetLogger(ctx).Errorf("err: %v", err)
		return
	}
	application.GetLogger(ctx).Infof("fields: %s", cUtils.ToString(field))
}
