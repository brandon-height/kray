package ude

import (
	"log"

	"github.com/brandon-height/kray/kube"
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

	// Create namespace
	if err := createNamespaceResource(g.Name, c.Client); err != nil {
		return err
	}
	// Create DNS Record for UDE
	// Create GCE DB Instance
	return nil
}

// Delete ...
func (g *GCP) Delete(c *kube.Config) error {
	// Delete namespace
	deleteNamespaceResource(g.Name, c.Client)
	// Delete DNS Record for UDE
	// Delete GCE DB Instance
	return nil
}

// List ...
func (g *GCP) List(name string, c *kube.Config) error {

	var err error

	log.Println("Listing resources in", g.Name)
	if g.PodList, err = listPodResource(g.Name, c.Client); err != nil {
		log.Fatalln(err)
	}
	// Verify we have resources to deal with.
	if len(g.PodList.Items) < 1 {
		log.Println("No pod resources found")
		return nil
	}
	// List out the pods
	log.Println("Pods:")
	for _, pod := range g.PodList.Items {
		log.Println(pod.GetName())
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
