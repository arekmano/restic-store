package exec

import (
    "os"
    "os/exec"
	"syscall"
	"github.com/sirupsen/logrus"
)

type ResticCommand struct {
	binaryPath string
	arguments []string
	environment []string
}

func InitCommand(arguments []string) *ResticCommand {
	command := ResticCommand{
		environment: os.Environ(),
		arguments: arguments,
	}
	binary, lookErr := exec.LookPath("restic")
    if lookErr != nil {
        panic(lookErr)
    }
	command.binaryPath = binary
	return &command
}

func (r *ResticCommand) Execute() {
    execErr := syscall.Exec(r.binaryPath, r.arguments, r.environment)
    if execErr != nil {
        panic(execErr)
    }
}


func (r *ResticCommand) Print() {
	logrus.
		WithField("binaryPath", r.binaryPath).
		WithField("arguments", r.arguments).
		WithField("environment", r.environment).
		Info("Restic Command Printed")
}
