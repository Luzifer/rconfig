package rconfig

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGeneralExecution(t *testing.T) {
	type test struct {
		Test        string `default:"foo" env:"shell" flag:"shell" description:"Test"`
		Test2       string `default:"blub" env:"testvar" flag:"testvar,t" description:"Test"`
		DefaultFlag string `default:"goo"`
		SadFlag     string
	}

	var (
		args []string
		cfg  test
	)

	exec := func(desc string, tests [][2]interface{}) {
		require.NoError(t, parse(&cfg, args))

		for _, test := range tests {
			assert.Equal(t, test[1], reflect.ValueOf(test[0]).Elem().Interface(), desc)
		}
	}

	cfg = test{}
	args = []string{
		"--shell=test23",
		"-t", "bla",
	}
	exec("defined arguments", [][2]interface{}{
		{&cfg.Test, "test23"},
		{&cfg.Test2, "bla"},
		{&cfg.SadFlag, ""},
		{&cfg.DefaultFlag, "goo"},
	})

	cfg = test{}
	args = []string{}
	exec("no arguments", [][2]interface{}{
		{&cfg.Test, "foo"},
	})

	cfg = test{}
	args = []string{}
	t.Setenv("shell", "test546")
	exec("no arguments and set env", [][2]interface{}{
		{&cfg.Test, "test546"},
	})
	require.NoError(t, os.Unsetenv("shell"))

	cfg = test{}
	args = []string{
		"--shell=test23",
		"-t", "bla",
		"positional1", "positional2",
	}
	exec("additional arguments", [][2]interface{}{
		{&cfg.Test, "test23"},
		{&cfg.Test2, "bla"},
		{&cfg.SadFlag, ""},
		{&cfg.DefaultFlag, "goo"},
	})

	assert.Equal(t, []string{"positional1", "positional2"}, Args())
}

func TestValidationIntegration(t *testing.T) {
	type tValidated struct {
		Test string `flag:"test" default:"" validate:"nonzero"`
	}

	var (
		cfgValidated = tValidated{}
		args         = []string{}
	)

	assert.Error(t, parseAndValidate(&cfgValidated, args))
}
