package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"net/http"
	"time"
)

// Resolver represents out resolver object
type Resolver struct {
	Client http.Client
}

// NewResolver sets up a new resolver
func NewResolver() *Resolver {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	return &Resolver{
		Client: client,
	}
}
