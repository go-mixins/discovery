// Package discovery provides service registration and de-registration
// primitives.
//
// Example usage (Consul):
//
//     import (
//         "log"
//
//         "github.com/go-mixins/discovery/consul"
//     )
//
//     registry, err := consul.New("localhost:8500", "")
//     if err != nil {
//         log.Fatalf("%+v", err)
//     }
//     err = registry.Register("service.worker1", "10.0.0.1:5001", "production")
//     if err != nil {
//         log.Fatalf("%+v", err)
//     }
//     defer registry.Close()
//
//     if err := registry.Deregister("service.worker1"); err != nil {
//         log.Printf("%+v", err)
//     }
//
package discovery

import "github.com/go-mixins/errors"

// Errors provide base error class
var Errors = errors.NewClass("discovery")

// Registrator publishes and unpublishes service
type Registrator interface {
	Register(ID, address string, tags ...string) error
	Deregister(ID string) error
}

//go:generate moq -out mock/discovery.go -pkg mock . Registrator
