package main

import (
	"context"
	"fmt"
	"reflect"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/constants"
)

type TestObject struct {
	ID     int64   `json:"_id,omitempty"`
	Text   string  `json:"testText"`
	Number int64   `json:"testNumber"`
}

func main() {
	app := application.NewApplication("c_e817e247740140e5a1b5", "b8168c1a017544498079871f5117f33c").Env(constants.PlatformEnvPRE)
	ctx := cUtils.SetDebugTypeToCtx(context.Background(), cConstants.DebugTypeLocal)

	var records []*TestObject
	app.Data.Object("testObject").FindAll(ctx, &records)
	var ids []int64
	for _, r := range records {
		ids = append(ids, r.ID)
		if len(ids) >= 500 {
			app.Data.Object("testObject").BatchDelete(ctx, ids)
			ids = []int64{}
		}
	}
	app.Data.Object("testObject").BatchDelete(ctx, ids)

	count, err := app.Data.Object("testObject").Count(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("count: %d\n", count)

	err = app.Data.Object("testObject").OrderBy("testNumber").Select("testText", "testNumber").FindStream(ctx, reflect.TypeOf( &TestObject{}), func(ctx context.Context, records interface{}) error {
		var rs []*TestObject
		for i := 0; i < reflect.ValueOf(records).Elem().Len(); i++ {
			o, ok := reflect.ValueOf(records).Elem().Index(i).Interface().(TestObject)
			if !ok {
				panic(fmt.Sprintf("should be TestObject, but %T", reflect.ValueOf(records).Elem().Index(i).Interface()))
			}
			rs = append(rs, &o)
		}

		if len(rs) > 0 {
			fmt.Printf("%d: %d~%d\n", len(rs), rs[0].Number, rs[len(rs)-1].Number)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	err = app.Data.Object("testObject").OrderBy("testNumber").Offset(3).Limit(302).Select("testText", "testNumber").FindStream(ctx, reflect.TypeOf( &TestObject{}), func(ctx context.Context, records interface{}) error {
		var rs []*TestObject
		for i := 0; i < reflect.ValueOf(records).Elem().Len(); i++ {
			o, ok := reflect.ValueOf(records).Elem().Index(i).Interface().(TestObject)
			if !ok {
				panic(fmt.Sprintf("should be TestObject, but %T", reflect.ValueOf(records).Elem().Index(i).Interface()))
			}
			rs = append(rs, &o)
		}

		if len(rs) > 0 {
			fmt.Printf("%d: %d~%d\n", len(rs), rs[0].Number, rs[len(rs)-1].Number)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
