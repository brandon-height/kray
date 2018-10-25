package ude

import (
	"testing"

	"github.com/brandon-height/kray/kube"
	"github.com/stretchr/testify/assert"
)

func Test_AWS_NewAWS(t *testing.T) {
	assert.Implements(t, (*Interface)(nil), NewAWS("aws"))
}

func Test_AWS_Create(t *testing.T) {
	ude := NewAWS("aws")
	assert.NoError(t, ude.Create(&kube.Config{}))
}

func Test_AWS_Delete(t *testing.T) {
	ude := NewAWS("aws")
	assert.NoError(t, ude.Delete(&kube.Config{}))
}

func Test_AWS_Up(t *testing.T) {
	ude := NewAWS("aws")
	assert.NoError(t, ude.Up(&kube.Config{}))
}

func Test_AWS_Down(t *testing.T) {
	ude := NewAWS("aws")
	assert.NoError(t, ude.Down(&kube.Config{}))
}
