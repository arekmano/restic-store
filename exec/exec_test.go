package exec_test

import (
	"testing"

	"github.com/arekmano/restic-store/exec"
)

func TestInitCommand(t *testing.T) {
	exec.InitCommand([]string{}) // Should not panic: assumes running on host with restic
}
