package image

import (
	"github.com/spf13/cobra"
)

// Cmd represents the image command
var Cmd = &cobra.Command{
	Use:   "image",
	Short: "tools to manipulate docker images",
	Long: `author: lwabish 
contact: imwubowen@gmail.com`,
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
