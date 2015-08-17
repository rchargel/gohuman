package main

// Client a client in the system.
type Client struct {
	ID          string
	Name        string
	Email       string
	ContactName string
}

// NewClient creates a new client.
func NewClient(name, contactName, email string) Client {
	return Client{
		Name:        name,
		ContactName: contactName,
		Email:       email,
	}
}
