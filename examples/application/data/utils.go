package main

import (
	"github.com/byted-apaas/server-common-go/utils"
)

//nolint: byted_json_accuracyloss_unknowstruct
func deepCopy(src map[string]interface{}) (map[string]interface{}, error) {
	var dst map[string]interface{}

	err := utils.Decode(src, &dst)
	if err != nil {
		return nil, err
	}

	return dst, nil
}
