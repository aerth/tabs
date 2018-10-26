package jsparse

import (
	"errors"
	"strings"
	"testing"
)

type testcase struct {
	name    string
	input   string
	gold    string
	wanterr error // strings match .Error() method
}

func TestParseJSON(t *testing.T) {
	for testnumber, tc := range []testcase{} {
		testnumber++ // we know counting does start at zero, but others may not :(
		b, err := Parse(strings.NewReader(tc.input))
		if (err == nil) != (tc.wanterr == nil) {
			t.Errorf("Test %v: want err == nil, got err=%v", testnumber, err)
			return
		}
		if string(b) != tc.gold {
			t.Errorf("Test %v: want %q, got %q", testnumber, tc.gold, string(b))
			return
		}
		// pass
	}
}

func errar(s string) error {
	return errors.New(s)
}
