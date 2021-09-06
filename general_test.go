package rconfig

import (
	"os"
	"reflect"
	"testing"
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
		if err := parse(&cfg, args); err != nil {
			t.Fatalf("Parsing options caused error: %s", err)
		}

		for _, test := range tests {
			if !reflect.DeepEqual(reflect.ValueOf(test[0]).Elem().Interface(), test[1]) {
				t.Errorf("%q expected value does not match: %#v != %#v", desc, test[0], test[1])
			}
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
	os.Setenv("shell", "test546")
	exec("no arguments and set env", [][2]interface{}{
		{&cfg.Test, "test546"},
	})
	os.Unsetenv("shell")

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

	if !reflect.DeepEqual(Args(), []string{"positional1", "positional2"}) {
		t.Errorf("expected positional arguments to match")
	}
}

func TestValidationIntegration(t *testing.T) {
	type tValidated struct {
		Test string `flag:"test" default:"" validate:"nonzero"`
	}

	var (
		cfgValidated = tValidated{}
		args         = []string{}
	)

	if err := parseAndValidate(&cfgValidated, args); err == nil {
		t.Errorf("Expected error, got none")
	}
}
