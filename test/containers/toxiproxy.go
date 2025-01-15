package containers

import (
	toxiproxy "github.com/Shopify/toxiproxy/v2/client"
)

// Upstream holds information about proxy.
type Upstream struct {
	proxy *toxiproxy.Proxy
}

// AddLatency add latency for given proxy connection.
func (p *Upstream) AddLatency(latency, jitter int) error {
	_, err := p.proxy.AddToxic(
		"latency_down",
		"latency",
		"downstream",
		1.0,
		toxiproxy.Attributes{
			"latency": latency,
			"jitter":  jitter,
		})
	return err
}
