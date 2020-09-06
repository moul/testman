package testpkg

import (
	"fmt"
	"math/rand"
	"testing"

	"moul.io/srand"
)

func ExampleAlwaysSucceed() {
	fmt.Println(AlwaysSucceed())
	// Output: <nil>
}

func TestStableAlwaysSucceed(t *testing.T) {
	if err := AlwaysSucceed(); err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestUnstableMaySucceed(t *testing.T) {
	rand.Seed(srand.Fast())
	if err := MaySucceed(); err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestBrokenAlwaysFailing(t *testing.T) {
	if err := AlwaysFailing(); err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
