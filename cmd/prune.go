package cmd


import (
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "stores a secret in the secret store",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	pruneCmd.Flags().StringVarP(&secretName, "secret-name", "s", "", "the name of the secret to fetch")
	pruneCmd.Flags().StringVarP(&inputDir, "input-directory", "i", "", "the input directory")
	pruneCmd.MarkFlagRequired("input-directory")
	pruneCmd.MarkFlagRequired("secret-name")
}
