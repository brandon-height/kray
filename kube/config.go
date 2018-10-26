package kube

import (
	"k8s.io/client-go/kubernetes"
)

// Config ...
type Config struct {
	Name   string
	Client kubernetes.Interface
}

// NewConfig returns a new Config type.
func NewConfig(name string, client kubernetes.Interface) *Config {
	return &Config{
		Name:   name,
		Client: client,
	}
}
