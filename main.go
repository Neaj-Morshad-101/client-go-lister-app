package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/appscodepc/.kube/config", "location to your config file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Printf("error %s while building config from flags\n", err.Error())

		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error getting InClusterConfig %s \n", err.Error())
			return
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s while creating clientset\n", err.Error())
		return
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s while listing all pods from default namespace\n", err.Error())
		return
	}
	fmt.Print(len(pods.Items))
	fmt.Println(" pods from default namespace")

	for _, pod := range pods.Items {
		fmt.Printf("%s\n", pod.Name)
	}
	fmt.Println("\n")
	//fmt.Println(pods)

	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s while listing all deployments from default namespace\n", err.Error())
		return
	}
	fmt.Print(len(deployments.Items))
	fmt.Println(" deployments from default namespace")
	for _, dep := range deployments.Items {
		fmt.Printf("%s\n", dep.Name)
	}

}
