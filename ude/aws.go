package ude

import (
	"log"

	"github.com/brandon-height/kray/kube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AWS ...
type AWS struct {
	*UDE
}

// NewAWS ...
func NewAWS(name string) *AWS {
	return &AWS{New(name)}
}

// Create ...
func (a *AWS) Create(c *kube.Config) error {

	namespace := &kube.Namespace{}

	log.Println("Creating", a.Name)

	// Create namespace
	namespace.Create(a.Name, c.Client)
	// Create DNS Record for UDE
	// Create DB Instance
	return nil

}

// Delete ...
func (a *AWS) Delete(c *kube.Config) error {

	namespace := &kube.Namespace{}

	log.Println("Deleting", a.Name)
	// Delete namespace
	n, err := c.Client.CoreV1().Namespaces().Get(a.Name, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	namespace.Delete(n, c.Client)
	// Delete DNS Record for UDE
	// Delete GCE DB Instance
	return nil

}

// List ...
func (a *AWS) List(name string, c *kube.Config) error {
	log.Println("Listing", a.Name)
	// List this UDE in AWS
	return nil
}

// Up ...
func (a *AWS) Up(c *kube.Config) error {
	log.Println("Turning up", a.Name)
	// Turn up the GCE instance
	// Turn up the Replication Controllers
	return nil
}

// Down ...
func (a *AWS) Down(c *kube.Config) error {
	log.Println("Turning down", a.Name)
	// Turn down the GCE instance
	// Turn down the Replication Controllers
	return nil

}
