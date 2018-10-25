package ude

import (
	"testing"

	"github.com/brandon-height/kray/kube"
	"github.com/stretchr/testify/assert"
)

func Test_NewAlicloud(t *testing.T) {
	assert.Implements(t, (*Interface)(nil), NewAlicloud("alicloud"))
}

func Test_Alicloud_Create(t *testing.T) {
	ude := NewAlicloud("alicloud")
	assert.NoError(t, ude.Create(&kube.Config{}))
}

func Test_Alicloud_Delete(t *testing.T) {
	ude := NewAlicloud("alicloud")
	assert.NoError(t, ude.Delete(&kube.Config{}))
}

func Test_Alicloud_Up(t *testing.T) {
	ude := NewAlicloud("alicloud")
	assert.NoError(t, ude.Up(&kube.Config{}))
}

func Test_Alicloud_Down(t *testing.T) {
	ude := NewAlicloud("alicloud")
	assert.NoError(t, ude.Down(&kube.Config{}))
}
