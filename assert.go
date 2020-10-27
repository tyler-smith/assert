package assert

import (
	"bytes"
	"testing"
)

func True(t *testing.T, given bool) {
	if !given {
		t.Fatal("Should be true but was false")
	}
}

func False(t *testing.T, given bool) {
	if given {
		t.Fatal("Should be false but was true")
	}
}

func Nil(t *testing.T, given interface{}) {
	if given != nil {
		t.Fatal("Should be nil.")
	}
}

func NotNil(t *testing.T, given interface{}) {
	if given == nil {
		t.Fatal("Should not be nil.")
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal("Error should be nil but was:", err.Error())
	}
}

func EqualError(t *testing.T, expected, given error) {
	if expected != given {
		t.Fatal("Errors not equal. Wanted:", expected, "\nGot:", given)
	}
}

func EqualString(t *testing.T, expected, given string) {
	if expected != given {
		t.Fatal("Strings not equal. Wanted:", expected, "\nGot:", given)
	}
}

func EqualInt(t *testing.T, expected, given int) {
	if expected != given {
		t.Fatal("Ints not equal. Wanted:", expected, "\nGot:", given)
	}
}

func EqualByteSlice(t *testing.T, expected, given []byte) {
	if !bytes.Equal(expected, given) {
		t.Fatal("Byte slices not equal. Wanted:", expected, "\nGot:", given)
	}
}

func IntsWithin(t *testing.T, expected, given, delta int) {
	expected -= given
	if expected < 0 {
		expected *= -1
	}
	if expected > delta {
		t.Fatal("Ints not close enough. Wanted", expected, "to be within", delta, "of", given)
	}
}
