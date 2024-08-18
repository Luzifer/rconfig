package rconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrecedence(t *testing.T) {
	type testcfg struct {
		A int `default:"1" vardefault:"a" env:"a" flag:"avar,a" description:"a"`
	}

	var (
		cfg         testcfg
		args        []string
		vardefaults map[string]string
	)

	exec := func(desc string, fn func() interface{}, exp interface{}) {
		cfg = testcfg{}
		SetVariableDefaults(vardefaults)
		assert.NoError(t, parse(&cfg, args), desc)

		assert.Equal(t, exp, fn())
	}

	// Provided: Flag, Env, Default, VarDefault
	args = []string{"-a", "5"}
	t.Setenv("a", "8")
	vardefaults = map[string]string{
		"a": "3",
	}

	exec("Provided: Flag, Env, Default, VarDefault", func() interface{} { return cfg.A }, 5)

	// Provided: Env, Default, VarDefault
	args = []string{}
	t.Setenv("a", "8")
	vardefaults = map[string]string{
		"a": "3",
	}

	exec("Provided: Env, Default, VarDefault", func() interface{} { return cfg.A }, 8)

	// Provided: Default, VarDefault
	args = []string{}
	require.NoError(t, os.Unsetenv("a"))
	vardefaults = map[string]string{
		"a": "3",
	}

	exec("Provided: Default, VarDefault", func() interface{} { return cfg.A }, 3)

	// Provided: Default
	args = []string{}
	require.NoError(t, os.Unsetenv("a"))
	vardefaults = map[string]string{}

	exec("Provided: Default", func() interface{} { return cfg.A }, 1)
}
