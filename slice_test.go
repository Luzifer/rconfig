package rconfig

import (
	"reflect"
	"testing"
)

func TestSliceParsing(t *testing.T) {
	var (
		args = []string{
			"--int=4,5", "-s", "hallo,welt",
		}
		cfg struct {
			Int         []int    `default:"1,2,3" flag:"int"`
			String      []string `default:"a,b,c" flag:"string"`
			IntP        []int    `default:"1,2,3" flag:"intp,i"`
			StringP     []string `default:"a,b,c" flag:"stringp,s"`
			EmptyString []string `default:""`
		}
	)

	if err := parse(&cfg, args); err != nil {
		t.Fatalf("Parsing options caused error: %s", err)
	}

	for _, test := range [][2]interface{}{
		{len(cfg.Int), 2},
		{cfg.Int, []int{4, 5}},

		{len(cfg.IntP), 3},
		{cfg.IntP, []int{1, 2, 3}},

		{len(cfg.String), 3},
		{cfg.String, []string{"a", "b", "c"}},

		{len(cfg.StringP), 2},
		{cfg.StringP, []string{"hallo", "welt"}},

		{len(cfg.EmptyString), 0},
	} {
		if !reflect.DeepEqual(test[0], test[1]) {
			t.Errorf("Expected value does not match: %#v != %#v", test[0], test[1])
		}
	}
}
