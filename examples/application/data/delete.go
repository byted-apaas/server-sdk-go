package main

import "github.com/byted-apaas/server-sdk-go/application"

func deleteRecord(id string) {
	application.GetLogger(ctx).Infof("=========== delete %d ==============", id)
	err := application.DataV3.Object("objectForAll").Delete(ctx, id)
	if err != nil {
		application.GetLogger(ctx).Errorf("delete record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("delete %d success", id)
}

func batchDelete() {
	// 先查
	var record []TestObjectV2
	err := application.DataV3.Object("objectForAll").Select("_id").Offset(4).Limit(2).Find(ctx, &record)
	if err != nil {
		application.GetLogger(ctx).Errorf("Find record error: %+v", err)
		return
	}

	var ids []string
	for _, r := range record {
		ids = append(ids, r.ID)
	}

	// 再删
	application.GetLogger(ctx).Infof("=========== batchDelete %+v ==============", ids)
	err = application.DataV3.Object("objectForAll").BatchDelete(ctx, ids)
	if err != nil {
		application.GetLogger(ctx).Errorf("delete record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("delete %d success", ids)
}

func batchDeleteWithResult() {
	// 先查
	var ids []string
	err := application.DataV3.Object("objectForAll").Select("_id").Offset(4).Limit(2).Find(ctx, &ids)
	if err != nil {
		application.GetLogger(ctx).Errorf("Find record error: %+v", err)
		return
	}

	// 再删
	var result interface{}
	application.GetLogger(ctx).Infof("=========== batchDelete %d ==============", ids)
	err = application.DataV3.Object("objectForAll").BatchDelete(ctx, ids, result)
	if err != nil {
		application.GetLogger(ctx).Errorf("delete record error: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("delete %d success", ids)
}
