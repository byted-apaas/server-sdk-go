package main

import (
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
	"github.com/byted-apaas/server-sdk-go/service/metadata/object/fields"
)

type TestObjectV2 struct {
	ID             int64                   `json:"_id"`
	Text           string                  `json:"text"`         // 文本
	BigintType     string                  `json:"bigintType"`   // 整数
	Number         float64                 `json:"number"`       // 浮点数
	DateType       string                  `json:"dateType"`     // 日期
	Datetime       int64                   `json:"datetimeType"` // 日期时间
	Phone          *faassdk.PhoneNumberV2  `json:"phone"`        // 手机号
	Email          string                  `json:"email"`        // 邮箱
	Option         *structs.OptionV2       `json:"option"`       // 选项
	Boolean        bool                    `json:"booleanType"`  // bool 值
	Avatar         *faassdk.AvatarV2       `json:"avatar"`       // 头像
	Multilingual   *faassdk.Multilingual   `json:"multilingual"` // 多语
	RichText       *faassdk.RichTextV2     `json:"richText"`
	Attachment     []*faassdk.AttachmentV2 `json:"attachment"`
	Autoid         string                  `json:"autoid,omitempty"`
	Formula        string                  `json:"formula,omitempty"`
	Lookup         interface{}             `json:"lookup"`                   // 可以用 interface 也可以用 *fields.Lookup
	ReferenceField *structs.LookupV2       `json:"referenceField,omitempty"` // 这个类型由引用的字段类型决定，这里我引用了一个“创建人”
	Decimal        string                  `json:"decimal"`
	Region         *faassdk.RegionV2       `json:"region"`
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
	//Rollup         fields.Rollup         `json:"rollup"`
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
