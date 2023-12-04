package main

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"context"
	"fmt"
	"reflect"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/opensdk"
)

type TestObject struct {
	ID     int64  `json:"_id,omitempty"`
	Text   string `json:"testText"`
	Number int64  `json:"testNumber"`
}

func main() {
	app := opensdk.NewApplication("***", "***").Env(constants.PlatformEnvUAT)
	ctx := cUtils.SetDebugTypeToCtx(context.Background(), cConstants.DebugTypeLocal)

	fmt.Println("Case 1:")
	err := app.Data.Object("testObject").OrderBy("testNumber").Select("testText", "testNumber").FindStream(ctx, reflect.TypeOf(&TestObject{}), func(ctx context.Context, records interface{}) error {
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

	fmt.Println("Case 2:")
	err = app.Data.Object("testObject").Select("testText", "testNumber").FindStream(ctx, reflect.TypeOf(&TestObject{}), nil,
		structs.FindStreamParam{
			IDGetter: func(record interface{}) (id int64, err error) {
				t, ok := record.(TestObject)
				if !ok {
					return 0, fmt.Errorf(fmt.Sprintf("should be TestObject, but %T", record))
				}
				return t.ID, nil
			},
			Handler: func(ctx context.Context, data *structs.FindStreamData) error {
				var rs []*TestObject
				for i := 0; i < reflect.ValueOf(data.Records).Elem().Len(); i++ {
					o, ok := reflect.ValueOf(data.Records).Elem().Index(i).Interface().(TestObject)
					if !ok {
						panic(fmt.Sprintf("should be TestObject, but %T", reflect.ValueOf(data.Records).Elem().Index(i).Interface()))
					}
					rs = append(rs, &o)
				}

				if len(rs) > 0 {
					fmt.Printf("%d: %d~%d\n", len(rs), rs[0].Number, rs[len(rs)-1].Number)
				}
				return nil
			},
		})
	if err != nil {
		panic(err)
	}

	fmt.Println("Case 3:")
	err = app.Data.Object("testObject").OrderBy("testNumber").Offset(3).Limit(302).Select("testText", "testNumber").FindStream(ctx, reflect.TypeOf(&TestObject{}), func(ctx context.Context, records interface{}) error {
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

	fmt.Println("Case 4:")
	err = app.Data.Object("testObject").Limit(302).Select("testText", "testNumber").FindStream(ctx, reflect.TypeOf(&TestObject{}), nil,
		structs.FindStreamParam{
			IDGetter: func(record interface{}) (id int64, err error) {
				t, ok := record.(TestObject)
				if !ok {
					return 0, fmt.Errorf(fmt.Sprintf("should be TestObject, but %T", record))
				}
				return t.ID, nil
			},
			Handler: func(ctx context.Context, data *structs.FindStreamData) error {
				var rs []*TestObject
				for i := 0; i < reflect.ValueOf(data.Records).Elem().Len(); i++ {
					o, ok := reflect.ValueOf(data.Records).Elem().Index(i).Interface().(TestObject)
					if !ok {
						panic(fmt.Sprintf("should be TestObject, but %T", reflect.ValueOf(data.Records).Elem().Index(i).Interface()))
					}
					rs = append(rs, &o)
				}

				if len(rs) > 0 {
					fmt.Printf("%d: %d~%d\n", len(rs), rs[0].Number, rs[len(rs)-1].Number)
				}
				return nil
			},
		})
	if err != nil {
		panic(err)
	}
}
