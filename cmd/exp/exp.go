package exp

import (
	"github.com/spf13/cobra"
)

// Cmd represents the exp command
var Cmd = &cobra.Command{
	Use:   "exp",
	Short: "temporary demo",
	Long: `author: lwabish 
contact: imwubowen@gmail.com`,
}

func init() {
	//rootCmd.AddCommand(expCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// expCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// expCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
