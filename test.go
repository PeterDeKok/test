package test

import (
	"reflect"
	"testing"
	"time"
)

func ExpectedError(t *testing.T, err error, expected string) {
	t.Helper()

	ExpectedErrorF(t, err, expected, "")
}

func ExpectedErrorF(t *testing.T, err error, expected, msg string) {
	t.Helper()

	var actual string

	if err == nil {
		actual = "<no error>"
	} else {
		actual = err.Error()
	}

	if len(msg) > 0 {
		msg = "\n" + msg
	}

	if actual != expected {
		t.Errorf("expected error\ngot:      \"%s\"\nexpected: %s\n%s", actual, expected, msg)

		t.Fail()
	}
}

func ExpectedNoError(t *testing.T, err error) {
	t.Helper()

	ExpectedNoErrorF(t, err, "")
}

func ExpectedNoErrorF(t *testing.T, err error, msg string) {
	t.Helper()

	if err == nil {
		return
	}

	if len(msg) > 0 {
		msg = "\n" + msg
	}

	t.Errorf("expected no error\ngot: %s\n%s", err, msg)

	t.Fail()
}

func ExpectedNotEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()

	ExpectedEqualF(t, actual, expected, true, "")
}

func ExpectedEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()

	ExpectedEqualF(t, actual, expected, false, "")
}

func ExpectedEqualF(t *testing.T, actual, expected interface{}, reverse bool, msg string) {
	t.Helper()

	equal := reflect.DeepEqual(actual, expected)

	if (!reverse && equal) || (reverse && !equal) {
		return
	}

	if len(msg) > 0 {
		msg = "\n" + msg
	}

	aType := reflect.TypeOf(actual)
	eType := reflect.TypeOf(expected)

	var not string

	if reverse {
		not = "not "
	}

	t.Errorf("expected values to %sbe equal\ngot:      %s (%s)\nexpected: %s (%s)%s", not, actual, aType, expected, eType, msg)

	t.Fail()
}

func ExpectedTime(t *testing.T, start, end time.Time, expected, delta time.Duration) {
	t.Helper()

	ExpectedTimeF(t, start, end, expected, delta, "")
}

func ExpectedTimeF(t *testing.T, start, end time.Time, expected, delta time.Duration, msg string) {
	t.Helper()

	actual := end.Sub(start).Truncate(time.Millisecond)

	if actual >= expected-delta && actual <= expected+delta {
		return
	}

	if len(msg) > 0 {
		msg = "\n" + msg
	}

	t.Errorf("expected duration to equal (+-%s)\ngot:      %s\nexpected: %s%s", delta, actual, expected, msg)

	t.Fail()
}

func ExpectedZeroValue(t *testing.T, actual interface{}) {
	t.Helper()

	ExpectedZeroValueF(t, actual, false, "")
}

func ExpectedNoZeroValue(t *testing.T, actual interface{}) {
	t.Helper()

	ExpectedZeroValueF(t, actual, true, "")
}

func ExpectedZeroValueF(t *testing.T, actual interface{}, reverse bool, msg string) {
	t.Helper()

	zeroVal := actual == nil || reflect.ValueOf(actual).IsZero()

	if (!reverse && zeroVal) || (reverse && !zeroVal) {
		return
	}

	if len(msg) > 0 {
		msg = "\n" + msg
	}

	var not string

	if reverse {
		not = "not "
	}

	t.Errorf("expected interface to %sbe zero value\ngot: %v, %s", not, actual, msg)

	t.Fail()
}
