package validators

import (
	"fmt"
	"reflect"
)

const (
	ErrorRequiredField    = "error: %v is a mandatory field"
	ErrorShouldNotBeEmpty = "error: %v is a mandatory field and should not be empty"
)

func checkRequirement(fieldName string, fieldValue interface{}) error {
	switch reflect.TypeOf(fieldValue).Kind() {
	case reflect.String:
		fv := reflect.ValueOf(fieldValue)
		if fv.IsZero() {
			return fmt.Errorf(ErrorRequiredField, fieldName)
		}
	case
		reflect.Slice,
		reflect.Array:
		fv := reflect.ValueOf(fieldValue)
		if fv.IsNil() {
			return fmt.Errorf(ErrorRequiredField, fieldName)
		}
		if fv.Len() == 0 {
			return fmt.Errorf(ErrorShouldNotBeEmpty, fieldName)
		}
	case reflect.Ptr:
		fv := reflect.ValueOf(fieldValue)
		if fv.IsNil() {
			return fmt.Errorf(ErrorRequiredField, fieldName)
		}

		switch reflect.TypeOf(fv).Kind() {
		case reflect.Struct:
			nfv := fv.Elem().Interface()
			return NewValidator(nfv).Validate()
		}
	}

	return nil
}
