// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package constants

const (
	OpTextType             = "text"
	OpRichTextType         = "richText"
	OpMultilingualType     = "multilingual"
	OpFloatType            = "float"
	OpEncryptNumberType    = "encryptNumber"
	OpEnumType             = "enum"
	OpLookupType           = "lookup"
	OpBooleanType          = "boolean"
	OpDatetimeType         = "datetime"
	OpDateType             = "date"
	OpPhoneNumberType      = "phone"
	OpEmailType            = "email"
	OpReferenceType        = "reference_field"
	OpAutoNumberType       = "auto_number"
	OpAvatarType           = "avatar"
	OpBackLookupType       = "back_lookup"
	OpAttachmentType       = "attachment"
	OpCompositeType        = "compositeType"
	OpExtractCompositeType = "extractCompositeType"
	OpFormulaType          = "formula"
	OpConditionType        = "condition"
	OpEnumMultiType        = "enum_multi"
	OpLookupMultiType      = "lookup_multi"
	OpInheritFieldType     = "inherit"
	OpGIDFieldType         = "gid"
	OpBigintFieldType      = "bigint"
	OpDecimalFieldType     = "decimal"
	OpRollupFieldType      = "rollup"
	OpRegionFieldType      = "region"
)

const (
	TextType             = "text"
	RichTextType         = "richText"
	MultilingualType     = "multilingual"
	FloatType            = "number"
	EncryptNumberType    = ""
	EnumType             = "option"
	LookupType           = "lookup"
	BooleanType          = "boolean"
	DatetimeType         = "dateTime"
	DateType             = "date"
	PhoneNumberType      = "mobileNumber"
	EmailType            = "email"
	ReferenceType        = "referenceField"
	AutoNumberType       = "autoId"
	AvatarType           = "avatarOrLogo"
	BackLookupType       = "backLookup"
	AttachmentType       = "file"
	CompositeType        = "compositeType"
	ExtractCompositeType = "extractSingleRecord"
	FormulaType          = "formula"
	ConditionType        = ""
	EnumMultiType        = ""
	LookupMultiType      = ""
	InheritFieldType     = ""
	GIDFieldType         = ""
	BigintFieldType      = "bigint"
	DecimalFieldType     = "decimal"
	RollupFieldType      = "rollup"
	RegionFieldType      = "region"
)

var (
	OpToProduct = map[string]string{
		OpTextType:             TextType,
		OpRichTextType:         RichTextType,
		OpMultilingualType:     MultilingualType,
		OpFloatType:            FloatType,
		OpEncryptNumberType:    EncryptNumberType,
		OpEnumType:             EnumType,
		OpLookupType:           LookupType,
		OpBooleanType:          BooleanType,
		OpDatetimeType:         DatetimeType,
		OpDateType:             DateType,
		OpPhoneNumberType:      PhoneNumberType,
		OpEmailType:            EmailType,
		OpReferenceType:        ReferenceType,
		OpAutoNumberType:       AutoNumberType,
		OpAvatarType:           AvatarType,
		OpBackLookupType:       BackLookupType,
		OpAttachmentType:       AttachmentType,
		OpCompositeType:        CompositeType,
		OpExtractCompositeType: ExtractCompositeType,
		OpFormulaType:          FormulaType,
		OpConditionType:        ConditionType,
		OpEnumMultiType:        EnumMultiType,
		OpLookupMultiType:      LookupMultiType,
		OpInheritFieldType:     InheritFieldType,
		OpGIDFieldType:         GIDFieldType,
		OpBigintFieldType:      BigintFieldType,
		OpDecimalFieldType:     DecimalFieldType,
		OpRollupFieldType:      RollupFieldType,
		OpRegionFieldType:      RegionFieldType,
	}
)

const (
	OpOptionTypeLocal  = "local"
	OpOptionTypeGlobal = "global"

	OptionTypeCustom = "custom"
	OptionTypeGlobal = "global"
)

const (
	UniqueTypeNone = iota + 1
	UniqueTypeMultilingualCaseSensitive
	UniqueTypeMultilingualCaseInsensitive
	UniqueTypeTextCaseSensitive
	UniqueTypeTextCaseInsensitive
	UniqueTypeEmail
	UniqueTypePhoneNumber
	UniqueTypeNumber
	UniqueTypeAutoNumberCaseSensitive
	UniqueTypeAutoNumberCaseInsensitive
)

var (
	ColorIDToName = map[string]string{
		"N": "grey",
		"B": "blue",
		"W": "cyan",
		"G": "green",
		"Y": "yellow",
		"O": "orange",
		"R": "red",
		"V": "magenta",
		"P": "purple",
		"I": "blueMagenta",
	}
	ColorNameToID = map[string]string{
		"grey":        "N",
		"blue":        "B",
		"cyan":        "W",
		"green":       "G",
		"yellow":      "Y",
		"orange":      "O",
		"red":         "R",
		"magenta":     "V",
		"purple":      "P",
		"blueMagenta": "I",
	}
)
