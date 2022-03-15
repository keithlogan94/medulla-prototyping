package kubernetes

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	v12 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"
)

func loadInCluster() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func getDeploymentClient() v12.DeploymentInterface {
	clientset := loadInCluster()
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	return deploymentsClient
}
