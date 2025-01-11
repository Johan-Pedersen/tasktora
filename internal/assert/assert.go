package assert

import (
	"database/sql"
	"reflect"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}

func NilError(t *testing.T, actual error) {
	t.Helper()
	if actual != nil {
		t.Errorf("got: %v; expected: nil", actual)
	}
}

func ErrNoRows(t *testing.T, actual error) {
	t.Helper()
	if actual != sql.ErrNoRows {
		t.Errorf("got: %v; expected: sql.ErrNoRows", actual)
	}
}
