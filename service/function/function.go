// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package function

import (
	"context"
)

type IFunction interface {
	Invoke(ctx context.Context, params map[string]interface{}, result interface{}) error
}
