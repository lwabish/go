package k8s

import (
	"github.com/lwabish/go-snippets/pkg/k8s"

	"github.com/spf13/cobra"
)

// nlcCmd represents the nlc command
var nlcCmd = &cobra.Command{
	Use:   "nlc",
	Short: "node label cleaning",
	Long: `author: lwabish 
contact: imwubowen@gmail.com`,
	Run: func(cmd *cobra.Command, args []string) {
		k8s.CleanNodeLabels(labelFilters, erase)
	},
}

var labelFilters []string
var erase bool

func init() {
	Cmd.AddCommand(nlcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nlcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nlcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	nlcCmd.Flags().StringSliceVarP(&labelFilters, "filters", "f", nil, "[keyword1,keyword2...]")
	err := nlcCmd.MarkFlagRequired("filters")
	nlcCmd.Flags().BoolVarP(&erase, "erase", "e", false, "erase matched label from nodes")
	if err != nil {
		return
	}
}
