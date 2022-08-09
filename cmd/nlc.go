package cmd

import (
	"github.com/lwabish/go-snippets/pkg/k8s"

	"github.com/spf13/cobra"
)

// nlcCmd represents the nlc command
var nlcCmd = &cobra.Command{
	Use:   "nlc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		k8s.CleanNodeLabels("")
	},
}

func init() {
	k8sCmd.AddCommand(nlcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nlcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nlcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
