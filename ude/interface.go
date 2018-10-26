package ude

import "github.com/brandon-height/kray/kube"

// Interface ...
type Interface interface {
	Create(c *kube.Config) error
	Delete(c *kube.Config) error
	List(name string, c *kube.Config) error
}

// Upper ...
type Upper interface {
	Up(c *kube.Config) error
}

// Downer ...
type Downer interface {
	Down(c *kube.Config) error
}
