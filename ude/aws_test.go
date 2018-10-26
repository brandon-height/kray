package ude

import (
	"testing"

	"github.com/brandon-height/kray/kube"
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func Test_AWS_NewAWS(t *testing.T) {
	assert.Implements(t, (*Interface)(nil), NewAWS("aws"))
}

func Test_AWS_Create(t *testing.T) {
	name := "derp"
	ude := NewAWS(name)
	config := kube.NewConfig(name, fake.NewSimpleClientset())
	assert.NoError(t, ude.Create(config))
}

func Test_AWS_Delete(t *testing.T) {
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "derp",
		},
	}

	name := "derp"
	ude := NewAWS(name)
	config := kube.NewConfig(name, fake.NewSimpleClientset(ns))
	assert.NoError(t, ude.Delete(config))
}

func Test_AWS_Up(t *testing.T) {
	name := "derp"
	ude := NewAWS(name)
	assert.NoError(t, ude.Up(&kube.Config{}))
}

func Test_AWS_Down(t *testing.T) {
	name := "derp"
	ude := NewAWS(name)
	assert.NoError(t, ude.Down(&kube.Config{}))
}
