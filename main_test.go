package main

import (
	"testing"

	"github.com/strowk/foxy-contexts/pkg/foxytest"
)

// TestWithFoxytest simply uses same testdata
// for both scenarios defining the mock and tests
func TestWithFoxytest(t *testing.T) {
	ts, err := foxytest.Read("testdata")
	if err != nil {
		t.Fatal(err)
	}
	ts.WithExecutable("go", []string{"run", "main.go", "serve", "testdata"})
	cntrl := foxytest.NewTestRunner(t)
	ts.WithLogging()
	ts.Run(cntrl)
	ts.AssertNoErrors(cntrl)
}
