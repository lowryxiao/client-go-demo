package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	// client
	client, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	// get a specific pod, need use resource as v1.Pod
	pod := v1.Pod{}
	err = client.Get().Namespace("kube-system").
		Resource("pods").Name("etcd-minikube").
		Do(context.TODO()).Into(&pod)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pod.Name)

	// get all pods under kube-system namespace
	pods := v1.PodList{}
	err = client.Get().Namespace("kube-system").Resource("pods").Do(context.TODO()).Into(&pods)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all pods under kube-system")
	for _, p := range pods.Items {
		fmt.Println(p.Name)
	}
}
