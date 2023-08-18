package std_record

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
	switch records.(type) {
	case []Record:
		rs := records.([]Record)
		for i := range rs {
			newRecords = append(newRecords, rs[i].Record)
		}
		return newRecords
	case []*Record:
		rs := records.([]*Record)
		for i := range rs {
			newRecords = append(newRecords, rs[i].Record)
		}
		return newRecords
	default:
		return records
	}
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
