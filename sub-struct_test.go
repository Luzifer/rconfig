package rconfig

import (
	"testing"
)

func TestSubStructParsing(t *testing.T) {
	var (
		args = []string{}
		cfg  struct {
			Test string `default:"blubb"`
			Sub  struct {
				Test string `default:"Hallo"`
			}
		}
	)

	if err := parse(&cfg, args); err != nil {
		t.Fatalf("Parsing options caused error: %s", err)
	}

	for _, test := range [][2]interface{}{
		{cfg.Test, "blubb"},
		{cfg.Sub.Test, "Hallo"},
	} {
		if test[0] != test[1] {
			t.Errorf("Expected value does not match: %#v != %#v", test[0], test[1])
		}
	}
}
