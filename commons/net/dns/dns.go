package dns

import (
	"strings"

	"github.com/miekg/dns"
	"github.com/pkg/errors"

	common_net "github.com/thn-dev/thn-go/commons/net"
)

// -----------------------------------------------------------------------------
// References
// - https://github.com/corny/dnscheck/blob/master/lookup.go
// -----------------------------------------------------------------------------

// LookupDomain - looks up domain name of given IP address
func LookupDomain(dnsServer string, ipAddress string) ([]string, error) {
	reverse, err := dns.ReverseAddr(ipAddress)
	if err == nil {
		dnsMessage := &dns.Msg{}
		dnsMessage.SetQuestion(reverse, dns.TypePTR)
		return lookup(dnsMessage, dnsServer)
	}

	return []string{}, errors.Wrapf(err, "No domains found for IP(%s)", ipAddress)
}

// LookupIP - returns IP address(es) for given domain name
func LookupIP(dnsServer string, domainName string, returnIPv6 bool) ([]string, error) {
	if returnIPv6 {
		return LookupIPv6(dnsServer, domainName)
	}

	return LookupIPv4(dnsServer, domainName)
}

// LookupIPv4 - return IPv4 address(es) for given domain name
func LookupIPv4(dnsServer string, domainName string) ([]string, error) {
	dnsMessage := &dns.Msg{}
	dnsMessage.SetQuestion(dns.Fqdn(domainName), dns.TypeA)
	return lookup(dnsMessage, dnsServer)
}

// LookupIPv6 - returns IPv6 address(es) for given domain name
func LookupIPv6(dnsServer string, domainName string) ([]string, error) {
	dnsMessage := &dns.Msg{}
	dnsMessage.SetQuestion(dns.Fqdn(domainName), dns.TypeAAAA)
	return lookup(dnsMessage, dnsServer)
}

// lookup - lookups for info requested in dns.Msg
func lookup(dnsMessage *dns.Msg, dnsServer string) ([]string, error) {
	dnsClient := &dns.Client{}

	// lookup for info
	info, _, errClient := dnsClient.Exchange(dnsMessage, formatDnsServer(dnsServer))
	if errClient != nil {
		return []string{}, errors.Wrap(errClient, "Unable to lookup for information")
	}

	if info == nil || info.Rcode != dns.RcodeSuccess {
		return []string{}, errors.New("No information found")
	}

	var records []string

	// DNS Resource Record
	for _, dnsRR := range info.Answer {
		switch dnsMessage.Question[0].Qtype {

		// handle IPv4 type
		case dns.TypeA:
			if record, ok := dnsRR.(*dns.A); ok {
				records = append(records, record.A.String())
			}

		// handle IPv6 type
		case dns.TypeAAAA:
			if record, ok := dnsRR.(*dns.AAAA); ok {
				records = append(records, record.AAAA.String())
			}

		// handle reverse-lookup type
		case dns.TypePTR:
			if record, ok := dnsRR.(*dns.PTR); ok {
				records = append(records, strings.TrimSuffix(record.Ptr, "."))
			}

		default:
			return nil, errors.Errorf("Unsupported type(%s)", dns.TypeToString[dnsMessage.Question[0].Qtype])
		}
	}

	return records, nil
}

func formatDnsServer(server string) string {
	return common_net.FormatServer(server, "53")
}
