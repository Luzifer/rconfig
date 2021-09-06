package rconfig

import (
	"os"
	"testing"
)

func TestPrecedence(t *testing.T) {
	type testcfg struct {
		A int `default:"1" vardefault:"a" env:"a" flag:"avar,a" description:"a"`
	}

	var (
		err         error
		cfg         testcfg
		args        []string
		vardefaults map[string]string
	)

	exec := func(desc string, fn func() interface{}, exp interface{}) {
		cfg = testcfg{}
		SetVariableDefaults(vardefaults)
		err = parse(&cfg, args)

		if err != nil {
			t.Errorf("%q parsing caused error: %s", desc, err)
		}

		if res := fn(); res != exp {
			t.Errorf("%q expected value does not match: %#v != %#v", desc, res, exp)
		}
	}

	// Provided: Flag, Env, Default, VarDefault
	args = []string{"-a", "5"}
	os.Setenv("a", "8")
	vardefaults = map[string]string{
		"a": "3",
	}

	exec("Provided: Flag, Env, Default, VarDefault", func() interface{} { return cfg.A }, 5)

	// Provided: Env, Default, VarDefault
	args = []string{}
	os.Setenv("a", "8")
	vardefaults = map[string]string{
		"a": "3",
	}

	exec("Provided: Env, Default, VarDefault", func() interface{} { return cfg.A }, 8)

	// Provided: Default, VarDefault
	args = []string{}
	os.Unsetenv("a")
	vardefaults = map[string]string{
		"a": "3",
	}

	exec("Provided: Default, VarDefault", func() interface{} { return cfg.A }, 3)

	// Provided: Default
	args = []string{}
	os.Unsetenv("a")
	vardefaults = map[string]string{}

	exec("Provided: Default", func() interface{} { return cfg.A }, 1)
}
