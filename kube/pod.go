package kube

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Pod defines a kubernetes Pod.
type Pod struct{}

// Create ...
func (p Pod) Create(name string, client kubernetes.Interface) error {
	/*
		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		}
		_, err := client.Core().Pods().Create(pod)
		return err
	*/
	return nil
}

// Delete ...
func (p Pod) Delete(namespace string, pod string, client kubernetes.Interface) error {
	return client.CoreV1().Pods(namespace).Delete(pod, &metav1.DeleteOptions{})
}

// List ...
func (p Pod) List(namespace string, client kubernetes.Interface) (*v1.PodList, error) {
	return client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
}
