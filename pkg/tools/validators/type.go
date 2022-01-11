package validators

import (
	"fmt"
	"reflect"
	"regexp"
)

const (
	emailValRegStr = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{1,64}@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	minEmailLength = 3
	maxEmailLength = 256

	//TODO: password supported characters regex string validation
	minPassLength = 8
)

var (
	rxEmail = regexp.MustCompile(emailValRegStr)
)

func checkEmail(fieldName string, fieldValue interface{}) (err error) {
	if err = minLength(fieldName, fieldValue, fmt.Sprintf("%d", minEmailLength)); err != nil {
		return err
	}
	if err = maxLength(fieldName, fieldValue, fmt.Sprintf("%d", maxEmailLength)); err != nil {
		return err
	}

	emailValue := reflect.ValueOf(fieldValue).String()
	if !rxEmail.MatchString(emailValue) {
		return fmt.Errorf("invalid email format")
	}

	return nil
}

func checkPassword(fieldName string, fieldValue interface{}) (err error) {
	if err = minLength(fieldName, fieldValue, fmt.Sprintf("%d", minPassLength)); err != nil {
		return err
	}

	return nil
}
