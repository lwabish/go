package image

import (
	"github.com/spf13/cobra"

	"github.com/lwabish/go/pkg/image"
)

// metaCmd represents the meta command
var metaCmd = &cobra.Command{
	Use:   "meta",
	Short: "write specific docker image git revisions to a text file",
	Long: `author: lwabish 
contact: imwubowen@gmail.com`,
	Run: func(cmd *cobra.Command, args []string) {
		image.ScanImageLabels(imageFilters)
	},
}

var imageFilters []string

func init() {
	Cmd.AddCommand(metaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	metaCmd.Flags().StringSliceVarP(&imageFilters, "filters", "f", nil, "[labelKey,labelValue,tagKeyword]")
	err := metaCmd.MarkFlagRequired("filters")
	if err != nil {
		return
	}
}
