package cmd

import (
	"github.com/lwabish/go-snippets/pkg/k8s"
	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

// k8sCmd represents the k8s command
var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "tools to manipulate kubernetes objects",
	Long: `author: lwabish 
contact: imwubowen@gmail.com`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(k8sCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	if home := homedir.HomeDir(); home != "" {
		k8s.KubeConfig = k8sCmd.PersistentFlags().String("kubeconfig", filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeConfig file")
	} else {
		k8s.KubeConfig = k8sCmd.PersistentFlags().String("kubeconfig", "",
			"absolute path to the kubeConfig file")
	}

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// k8sCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	k8s.InitClient()
}
