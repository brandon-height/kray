package ude

import (
	"log"

	"github.com/brandon-height/kray/kube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GCP ...
type GCP struct {
	*UDE
}

// NewGCP ...
func NewGCP(name string) *GCP {
	return &GCP{New(name)}
}

// Create ...
func (g *GCP) Create(c *kube.Config) error {
	log.Println("Creating", g.Name)
	// Create namespace
	namespace := &kube.Namespace{}
	namespace.Create(g.Name, c.Client)
	// Create DNS Record for UDE
	// Create GCE DB Instance
	return nil
}

// Delete ...
func (g *GCP) Delete(c *kube.Config) error {
	log.Println("Deleting", g.Name)
	// Delete namespace
	namespace := &kube.Namespace{}
	n, err := c.Client.CoreV1().Namespaces().Get(g.Name, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	namespace.Delete(n, c.Client)
	// Delete DNS Record for UDE
	// Delete GCE DB Instance
	return nil
}

// List ...
func (g *GCP) List(name string, c *kube.Config) error {
	log.Println("Listing resources in", g.Name)
	// initialize namespace object.
	namespace := &kube.Namespace{}
	namespaces, err := namespace.List(g.Name, c.Client)
	if err != nil {
		return err
	}
	if len(namespaces.Items) < 1 {
		log.Println("No pod resources found")
		return nil
	}
	// List out the pods
	log.Println("Pods:")
	for _, pod := range namespaces.Items {
		log.Println(pod)
	}
	// List other resources
	// List DNS Record for UDE
	// List GCE DB Instance
	return nil
}

// Up ...
func (g *GCP) Up(c *kube.Config) error {
	log.Println("Turning up", g.Name)
	// Turn up the GCE DB Instance
	// Turn up the Replication Controllers
	return nil
}

// Down ...
func (g *GCP) Down(c *kube.Config) error {
	log.Println("Turning down", g.Name)
	// Turn down the GCE DB Instance
	// Turn down the Replication Controllers
	return nil
}
