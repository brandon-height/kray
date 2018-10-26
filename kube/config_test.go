package kube

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func Test_NewClient(t *testing.T) {
	c, err := NewClient()
	assert.NoError(t, err)
	assert.Implements(t, (*kubernetes.Interface)(nil), c)
}

func Test_NewConfig(t *testing.T) {
	c := NewConfig("derp", fake.NewSimpleClientset())
	assert.IsType(t, (*Config)(nil), c)
}
