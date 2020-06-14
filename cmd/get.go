package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var outputDir string
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetches a secret from the secret store",
	Long: `restic-secret-store builds on top of restic's backup/restore capabilities to retrieve to directories, based off the a given secret name.

	This command works by calling "restic restore" under the hood. Use the "dryrun" flag to see what the restic command will be.`,
	Example: "restic-secret-store get --repository ./encrypted-restic-repo -o ./unencypted-secret -s test-secret",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := createStore()
		options := createOptions()
		options.Tags = append(options.Tags, secretName)
		command := s.Get(outputDir, options)
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
	getCmd.Flags().StringVarP(&secretName, "secret-name", "s", "", "the name of the secret to fetch")
	getCmd.Flags().StringVarP(&outputDir, "output-directory", "o", "", "the output directory")
	getCmd.MarkFlagRequired("output-directory")
	getCmd.MarkFlagRequired("secret-name")

}
