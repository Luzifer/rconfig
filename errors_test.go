package rconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	for test, parsable := range map[string]interface{}{
		"use string as default to int": struct {
			A int `default:"a"` //revive:disable-line:struct-tag // Intentional error for testing
		}{},
		"use string as default to float": struct {
			A float32 `default:"a"`
		}{},
		"use string as default to uint": struct {
			A uint `default:"a"`
		}{},
		"use string as default to uint in sub-struct": struct {
			B struct {
				A uint `default:"a"`
			}
		}{},
		"use string list as default to int slice": struct {
			A []int `default:"a,b"`
		}{},
	} {
		assert.Error(t, parse(&parsable, nil), test) //#nosec:G601 // Fine for this test
	}

	assert.Error(t, parse(struct {
		A string `default:"a"`
	}{}, nil), "feeding non-pointer to parse")

	assert.Error(t, parse("test", nil), "feeding non-pointer string to parse")
}
