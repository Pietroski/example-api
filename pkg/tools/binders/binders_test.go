package binders_test

import (
	"encoding/json"
	"github.com/SimpleOpenBadge/authentication-api/pkg/tools/binders"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShouldBindJSON(t *testing.T) {
	t.Run("Test unique case", func(t *testing.T) {
		type ExampleStruct struct {
			Field1 string `json:"field1,omitempty"`
			Field2 string `json:"field2,omitempty"`
			Field3 int    `json:"field3,omitempty"`
		}

		es := &ExampleStruct{
			Field1: "field1 value",
			Field2: "field2 value",
			Field3: 1234,
		}

		bes, err := json.Marshal(es)
		require.NoError(t, err)

		sbes := string(bes)
		var newES ExampleStruct
		err = binders.ShouldBindJSON(sbes, &newES)
		require.NoError(t, err)
		require.Equal(t, es.Field1, newES.Field1)
		require.Equal(t, es.Field2, newES.Field2)
		require.Equal(t, es.Field3, newES.Field3)
	})

	t.Run("Test multiple cases", func(t *testing.T) {
		type ExampleStruct struct {
			Field1 string `json:"field1,omitempty"`
			Field2 string `json:"field2,omitempty"`
			Field3 int    `json:"field3,omitempty"`
		}

		es := &ExampleStruct{
			Field1: "field1 value",
			Field2: "field2 value",
			Field3: 1234,
		}
		bes, err := json.Marshal(es)
		require.NoError(t, err)

		es2 := &ExampleStruct{
			Field1: "field1 new value",
			Field2: "field2 old value",
			Field3: 4321,
		}
		bes2, err := json.Marshal(es2)
		require.NoError(t, err)



		sbes := string(bes)
		var newES ExampleStruct
		err = binders.ShouldBindJSON(sbes, &newES)
		require.NoError(t, err)
		require.Equal(t, es.Field1, newES.Field1)
		require.Equal(t, es.Field2, newES.Field2)
		require.Equal(t, es.Field3, newES.Field3)

		sbes2 := string(bes2)
		var newES2 ExampleStruct
		err = binders.ShouldBindJSON(sbes2, &newES2)
		require.NoError(t, err)
		require.Equal(t, es2.Field1, newES2.Field1)
		require.Equal(t, es2.Field2, newES2.Field2)
		require.Equal(t, es2.Field3, newES2.Field3)
	})
}
