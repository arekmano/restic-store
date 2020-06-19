package store

import (
	"fmt"

	"github.com/arekmano/restic-store/exec"
)

// ResticStore represents a restic secret store.
type ResticStore struct {
	Repository string
	Host       string
}

// ResticConfiguration represents all configuration required to create a
// ResticStore
type ResticConfiguration struct {
	Repository string
	Tags       []string
	Region     string
	Host       string
}

// ResticOptions represents all the keyword arguments that will go into the execution of the restic command
type ResticOptions struct {
	Tags    []string
	Options map[string]string
}

// NewRestic Creates a ResticStore, based on the given ResticConfiguration
func NewRestic(config *ResticConfiguration) *ResticStore {
	return &ResticStore{
		Host:       config.Host,
		Repository: config.Repository,
	}
}

// Put will insert a secret into the restic repository, by crafting the
// relevant restic command.
func (r *ResticStore) Put(inputDir string, options *ResticOptions) *exec.ResticCommand {
	args := r.prepareArguments(options)
	args = append(
		args,
		"backup",
		inputDir,
	)

	return exec.InitCommand(args)
}

func (r *ResticStore) ListSnapshots(options *ResticOptions) *exec.ResticCommand {
	args := r.prepareArguments(options)
	args = append(
		args,
		"snapshots",
	)

	return exec.InitCommand(args)
}

// Get will retrieve a secret from the restic repository, by crafting the
// relevant restic command.
func (r *ResticStore) Get(destDir string, options *ResticOptions, snapshotID string) *exec.ResticCommand {

	args := r.prepareArguments(options)
	args = append(
		args,
		"restore",
		snapshotID,
		"--target",
		destDir,
	)

	return exec.InitCommand(args)
}
func (r *ResticStore) prepareArguments(options *ResticOptions) []string {
	args := []string{
		"--repo",
		r.Repository,
		"--host",
		r.Host,
		"--verbose",
		"--json",
	}
	for option, val := range options.Options {
		args = append(args, "--option", fmt.Sprintf("%s=%s", option, val))
	}
	for _, tag := range options.Tags {
		args = append(args, "--tag", tag)
	}
	return args
}
