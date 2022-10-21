// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package user

import (
	"context"
	"os"
	"testing"

	cUtils "github.com/byted-apaas/server-common-go/utils"
)

var (
	ctx  = context.Background()
	user = NewUser(nil)
)

func Init() {
	//ctx = context.WithValue(ctx, cConstants.CtxKeyDebugType, cConstants.DebugTypeLocal)
	_ = os.Setenv("ENV", "xxx")
	_ = os.Setenv("KTenantName", "xxx")
	_ = os.Setenv("KNamespace", "xxx")
	_ = os.Setenv("KSvcID", "xxx")
	_ = os.Setenv("KClientID", "xxx")
	_ = os.Setenv("KClientSecret", "xxx")
	_ = os.Setenv("KOpenApiDomain", "xxx")
	_ = os.Setenv("KFaaSInfraDomain", "xxx")
	//_ = os.Setenv("KFaaSInfraDomain", "xxx")
	//_ = os.Setenv("KFaaSInfraDomain", "xxx")
	_ = os.Setenv("CONSUL_HTTP_HOST", "xxx")
	_ = os.Setenv("RUNTIME_IDC_NAME", "xxx")
}

func TestMain(m *testing.M) {
	Init()
	m.Run()
}

func TestGetLocale(t *testing.T) {
	userID := user.GetUserID(ctx)
	cUtils.PrintLog(userID)

	locale, err := user.GetLocaleByUserID(ctx, userID)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(locale)

	locales, err := user.GetLocaleByUserIDList(ctx, []int64{userID, 1704448044225543})
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog(locales)
}
