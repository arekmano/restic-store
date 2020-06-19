package exec

import (
	OSExec "os/exec"

	"github.com/sirupsen/logrus"
)

// ResticCommand represents the restic command that can be executed
type ResticCommand struct {
	BinaryPath string
	Arguments  []string
}

// InitCommand will create a new ResticCommand.
// This function will panic if restic is not installed and found in the PATH
func InitCommand(Arguments []string) (*ResticCommand, error) {
	command := ResticCommand{
		Arguments: Arguments,
	}
	binary, lookErr := OSExec.LookPath("restic")
	if lookErr != nil {
		return nil, lookErr
	}
	command.BinaryPath = binary
	return &command, nil
}

// Execute will execute the ResticCommand. Returns the output of the command in
// bytes.
func (r *ResticCommand) Execute() ([]byte, error) {
	logrus.
		WithField("Arguments", r.Arguments).
		WithField("BinaryPath", r.BinaryPath).
		Info("Executing Command")
	command := OSExec.Command(r.BinaryPath, r.Arguments...)
	return command.CombinedOutput()
}

// Print logs the command using the logger. Useful for debugging.
func (r *ResticCommand) Print() {
	logrus.
		WithField("BinaryPath", r.BinaryPath).
		WithField("Arguments", r.Arguments).
		Info("Restic Command Printed")
}
