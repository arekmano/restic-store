package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var inputDir string
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Stores a secret in the secret store",
	Long: `restic-secret-store builds on top of restic's backup/restore capabilities to retrieve to directories, based off the a given secret name.

	This command works by calling "restic backup" under the hood. Use the "dryrun" flag to see what the restic command will be.`,
	Example: "restic-secret-store put --repository ./encrypted-restic-repo -i ./unencypted-secret -s test-secret",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := createStore()
		options := createOptions()
		options.Tags = append(options.Tags, secretName)

		command := s.Put(inputDir, options)
		if dryRun {
			command.Print()
		} else {
			output, err := command.Execute()
			if err != nil {
				logrus.Fatal(err)
			}
			fmt.Println(string(output))
		}
		return nil
	},
}

func init() {
	putCmd.Flags().StringVarP(&secretName, "secret-name", "s", "", "the name of the secret being stored")
	putCmd.Flags().StringVarP(&inputDir, "input-directory", "i", "", "the directory to store")
	putCmd.MarkFlagRequired("input-directory")
	putCmd.MarkFlagRequired("secret-name")
}
