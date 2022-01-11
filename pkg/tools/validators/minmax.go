package validators

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	WrongMinFormatType = "error: %v field has wrong min tag type format: %w"
	WrongMaxFormatType = "error: %v field has wrong max tag type format: %w"

	ErrorSmallerThan = "error: %v has field value smaller than minimum tag validation value"
	ErrorGreaterThan = "error: %v has field value greater than maximum tag validation value"
)

func minLength(fieldName string, fieldValue interface{}, strMinLength string) error {
	minValue, err := strconv.Atoi(strMinLength)
	if err != nil {
		return fmt.Errorf(WrongMinFormatType, fieldName, err)
	}

	switch reflect.TypeOf(fieldValue).Kind() {
	case
		reflect.String,
		reflect.Slice,
		reflect.Array:
		fv := reflect.ValueOf(fieldValue)
		if fv.Len() < minValue {
			return fmt.Errorf(ErrorSmallerThan, fieldName)
		}
	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		fv := reflect.ValueOf(fieldValue).Int()
		if fv < int64(minValue) {
			return fmt.Errorf(ErrorSmallerThan, fieldName)
		}
	case
		reflect.Float32,
		reflect.Float64:
		fv := reflect.ValueOf(fieldValue).Float()
		if fv < float64(minValue) {
			return fmt.Errorf(ErrorSmallerThan, fieldName)
		}
	}

	return nil
}

func maxLength(fieldName string, fieldValue interface{}, strMinLength string) error {
	maxValue, err := strconv.Atoi(strMinLength)
	if err != nil {
		return fmt.Errorf(WrongMaxFormatType, fieldName, err)
	}

	switch reflect.TypeOf(fieldValue).Kind() {
	case
		reflect.String,
		reflect.Slice,
		reflect.Array:
		fv := reflect.ValueOf(fieldValue)
		if fv.Len() > maxValue {
			return fmt.Errorf(ErrorGreaterThan, fieldName)
		}
	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		fv := reflect.ValueOf(fieldValue).Int()
		if fv > int64(maxValue) {
			return fmt.Errorf(ErrorGreaterThan, fieldName)
		}
	case
		reflect.Float32,
		reflect.Float64:
		fv := reflect.ValueOf(fieldValue).Float()
		if fv > float64(maxValue) {
			return fmt.Errorf(ErrorGreaterThan, fieldName)
		}
	}

	return nil
}
