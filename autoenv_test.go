package rconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeriveEnvVarName(t *testing.T) {
	for test, expect := range map[string]string{
		"1Foobar":              "1_FOOBAR",
		"BC1":                  "BC1",
		"BIGCase1":             "BIG_CASE1",
		"BIGCase":              "BIG_CASE",
		"Camel1":               "CAMEL1",
		"camel":                "CAMEL",
		"Camel":                "CAMEL",
		"CAMEL":                "CAMEL",
		"CamelCase":            "CAMEL_CASE",
		"_foobar":              "FOOBAR",
		"ILoveGoAndJSONSoMuch": "I_LOVE_GO_AND_JSON_SO_MUCH",
		"mrT":                  "MR_T",
		"my_case1":             "MY_CASE1",
		"MyFieldName":          "MY_FIELD_NAME",
		"SmallCASE":            "SMALL_CASE",
	} {
		assert.Equal(t, expect, deriveEnvVarName(test))
	}
}
