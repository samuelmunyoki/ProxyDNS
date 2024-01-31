package handlers

import (
	"fmt"

	"github.com/miekg/dns"
)

type Forwarder struct {
	ForwardAddr string
}
func (f *Forwarder) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	c := new(dns.Client)
	c.Net = "udp"

	resp, _, err := c.Exchange(r, f.ForwardAddr)
	if err != nil {
		fmt.Printf("Error forwarding DNS request: %v\n", err)
		return
	}

	// Write the response back to the original client
	w.WriteMsg(resp)
}

type HandleUDPReq struct {
	Forwarder *Forwarder
}

func (h *HandleUDPReq) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	// fmt.Printf("UDP Req %+v\n\n", r)
	h.Forwarder.ServeDNS(w, r)
}

type HandleTCPReq struct {
	Forwarder *Forwarder
}

func (h *HandleTCPReq) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	// fmt.Printf("TCP Req %+v\n\n", r)
	h.Forwarder.ServeDNS(w, r)
}