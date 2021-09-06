package rconfig

import (
	"testing"
)

func TestErrors(t *testing.T) {
	for test, parsable := range map[string]interface{}{
		"use string as default to int": struct {
			A int `default:"a"`
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
		if err := parse(&parsable, nil); err == nil {
			t.Errorf("Expected error but got none. Test: %s", test)
		}
	}

	if err := parse(struct {
		A string `default:"a"`
	}{}, nil); err == nil {
		t.Errorf("Expected error when feeding non-pointer struct to parse")
	}

	if err := parse("test", nil); err == nil {
		t.Errorf("Expected error when feeding non-pointer string to parse")
	}
}
