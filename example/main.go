package main

import (
	"log"

	"github.com/go-mixins/discovery/consul"
)

func main() {
	registry, err := consul.New("localhost:8500", "")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	err = registry.Register("service.worker1", "10.0.0.1:5001", "production")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer func() {
		if err := registry.Deregister("service.worker1"); err != nil {
			log.Printf("%+v", err)
		}
	}()
}
