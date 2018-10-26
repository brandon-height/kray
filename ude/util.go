package ude

import (
	"log"

	"github.com/brandon-height/kray/kube"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createNamespaceResource(name string, client kubernetes.Interface) error {

	namespace := &kube.Namespace{}

	log.Println("Creating", name)

	// Create namespace
	namespace.Create(name, client)

	return nil

}

func deleteNamespaceResource(name string, client kubernetes.Interface) error {
	namespace := &kube.Namespace{}

	log.Println("Deleting", name)
	// Delete namespace
	n, err := client.CoreV1().Namespaces().Get(name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	namespace.Delete(n, client)
	return nil

}

func listPodResource(name string, client kubernetes.Interface) (*v1.PodList, error) {
	pod := &kube.Pod{}
	return pod.List(name, client)
}
