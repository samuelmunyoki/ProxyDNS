package utils

import (
	"net"

	"github.com/miekg/dns"
)

// helper function to get keys from a map
func Keys(m map[string]struct{}) []string {
	var result []string
	for k := range m {
		result = append(result, k)
	}
	return result
}

func ExtractClientIP(w dns.ResponseWriter) string {
	switch addr := w.RemoteAddr().(type) {
	case *net.UDPAddr:
		return addr.IP.String()
	case *net.TCPAddr:
		return addr.IP.String()
	default:
		return "unknown"
	}
}