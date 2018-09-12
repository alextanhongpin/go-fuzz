package bug_test

import (
	"testing"
	"testing/quick"

	"github.com/alextanhongpin/go-fuzz"
	fuzz "github.com/google/gofuzz"
)

func TestPanicWhenGreaterThan1000000(t *testing.T) {
	f := fuzz.New()
	var myInt int
	for i := 0; i < 1000; i += 1 {
		f.Fuzz(&myInt)
		bug.FaintHeart(myInt)
	}
	t.Log("done")
}

func TestQuickCheck(t *testing.T) {
	faint := func(i int) bool {
		bug.FaintHeart(i)
		return true
	}

	if err := quick.Check(faint, nil); err != nil {
		t.Error(err)
	}
}

func TestEqual(t *testing.T) {
	m := make(map[string]int)
	set := func(k string, v int) interface{} {
		m[k] = v
		return v
	}

	get := func(k string, v int) interface{} {
		// vv := m[k]
		return 0
	}

	if err := quick.CheckEqual(set, get, nil); err != nil {
		t.Error(err)
	}
}
