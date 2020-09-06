package main

import (
	"flag"
	"testing"

	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	err := run([]string{"testman"})
	if err != flag.ErrHelp {
		t.Fatalf("err should be flag.ErrHelp: %v", err)
	}
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
