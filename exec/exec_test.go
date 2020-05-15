package exec_test

import (
	"restic-secret-store/exec"
	"testing"
)

func TestInitCommand(t *testing.T) {
	exec.InitCommand([]string{}) // Should not panic: assumes running on host with restic
}
