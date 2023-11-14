package flow

import (
	"context"
	"encoding/json"
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
