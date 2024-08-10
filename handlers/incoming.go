package handlers

import (
	"math/rand"
	"time"

	"github.com/miekg/dns"
	//"github.com/samuelmunyoki/ProxyDNS/firebase"
	"github.com/samuelmunyoki/ProxyDNS/utils"
)

type Forwarder struct {
	DNS           []string
	LastDomain     string
	LastDomainTime time.Time

}

func NewForwarder() *Forwarder {
	return &Forwarder{
		DNS:         []string{"8.8.8.8:53", "8.8.4.4:53", "1.1.1.1:53", "1.0.0.1:53"},
	}
}

func (f *Forwarder) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	c := new(dns.Client)
	c.Net = "udp"

	resp, _, err := c.Exchange(r, f.DNS[rand.Intn(len(f.DNS))])

	if err != nil {
		utils.Flog.Error(err)
		// fmt.Printf("Error forwarding DNS request: %v\n", err)
		return
	}

	for _, question := range r.Question {

		if question.Qtype == dns.TypeA {
			fmt.Printf("Question: %s\n", question.Name)
			for _, answer := range resp.Answer {
				switch answer.(type) {
				case *dns.A:
					if question.Name != f.LastDomain || time.Since(f.LastDomainTime) >= 5*time.Second {
						f.LastDomain = question.Name
						f.LastDomainTime = time.Now()
						//go firebase.AddLog(utils.ExtractClientIP(w), question.Name)
					}
				default:
				}
			}
		}
	}

	// Write the response back to the original client
	w.WriteMsg(resp)

}

type HandleUDPReq struct {
	Forwarder *Forwarder
}

func (h *HandleUDPReq) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	h.Forwarder.ServeDNS(w, r)
}

type HandleTCPReq struct {
	Forwarder *Forwarder
}

func (h *HandleTCPReq) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	h.Forwarder.ServeDNS(w, r)
}
