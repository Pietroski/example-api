package validators_test

import (
	"github.com/SimpleOpenBadge/authentication-api/pkg/tools/validators"
	"github.com/stretchr/testify/require"
	"testing"
)

type (
	TestStruct struct {
		Field1 string `validation:"required,type=email,min=5,max=100"`
		Field2 int    `validation:"required,min=5,max=10"`
		Field3 []int  `validation:"required,min=2,max=10"`
		Field4 bool
		Field5 string `validation:"required"`
		Field6 bool
		Field7 []int          `validation:"required"`
		Field8 []int          `validation:"required"`
		Field9 *TestSubStruct `validation:"required"`
	}

	TestSubStruct struct {
		Field1 string `validation:"required,type=email,min=5,max=100"`
		Field2 int    `validation:"required,min=5,max=10"`
		Field3 []int  `validation:"required,min=2,max=10"`
		Field4 bool
		Field5 string `validation:"required"`
		Field6 bool
		Field7 []int `validation:"required"`
		Field8 []int `validation:"required"`
	}
)

func TestValidate(t *testing.T) {
	ts := TestStruct{
		Field1: "email Field 1",
		Field2: 7,
		Field3: []int{4, 5, 6},
		Field4: true,
		Field5: "asd",
		Field6: false,
		Field7: []int{0},
		Field8: []int{0},
		Field9: &TestSubStruct{
			Field1: "email Field 1 sub",
			Field2: 8,
			Field3: []int{3, 6, 9},
			Field4: false,
			Field5: "hello 5",
			Field6: false,
			Field7: []int{7},
			Field8: []int{8},
		},
	}
	err := validators.NewValidator(ts).Validate()
	require.NoError(t, err)
}
