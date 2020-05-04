package cmd

import (
	"restic-secret-store/store"
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

const defaultRepository = "s3:s3.amazonaws.com/restic-secret-store-repository"
const defaultRegion = "us-west-2"
const secretTag = "secrets"

var rootCmd = &cobra.Command{
	Use:   "restic-secret-store",
	Short: "restic-secret-store is a wrapper over restic, used to read/write secrets",
}

var dryRun bool
var secretName string
var region string
var repository string

func init() {
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dryrun", "d", false, "If specified, will only print out the command")
	rootCmd.PersistentFlags().StringVarP(&repository, "repository", "r", defaultRepository, "Uses the specified repository")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "z", defaultRegion, "Uses the specified AWS region")
	rootCmd.AddCommand(putCmd, getCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
}

func createStore() *store.ResticStore {
	return store.NewRestic(&store.ResticConfiguration{
		Host: "Restic-Store",
		IsDryRun: dryRun,
		Region: region,
		Repository: repository,
		Tags: []string{secretTag, secretName},
	})
}