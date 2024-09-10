package main

import "github.com/byted-apaas/server-sdk-go/application"

func deleteRecord(id int64) {
	application.GetLogger(ctx).Infof("=========== delete %d ==============", id)
	err := application.DataV2.Object("objectForAll").Delete(ctx, id)
	if err != nil {
		application.GetLogger(ctx).Errorf("delete record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("delete %d success", id)
}
