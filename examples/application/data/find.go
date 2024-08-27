package main

import (
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/application"
)

func find() {
	var record1 []TestObject
	err := application.DataV2.Object("objectForAll").
		Offset(0).Limit(10).
		Select(AllFieldAPINames...).
		Find(ctx, &record1)
	if err != nil {
		panic(err)
	}
	application.GetLogger(ctx).Infof("record1: %s", cUtils.ToString(record1))
}

func findOne() {
	var record1 TestObject
	err := application.DataV2.Object("objectForAll").
		Offset(0).Limit(10).
		Select(AllFieldAPINames...).
		FindOne(ctx, &record1)
	if err != nil {
		panic(err)
	}
}

func findStream() {

}
