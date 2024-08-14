package test

import (
	"testing"
)

func Equal(t *testing.T, a, b any) {
	if a != b {
		t.Fatalf("expected %v, got %v", a, b)
	}
}
