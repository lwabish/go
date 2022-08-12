package exp

import (
	"github.com/lwabish/go-snippets/pkg/exp"
	"github.com/spf13/cobra"
)

// hugepageCmd represents the hugepage command
var hugepageCmd = &cobra.Command{
	Use:   "hugepage",
	Short: "go demo to use linux hugepage",
	Long:  `inspired by kernel demo of hugepage`,
	Run: func(cmd *cobra.Command, args []string) {
		exp.Run(false)
	},
}

func init() {
	Cmd.AddCommand(hugepageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hugepageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hugepageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
