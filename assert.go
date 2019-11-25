package assert

import (
	"testing"
)

func True(t *testing.T, a bool) {
	if !a {
		t.Fatal("Should be true but was false")
	}
}

func False(t *testing.T, a bool) {
	True(t, !a)
}

func EqualString(t *testing.T, a, b string) {
	if a != b {
		t.Fatal("Strings not equal. Wanted:", a, "\nGot:", b)
	}
}

func EqualError(t *testing.T, a, b error) {
	if a != b {
		t.Fatal("Errors not equal. Wanted:", a, "\nGot:", b)
	}
}

func EqualInt(t *testing.T, a, b int) {
	if a != b {
		t.Fatal("Ints not equal. Wanted:", a, "\nGot:", b)
	}
}

func IntsWithin(t *testing.T, a, b, delta int) {
	a -= b
	if a < 0 {
		a *= -1
	}
	if a > delta {
		t.Fatal("Ints not close enough. Wanted", a, "to be within", delta, "of", b)
	}
}

func NotNil(t *testing.T, a interface{}) {
	if a == nil {
		t.Fatal("Should not be nil.")
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal("Error should be nil but was:", err.Error())
	}
}
