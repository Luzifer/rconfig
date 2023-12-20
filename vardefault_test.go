package rconfig

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVardefaultParsing(t *testing.T) {
	type test struct {
		MySecretValue string `default:"secret" env:"foo" vardefault:"my_secret_value"`
		MyUsername    string `default:"luzifer" vardefault:"username"`
		SomeVar       string `flag:"var" description:"some variable"`
		IntVar        int64  `vardefault:"int_var" default:"23"`
	}

	var (
		cfg         test
		args        = []string{}
		err         error
		vardefaults = map[string]string{
			"my_secret_value": "veryverysecretkey",
			"unkownkey":       "hi there",
			"int_var":         "42",
		}
	)

	exec := func(desc string, tests [][2]interface{}) {
		require.NoError(t, parse(&cfg, args))

		for _, test := range tests {
			assert.Equal(t, test[1], reflect.ValueOf(test[0]).Elem().Interface(), desc)
		}
	}

	SetVariableDefaults(vardefaults)
	exec("manually provided variables", [][2]interface{}{
		{&cfg.IntVar, int64(42)},
		{&cfg.MySecretValue, "veryverysecretkey"},
		{&cfg.MyUsername, "luzifer"},
		{&cfg.SomeVar, ""},
	})

	SetVariableDefaults(VarDefaultsFromYAML([]byte("---\nmy_secret_value: veryverysecretkey\nunknownkey: hi there\nint_var: 42\n")))
	exec("defaults from YAML data", [][2]interface{}{
		{&cfg.IntVar, int64(42)},
		{&cfg.MySecretValue, "veryverysecretkey"},
		{&cfg.MyUsername, "luzifer"},
		{&cfg.SomeVar, ""},
	})

	tmp, _ := os.CreateTemp("", "")
	t.Cleanup(func() {
		tmp.Close()           //nolint:errcheck,gosec,revive // Just cleanup, will be closed automatically
		os.Remove(tmp.Name()) //nolint:errcheck,gosec,revive // Just cleanup of tmp-file
	})
	yamlData := "---\nmy_secret_value: veryverysecretkey\nunknownkey: hi there\nint_var: 42\n"
	_, err = tmp.WriteString(yamlData)
	require.NoError(t, err)
	SetVariableDefaults(VarDefaultsFromYAMLFile(tmp.Name()))
	exec("defaults from YAML file", [][2]interface{}{
		{&cfg.IntVar, int64(42)},
		{&cfg.MySecretValue, "veryverysecretkey"},
		{&cfg.MyUsername, "luzifer"},
		{&cfg.SomeVar, ""},
	})

	SetVariableDefaults(VarDefaultsFromYAML([]byte("---\nmy_secret_value = veryverysecretkey\nunknownkey = hi there\nint_var = 42\n")))
	exec("defaults from invalid YAML data", [][2]interface{}{
		{&cfg.IntVar, int64(23)},
		{&cfg.MySecretValue, "secret"},
		{&cfg.MyUsername, "luzifer"},
		{&cfg.SomeVar, ""},
	})

	SetVariableDefaults(VarDefaultsFromYAMLFile("/tmp/this_file_should_not_exist_146e26723r"))
	exec("defaults from non-existing YAML file", [][2]interface{}{
		{&cfg.IntVar, int64(23)},
		{&cfg.MySecretValue, "secret"},
		{&cfg.MyUsername, "luzifer"},
		{&cfg.SomeVar, ""},
	})
}
