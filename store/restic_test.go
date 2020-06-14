package store_test

import (
	"math/rand"
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

var randomOption = store.ResticOptions{
	Tags: []string{
		strconv.FormatInt(rand.Int63(), 36),
		strconv.FormatInt(rand.Int63(), 36),
	},
	Options: map[string]string{
		"s3.region": strconv.FormatInt(rand.Int63(), 36),
	},
}

func TestPut(t *testing.T) {
	// Test Data
	inputDir := strconv.FormatInt(rand.Int63(), 36)
	expectedArgs := []string{
		"--repo",
		randomStore.Repository,
		"--host",
		randomStore.Host,
		"--verbose",
		"--json",
		"--option",
		"s3.region=" + randomOption.Options["s3.region"],
		"--tag",
		randomOption.Tags[0],
		"--tag",
		randomOption.Tags[1],
		"backup",
		inputDir,
	}

	// Execute
	command := randomStore.Put(inputDir, &randomOption)

	// Verify
	require.NotNil(t, command.BinaryPath)
	require.Equal(t, expectedArgs, command.Arguments)
}

func TestGet(t *testing.T) {
	// Test Data
	outputDir := strconv.FormatInt(rand.Int63(), 36)
	expectedArgs := []string{
		"--repo",
		randomStore.Repository,
		"--host",
		randomStore.Host,
		"--verbose",
		"--json",
		"--option",
		"s3.region=" + randomOption.Options["s3.region"],
		"--tag",
		randomOption.Tags[0],
		"--tag",
		randomOption.Tags[1],
		"restore",
		"latest",
		"--target",
		outputDir,
	}

	// Execute
	command := randomStore.Get(outputDir, &randomOption)

	// Verify
	require.NotNil(t, command.BinaryPath)
	require.Equal(t, command.Arguments, expectedArgs)
}
