package rconfig

import (
	"testing"
	"time"
)

func TestDurationParsing(t *testing.T) {
	var (
		args = []string{
			"--duration=23s", "-o", "45m",
		}
		cfg struct {
			Test    time.Duration `flag:"duration"`
			TestS   time.Duration `flag:"other-duration,o"`
			TestDef time.Duration `default:"30h"`
		}
	)

	if err := parse(&cfg, args); err != nil {
		t.Fatalf("Parsing options caused error: %s", err)
	}

	for _, test := range [][2]interface{}{
		{cfg.Test, 23 * time.Second},
		{cfg.TestS, 45 * time.Minute},

		{cfg.TestDef, 30 * time.Hour},
	} {
		if test[0] != test[1] {
			t.Errorf("Expected value does not match: %#v != %#v", test[0], test[1])
		}
	}
}
