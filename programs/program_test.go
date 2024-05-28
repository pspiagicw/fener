package program

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestPrograms(t *testing.T) {

	// t.Skipf("Skipping program-based test")
	files, err := filepath.Glob("*.fn")

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			cmd := exec.Command("../fener", "run", file)

			_, err := cmd.Output()
			if err != nil {
				t.Errorf("Error for %s: %v", file, err)
			}
		})
	}
}
