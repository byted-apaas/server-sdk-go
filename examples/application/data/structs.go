package main

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
	"github.com/byted-apaas/server-sdk-go/service/metadata/object/fields"
)

type TestTransactionCreateObjectV2 struct {
	// 创建得传 int64 不然会报错。。。或者不传也可以
	ID             string                       `json:"_id"`
	Text           string                       `json:"text"`         // 文本
	Text2          string                       `json:"text2"`        // 文本2
	BigintType     string                       `json:"bigintType"`   // 整数
	Number         float64                      `json:"number"`       // 浮点数
	DateType       string                       `json:"dateType"`     // 日期
	Datetime       int64                        `json:"datetimeType"` // 日期时间
	Phone          *faassdk.PhoneNumberV3       `json:"phone"`        // 手机号
	Email          string                       `json:"email"`        // 邮箱
	Option         string                       `json:"option"`       // 选项  transaction 只能传 string，object 能传 *structs.OptionV3 | string， 传 string 需要传 api_name
	Option2        []string                     `json:"option2"`      // 选项  多选
	Boolean        bool                         `json:"booleanType"`  // bool 值
	Avatar         *faassdk.AvatarV3            `json:"avatar"`       // 头像
	Multilingual   *structs.Multilingual        `json:"multilingual"` // 多语
	RichText       *faassdk.RichTextV3          `json:"richText"`     // 支持  输入支持：*faassdk.RichTextV3 ｜ string，结果返回的是 string
	Attachment     []*faassdk.AttachmentModelV3 `json:"attachment"`
	Autoid         string                       `json:"autoid,omitempty"`
	Formula        string                       `json:"formula,omitempty"`
	Lookup         interface{}                  `json:"lookup"`                   // 可以用 interface 也可以用 *fields.Lookup
	ReferenceField *faassdk.LookupV3            `json:"referenceField,omitempty"` // 这个类型由引用的字段类型决定，这里我引用了一个“创建人”
	Decimal        string                       `json:"decimal"`
	Region         *faassdk.RegionV3            `json:"region"`
	//Rollup         fields.Rollup         `json:"rollup"`
	//ExtractSubObject *SubObject            `json:"extractSubObject,omitempty"`
}

type TestCreateObjectV2 struct {
	// 创建得传 int64 不然会报错。。。或者不传也可以
	ID             int64                        `json:"_id"`
	Text           string                       `json:"text"`         // 文本
	Text2          string                       `json:"text2"`        // 文本2
	BigintType     string                       `json:"bigintType"`   // 整数
	Number         float64                      `json:"number"`       // 浮点数
	DateType       string                       `json:"dateType"`     // 日期
	Datetime       int64                        `json:"datetimeType"` // 日期时间
	Phone          *faassdk.PhoneNumberV3       `json:"phone"`        // 手机号
	Email          string                       `json:"email"`        // 邮箱
	Option         *faassdk.OptionV3            `json:"option"`       // 选项  *structs.OptionV3 | string， 传 string 需要传 api_name
	Boolean        bool                         `json:"booleanType"`  // bool 值
	Avatar         *faassdk.AvatarV3            `json:"avatar"`       // 头像
	Multilingual   *structs.Multilingual        `json:"multilingual"` // 多语
	RichText       *faassdk.RichTextV3          `json:"richText"`     // 支持  输入支持：*faassdk.RichTextV3 ｜ string，结果返回的是 string
	Attachment     []*faassdk.AttachmentModelV3 `json:"attachment"`
	Autoid         string                       `json:"autoid,omitempty"`
	Formula        string                       `json:"formula,omitempty"`
	Lookup         interface{}                  `json:"lookup"`                   // 可以用 interface 也可以用 *fields.Lookup
	ReferenceField *faassdk.LookupV3            `json:"referenceField,omitempty"` // 这个类型由引用的字段类型决定，这里我引用了一个“创建人”
	Decimal        string                       `json:"decimal"`
	Region         *faassdk.RegionV3            `json:"region"`
	//Rollup         fields.Rollup         `json:"rollup"`
	//ExtractSubObject *SubObject            `json:"extractSubObject,omitempty"`
}

type TestObjectV2 struct {
	ID             string                       `json:"_id"`
	Text           string                       `json:"text"`         // 文本
	Text2          string                       `json:"text2"`        // 文本2
	BigintType     string                       `json:"bigintType"`   // 整数
	Number         float64                      `json:"number"`       // 浮点数
	DateType       string                       `json:"dateType"`     // 日期
	Datetime       int64                        `json:"datetimeType"` // 日期时间
	Phone          *faassdk.PhoneNumberV3       `json:"phone"`        // 手机号
	Email          string                       `json:"email"`        // 邮箱
	Option         *faassdk.OptionV3            `json:"option"`       // 选项  *structs.OptionV3 | string， 传 string 需要传 api_name
	Boolean        bool                         `json:"booleanType"`  // bool 值
	Avatar         *faassdk.AvatarV3            `json:"avatar"`       // 头像
	Multilingual   *structs.Multilingual        `json:"multilingual"` // 多语
	RichText       interface{}                  `json:"richText"`     // 支持  输入支持：*faassdk.RichTextV3 ｜ string，结果返回的是 string
	Attachment     []*faassdk.AttachmentModelV3 `json:"attachment"`
	Autoid         string                       `json:"autoid,omitempty"`
	Formula        string                       `json:"formula,omitempty"`
	Lookup         interface{}                  `json:"lookup"`                   // 可以用 interface 也可以用 *fields.Lookup
	ReferenceField *faassdk.LookupV3            `json:"referenceField,omitempty"` // 这个类型由引用的字段类型决定，这里我引用了一个“创建人”
	Decimal        string                       `json:"decimal"`
	Region         *faassdk.RegionV3            `json:"region"`
	//Rollup         fields.Rollup         `json:"rollup"`
	//ExtractSubObject *SubObject            `json:"extractSubObject,omitempty"`
}

type TestObject struct {
	ID             int64                 `json:"_id"`
	Text           string                `json:"text"`         // 文本
	BigintType     string                `json:"bigintType"`   // 整数
	Number         float64               `json:"number"`       // 浮点数
	DateType       string                `json:"dateType"`     // 日期
	Datetime       int64                 `json:"datetimeType"` // 日期时间
	Phone          *faassdk.PhoneNumber  `json:"phone"`        // 手机号
	Email          string                `json:"email"`        // 邮箱
	Option         string                `json:"option"`       // 选项
	Boolean        bool                  `json:"booleanType"`  // bool 值
	Avatar         *faassdk.Avatar       `json:"avatar"`       // 头像
	Multilingual   *faassdk.Multilingual `json:"multilingual"` // 多语
	RichText       *faassdk.RichText     `json:"richText"`
	Attachment     []*structs.Attachment `json:"attachment"`
	Autoid         string                `json:"autoid,omitempty"`
	Formula        string                `json:"formula,omitempty"`
	Lookup         interface{}           `json:"lookup"`                   // 可以用 interface 也可以用 *fields.Lookup
	ReferenceField *fields.Lookup        `json:"referenceField,omitempty"` // 这个类型由引用的字段类型决定，这里我引用了一个“创建人”
	Decimal        string                `json:"decimal"`
	Region         *faassdk.Region       `json:"region"`
	Rollup         fields.Rollup         `json:"rollup"`
	//ExtractSubObject *SubObject            `json:"extractSubObject,omitempty"`
}

type SubObject struct {
	Text         string                `json:"text"`
	Bigint       string                `json:"bigint"`
	Number       float64               `json:"number"`
	Decimal      string                `json:"decimal"`
	Date         string                `json:"date"`
	Datetime     int64                 `json:"datetime"`
	Phone        *faassdk.PhoneNumber  `json:"phone"`
	Email        string                `json:"email"`
	Option       string                `json:"option"`
	Boolean      bool                  `json:"boolean"`
	Multilingual *faassdk.Multilingual `json:"multilingual"`
	Formula      string                `json:"formula,omitempty"`
	Attachment   []*structs.Attachment `json:"attachment"`
	Lookup       interface{}           `json:"lookup"`
	Region       *faassdk.Region       `json:"region"`
}

var (
	AllFieldAPINames = []string{
		"_id",
		"text",
		"bigintType",
		"number",
		"dateType",
		"datetimeType",
		"phone",
		"email",
		"option",
		"booleanType",
		"avatar",
		"multilingual",
		"richText",
		"attachment",
		"autoid",
		"formula",
		"lookup",
		"referenceField",
		"decimal",
		"region",
	}
)
