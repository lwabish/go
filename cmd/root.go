package cmd

import (
	"fmt"
	"github.com/lwabish/go-snippets/cmd/exp"
	"github.com/lwabish/go-snippets/cmd/image"
	"github.com/lwabish/go-snippets/cmd/k8s"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
)

var version = "v1.0.4"
var genDoc = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "lwabish",
	Version: version,
	Short:   "command line tools written by go",
	Long: `author: lwabish 
contact: imwubowen@gmail.com`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if !genDoc {
			fmt.Println("add -h to see help message")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	generateDoc()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(exp.Cmd)
	rootCmd.AddCommand(image.Cmd)
	rootCmd.AddCommand(k8s.Cmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-snippets.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&genDoc, "gen-doc", "g", false, "generate doc and exit")
}

func generateDoc() {
	if genDoc {
		fmt.Println("Updating doc tree...")
		_ = doc.GenMarkdownTree(rootCmd, "./docs/")
		fmt.Printf("Done...\n")
		os.Exit(0)
	}
}
