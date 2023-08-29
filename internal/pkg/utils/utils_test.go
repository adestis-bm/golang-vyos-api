package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheck(t *testing.T) {
	uhd, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("os.UserHomeDir() failed with: %s", err)
	}

	for _, chk := range []struct {
		in string
		ex string
	}{
		{in: "", ex: ""},
		{in: "/tmp/xyz", ex: "/tmp/xyz"},
		{in: "~/xyz", ex: filepath.Join(uhd, "xyz")},
	} {
		re, err := Check(chk.in)
		if err != nil {
			t.Errorf("Check(%q) failed with: %s", chk.in, err)
			continue
		}
		if re != chk.ex {
			t.Errorf("Check(%q) is %q, expected %q", chk.in, re, chk.ex)
		}
	}
}
