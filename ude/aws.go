package ude

import (
	"log"

	"github.com/brandon-height/kray/kube"
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
	log.Println("Creating", a.Name)
	// Create this UDE in AWS
	return nil
}

// Delete ...
func (a *AWS) Delete(c *kube.Config) error {
	log.Println("Deleting", a.Name)
	// Delete this UDE in AWS
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
