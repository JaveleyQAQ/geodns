package query

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/JaveleyQAQ/geodns/internal/config"
	"github.com/miekg/dns"
)

type DNSQuery struct {
	provider *config.DNSProvider
}

func NewDNSQuery() *DNSQuery {
	return &DNSQuery{
		provider: config.Provider,
	}
}

func (dq *DNSQuery) BuildQuery(domain string, recordType uint16) string {
	m := new(dns.Msg)
	m.Id = dns.Id()
	m.RecursionDesired = true
	m.Question = make([]dns.Question, 1)
	m.Question[0] = dns.Question{
		Name:   dns.Fqdn(domain),
		Qtype:  recordType,
		Qclass: dns.ClassINET,
	}

	buf, err := m.Pack()
	if err != nil {
		log.Fatalf("Error packing DNS query: %v", err)
	}
	return base64.RawURLEncoding.EncodeToString(buf)
}

func RandFloat() float64 {
	return float64(time.Now().UnixNano()%1e9) / 1e9
}
