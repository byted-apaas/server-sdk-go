// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package parser

import (
	"github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/metadata/object/fields"
)

func parseBase(f *structs.Field) fields.FieldBase {
	return fields.FieldBase{
		Type:    constants.OpToProduct[f.Type.Name],
		APIName: f.APIName,
		Label:   f.Label,
	}
}

func transUnique(f *structs.Field) bool {
	return f.UniqueType > 1
}

func transCaseSensitive(f *structs.Field) bool {
	switch f.UniqueType {
	case constants.UniqueTypeMultilingualCaseSensitive,
		constants.UniqueTypeTextCaseSensitive,
		constants.UniqueTypeAutoNumberCaseSensitive:
		return true
	default:
		return false
	}
}

func parseText(f *structs.Field) (*fields.Text, error) {
	s := &structs.NestedTextSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.Text{
		FieldBase:       parseBase(f),
		Required:        f.Required,
		Unique:          transUnique(f),
		CaseSensitive:   transCaseSensitive(f),
		Multiline:       s.Multiline,
		MaxLength:       s.MaxLength,
		ValidationRegex: "", // TODO 未返回
		ErrorMsg:        "",
	}, nil
}

func parseRichText(f *structs.Field) (*fields.RichText, error) {
	s := &structs.NestedRichTextSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.RichText{
		FieldBase: parseBase(f),
		Required:  f.Required,
		MaxLength: s.MaxLength,
	}, nil
}

func parseMultilingual(f *structs.Field) (*fields.Multilingual, error) {
	s := &structs.NestedMultilingualSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.Multilingual{
		FieldBase:     parseBase(f),
		Required:      f.Required,
		Unique:        transUnique(f),
		CaseSensitive: transCaseSensitive(f),
		Multiline:     s.Multiline,
		MaxLength:     s.MaxLength,
	}, nil
}

func parseFloat(f *structs.Field) (*fields.Number, error) {
	s := &structs.NestedFloatSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.Number{
		FieldBase:           parseBase(f),
		Required:            f.Required,
		Unique:              transUnique(f),
		DisplayAsPercentage: s.DisplayAsPercentage,
		DecimalPlacesNumber: s.DecimalPlaces,
	}, nil
}

func transOptionSource(s *structs.NestedEnumSetting) (res string) {
	switch s.OptionType {
	case constants.OpOptionTypeLocal:
		res = constants.OptionTypeCustom
	case constants.OpOptionTypeGlobal:
		res = constants.OptionTypeGlobal
	default:
		res = constants.OptionTypeCustom
	}
	return
}

func transColor(s structs.Option) string {
	return constants.ColorIDToName[s.ColorID]
}

func transOptionList(s *structs.NestedEnumSetting) (res []*fields.OptionItem) {
	res = make([]*fields.OptionItem, len(*(s.Options)))
	for i, option := range *(s.Options) {
		res[i] = &fields.OptionItem{
			Label:       option.Label,
			APIName:     option.APIName,
			Description: option.Description,
			Color:       transColor(option),
			Active:      option.Active > 0,
		}
	}
	return
}

func parseEnum(f *structs.Field) (*fields.Option, error) {
	s := &structs.NestedEnumSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.Option{
		FieldBase:           parseBase(f),
		Required:            f.Required,
		Multiple:            s.IsArray,
		OptionSource:        transOptionSource(s),
		GlobalOptionAPIName: s.RelatedToGlobalOption.APIName,
		OptionList:          transOptionList(s),
	}, nil

}

func transSort(s *structs.NestedLookupSetting) []fields.SortCondition {
	sorts := make([]fields.SortCondition, len(s.DisplayOrder))
	for i, s := range s.DisplayOrder {
		sorts[i] = fields.SortCondition{
			FieldAPIName: s.FieldAPIName,
			Sort:         s.Direction,
		}
	}
	return sorts
}

func parseLookup(f *structs.Field) (*fields.Lookup, error) {
	s := &structs.NestedLookupSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.Lookup{
		FieldBase:      parseBase(f),
		Required:       f.Required,
		Multiple:       s.IsArray,
		ObjectAPIName:  s.ReferencedObjectAPIName,
		Hierarchy:      s.IsHierarchy,
		DisplayStyle:   s.DisplayStyle,
		SortConditions: transSort(s),
		Filter:         nil, // TODO 未返回
	}, nil
}

func parseBoolean(f *structs.Field) (*fields.Boolean, error) {
	s := &structs.NestedBooleanSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.Boolean{
		FieldBase:            parseBase(f),
		DescriptionWhenTrue:  s.Description4True,
		DescriptionWhenFalse: s.Description4False,
		DefaultValue:         s.DefaultValue,
	}, nil
}

func parseDateTime(f *structs.Field) (*fields.DateTime, error) {
	return &fields.DateTime{
		FieldBase: parseBase(f),
		Required:  f.Required,
	}, nil
}

func parseDate(f *structs.Field) (*fields.Date, error) {
	return &fields.Date{
		FieldBase: parseBase(f),
		Required:  f.Required,
	}, nil
}

func parsePhoneNumber(f *structs.Field) (*fields.MobileNumber, error) {
	return &fields.MobileNumber{
		FieldBase: parseBase(f),
		Required:  f.Required,
		Unique:    transUnique(f),
	}, nil
}

func parseEmail(f *structs.Field) (*fields.Email, error) {
	return &fields.Email{
		FieldBase: parseBase(f),
		Required:  f.Required,
		Unique:    transUnique(f),
	}, nil
}

func parseReference(f *structs.Field) (*fields.ReferenceField, error) {
	s := &structs.NestedReferenceFieldSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.ReferenceField{
		FieldBase:         parseBase(f),
		GuideFieldAPIName: s.TargetReferencedObjectAPIName,
		FieldAPIName:      s.TargetReferenceFieldAPIName,
	}, nil
}

func parseAutoNumber(f *structs.Field) (*fields.AutoID, error) {
	s := &structs.NestedAutoNumberSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.AutoID{
		FieldBase:      parseBase(f),
		GenerateMethod: s.GenerationMethod,
		DigitsNumber:   s.Digits,
		Prefix:         s.Prefix,
		Suffix:         s.Suffix,
	}, nil
}

func parseAvatar(f *structs.Field) (*fields.AvatarOrLogo, error) {
	s := &structs.NestedAvatarSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.AvatarOrLogo{
		FieldBase:    parseBase(f),
		DisplayStyle: s.DisplayStyle,
	}, nil
}

func parseAttachment(f *structs.Field) (*fields.File, error) {
	s := &structs.NestedAttachmentSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.File{
		FieldBase: parseBase(f),
		Required:  f.Required,
		Multiple:  s.Multiple,
		FileTypes: s.MimeTypes,
	}, nil
}

func parseComposite(f *structs.Field) (interface{}, error) {
	s := &structs.NestedCompositeSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	parseFields, err := ParseFields(s.RelatedToFields)
	if err != nil {
		return nil, err
	}
	if s.IsExtract {
		field := &fields.ExtractSingleRecord{
			FieldBase:            parseBase(f),
			CompositeTypeAPIName: s.RelatedToAPIName,
			SubFields:            parseFields,
			Filter:               s.Filter,
			SortConditions:       s.Sort,
			RecordPosition:       transRecordPosition(s),
		}
		field.Type = constants.ExtractCompositeType
		return field, nil
	} else {
		return &fields.CompositeType{
			FieldBase:            parseBase(f),
			CompositeTypeAPIName: s.RelatedToAPIName,
			Required:             f.Required,
			Multiple:             s.Multiple,
			SubFields:            parseFields,
		}, nil
	}
}

func transRecordPosition(s *structs.NestedCompositeSetting) int64 {
	var pos int64
	switch s.FilterRecord.Type {
	case "first":
		pos = 1
		break
	case "last":
		pos = -1
		break
	case "specified":
		pos = s.FilterRecord.Index
	}
	return pos
}

func parseExtractComposite(f *structs.Field) (*fields.ExtractSingleRecord, error) {
	s := &structs.NestedCompositeSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	parseFields, err := ParseFields(s.RelatedToFields)
	if err != nil {
		return nil, err
	}
	return &fields.ExtractSingleRecord{
		FieldBase:            parseBase(f),
		CompositeTypeAPIName: s.RelatedToAPIName,
		SubFields:            parseFields,
		Filter:               s.Filter,
		SortConditions:       s.Sort,
		RecordPosition:       transRecordPosition(s),
	}, nil
}

func parseFormula(f *structs.Field) (*fields.Formula, error) {
	s := &structs.NestedFormulaSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.Formula{
		FieldBase:  parseBase(f),
		ReturnType: s.ReturnType,
		Formula:    s.Formulas,
	}, nil
}

func parseBigint(f *structs.Field) (*fields.Bigint, error) {
	s := &structs.NestedFloatSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.Bigint{
		FieldBase: parseBase(f),
		Required:  f.Required,
		Unique:    transUnique(f),
	}, nil
}

func parseDecimal(f *structs.Field) (*fields.Decimal, error) {
	s := &structs.NestedFloatSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.Decimal{
		FieldBase:           parseBase(f),
		Required:            f.Required,
		Unique:              transUnique(f),
		DisplayAsPercentage: s.DisplayAsPercentage,
		DecimalPlacesNumber: s.DecimalPlaces,
	}, nil
}

func parseRollup(f *structs.Field) (*fields.Rollup, error) {
	s := &structs.NestedRollupSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.Rollup{
		FieldBase:                parseBase(f),
		RollupType:               s.RollupFunctionType,
		RollupObjectApiName:      s.RollupedObject.ApiName,
		RollupFieldApiName:       s.RollupedField.ApiName,
		RollupLookupFieldApiName: s.RollupedLookupField.ApiName,
		Filter:                   s.RollupRangeFilter,
	}, nil
}

func parseRegion(f *structs.Field) (*fields.Region, error) {
	s := structs.NestedRegionSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.Region{
		FieldBase:   parseBase(f),
		Required:    f.Required,
		Multiple:    s.Multiple,
		OptionLevel: s.OptionLevel,
		StrictLevel: s.StrictLevel,
		Filter:      s.Filter,
	}, nil
}

func ParseField(f *structs.Field) (res interface{}, err error) {
	switch f.Type.Name {
	case constants.OpTextType:
		res, err = parseText(f)
		break
	case constants.OpRichTextType:
		res, err = parseRichText(f)
		break
	case constants.OpMultilingualType:
		res, err = parseMultilingual(f)
		break
	case constants.OpFloatType:
		res, err = parseFloat(f)
		break
	case constants.OpEncryptNumberType:
		break
	case constants.OpEnumType:
		res, err = parseEnum(f)
		break
	case constants.OpLookupType:
		res, err = parseLookup(f)
		break
	case constants.OpBooleanType:
		res, err = parseBoolean(f)
		break
	case constants.OpDatetimeType:
		res, err = parseDateTime(f)
		break
	case constants.OpDateType:
		res, err = parseDate(f)
		break
	case constants.OpPhoneNumberType:
		res, err = parsePhoneNumber(f)
		break
	case constants.OpEmailType:
		res, err = parseEmail(f)
		break
	case constants.OpReferenceType:
		res, err = parseReference(f)
		break
	case constants.OpAutoNumberType:
		res, err = parseAutoNumber(f)
		break
	case constants.OpAvatarType:
		res, err = parseAvatar(f)
		break
	case constants.OpBackLookupType:
		break
	case constants.OpAttachmentType:
		res, err = parseAttachment(f)
		break
	case constants.OpCompositeType:
		res, err = parseComposite(f)
		break
	case constants.OpExtractCompositeType:
		res, err = parseExtractComposite(f)
		break
	case constants.OpFormulaType:
		res, err = parseFormula(f)
		break
	case constants.OpConditionType:
		break
	case constants.OpEnumMultiType:
		break
	case constants.OpLookupMultiType:
		break
	case constants.OpInheritFieldType:
		break
	case constants.OpGIDFieldType:
		break
	case constants.OpBigintFieldType:
		res, err = parseBigint(f)
		break
	case constants.OpDecimalFieldType:
		res, err = parseDecimal(f)
		break
	case constants.OpRollupFieldType:
		res, err = parseRollup(f)
		break
	case constants.OpRegionFieldType:
		res, err = parseRegion(f)
		break
	}
	return
}

func ParseFields(fs []*structs.Field) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, len(fs))
	var tmp interface{}
	for _, f := range fs {
		tmp, err = ParseField(f)
		if err != nil {
			return nil, err
		}
		res[f.APIName] = tmp
	}
	return
}
