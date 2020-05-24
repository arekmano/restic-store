package cmd

import (
	"fmt"
	"os"

	"github.com/arekmano/restic-store/store"

	"github.com/spf13/cobra"
)

const defaultRepository = "s3:s3.amazonaws.com/restic-secret-store-repository"
const defaultRegion = "us-west-2"
const secretTag = "secrets"

var rootCmd = &cobra.Command{
	Use:   "restic-secret-store",
	Short: "restic-secret-store is a wrapper over restic, used to read/write secrets",
	Long:  `restic-secret-store builds on top of restic's backup capabilities to read/write directories to a repository, based off the secret name.`,
}

var dryRun bool
var secretName string
var region string
var repository string

func init() {
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dryrun", "d", false, "If specified, will only print out the command")
	rootCmd.PersistentFlags().StringVarP(&repository, "repository", "r", defaultRepository, "The restic repository to use for the secret store. See https://restic.readthedocs.io/en/stable/030_preparing_a_new_repo.html for more information.")
	// (assumes S3 repository TODO: refactor this away)
	rootCmd.PersistentFlags().StringVarP(&region, "region", "z", defaultRegion, "Use the specified AWS region")
	rootCmd.AddCommand(putCmd, getCmd)
}

// Execute will be run when restic-secret-store is executed.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createStore() *store.ResticStore {
	return store.NewRestic(&store.ResticConfiguration{
		Host:       "Restic-Store",
		Region:     region,
		Repository: repository,
		Tags:       []string{secretTag},
	})
}
