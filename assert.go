package sensitive

import (
	"reflect"
	"testing"
)

func assertTrue(t *testing.T, expr bool, msg string) {
	t.Helper()

	if !expr {
		t.Error(msg)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf(`no error expected, but was "%s"`, err)
	}
}

func assertDeepEquals(t *testing.T, expected, actual interface{}, msg string) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Error(msg)
	}
}
