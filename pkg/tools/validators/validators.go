package validators

import "regexp"

const (
	emailValRegStr = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{1,64}@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	//TODO: password supported characters regex string validation
)

func IsEmailValid(email string) bool {
	var rxEmail = regexp.MustCompile(emailValRegStr)
	if len(email) < 3 || len(email) > 254 || !rxEmail.MatchString(email) {
		return false
	}

	return true
}

func IsPasswordValid(password string) bool {
	if len(password) < 6 {
		return false
	}

	return true
}
