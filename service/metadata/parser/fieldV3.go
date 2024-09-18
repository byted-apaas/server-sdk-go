package parser

import (
	"github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/metadata/object/fields"
)

func parseBaseV3(f *structs.Field) fields.FieldBaseV3 {
	return fields.FieldBaseV3{
		Type:    constants.OpToProduct[f.Type.Name],
		APIName: f.APIName,
		Label:   f.Label.TransToMultilingualV3(),
	}
}

func parseTextV3(f *structs.Field) (*fields.TextV3, error) {
	s := &structs.NestedTextSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.TextV3{
		FieldBaseV3:     parseBaseV3(f),
		Required:        f.Required,
		Unique:          transUnique(f),
		CaseSensitive:   transCaseSensitive(f),
		Multiline:       s.Multiline,
		MaxLength:       s.MaxLength,
		ValidationRegex: "", // TODO 未返回
		ErrorMsg:        "",
	}, nil
}

func parseRichTextV3(f *structs.Field) (*fields.RichTextV3, error) {
	s := &structs.NestedRichTextSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.RichTextV3{
		FieldBaseV3: parseBaseV3(f),
		Required:    f.Required,
		MaxLength:   s.MaxLength,
	}, nil
}

func parseMultilingualV3(f *structs.Field) (*fields.MultilingualV3, error) {
	s := &structs.NestedMultilingualSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.MultilingualV3{
		FieldBaseV3:   parseBaseV3(f),
		Required:      f.Required,
		Unique:        transUnique(f),
		CaseSensitive: transCaseSensitive(f),
		Multiline:     s.Multiline,
		MaxLength:     s.MaxLength,
	}, nil
}

func parseFloatV3(f *structs.Field) (*fields.NumberV3, error) {
	s := &structs.NestedFloatSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.NumberV3{
		FieldBaseV3:         parseBaseV3(f),
		Required:            f.Required,
		Unique:              transUnique(f),
		DisplayAsPercentage: s.DisplayAsPercentage,
		DecimalPlacesNumber: s.DecimalPlaces,
	}, nil
}

func transOptionListV3(s *structs.NestedEnumSetting) (res []*fields.OptionItemV3) {
	res = make([]*fields.OptionItemV3, len(*(s.Options)))
	for i, option := range *(s.Options) {
		res[i] = &fields.OptionItemV3{
			Label:       option.Label.TransToMultilingualV3(),
			APIName:     option.APIName,
			Description: option.Description.TransToMultilingualV3(),
			Color:       transColor(option),
			Active:      option.Active > 0,
		}
	}
	return
}

func parseEnumV3(f *structs.Field) (*fields.OptionV3, error) {
	s := &structs.NestedEnumSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.OptionV3{
		FieldBaseV3:         parseBaseV3(f),
		Required:            f.Required,
		Multiple:            s.IsArray,
		OptionSource:        transOptionSource(s),
		GlobalOptionAPIName: s.RelatedToGlobalOption.APIName,
		OptionList:          transOptionListV3(s),
	}, nil

}

func parseLookupV3(f *structs.Field) (*fields.LookupV3, error) {
	s := &structs.NestedLookupSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.LookupV3{
		FieldBaseV3:    parseBaseV3(f),
		Required:       f.Required,
		Multiple:       s.IsArray,
		ObjectAPIName:  s.ReferencedObjectAPIName,
		Hierarchy:      s.IsHierarchy,
		DisplayStyle:   s.DisplayStyle,
		SortConditions: transSort(s),
		Filter:         nil, // TODO 未返回
	}, nil
}

func parseBooleanV3(f *structs.Field) (*fields.BooleanV3, error) {
	s := &structs.NestedBooleanSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.BooleanV3{
		FieldBaseV3:          parseBaseV3(f),
		DescriptionWhenTrue:  s.Description4True.TransToMultilingualV3(),
		DescriptionWhenFalse: s.Description4False.TransToMultilingualV3(),
		DefaultValue:         s.DefaultValue,
	}, nil
}

func parseDateTimeV3(f *structs.Field) (*fields.DateTimeV3, error) {
	return &fields.DateTimeV3{
		FieldBaseV3: parseBaseV3(f),
		Required:    f.Required,
	}, nil
}

func parseDateV3(f *structs.Field) (*fields.DateV3, error) {
	return &fields.DateV3{
		FieldBaseV3: parseBaseV3(f),
		Required:    f.Required,
	}, nil
}

func parsePhoneNumberV3(f *structs.Field) (*fields.MobileNumberV3, error) {
	return &fields.MobileNumberV3{
		FieldBaseV3: parseBaseV3(f),
		Required:    f.Required,
		Unique:      transUnique(f),
	}, nil
}

func parseEmailV3(f *structs.Field) (*fields.EmailV3, error) {
	return &fields.EmailV3{
		FieldBaseV3: parseBaseV3(f),
		Required:    f.Required,
		Unique:      transUnique(f),
	}, nil
}

func parseReferenceV3(f *structs.Field) (*fields.ReferenceFieldV3, error) {
	s := &structs.NestedReferenceFieldSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.ReferenceFieldV3{
		FieldBaseV3:       parseBaseV3(f),
		GuideFieldAPIName: s.TargetReferencedObjectAPIName,
		FieldAPIName:      s.TargetReferenceFieldAPIName,
	}, nil
}

func parseAutoNumberV3(f *structs.Field) (*fields.AutoIDV3, error) {
	s := &structs.NestedAutoNumberSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.AutoIDV3{
		FieldBaseV3:    parseBaseV3(f),
		GenerateMethod: s.GenerationMethod,
		DigitsNumber:   s.Digits,
		Prefix:         s.Prefix,
		Suffix:         s.Suffix,
	}, nil
}

func parseAvatarV3(f *structs.Field) (*fields.AvatarOrLogoV3, error) {
	s := &structs.NestedAvatarSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.AvatarOrLogoV3{
		FieldBaseV3:  parseBaseV3(f),
		DisplayStyle: s.DisplayStyle,
	}, nil
}

func parseAttachmentV3(f *structs.Field) (*fields.FileV3, error) {
	s := &structs.NestedAttachmentSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.FileV3{
		FieldBaseV3: parseBaseV3(f),
		Required:    f.Required,
		Multiple:    s.Multiple,
		FileTypes:   s.MimeTypes,
	}, nil
}

func parseCompositeV3(f *structs.Field) (interface{}, error) {
	s := &structs.NestedCompositeSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	parseFields, err := ParseFieldsV3(s.RelatedToFields)
	if err != nil {
		return nil, err
	}
	if s.IsExtract {
		field := &fields.ExtractSingleRecordV3{
			FieldBaseV3:          parseBaseV3(f),
			CompositeTypeAPIName: s.RelatedToAPIName,
			SubFields:            parseFields,
			Filter:               s.Filter,
			SortConditions:       s.Sort,
			RecordPosition:       transRecordPosition(s),
		}
		field.Type = constants.ExtractCompositeType
		return field, nil
	} else {
		return &fields.CompositeTypeV3{
			FieldBaseV3:          parseBaseV3(f),
			CompositeTypeAPIName: s.RelatedToAPIName,
			Required:             f.Required,
			Multiple:             s.Multiple,
			SubFields:            parseFields,
		}, nil
	}
}

func parseExtractCompositeV3(f *structs.Field) (*fields.ExtractSingleRecordV3, error) {
	s := &structs.NestedCompositeSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	parseFields, err := ParseFieldsV3(s.RelatedToFields)
	if err != nil {
		return nil, err
	}
	return &fields.ExtractSingleRecordV3{
		FieldBaseV3:          parseBaseV3(f),
		CompositeTypeAPIName: s.RelatedToAPIName,
		SubFields:            parseFields,
		Filter:               s.Filter,
		SortConditions:       s.Sort,
		RecordPosition:       transRecordPosition(s),
	}, nil
}

func parseFormulaV3(f *structs.Field) (*fields.FormulaV3, error) {
	s := &structs.NestedFormulaSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.FormulaV3{
		FieldBaseV3: parseBaseV3(f),
		ReturnType:  s.ReturnType,
		Formula:     s.Formulas.TransToMultilingualV3(),
	}, nil
}

func parseBigintV3(f *structs.Field) (*fields.BigintV3, error) {
	s := &structs.NestedFloatSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.BigintV3{
		FieldBaseV3: parseBaseV3(f),
		Required:    f.Required,
		Unique:      transUnique(f),
	}, nil
}

func parseDecimalV3(f *structs.Field) (*fields.DecimalV3, error) {
	s := &structs.NestedFloatSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.DecimalV3{
		FieldBaseV3:         parseBaseV3(f),
		Required:            f.Required,
		Unique:              transUnique(f),
		DisplayAsPercentage: s.DisplayAsPercentage,
		DecimalPlacesNumber: s.DecimalPlaces,
	}, nil
}

func parseRollupV3(f *structs.Field) (*fields.RollupV3, error) {
	s := &structs.NestedRollupSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	return &fields.RollupV3{
		FieldBaseV3:              parseBaseV3(f),
		RollupType:               s.RollupFunctionType,
		RollupObjectApiName:      s.RollupedObject.ApiName,
		RollupFieldApiName:       s.RollupedField.ApiName,
		RollupLookupFieldApiName: s.RollupedLookupField.ApiName,
		Filter:                   s.RollupRangeFilter,
	}, nil
}

func parseRegionV3(f *structs.Field) (*fields.RegionV3, error) {
	s := structs.NestedRegionSetting{}
	if err := utils.Decode(f.Type.Settings, &s); err != nil {
		return nil, err
	}
	filterV3 := []*structs.RegionFilterV3{}
	for _, filter := range s.Filter {
		filterV3 = append(filterV3, filter.ToV3())
	}
	return &fields.RegionV3{
		FieldBaseV3: parseBaseV3(f),
		Required:    f.Required,
		Multiple:    s.Multiple,
		OptionLevel: s.OptionLevel,
		StrictLevel: s.StrictLevel,
		Filter:      filterV3,
	}, nil
}

func ParseFieldV3(f *structs.Field) (res interface{}, err error) {
	switch f.Type.Name {
	case constants.OpTextType:
		res, err = parseTextV3(f)
		break
	case constants.OpRichTextType:
		res, err = parseRichTextV3(f)
		break
	case constants.OpMultilingualType:
		res, err = parseMultilingualV3(f)
		break
	case constants.OpFloatType:
		res, err = parseFloatV3(f)
		break
	case constants.OpEncryptNumberType:
		break
	case constants.OpEnumType:
		res, err = parseEnumV3(f)
		break
	case constants.OpLookupType:
		res, err = parseLookupV3(f)
		break
	case constants.OpBooleanType:
		res, err = parseBooleanV3(f)
		break
	case constants.OpDatetimeType:
		res, err = parseDateTimeV3(f)
		break
	case constants.OpDateType:
		res, err = parseDateV3(f)
		break
	case constants.OpPhoneNumberType:
		res, err = parsePhoneNumberV3(f)
		break
	case constants.OpEmailType:
		res, err = parseEmailV3(f)
		break
	case constants.OpReferenceType:
		res, err = parseReferenceV3(f)
		break
	case constants.OpAutoNumberType:
		res, err = parseAutoNumberV3(f)
		break
	case constants.OpAvatarType:
		res, err = parseAvatarV3(f)
		break
	case constants.OpBackLookupType:
		break
	case constants.OpAttachmentType:
		res, err = parseAttachmentV3(f)
		break
	case constants.OpCompositeType:
		res, err = parseCompositeV3(f)
		break
	case constants.OpExtractCompositeType:
		res, err = parseExtractCompositeV3(f)
		break
	case constants.OpFormulaType:
		res, err = parseFormulaV3(f)
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
		res, err = parseBigintV3(f)
		break
	case constants.OpDecimalFieldType:
		res, err = parseDecimalV3(f)
		break
	case constants.OpRollupFieldType:
		res, err = parseRollupV3(f)
		break
	case constants.OpRegionFieldType:
		res, err = parseRegionV3(f)
		break
	}
	return
}

func ParseFieldsV3(fs []*structs.Field) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, len(fs))
	var tmp interface{}
	for _, f := range fs {
		tmp, err = ParseFieldV3(f)
		if err != nil {
			return nil, err
		}
		res[f.APIName] = tmp
	}
	return
}
