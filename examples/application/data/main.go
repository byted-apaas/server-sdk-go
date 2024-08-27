package main

import (
	"context"
	"os"

	cConstants "github.com/byted-apaas/server-common-go/constants"
)

var (
	ctx = context.Background()
)

func init() {
	// 环境变量可参考：https://cloud-boe.bytedance.net/faas/function/scmrhatr/cluster/config?region=cn-north&cluster=faas-cn-north&page=1&page_size=10
	ctx = context.WithValue(ctx, cConstants.CtxKeyEnvBoe, "boe")
	_ = os.Setenv("ENV", "staging")
	_ = os.Setenv("KTenantName", "lowcode_v2-dev347")
	_ = os.Setenv("KNamespace", "package_76452c__c")
	_ = os.Setenv("KSvcID", "byf_ny52kh7goc")
	_ = os.Setenv("KClientID", "293edb95e4134ec379f65cebf3130c3f597170384875e592eeed2eb32413fcc2")
	_ = os.Setenv("KClientSecret", "0cbf367f09ce47c6e94e97de5fecd3d42a5b94084bdfda26a33d72bcfc75ac15af1c2ef6c868d39a19e996e581dc92d8")
	_ = os.Setenv("KOpenApiDomain", "http://oapi-kunlun-staging-boe.byted.org")
	_ = os.Setenv("KOpenApiDomain", "https://apaas-innerapi-boe.bytedance.net")
	_ = os.Setenv("KFaaSInfraDomain", "http://apaas-faasinfra-staging-boe.bytedance.net")
}

func main() {
	find()
	findOne()
}
