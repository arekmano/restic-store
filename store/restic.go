package store

import (
	"restic-secret-store/exec"
)

type ResticStore struct {
	repository string
	isDryRun bool
	tags []string
	region string
	host string
}

type ResticConfiguration struct {
	Repository string
	IsDryRun bool
	Tags []string
	Region string
	Host string
}

func NewRestic(config *ResticConfiguration) *ResticStore {
	return &ResticStore{
		host: config.Host,
		isDryRun: config.IsDryRun,
		region: config.Region,
		repository: config.Repository,
		tags: config.Tags,
	}
}

func (r *ResticStore) Put(secretName, inputDir string) {
	args := []string{
		"restic",
		"--repo",
		r.repository,
		"--host",
		r.host,
		"--option",
		"s3.region="+r.region,
		"--verbose",
	}
	for _, tag := range r.tags {
		args = append(args, "--tag", tag)
	}
	args = append(
		args,
		"backup",
		inputDir,
	)
	command := exec.InitCommand(args)
	if r.isDryRun {
		command.Print()
	} else {
		command.Execute()
	}
}


func (r *ResticStore) Get(secretName, destDir string) {
	args := []string{
		"restic",
		"--repo",
		r.repository,
		"--host",
		r.host,
		"--option",
		"s3.region="+r.region,
		"--verbose",
		"--tag",
		secretName,
	}
	args = append(
		args,
		"restore",
		"latest",
		"--target",
		destDir,
	)
	command := exec.InitCommand(args)
	if r.isDryRun {
		command.Print()
	} else {
		command.Execute()
	}

}
