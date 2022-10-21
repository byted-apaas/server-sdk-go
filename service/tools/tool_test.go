// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package tools

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cUtils "github.com/byted-apaas/server-common-go/utils"
)

var (
	ctx = context.Background()
)

func Init() {
	ctx = context.WithValue(ctx, cConstants.CtxKeyDebugType, cConstants.DebugTypeLocal)
	_ = os.Setenv("ENV", "development")
	_ = os.Setenv("KTenantName", "xxx")
	_ = os.Setenv("KNamespace", "xxx")
	_ = os.Setenv("KSvcID", "xxx")
	_ = os.Setenv("KClientID", "xxx")
	_ = os.Setenv("KClientSecret", "xxx")
	_ = os.Setenv("KOpenApiDomain", "xxx")
	_ = os.Setenv("KFaaSInfraDomain", "xxx")
	_ = os.Setenv("CONSUL_HTTP_HOST", "xxx")
	_ = os.Setenv("RUNTIME_IDC_NAME", "boe")
	//_ = os.Setenv("KFaaSInfraDomain", "xxx")
	//_ = os.Setenv("KFaaSInfraDomain", "xxx")
}

func TestMain(m *testing.M) {
	Init()
	m.Run()
}

var (
	tools = &Tools{}
)

func Test_GetTenantInfo(t *testing.T) {
	tenantInfo, err := tools.GetTenantInfo(ctx)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(tenantInfo)
}

func Test_Retry(t *testing.T) {
	err := tools.Retry(func() error {
		e := errors.New("err is not nil")
		fmt.Printf("err: %v\n", e)
		return e
	}, NewRetryOption(2, time.Second))
	cUtils.PrintLog(err)
}
