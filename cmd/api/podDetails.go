package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"os"
	"time"
)

func podDetails() {
	podName := os.Getenv("POD_NAME")
	namespace := os.Getenv("NAMESPACE")

	config, err := rest.InClusterConfig()
	if err != nil {
		log.Println(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	pod, err := clientset.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	log.Println(pod)

	defer cancel()

}
