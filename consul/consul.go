// Package consul provides simple Consul client API wrapper that handles
// registering and deregistering of service
package consul

import (
	"net"
	"strconv"
	"strings"

	"github.com/hashicorp/consul/api"

	"github.com/go-mixins/discovery"
)

// Errors makes error class "discovery" more specific
var Errors = discovery.Errors.Sub("consul")

// Registrator implements discovery.Registrator for Consul service
type Registrator struct {
	client *api.Client
}

var _ discovery.Registrator = (*Registrator)(nil)

// New returns new instance of the Registrator
func New(address, datacenter string) (res *Registrator, err error) {
	res = new(Registrator)
	config := api.DefaultConfig()
	config.Address = address
	config.Datacenter = datacenter
	if res.client, err = api.NewClient(config); err != nil {
		err = Errors.Wrap(err, "connecting to consul")
		return
	}
	return
}

// Register is called to register service in Consul with specific name,
// address and tags.
// If ID is "name.id" pair, name will be separated from ID and used as DNS
// service name for later discovery.
// If address is not an empty string it will be parsed as "host:port" pair.
func (r *Registrator) Register(ID, address string, tags ...string) error {
	reg := api.AgentServiceRegistration{
		Tags: tags,
	}
	nameID := strings.SplitN(ID, ".", 2)
	if len(nameID) == 1 {
		reg.Name = ID
	} else {
		reg.Name = nameID[0]
		reg.ID = ID
	}
	if address != "" {
		host, portstr, err := net.SplitHostPort(address)
		if err != nil {
			return Errors.Wrap(err, "extracting port")
		}
		reg.Address = host
		reg.Port, err = strconv.Atoi(portstr)
		if err != nil {
			return Errors.Wrap(err, "parsing port number")
		}
	}
	return Errors.Wrap(r.client.Agent().ServiceRegister(&reg), "registering service")
}

// Deregister removes service ID from the registry
func (r *Registrator) Deregister(id string) error {
	return Errors.Wrap(r.client.Agent().ServiceDeregister(id), "deregistering service")
}

// Close must be called to close connection to Consul. Note that service ids
// are not deregistered implicitly.
func (r *Registrator) Close() error {
	return Errors.Wrap(r.Close(), "closing r")
}
