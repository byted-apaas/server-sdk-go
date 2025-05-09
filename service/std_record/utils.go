package std_record

import "github.com/byted-apaas/server-common-go/utils"

func ConvertStdRecord(record interface{}) (newRecord interface{}) {
	if record == nil {
		return nil
	}

	switch record.(type) {
	case Record:
		newRecord = record.(Record).Record
	case *Record:
		newRecord = record.(*Record).Record
	default:
		newRecord = record
	}

	return newRecord
}

func ConvertStdRecords(records interface{}) interface{} {
	if records == nil {
		return nil
	}

	var newRecords []interface{}
	if rs, ok := records.([]interface{}); ok {
		for _, r := range rs {
			newRecords = append(newRecords, ConvertStdRecord(r))
		}
	} else if rs, ok := records.([]Record); ok {
		for i := range rs {
			newRecords = append(newRecords, rs[i].Record)
		}
	} else if rs, ok := records.([]*Record); ok {
		for i := range rs {
			newRecords = append(newRecords, rs[i].Record)
		}
	} else {
		return records
	}

	return newRecords
}

func ConvertStdRecordsFromMap(records map[int64]interface{}) map[int64]interface{} {
	if records == nil || len(records) == 0 {
		return nil
	}

	newRecords := map[int64]interface{}{}
	for key, record := range records {
		newRecords[key] = ConvertStdRecord(record)
	}
	return newRecords
}

func ConvertStdRecordsFromMapV3(records map[string]interface{}) ([]interface{}, error) {
	if records == nil || len(records) == 0 {
		return nil, nil
	}

	var newRecords []interface{}
	for key, record := range records {

		tmp := ConvertStdRecord(record)

		newRecord, err := decodeRecord(tmp)
		if err != nil {
			return nil, err
		}

		newRecord["_id"] = key
		newRecords = append(newRecords, newRecord)
	}
	return newRecords, nil
}

// decodeRecord 将数据解码为map[string]interface{}，record 的类型可能是 map[string]interface{} 或 struct
//nolint: byted_json_accuracyloss_unknowstruct
func decodeRecord(record interface{}) (map[string]interface{}, error) {
	var newRecord map[string]interface{}
	err := utils.Decode(record, &newRecord)
	if err != nil {
		return nil, err
	}
	return newRecord, nil
}
