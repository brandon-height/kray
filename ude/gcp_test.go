package ude

import (
	"testing"

	"github.com/brandon-height/kray/kube"
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func Test_NewGCP(t *testing.T) {
	assert.Implements(t, (*Interface)(nil), NewGCP("gcp"))
}

func Test_Create(t *testing.T) {
	name := "derp"
	ude := NewGCP(name)
	config := kube.NewConfig(name, fake.NewSimpleClientset())
	assert.NoError(t, ude.Create(config))
}

func Test_Delete(t *testing.T) {
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "derp",
		},
	}

	name := "derp"
	ude := NewGCP("derp")
	config := kube.NewConfig(name, fake.NewSimpleClientset(ns))
	assert.NoError(t, ude.Delete(config))
}

func Test_Up(t *testing.T) {
	ude := NewGCP("gcp")
	assert.NoError(t, ude.Up(&kube.Config{}))
}

func Test_Down(t *testing.T) {
	ude := NewGCP("gcp")
	assert.NoError(t, ude.Down(&kube.Config{}))
}
