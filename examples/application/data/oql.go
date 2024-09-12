package main

import "github.com/byted-apaas/server-sdk-go/application"

func oql() {
	var record1 []TestObject
	_ = application.DataV3.Oql("select _id from objectForAll limit 3").Execute(ctx, &record1)
	application.GetLogger(ctx).Infof("res: %+v", record1)
}
