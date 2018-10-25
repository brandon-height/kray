package ude

// UDE ...
type UDE struct {
	Name string
}

// New returns a new instance of type UDE.
func New(name string) *UDE {
	return &UDE{
		Name: name,
	}
}
