package utils

import (
	"reflect"
	"testing"
)

func Expect[T comparable](t *testing.T, got T, want T, msg string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%s: expected %v, got %v", msg, want, got)
	}
}
