package option_test

import (
	"testing"

	"github.com/duythinht/option"
)

func TestSomeInt(t *testing.T) {
	five := option.Some(5)

	switch s := five.(type) {
	case option.SomeOf[int]:
		t.Logf("s %d", s.Unwrap())
	default:
		t.Logf("s %v is not Some[int]", s)
		t.Fail()
	}

	if five.UnwrapOr(10) != 5 {
		t.Fail()
		t.Logf("Should be 5")
	}
}

func TestNone(t *testing.T) {

	nop := option.None[int]()

	if !nop.IsNone() {
		t.Fail()
	}

	if nop.UnwrapOr(10) != 10 {
		t.Fail()
	}
}
