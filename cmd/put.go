package cmd


import (
	"github.com/spf13/cobra"
)

var inputDir string
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "stores a secret in the secret store",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := createStore()
		s.Put(secretName, inputDir)
		return nil
	},
}

func init() {
	putCmd.Flags().StringVarP(&secretName, "secret-name", "s", "", "the name of the secret to fetch")
	putCmd.Flags().StringVarP(&inputDir, "input-directory", "i", "", "the input directory")
	putCmd.MarkFlagRequired("input-directory")
	putCmd.MarkFlagRequired("secret-name")
}
