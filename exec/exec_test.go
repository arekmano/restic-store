package exec_test

import (
	"testing"

	"github.com/arekmano/restic-store/exec"
	"github.com/stretchr/testify/require"
)

func TestInitCommand(t *testing.T) {
	exec.InitCommand([]string{}) // Should not panic: assumes running on host with restic
}

func TestExecute(t *testing.T) {
	// Test data
	cmd := exec.ResticCommand{
		BinaryPath: "/usr/bin/echo",
		Arguments:  []string{"abc"},
	}

	// Execute
	result, err := cmd.Execute()

	// Verify
	require.NoError(t, err)
	require.Equal(t, "abc\n", string(result))
}
