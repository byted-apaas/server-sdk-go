package std_record

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/byted-apaas/server-sdk-go/common/exceptions"
	"github.com/byted-apaas/server-sdk-go/common/utils"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
)

type IRecord interface {
	GetFieldValue(fieldAPIName string) (value interface{}, err error)    // 获取字段值
	DecodeFieldValue(fieldAPIName string, value interface{}) (err error) // 解析字段值
	DecodeRecordValue(value interface{}) (err error)                     // 解析记录值

	GetFieldValueInt64(fieldAPIName string) (value int64, err error)                       // 字段类型：浮点数（小数个数为0）
	GetFieldValueFloat64(fieldAPIName string) (value float64, err error)                   // 字段类型：浮点数
	GetFieldValueBool(fieldAPIName string) (value bool, err error)                         // 字段类型：布尔
	GetFieldValueString(fieldAPIName string) (value string, err error)                     // 字段类型：字符串
	GetFieldValueEmail(fieldAPIName string) (value string, err error)                      // 字段类型：邮箱
	GetFieldValueDate(fieldAPIName string) (value string, err error)                       // 字段类型：日期
	GetFieldValueDatetime(fieldAPIName string) (value int64, err error)                    // 字段类型：日期时间
	GetFieldValueOptionSingle(fieldAPIName string) (value string, err error)               // 字段类型：单值选项
	GetFieldValueOptionMulti(fieldAPIName string) (value []string, err error)              // 字段类型：多值选项
	GetFieldValueMultilingual(fieldAPIName string) (value faassdk.Multilingual, err error) // 字段类型：多语
	GetFieldValuePhoneNumber(fieldAPIName string) (value *faassdk.PhoneNumber, err error)  // 字段类型：电话号码
	GetFieldValueAvatar(fieldAPIName string) (value *faassdk.Avatar, err error)            // 字段类型：头像
	GetFieldValueAttachment(fieldAPIName string) (value []*faassdk.Attachment, err error)  // 字段类型：文件
	GetFieldValueRichText(fieldAPIName string) (value *faassdk.RichText, err error)        // 字段类型：富文本

	SetFieldValue(fieldAPIName string, value interface{})

	GetUnauthFields() []string // 获取无权限字段
}

type Record struct {
	Record       map[string]interface{} `json:"record"`
	UnauthFields []string               `json:"_unauthFields"`
}

func (r *Record) GetFieldValue(fieldAPIName string) (value interface{}, err error) {
	if r == nil || r.Record == nil {
		return nil, exceptions.ErrRecordIsEmpty
	}

	if value, ok := r.Record[fieldAPIName]; ok {
		return value, nil
	}

	// 当值为空时，需要判断是否是无权限字段
	if value == nil && utils.StrInStrs(r.UnauthFields, fieldAPIName) {
		return nil, exceptions.ErrFieldNoPermission
	}

	return nil, exceptions.ErrFieldNotFound
}

func (r *Record) GetFieldValueInt64(fieldAPIName string) (value int64, err error) {
	if v, err := r.GetFieldValue(fieldAPIName); err != nil {
		return 0, err
	} else if value, ok := v.(int64); ok {
		return value, nil
	}

	return 0, exceptions.ErrFieldTypeNotMatch
}

func (r *Record) GetFieldValueFloat64(fieldAPIName string) (value float64, err error) {
	if v, err := r.GetFieldValue(fieldAPIName); err != nil {
		return 0, err
	} else if value, ok := v.(float64); ok {
		return value, nil
	}

	return 0, exceptions.ErrFieldTypeNotMatch
}

func (r *Record) GetFieldValueBool(fieldAPIName string) (value bool, err error) {
	if v, err := r.GetFieldValue(fieldAPIName); err != nil {
		return false, err
	} else if value, ok := v.(bool); ok {
		return value, nil
	}

	return false, exceptions.ErrFieldTypeNotMatch
}

func (r *Record) GetFieldValueString(fieldAPIName string) (value string, err error) {
	if v, err := r.GetFieldValue(fieldAPIName); err != nil {
		return "", err
	} else if value, ok := v.(string); ok {
		return value, nil
	}

	return "", exceptions.ErrFieldTypeNotMatch
}

func (r *Record) GetFieldValueDate(fieldAPIName string) (value string, err error) {
	return r.GetFieldValueString(fieldAPIName)
}

func (r *Record) GetFieldValueDatetime(fieldAPIName string) (value int64, err error) {
	return r.GetFieldValueInt64(fieldAPIName)
}

func (r *Record) GetFieldValueMultilingual(fieldAPIName string) (value faassdk.Multilingual, err error) {
	err = r.DecodeFieldValue(fieldAPIName, value)
	return
}

func (r *Record) GetFieldValuePhoneNumber(fieldAPIName string) (value *faassdk.PhoneNumber, err error) {
	err = r.DecodeFieldValue(fieldAPIName, value)
	return
}

func (r *Record) GetFieldValueEmail(fieldAPIName string) (value string, err error) {
	return r.GetFieldValueString(fieldAPIName)
}

func (r *Record) GetFieldValueOptionSingle(fieldAPIName string) (value string, err error) {
	return r.GetFieldValueString(fieldAPIName)
}

func (r *Record) GetFieldValueOptionMulti(fieldAPIName string) (value []string, err error) {
	if v, err := r.GetFieldValue(fieldAPIName); err != nil {
		return nil, err
	} else if v == nil {
		return nil, exceptions.ErrFieldIsEmpty
	} else if value, ok := v.([]string); ok {
		return value, nil
	}
	return nil, exceptions.ErrFieldTypeNotMatch
}

func (r *Record) GetFieldValueAvatar(fieldAPIName string) (value *faassdk.Avatar, err error) {
	err = r.DecodeFieldValue(fieldAPIName, value)
	return
}

func (r *Record) GetFieldValueAttachment(fieldAPIName string) (value []*faassdk.Attachment, err error) {
	err = r.DecodeFieldValue(fieldAPIName, value)
	return
}

func (r *Record) GetFieldValueRichText(fieldAPIName string) (value *faassdk.RichText, err error) {
	err = r.DecodeFieldValue(fieldAPIName, value)
	return
}

func (r *Record) DecodeFieldValue(fieldAPIName string, value interface{}) (err error) {
	if v, err := r.GetFieldValue(fieldAPIName); err != nil {
		return err
	} else if v == nil {
		return exceptions.ErrFieldIsEmpty
	} else if err = mapstructure.Decode(v, &value); err != nil {
		fmt.Printf("%+v\n", err)
		return exceptions.ErrFieldTypeNotMatch
	}

	return nil
}

// DecodeRecordValue 将 record 解码成指定结构
func (r *Record) DecodeRecordValue(value interface{}) (err error) {
	if r == nil || r.Record == nil {
		return exceptions.ErrRecordIsEmpty
	}

	return mapstructure.Decode(r.Record, &value)
}

func (r *Record) GetUnauthFields() []string {
	if r != nil {
		return r.UnauthFields
	}
	return nil
}
