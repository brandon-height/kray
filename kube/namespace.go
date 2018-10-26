package kube

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Namespace defines a kubernetes Namespace.
type Namespace struct{}

// Create ...
func (n Namespace) Create(name string, client kubernetes.Interface) error {
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	_, err := client.Core().Namespaces().Create(ns)
	return err
}

// Delete ...
func (n Namespace) Delete(ns *v1.Namespace, client kubernetes.Interface) error {
	return client.CoreV1().Namespaces().Delete(ns.ObjectMeta.Name, &metav1.DeleteOptions{})
}

// List ...
func (n Namespace) List(name string, client kubernetes.Interface) (*v1.PodList, error) {
	return client.CoreV1().Pods(name).List(metav1.ListOptions{})
}
