package main

import (
	"github.com/byted-apaas/server-sdk-go/application/operator"
	"github.com/mitchellh/mapstructure"

	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/service/std_record"
)

func find() {
	var record1 []TestObjectV2
	err := application.DataV2.Object("objectForAll").
		Offset(0).Limit(10).
		//Select(AllFieldAPINames...).
		Select("phone", "option", "email").
		Find(ctx, &record1)
	if err != nil {
		application.GetLogger(ctx).Errorf("err: %v", err)
	}
	application.GetLogger(ctx).Infof("record1: %s", cUtils.ToString(record1))
}

func findWithFilter() {
	//cond := operator.And(
	//	operator.Gt("bigintType", 0),
	//	operator.NotEmpty("text"),
	//	operator.NotEmpty("email"))

	cond := operator.Or(
		operator.Eq("_id", 1808488296249436),
		operator.Eq("_id", 1809101723792428),
	)

	var record1 []std_record.Record
	err := application.DataV2.Object("objectForAll").
		Offset(0).Limit(10).
		//Select(AllFieldAPINames...).
		Select("phone", "option", "email").
		Where(cond).
		Find(ctx, &record1)
	if err != nil {
		application.GetLogger(ctx).Errorf("err: %v", err)
	}
	application.GetLogger(ctx).Infof("total: %v", len(record1))
	application.GetLogger(ctx).Infof("record1: %s", cUtils.ToString(record1))
}

func findOne() {
	var record1 interface{}
	err := application.DataV2.Object("objectForAll").
		Offset(0).Limit(10).
		Select(AllFieldAPINames...).
		FindOne(ctx, &record1)
	if err != nil {
		panic(err)
	}
	application.GetLogger(ctx).Infof("record1: %s", cUtils.ToString(record1))

	r, ok := record1.(map[string]interface{})
	if ok {
		application.GetLogger(ctx).Infof("convert: %+v", cUtils.ToString(r))
	}

	// 将 interface 转为 struct
	var testObject = TestObject{}
	mapstructure.Decode(r, &testObject)
	application.GetLogger(ctx).Infof("testObject: %+v", testObject)
}

func findStream() {

}

// Record 类型转换
func findOneRecordStruct1() {
	var record1 std_record.Record
	err := application.DataV2.Object("objectForAll").
		Offset(0).Limit(10).
		Select(AllFieldAPINames...).
		FindOne(ctx, &record1)
	if err != nil {
		panic(err)
	}
	application.GetLogger(ctx).Infof("record1: %s", cUtils.ToString(record1))

	// 将 interface 转为 struct
	var testObject = TestObject{}
	err = record1.DecodeRecordValue(&testObject)
	if err != nil {
		application.GetLogger(ctx).Infof("err: %+v", err)
	}
	application.GetLogger(ctx).Infof("testObject: %+v", testObject)
}

func findOneRecordStruct() {
	var record1 TestObject
	var unauthFiled [][]string
	err := application.DataV2.Object("objectForAll").
		Offset(0).Limit(10).
		Select(AllFieldAPINames...).
		FindOne(ctx, &record1, &unauthFiled)
	if err != nil {
		panic(err)
	}
	application.GetLogger(ctx).Infof("record1: %s", cUtils.ToString(record1))
	application.GetLogger(ctx).Infof("unauthFiled: %s", cUtils.ToString(unauthFiled))

	//// 将 interface 转为 struct
	//var testObject = TestObject{}
	//err = record1.DecodeRecordValue(&testObject)
	//if err != nil {
	//	application.GetLogger(ctx).Infof("err: %+v", err)
	//}
	//application.GetLogger(ctx).Infof("testObject: %+v", testObject)
}
