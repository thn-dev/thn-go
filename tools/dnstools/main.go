package main

import (
	"fmt"
	"net"
	"os"

	log "github.com/ccpaging/log4go"
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	common_dns "github.com/thn-dev/thn-go/commons/net/dns"
)

var (
	app = kingpin.New("dnstools", "DNS Lookup Tools")

	ip2domain    = app.Command("ip2domain", "Lookup for Domain Name(s)")
	ip2domainDns = ip2domain.Arg("dns", "DNS server").Required().String()
	ip2domainIP  = ip2domain.Arg("ip", "IP address").Required().IP()

	domain2ip       = app.Command("domain2ip", "Lookup for IP address(es")
	domain2ipDns    = domain2ip.Arg("dns", "DNS server").Required().String()
	domain2ipDomain = domain2ip.Arg("domain", "Domain Name").Required().String()
)

func main() {
	inputCommand := kingpin.MustParse(app.Parse(os.Args[1:]))
	switch inputCommand {

	case ip2domain.FullCommand():
		convertIPToDomain(*ip2domainDns, *ip2domainIP)
		break

	case domain2ip.FullCommand():
		convertDomainToIP(*domain2ipDns, *domain2ipDomain)
		break
	}
}

func convertIPToDomain(dnsServer string, ip net.IP) {
	domains, err := common_dns.LookupDomain(dnsServer, ip.String())
	if err != nil {
		log.Exit(err)
	}

	fmt.Printf("IP(%s) -> Domains(%v)\n", ip.String(), domains)
}

func convertDomainToIP(dnsServer string, domainName string) {
	ipv4List, err := common_dns.LookupIPv4(dnsServer, domainName)
	if err != nil {
		log.Exit(err)
	}

	ipv6List, err := common_dns.LookupIPv6(dnsServer, domainName)
	if err != nil {
		log.Exit(err)
	}

	fmt.Printf("Domain(%s)\n", domainName)

	if len(ipv4List) > 0 {
		fmt.Printf("-> IPv4(%v)\n", ipv4List)
	}

	if len(ipv6List) > 0 {
		fmt.Printf("-> IPv6(%v)\n", ipv6List)
	}
}
