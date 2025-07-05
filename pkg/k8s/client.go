package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var KubeConfig *string

var client *kubernetes.Clientset

func InitClient() {
	var config *rest.Config
	if c, err := rest.InClusterConfig(); err == nil {
		config = c
	} else if c, err = clientcmd.BuildConfigFromFlags("", *KubeConfig); err == nil {
		config = c
	} else {
		panic(err)
	}

	cli, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	client = cli
}

func GetClient() *kubernetes.Clientset {
	return client
}
