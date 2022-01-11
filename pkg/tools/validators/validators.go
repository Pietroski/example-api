package validators

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type (
	Validator interface {
		Validate() (err error)
	}

	validation struct {
		obj         interface{}
		fieldNumber int
		fieldNames
		fieldTags
		fieldValues
	}

	fieldTags   map[string]string
	fieldValues map[string]interface{}
	fieldNames  []string
)

func NewValidator(obj interface{}) Validator {
	objValueLen := reflect.ValueOf(obj).NumField()

	return &validation{
		obj:         obj,
		fieldNumber: objValueLen,
		fieldNames:  make(fieldNames, objValueLen),
		fieldTags:   make(fieldTags, objValueLen),
		fieldValues: make(fieldValues, objValueLen),
	}
}

func (v *validation) Validate() (err error) {
	err = v.extractFieldNames()
	err = v.extractStructFieldTags()
	err = v.extractStructFieldValues()

	return v.validate()
}

func (v *validation) extractFieldNames() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered error %v", r)
		}
	}()

	typeOfObj := reflect.TypeOf(v.obj)
	for idx := 0; idx < v.fieldNumber; idx++ {
		v.fieldNames[idx] = typeOfObj.Field(idx).Name
	}

	return
}

func (v *validation) extractStructFieldTags() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered error %v", r)
		}
	}()

	typeOfObj := reflect.TypeOf(v.obj)
	for _, fieldName := range v.fieldNames {
		fieldValue, validFieldName := typeOfObj.FieldByName(fieldName)
		if !validFieldName {
			continue
		}
		if structTag, lookupOk := fieldValue.Tag.Lookup("validation"); lookupOk {
			v.fieldTags[fieldName] = structTag
		}
	}

	// TODO: remove it!!
	//printIndented(v.fieldTags)

	return nil
}

func (v *validation) extractStructFieldValues() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered error %v", r)
		}
	}()

	objValue := reflect.ValueOf(v.obj)
	for _, fieldName := range v.fieldNames {
		v.fieldValues[fieldName] = objValue.FieldByName(fieldName).Interface()
	}

	// TODO: remove it!!
	//printIndented(v.fieldValues)

	return nil
}

func (v *validation) validate() (err error) {
	for fieldName, fieldValue := range v.fieldValues {
		caseTypesMap := splitStrIntoStrMap(v.fieldTags[fieldName])
		if err = v.validationMapIterator(fieldName, fieldValue, caseTypesMap); err != nil {
			return err
		}
	}

	return nil
}

func (v *validation) validationMapIterator(
	fieldName string,
	fieldValue interface{},
	caseTypesMap map[string]string,
) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered error %v", r)
		}
	}() // recovery

	{ // mandatory validator
		if _, required := caseTypesMap["required"]; required {
			if err = checkRequirement(fieldName, fieldValue); err != nil {
				return
			}
		}
	}

	{ // length validators
		if minValue, hasMin := caseTypesMap["min"]; hasMin {
			if err = minLength(fieldName, fieldValue, minValue); err != nil {
				return
			}
		}

		if maxValue, hasMax := caseTypesMap["max"]; hasMax {
			if err = maxLength(fieldName, fieldValue, maxValue); err != nil {
				return
			}
		}
	}

	{ // TODO: add custom validators
		//
	}

	return
}

func splitStrIntoStrMap(s string) map[string]string {
	xs := strings.Split(s, ",")
	ssLen := len(xs)

	mss := make(map[string]string, ssLen)
	for _, v := range xs {
		st := strings.Split(v, "=")
		if len(st) == 2 {
			mss[st[0]] = st[1]
			continue
		}
		mss[st[0]] = ""
	}

	return mss
}

func printIndented(obj interface{}) (err error) {
	bfv, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println("Field values: ", string(bfv))

	return
}
