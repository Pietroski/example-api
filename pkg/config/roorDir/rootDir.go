package rootDir

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// SearchRootDir searches for the project's root directory.
func SearchRootDir() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get root path: %s", err.Error())
	}

	separator := defineSeparator(pwd)
	pwdSlice := strings.Split(pwd, separator)
	pwdSliceLen := len(pwdSlice)

	for i := pwdSliceLen; i > 1; i-- {
		slicedSlice := pwdSlice[:i]
		newpwd := strings.Join(slicedSlice, separator)
		if rootDir, ok, err := walkByDir(newpwd); err != nil {
			return "", fmt.Errorf("could not get walk dir: %s", err.Error())
		} else if ok {
			return rootDir, nil
		}
	}

	return "", nil
}

// walkByDir Scans all the files from a given dir and checks for a .env
func walkByDir(pwd string) (string, bool, error) {
	files, err := ioutil.ReadDir(pwd)
	if err != nil {
		return "", false, err
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".env") {
			return pwd, true, nil
		}
	}

	return "", false, nil
}

// defineSeparator defines which one is the OS dir separator
func defineSeparator(str string) string {
	var separator string
	switch {
	case strings.Contains(str, "/"):
		separator = "/"
	case strings.Contains(str, "\\"):
		separator = "\\"
	}

	return separator
}
