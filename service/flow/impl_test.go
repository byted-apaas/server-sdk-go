package flow

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/tools"
)

var (
	ctx = context.Background()
	wf  = &Flow{}
)

func init() {
	os.Setenv("KClientID", "d9f2925541ccf9d4b6f1563786989075adc39d40955d67baeb0f213d3bcfa244")
	os.Setenv("KClientSecret", "99c44d6a066d15f26541c10cc2095c1e49bc3cafb7fcb2e4bf1a847fbb81a7eef5bbd4fccd5ff77ca9e510a5a564fe93")
	os.Setenv("KInnerAPIDomain", "https://apaas-innerapi-boe.bytedance.net")
	os.Setenv("KOpenApiDomain", "http://oapi-kunlun-staging-boe.byted.org")
	os.Setenv("KFaaSInfraDomain", "http://apaas-faasinfra-staging-boe.bytedance.net")
	os.Setenv("KTenantName", "xuzhaoning-dev702")
	os.Setenv("KNamespace", "package_55d00d__c")
	ctx = cUtils.SetTTEnvToCtx(ctx, "boe_openapi_workflow")
}

func TestGetUserTaskInfoByInstance(t *testing.T) {
	taskInfoList, err := wf.GetExecutionUserTaskInfo(ctx, 1744480312321054)
	cUtils.PrintLog(taskInfoList, err)
	assert.Empty(t, err)
}

func TestGetExecutionInfo(t *testing.T) {
	executionInfo, err := wf.GetExecutionInfo(ctx, 1746198129275917)
	cUtils.PrintLog(executionInfo, err)
	assert.Empty(t, err)
}

func TestInvoke(t *testing.T) {
	executionInfo, err := wf.Execute(ctx, "require", &structs.ExecuteOptions{Params: map[string]interface{}{
		"variable_i8qembs": []int64{1},
		"variable_t9psayc": []int64{1},
	}})
	cUtils.PrintLog(executionInfo, err)
	GPJSONStr(executionInfo)
	pErr := tools.NewTools(nil).ParseErr(ctx, err)
	cUtils.PrintLog(executionInfo, err, pErr)
	assert.Empty(t, err)
}

func TestTerminate(t *testing.T) {
	err := wf.RevokeExecution(ctx, 1744480312321054, &structs.RevokeOptions{
		Reason: structs.Multilingual{
			En: "reason",
		},
	})
	assert.Empty(t, err)
}

func TestExecute_outParams(t *testing.T) {
	executionInfo, err := wf.Execute(ctx, "outParams", &structs.ExecuteOptions{Params: map[string]interface{}{"variable_vlwhwc9": "loop"}})
	cUtils.PrintLog(executionInfo, err)
	pErr := tools.NewTools(nil).ParseErr(ctx, err)
	cUtils.PrintLog(executionInfo, err, pErr)
	assert.Empty(t, err)
}

func TestTotal(t *testing.T) {
	executionInfo, err := wf.Execute(ctx, "task", nil)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog("invoke: ", executionInfo)
	info, err := wf.GetExecutionInfo(ctx, executionInfo.ExecutionID)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog("info: ", info)
	taskInfo, err := wf.GetExecutionUserTaskInfo(ctx, executionInfo.ExecutionID)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog("taskInfo: ", taskInfo)
	err = wf.RevokeExecution(ctx, executionInfo.ExecutionID, &structs.RevokeOptions{Reason: structs.Multilingual{En: "en", Zh: "中"}})
	if err != nil {
		panic(err)
	}
	info, err = wf.GetExecutionInfo(ctx, executionInfo.ExecutionID)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog("info1: ", info)
	taskInfo, err = wf.GetExecutionUserTaskInfo(ctx, executionInfo.ExecutionID)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog("taskInfo1: ", taskInfo)
}

func GPJSONStr(params interface{}) string {
	marshal, _ := json.Marshal(params)
	str := string(marshal)
	return str
}

func TestGetApprovalInstanceList(t *testing.T) {
	approvalInstanceList, err := wf.GetApprovalInstanceList(ctx, nil)
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog("approvalInstanceList: ", approvalInstanceList)

	approvalInstance, err := wf.GetApprovalInstance(ctx, &structs.GetApprovalInstanceOptions{ApprovalInstanceId: approvalInstanceList.ApprovalInstanceIDs[0]})
	if err != nil {
		panic(err)
	}
	cUtils.PrintLog("approvalInstance: ", approvalInstance)
}
