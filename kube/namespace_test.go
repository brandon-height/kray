package kube

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func Test_CreateNamespace(t *testing.T) {
	f := fake.NewSimpleClientset()
	n := &Namespace{}
	n.Create("derp", f)

	nsList, _ := f.CoreV1().Namespaces().List(metav1.ListOptions{})
	for _, namespace := range nsList.Items {
		assert.Equal(t, namespace.GetName(), "derp")
	}

}

func Test_DeleteNamespace(t *testing.T) {
	n := &Namespace{}
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "derp",
		},
	}
	f := fake.NewSimpleClientset(ns)
	// Verify the namespace exists now.
	nsList, _ := f.CoreV1().Namespaces().List(metav1.ListOptions{})
	for _, namespace := range nsList.Items {
		assert.Equal(t, namespace.GetName(), "derp")
	}

	// Now delete it
	n.Delete(ns, f)
	nsList, _ = f.CoreV1().Namespaces().List(metav1.ListOptions{})
	for _, namespace := range nsList.Items {
		assert.NotEqual(t, namespace.GetName(), "derp")
	}

}
