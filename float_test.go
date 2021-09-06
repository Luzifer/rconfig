package rconfig

import (
	"testing"
)

func TestFloatParsing(t *testing.T) {
	var (
		args = []string{
			"--float32=5.5", "-3", "6.6",
			"--float64=7.7", "-6", "8.8",
		}
		cfg struct {
			Test32  float32 `flag:"float32"`
			Test32P float32 `flag:"float32p,3"`
			Test64  float64 `flag:"float64"`
			Test64P float64 `flag:"float64p,6"`
			TestDef float32 `default:"66.256"`
		}
	)

	if err := parse(&cfg, args); err != nil {
		t.Fatalf("Parsing options caused error: %s", err)
	}

	for _, test := range [][2]interface{}{
		{cfg.Test32, float32(5.5)},
		{cfg.Test32P, float32(6.6)},
		{cfg.Test64, float64(7.7)},
		{cfg.Test64P, float64(8.8)},

		{cfg.TestDef, float32(66.256)},
	} {
		if test[0] != test[1] {
			t.Errorf("Expected value does not match: %#v != %#v", test[0], test[1])
		}
	}
}
