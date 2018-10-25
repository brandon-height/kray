package ude

import (
	"log"

	"github.com/brandon-height/kray/kube"
)

// Alicloud ...
type Alicloud struct {
	*UDE
}

// NewAlicloud ...
func NewAlicloud(name string) *Alicloud {
	return &Alicloud{New(name)}
}

// Create ...
func (a *Alicloud) Create(c *kube.Config) error {
	log.Println("Creating", a.Name)
	// Create this UDE in Alicloud
	return nil
}

// Delete ...
func (a *Alicloud) Delete(c *kube.Config) error {
	log.Println("Deleting", a.Name)
	// Delete this UDE in Alicloud
	return nil
}

// Up ...
func (a *Alicloud) Up(c *kube.Config) error {
	log.Println("Turning up", a.Name)
	// Turn up the GCE instance
	// Turn up the Replication Controllers
	return nil
}

// Down ...
func (a *Alicloud) Down(c *kube.Config) error {
	log.Println("Turning down", a.Name)
	// Turn down the GCE instance
	// Turn down the Replication Controllers
	return nil

}
