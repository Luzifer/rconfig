package rconfig

import (
	"testing"
)

func TestUintParsing(t *testing.T) {
	var (
		args = []string{
			"--int=1", "-i", "2",
			"--int8=3", "-8", "4",
			"--int32=5", "-3", "6",
			"--int64=7", "-6", "8",
			"--int16=9", "-1", "10",
		}
		cfg struct {
			Test    uint   `flag:"int"`
			TestP   uint   `flag:"intp,i"`
			Test8   uint8  `flag:"int8"`
			Test8P  uint8  `flag:"int8p,8"`
			Test16  uint16 `flag:"int16"`
			Test16P uint16 `flag:"int16p,1"`
			Test32  uint32 `flag:"int32"`
			Test32P uint32 `flag:"int32p,3"`
			Test64  uint64 `flag:"int64"`
			Test64P uint64 `flag:"int64p,6"`
			TestDef uint8  `default:"66"`
		}
	)

	if err := parse(&cfg, args); err != nil {
		t.Fatalf("Parsing options caused error: %s", err)
	}

	for _, test := range [][2]interface{}{
		{cfg.Test, uint(1)},
		{cfg.TestP, uint(2)},
		{cfg.Test8, uint8(3)},
		{cfg.Test8P, uint8(4)},
		{cfg.Test32, uint32(5)},
		{cfg.Test32P, uint32(6)},
		{cfg.Test64, uint64(7)},
		{cfg.Test64P, uint64(8)},
		{cfg.Test16, uint16(9)},
		{cfg.Test16P, uint16(10)},

		{cfg.TestDef, uint8(66)},
	} {
		if test[0] != test[1] {
			t.Errorf("Expected value does not match: %#v != %#v", test[0], test[1])
		}
	}
}
