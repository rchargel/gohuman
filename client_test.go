package main

import "testing"

func TestNewClient(t *testing.T) {
	client := NewClient("Client", "John Smith", "jsmith@company.com")
	if client.ContactName != "John Smith" {
		t.Errorf("Invalid contact name: %v\n", client.ContactName)
	}
	if client.Name != "Client" {
		t.Errorf("Invalid name: %v\n", client.Name)
	}
	if client.Email != "jsmith@company.com" {
		t.Errorf("Invalid email: %v\n", client.Email)
	}
	if len(client.ID) == 0 {
		t.Errorf("Invalid ID: %v\n", client.ID)
	}
}
