package rconfig

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
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
		vardefaults = map[string]string{
			"my_secret_value": "veryverysecretkey",
			"unkownkey":       "hi there",
			"int_var":         "42",
		}
	)

	exec := func(desc string, tests [][2]interface{}) {
		if err := parse(&cfg, args); err != nil {
			t.Fatalf("Parsing options caused error: %s", err)
		}

		for _, test := range tests {
			if !reflect.DeepEqual(reflect.ValueOf(test[0]).Elem().Interface(), test[1]) {
				t.Errorf("%q expected value does not match: %#v != %#v", desc, test[0], test[1])
			}
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

	tmp, _ := ioutil.TempFile("", "")
	defer func() {
		tmp.Close()
		os.Remove(tmp.Name())
	}()
	yamlData := "---\nmy_secret_value: veryverysecretkey\nunknownkey: hi there\nint_var: 42\n"
	tmp.WriteString(yamlData)
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
