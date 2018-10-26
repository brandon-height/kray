package kube

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

/*
func Test_CreatePod(t *testing.T) {
	f := fake.NewSimpleClientset()
	p := &Pod{}
	p.Create("derp", f)

	podList, _ := f.CoreV1().Pods("derp").List(metav1.ListOptions{})
	for _, pod := range podList.Items {
		assert.Equal(t, pod.GetName(), "derp")
	}

}
*/

func Test_DeletePod(t *testing.T) {
	name, namespace := "derp", "derp"
	p := &Pod{}
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	f := fake.NewSimpleClientset(pod)
	// Verify the pod exists now.
	podList, _ := f.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	for _, pp := range podList.Items {
		assert.Equal(t, pp.GetName(), name)
	}

	// Now delete it
	p.Delete(namespace, pod.GetName(), f)
	podList, _ = f.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	for _, pp := range podList.Items {
		assert.NotEqual(t, pp.GetName(), name)
	}

}
