package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var KubeConfig *string
var client *kubernetes.Clientset

func InitClient() {
	// todo:support in cluster sa
	// use the current context in kubeConfig
	config, err := clientcmd.BuildConfigFromFlags("", *KubeConfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientSet
	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}
