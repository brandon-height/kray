package ude

import "k8s.io/api/core/v1"

// UDE ...
type UDE struct {
	Name          string
	PodList       *v1.PodList
	ConfigMapList *v1.ConfigMapList
	SecretList    *v1.SecretList
}

// New returns a new instance of type UDE.
func New(name string) *UDE {
	return &UDE{
		Name: name,
	}
}
