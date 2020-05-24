package store_test

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/arekmano/restic-store/store"

	"github.com/stretchr/testify/require"
)

var randomStore = store.NewRestic(&store.ResticConfiguration{
	Host:       strconv.FormatInt(rand.Int63(), 36),
	Region:     strconv.FormatInt(rand.Int63(), 36),
	Repository: strconv.FormatInt(rand.Int63(), 36),
	Tags: []string{
		strconv.FormatInt(rand.Int63(), 36),
		strconv.FormatInt(rand.Int63(), 36),
	},
})

func TestPut(t *testing.T) {
	// Test Data
	inputDir := strconv.FormatInt(rand.Int63(), 36)
	secretName := strconv.FormatInt(rand.Int63(), 36)
	expectedArgs := []string{"restic",
		"--repo",
		randomStore.Repository,
		"--host",
		randomStore.Host,
		"--option",
		"s3.region=" + randomStore.Region,
		"--verbose",
		"--tag",
		secretName,
		"--tag",
		randomStore.Tags[0],
		"--tag",
		randomStore.Tags[1],
		"backup",
		inputDir,
	}

	// Execute
	command := randomStore.Put(secretName, inputDir)

	// Verify
	require.NotNil(t, command.BinaryPath)
	require.Equal(t, command.Arguments, expectedArgs)
	require.Equal(t, command.Environment, os.Environ())
}

func TestGet(t *testing.T) {
	// Test Data
	outputDir := strconv.FormatInt(rand.Int63(), 36)
	secretName := strconv.FormatInt(rand.Int63(), 36)
	expectedArgs := []string{"restic",
		"--repo",
		randomStore.Repository,
		"--host",
		randomStore.Host,
		"--option",
		"s3.region=" + randomStore.Region,
		"--verbose",
		"--tag",
		secretName,
		"restore",
		"latest",
		"--target",
		outputDir,
	}

	// Execute
	command := randomStore.Get(secretName, outputDir)

	// Verify
	require.NotNil(t, command.BinaryPath)
	require.Equal(t, command.Arguments, expectedArgs)
	require.Equal(t, command.Environment, os.Environ())

}
