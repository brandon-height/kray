package kube

import (
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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

// NewClient ...
func NewClient() (kubernetes.Interface, error) {
	var (
		config *rest.Config
		err    error
	)

	if _, err = os.Stat(os.Getenv("HOME") + "/.kube/config"); err == nil {
		if config, err = clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config"); err != nil {
			log.Fatalln(err)
		}
	} else {
		if config, err = rest.InClusterConfig(); err != nil {
			log.Fatalln(err)
		}
	}

	return kubernetes.NewForConfig(config)
}
