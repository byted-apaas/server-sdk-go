package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/byted-apaas/server-sdk-go/application/operator"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/mitchellh/mapstructure"

	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/service/std_record"
)

func find() {
	application.GetLogger(ctx).Infof("=============== find ==================")
	var record1 []TestObjectV2
	err := application.DataV3.Object("objectForAll").
		Offset(0).Limit(10).
		Select(AllFieldAPINames...).
		//Select("_id", "phone", "option", "email").
		Find(ctx, &record1)
	if err != nil {
		application.GetLogger(ctx).Errorf("err: %v", err)
		return
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
	err := application.DataV3.Object("objectForAll").
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
	err := application.DataV3.Object("objectForAll").
		Where(operator.Eq("_id", 1808488296249436)).
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
	var testObject = TestObjectV2{}
	mapstructure.Decode(r, &testObject)
	application.GetLogger(ctx).Infof("testObject: %+v", testObject)
}

func findStream() {
	application.GetLogger(ctx).Infof(" ============ findstream =============")
	err := application.DataV3.Object("objectForAll").
		Select(AllFieldAPINames...).
		FindStream(ctx, reflect.TypeOf(&TestObjectV2{}), nil, structs.FindStreamParam{
			Handler: func(ctx context.Context, data *structs.FindStreamData) (err error) {
				var rs []*TestObjectV2
				for i := 0; i < reflect.ValueOf(data.Records).Elem().Len(); i++ {
					o, ok := reflect.ValueOf(data.Records).Elem().Index(i).Interface().(TestObjectV2)
					if !ok {
						panic(fmt.Sprintf("should be AllFieldObject, but %T", reflect.ValueOf(data.Records).Elem().Index(i).Interface()))
					}
					rs = append(rs, &o)
				}

				application.GetLogger(ctx).Infof("count: %d", len(rs))
				application.GetLogger(ctx).Infof("record1: %s", cUtils.ToString(rs[0]))
				return nil
			},
			PageLimit: 20,
		})
	if err != nil {
		application.GetLogger(ctx).Errorf("findstream failed: %+v", err)
		return
	}

}

func findWithFuzzySearch() {
	application.GetLogger(ctx).Infof(" ============ findWithFuzzySearch =============")
	var record1 []TestObjectV2
	err := application.DataV3.Object("objectForAll").
		Select(AllFieldAPINames...).
		FuzzySearch("update", []string{"text", "phone"}).
		Find(ctx, &record1)
	if err != nil {
		application.GetLogger(ctx).Errorf("[findWithFuzzySearch] err: %v", err)
		return
	}
	application.GetLogger(ctx).Infof("count: %d, record1: %s", len(record1), cUtils.ToString(record1))
}

func getCount() {
	application.GetLogger(ctx).Infof(" ============ count =============")
	count, err := application.DataV3.Object("objectForAll").Select(AllFieldAPINames...).Count(ctx)
	if err != nil {
		application.GetLogger(ctx).Infof("err: %+v", err)
		return
	}
	application.GetLogger(ctx).Infof("count: %d", count)
}

// ============================================ 以下为 datav1 的请求 =======================================
// Record 类型转换
func findOneRecordStruct1() {
	var record1 std_record.Record
	err := application.Data.Object("objectForAll").
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
	err := application.Data.Object("objectForAll").
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

func findWithFilterV1() {
	//cond := operator.And(
	//	operator.Gt("bigintType", 0),
	//	operator.NotEmpty("text"),
	//	operator.NotEmpty("email"))

	cond := operator.Or(
		operator.Eq("_id", 1808488296249436),
		operator.Eq("_id", 1809101723792428),
	)

	var record1 []std_record.Record
	err := application.Data.Object("objectForAll").
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
