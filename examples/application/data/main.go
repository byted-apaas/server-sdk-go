package main

import (
	"context"
	"os"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	"github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/application"
)

var (
	ctx = context.Background()
)

func init() {
	// 环境变量可参考：https://cloud-boe.bytedance.net/faas/function/scmrhatr/cluster/config?region=cn-north&cluster=faas-cn-north&page=1&page_size=10
	// 应用：https://lowcode_v2-dev347.aedev.feishuapp.bytedance.net/ae/builder/v3/package_76452c__c/pc?lane_id=develop
	ctx = context.WithValue(ctx, cConstants.CtxKeyEnvBoe, "boe")
	//ctx = context.WithValue(ctx, cConstants.CtxKeyUser, int64(1808486394085627))
	ctx = context.WithValue(ctx, cConstants.CtxKeyLogID, genLogID())
	ctx = context.WithValue(ctx, cConstants.CtxKeyTTEnv, "boe_faas_sdk")
	_ = os.Setenv("ENV", "staging")
	_ = os.Setenv("KTenantName", "lowcode_v2-dev347")
	_ = os.Setenv("KNamespace", "package_76452c__c")
	_ = os.Setenv("KSvcID", "byf_ny52kh7goc")
	_ = os.Setenv("KSvcSecret", "0f99e213c471c567928c0338c89e3487ddd897665ac2a98122254fa834dfcfe2215401e390cef9f99bef7ef45387341d")
	_ = os.Setenv("KClientID", "293edb95e4134ec379f65cebf3130c3f597170384875e592eeed2eb32413fcc2")
	_ = os.Setenv("KClientSecret", "0cbf367f09ce47c6e94e97de5fecd3d42a5b94084bdfda26a33d72bcfc75ac15af1c2ef6c868d39a19e996e581dc92d8")
	//_ = os.Setenv("KOpenApiDomain", "http://oapi-kunlun-staging-boe.byted.org")
	_ = os.Setenv("KOpenApiDomain", "https://ae-openapi.feishu-boe.cn")
	_ = os.Setenv("KInnerAPIDomain", "https://apaas-innerapi-boe.bytedance.net")
	_ = os.Setenv("KFaaSInfraDomain", "http://apaas-faasinfra-staging-boe.bytedance.net")
	//_ = os.Setenv("KLGWPSM", "lark.apigw.apigw_biz")
	//_ = os.Setenv("KLGWCluster", "apaas_openapi")
}

func main() {
	application.GetLogger(ctx).Infof("logid: %s", utils.GetLogIDFromCtx(ctx))

	//delID, updateID := createRecord()
	//deleteRecord(delID)
	//update(updateID)
	////update2(updateID)
	//update3(updateID)
	//
	//find()
	//findWithFilter()
	//findStream()
	//findOne()
	//getCount()
	//findWithFuzzySearch()

	//findWithFilterV1()

	//batchDelete()
	//batchDeleteWithResult()

	//batchCreate()
	//batchUpdate()

	getFields()
	//getField()
	//oql()

	//transactionV3()
	//batchTransactionV3()
}
