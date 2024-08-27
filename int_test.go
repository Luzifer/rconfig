package rconfig

import (
	"testing"
)

func TestIntParsing(t *testing.T) {
	var (
		args = []string{
			"--int=1", "-i", "2",
			"--int8=3", "-8", "4",
			"--int16=2", "-1", "3",
			"--int32=5", "-3", "6",
			"--int64=7", "-6", "8",
		}
		cfg struct {
			Test    int   `flag:"int"`
			TestP   int   `flag:"intp,i"`
			Test8   int8  `flag:"int8"`
			Test8P  int8  `flag:"int8p,8"`
			Test16  int16 `flag:"int16"`
			Test16P int16 `flag:"int16p,1"`
			Test32  int32 `flag:"int32"`
			Test32P int32 `flag:"int32p,3"`
			Test64  int64 `flag:"int64"`
			Test64P int64 `flag:"int64p,6"`
			TestDef int8  `default:"66"`
		}
	)

	if err := parse(&cfg, args); err != nil {
		t.Fatalf("Parsing options caused error: %s", err)
	}

	for _, test := range [][2]interface{}{
		{cfg.Test, int(1)},
		{cfg.TestP, 2},
		{cfg.Test8, int8(3)},
		{cfg.Test8P, int8(4)},
		{cfg.Test16, int16(2)},
		{cfg.Test16P, int16(3)},
		{cfg.Test32, int32(5)},
		{cfg.Test32P, int32(6)},
		{cfg.Test64, int64(7)},
		{cfg.Test64P, int64(8)},
		{cfg.TestDef, int8(66)},
	} {
		if test[0] != test[1] {
			t.Errorf("Expected value does not match: %#v != %#v", test[0], test[1])
		}
	}
}
