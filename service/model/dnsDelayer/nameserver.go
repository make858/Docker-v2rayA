// +build !confonly

package dnsDelayer

import (
	"context"
	"errors"
	"log"

	"v2ray.com/core/common/net"
	"v2ray.com/core/features/dns/localdns"
)

// IPOption is an object for IP query options.
type IPOption struct {
	IPv4Enable bool
	IPv6Enable bool
}

// Client is the interface for DNS client.
type Client interface {
	// Name of the Client.
	Name() string

	// QueryIP sends IP queries to its configured server.
	QueryIP(ctx context.Context, domain string, option IPOption) ([]net.IP, error)
}

type localNameServer struct {
	client *localdns.Client
}

func (s *localNameServer) QueryIP(ctx context.Context, domain string, option IPOption) ([]net.IP, error) {
	if option.IPv4Enable && option.IPv6Enable {
		return s.client.LookupIP(domain)
	}

	if option.IPv4Enable {
		return s.client.LookupIPv4(domain)
	}

	if option.IPv6Enable {
		return s.client.LookupIPv6(domain)
	}

	return nil, errors.New("neither IPv4 nor IPv6 is enabled")
}

func (s *localNameServer) Name() string {
	return "localhost"
}

func NewLocalNameServer() *localNameServer {
	log.Println(errors.New("DNS: created localhost client"))
	return &localNameServer{
		client: localdns.New(),
	}
}