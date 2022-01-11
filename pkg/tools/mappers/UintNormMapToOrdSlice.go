/*
	NormalisedMapToOrderedSlice
*/

package mappers

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	ErrorInvalidArgumentType = "error: invalid argument type; the argument should be of type map[int]interface{}"
	ErrorInvalidMapKeyType   = "error: invalid map key type; the map key type should be an int and all the map keys should start at zero or one and it should be a sequence"
	ErrorEmptyMap            = "error: map cannot be empty"
)

func UintNormMapToOrdSlice(normMap interface{}) ([]interface{}, error) {
	if reflect.TypeOf(normMap).Kind() != reflect.Map {
		return nil, fmt.Errorf(ErrorInvalidArgumentType)
	}

	limit := mapLen(normMap)
	if limit == 0 {
		return nil, fmt.Errorf(ErrorEmptyMap)
	}

	var sIdx int
	init, err := initialMapNumberValue(normMap)
	if err != nil {
		return nil, err
	}

	ordSlice := make([]interface{}, limit)
	if init == 1 {
		sIdx--
	}

	for idx := init; idx < limit; idx++ {
		if ordSlice[sIdx], err = mapValue(normMap, idx); err != nil {
			return nil, err
		}
		sIdx++
	}

	return ordSlice, nil
}

func initialMapNumberValue(normMap interface{}) (int, error) {
	if reflect.ValueOf(normMap).MapIndex(reflect.ValueOf(0)).IsValid() {
		return 0, nil
	}
	if reflect.ValueOf(normMap).MapIndex(reflect.ValueOf(1)).IsValid() {
		return 1, nil
	}

	return 0, errors.New(ErrorInvalidMapKeyType)
}

func mapLen(normMap interface{}) int {
	return len(reflect.ValueOf(normMap).MapKeys())
}

func mapValue(normMap interface{}, idx int) (value interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered error: map value not found: %v", r)
		}
	}()

	value = reflect.ValueOf(normMap).MapIndex(reflect.ValueOf(idx)).Interface()

	return
}
