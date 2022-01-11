package mappers_test

import (
	"fmt"
	"github.com/SimpleOpenBadge/authentication-api/pkg/tools/mappers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNormMapToOrdSlice(t *testing.T) {
	testMap := map[int]interface{}{
		0:  "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: struct {
			test string
		}{
			test: "test",
		},
	}

	ordSlice, err := mappers.UintNormMapToOrdSlice(testMap)
	require.NoError(t, err)
	fmt.Println("Ordinated slice: ", ordSlice, ordSlice[0], ordSlice[10])
}
