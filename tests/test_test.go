package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestTest(t *testing.T) {
	if true != true {
		t.Fatalf("Test system failing!")
	}
}
