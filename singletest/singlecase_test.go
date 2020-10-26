package singletest

import (
	"testing"
)

func TestIntMin(t *testing.T) {
	intMin := IntMin(4, 7)
	if intMin != 4 {
		t.Errorf("IntMin(4, 7) = %d; want 6", intMin)
	}
}
