package structs

type ApprovalInstanceListOptions struct {
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	PageSize  int64  `json:"pageSize"`
	PageToken string `json:"pageToken"`
}

type ApprovalInstanceList struct {
	ApprovalInstanceIDs []int64 `json:"approvalInstanceIds"`
	PageToken           string  `json:"pageToken"`
	Count               int64   `json:"count"`
	HasMore             bool    `json:"hasMore"`
}

type GetApprovalInstanceOptions struct {
	ApprovalInstanceId int64 `json:"approvalInstanceId"`
	IncludeFormData    bool  `json:"includeFormData"`
	IncludeContent     bool  `json:"includeContent"`
}

type ApprovalInstance struct {
	// 审批实例 ID
	ID int64 `json:"id"`
	// 审批实例名称
	Label map[string]string `json:"label"`
	// 审批发起人Id
	Initiator int64 `json:"initiator"`
	// 发起时间
	InstanceStartTime int64 `json:"instanceStartTime"`
	// 审批流程状态
	Status string `json:"status"`
	// 审批任务列表
	Tasks []*ApprovalTask `json:"tasks"`
	// 评论列表
	Comments []*ApprovalComment `json:"comments"`
}

type ApprovalTask struct {
	// 审批任务ID
	ID int64 `json:"id"`
	// 任务类型，审批、填写任务、抄送
	TaskType string `json:"taskType"`
	// 审批任务状态 in_process agreed rejected canceled failed completed unstarted auto_end terminated
	TaskStatus string `json:"taskStatus"`
	// 任务开始时间
	TaskStartTime int64 `json:"taskStartTime"`
	// 任务结束时间
	TaskEndTime int64 `json:"taskEndTime"`
	// 任务表单数据，默认不传递，除非请求的 include 参数中包含 ApprovalTask_FormData
	FormData string `json:"formData"`
	// 任务类型  all-会签 first-或签 orderly-依次审批
	ApprovalLogic string `json:"approvalLogic"`
	// 任务已办人
	Approvers []int64 `json:"approvers"`
	// 任务指派人
	Assigners []int64 `json:"assigners"`
	// 审批任务链接
	TaskURL string `json:"taskURL"`
}

type ApprovalComment struct {
	// 审批评论ID
	ID int64 `json:"id"`
	// 评论人
	Commenter int64 `json:"commenter"`
	// 评论内容（富文本html），默认不传递，除非请求的 include 参数中包含 ApprovalComment_Content
	Content string `json:"content"`
	// 评论创建时间
	CreateAt int64 `json:"createAt"`
	// 评论更新时间
	UpdateAt int64 `json:"updateAt"`
}

type GetApprovalInstanceListResp struct {
	ApprovalInstanceIDs []int64 `json:"approval_instance_ids"`
	PageToken           string  `json:"page_token"`
	Count               int64   `json:"count"`
	HasMore             bool    `json:"has_more"`
}

type GetApprovalInstanceResp struct {
	ApprovalInstance struct {
		// 审批实例 ID
		ID int64 `json:"id"`
		// 审批实例名称
		Label map[string]string `json:"label"`
		// 审批发起人Id
		Initiator int64 `json:"initiator"`
		// 发起时间
		InstanceStartTime int64 `json:"instance_start_time"`
		// 审批流程状态
		Status string `json:"status"`
		// 审批任务列表
		Tasks []*struct {
			// 审批任务ID
			ID int64 `json:"label"`
			// 任务类型，审批、填写任务、抄送
			TaskType string `json:"task_type"`
			// 审批任务状态 in_process agreed rejected canceled failed completed unstarted auto_end terminated
			TaskStatus string `json:"task_status"`
			// 任务开始时间
			TaskStartTime int64 `json:"task_start_time"`
			// 任务结束时间
			TaskEndTime int64 `json:"task_end_time"`
			// 任务表单数据，默认不传递，除非请求的 include 参数中包含 ApprovalTask_FormData
			FormData string `json:"form_data"`
			// 任务类型  all-会签 first-或签 orderly-依次审批
			ApprovalLogic string `json:"approval_logic"`
			// 任务已办人
			Approvers []int64 `json:"approvers"`
			// 任务指派人
			Assigners []int64 `json:"assigners"`
			// 审批任务链接
			TaskURL string `json:"task_url"`
		} `json:"tasks"`
		// 评论列表
		Comments []*struct {
			// 审批评论ID
			ID int64 `json:"id"`
			// 评论人
			Commenter int64 `json:"commenter"`
			// 评论内容（富文本html），默认不传递，除非请求的 include 参数中包含 ApprovalComment_Content
			Content string `json:"content"`
			// 评论创建时间
			CreateAt int64 `json:"create_at"`
			// 评论更新时间
			UpdateAt int64 `json:"update_at"`
		} `json:"comments"`
	} `json:"approval_instance"`
}

func (r *GetApprovalInstanceResp) ToApprovalInstance() *ApprovalInstance {
	if r == nil {
		return nil
	}

	result := &ApprovalInstance{
		ID:                r.ApprovalInstance.ID,
		Label:             r.ApprovalInstance.Label,
		Initiator:         r.ApprovalInstance.Initiator,
		InstanceStartTime: r.ApprovalInstance.InstanceStartTime,
		Status:            r.ApprovalInstance.Status,
	}

	for _, task := range r.ApprovalInstance.Tasks {
		if task == nil {
			continue
		}
		result.Tasks = append(result.Tasks, &ApprovalTask{
			ID:            task.ID,
			TaskType:      task.TaskType,
			TaskStatus:    task.TaskStatus,
			TaskStartTime: task.TaskStartTime,
			TaskEndTime:   task.TaskEndTime,
			FormData:      task.FormData,
			ApprovalLogic: task.ApprovalLogic,
			Approvers:     task.Approvers,
			Assigners:     task.Assigners,
			TaskURL:       task.TaskURL,
		})
	}

	for _, comment := range r.ApprovalInstance.Comments {
		if comment == nil {
			continue
		}
		result.Comments = append(result.Comments, &ApprovalComment{
			ID:        comment.ID,
			Commenter: comment.Commenter,
			Content:   comment.Content,
			CreateAt:  comment.CreateAt,
			UpdateAt:  comment.UpdateAt,
		})
	}

	return result
}
