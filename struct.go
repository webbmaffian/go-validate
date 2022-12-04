package validate

import (
	"reflect"
	"strings"

	"github.com/webbmaffian/go-validate/utils"
)

const tagSeparator = ","
const kvSeparator = "="

func Struct(v any, options ...utils.Options) (valid bool, errors ValidationErrors) {
	var opt utils.Options

	if len(options) != 0 {
		opt = options[0]
	}

	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	if typ.Kind() == reflect.Pointer {
		val = val.Elem()
		typ = val.Type()
	}

	if typ.Kind() != reflect.Struct {
		return false, nil
	}

	errors = make(ValidationErrors, 0, 5)
	path := "$"

	iterateStructFields(typ, val, &opt, path, &errors)

	valid = len(errors) == 0

	return
}

func iterateStructFields(typ reflect.Type, val reflect.Value, opt *utils.Options, path string, errors *ValidationErrors) {
	numFields := typ.NumField()

	for i := 0; i < numFields; i++ {
		fld := typ.Field(i)

		if !fld.IsExported() {
			continue
		}

		fldVal := getValue(val.Field(i), opt)

		if zero, ok := fldVal.Interface().(IsZeroer); ok && opt.SkipNil && zero.IsZero() {
			continue
		}

		if tagStr := fld.Tag.Get("validate"); tagStr != "" {
			tags := strings.Split(tagStr, tagSeparator)

			for _, tag := range tags {
				tag, arg, _ := strings.Cut(tag, kvSeparator)

				if validator, exists := registeredValidators[tag]; exists {
					if !validator.acceptZero && (!fldVal.IsValid() || fldVal.IsZero()) {
						continue
					}

					if err := validator.validator(fldVal, val, arg, opt); err != nil {
						valErr := ValidationError{
							Tag:     tag,
							Message: err.Error(),
							Path:    path + "." + fieldName(fld),
						}

						if fldVal.IsValid() && fldVal.CanInterface() {
							valErr.Value = fldVal.Interface()
						}

						*errors = append(*errors, valErr)
					}
				}
			}
		}

		if fld.Type.Kind() == reflect.Pointer {
			fldVal = fldVal.Elem()
		}

		if fldVal.Kind() == reflect.Struct {
			iterateStructFields(fldVal.Type(), fldVal, opt, path+"."+fieldName(fld), errors)
		}
	}
}

func fieldName(fld reflect.StructField) string {
	if tag := fld.Tag.Get("json"); tag != "" {
		name, _, _ := strings.Cut(tag, ",")

		return name
	}

	return fld.Name
}
