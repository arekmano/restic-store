package cmd


import (
	"github.com/spf13/cobra"
)


var outputDir string
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "fetches a secret from the secret store",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := createStore()
		s.Get(secretName, outputDir)
		return nil
	},
}

func init() {
	getCmd.Flags().StringVarP(&secretName, "secret-name", "s", "", "the name of the secret to fetch")
	getCmd.Flags().StringVarP(&outputDir, "output-directory", "o", "", "the output directory")
	getCmd.MarkFlagRequired("output-directory")
	getCmd.MarkFlagRequired("secret-name")

}
