package rconfig

import (
	"testing"
)

func TestBoolParsing(t *testing.T) {
	var (
		args = []string{
			"--test2",
			"-t",
		}
		cfg struct {
			Test1          bool `default:"true"`
			Test2          bool `default:"false" flag:"test2"`
			Test3          bool `default:"true" flag:"test3,t"`
			Test4          bool `flag:"test4"`
			TestEnvDefault bool `default:"true" env:"AN_ENV_VARIABLE_HOPEFULLY_NEVER_SET_DSFGDF"`
		}
	)

	if err := parse(&cfg, args); err != nil {
		t.Fatalf("Parsing options caused error: %s", err)
	}

	for _, test := range [][2]interface{}{
		{cfg.Test1, true},
		{cfg.Test2, true},
		{cfg.Test3, true},
		{cfg.Test4, false},
		{cfg.TestEnvDefault, true},
	} {
		if test[0] != test[1] {
			t.Errorf("Expected value does not match: %#v != %#v", test[0], test[1])
		}
	}
}
