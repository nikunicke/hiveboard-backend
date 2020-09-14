package mongodb

import (
	"testing"
)

func TestReturnTwo(t *testing.T) {
	got := ReturnTwo()
	if got != 2 {
		t.Errorf("ReturnTwo() = %d; want 2", got)
	}
}
