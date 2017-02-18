package dns_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thn-dev/thn-go/commons/net/dns"
)

const (
	validDnsServer   = "8.8.8.8"
	invalidDnsServer = "8.8.8.9"
)

var validDomainName string
var validDomainNames []string
var validIPv4Array []string
var validIPv6Array []string

func init() {
	validDomainName = "www.google.com"
	validDomainNames = []string{"iad30s08-in-f4.1e100.net", "iad30s08-in-f132.1e100.net", "lga15s43-in-x04.1e100.net"}
	validIPv4Array = []string{"172.217.7.132", "172.217.3.36", "172.217.5.228"}
	validIPv6Array = []string{"2607:f8b0:4006:807::2004", "2607:f8b0:4004:805::2004"}
}

func TestLookupDomain_Invalid_DNSServer(t *testing.T) {
	validIP := "8.8.8.8"
	actual, err := dns.LookupDomain(invalidDnsServer, validIP)
	assert.NotNil(t, err, "expected an error")
	assert.Equal(t, 0, len(actual))
}

func TestLookupDomain_Valid_DNSServer(t *testing.T) {
	validIP := "8.8.8.8"
	expected := []string{"google-public-dns-a.google.com"}
	actual, err := dns.LookupDomain(validDnsServer, validIP)
	assert.Nil(t, err, "expected no error")
	assert.Equal(t, expected, actual)
}

func TestLookupDomain_IPv4(t *testing.T) {
	validIP := validIPv4Array[0]
	expected := validDomainNames

	actual, err := dns.LookupDomain(validDnsServer, validIP)
	assert.Nil(t, err, "expected no error")
	assert.Equal(t, expected, actual)
	assert.Contains(t, expected, actual[0])

	invalidIP := "172.30.192.a"
	actual, err = dns.LookupDomain(validDnsServer, invalidIP)
	assert.NotNil(t, err, "expected an error")
	assert.Equal(t, 0, len(actual))

	unknownIP := "172.30.192.120"
	actual, err = dns.LookupDomain(validDnsServer, unknownIP)
	assert.NotNil(t, err, "expected an error")
	assert.Equal(t, 0, len(actual))
}

func TestLookupDomain_IPv6(t *testing.T) {
	validShortIP := "2607:f8b0:4006:807::2004"
	expected := validDomainNames
	actual, err := dns.LookupDomain(validDnsServer, validShortIP)
	assert.Nil(t, err, "expected no error")
	assert.Contains(t, expected, actual[0])

	validLongIP := "2607:f8b0:4006:0807:0000:0000:0000:2004"
	actual, err = dns.LookupDomain(validDnsServer, validLongIP)
	assert.Nil(t, err, "expected no error")
	assert.Contains(t, expected, actual[0])

	invalidIP := "2607:f8b0:4006:0807:a::2004"
	actual, err = dns.LookupDomain(validDnsServer, invalidIP)
	assert.NotNil(t, err, "expected an error")
	assert.Equal(t, 0, len(actual))
}

func TestLookupIP(t *testing.T) {
	expectedIPv4 := validIPv4Array

	actualIPv4, err := dns.LookupIP(validDnsServer, validDomainName, false)
	assert.Nil(t, err, "expected no error")
	assert.Equal(t, 1, len(actualIPv4))
	assert.Contains(t, expectedIPv4, actualIPv4[0])

	expectedIPv6 := validIPv6Array

	actualIPv6, err := dns.LookupIP(validDnsServer, validDomainName, true)
	assert.Nil(t, err, "expected no error")
	assert.Equal(t, 1, len(actualIPv6))
	assert.Contains(t, expectedIPv6, actualIPv6[0])
}

func TestLookupIPv4_Valid(t *testing.T) {
	expectedIPv4 := validIPv4Array

	actualIPv4, err := dns.LookupIP(validDnsServer, validDomainName, false)
	assert.Nil(t, err, "expected no error")
	assert.Equal(t, 1, len(actualIPv4))
	assert.Contains(t, expectedIPv4, actualIPv4[0])
}

func TestLookupIPv6_Valid(t *testing.T) {
	expectedIPv6 := validIPv6Array

	actualIPv6, err := dns.LookupIP(validDnsServer, validDomainName, true)
	assert.Nil(t, err, "expected no error")
	assert.Equal(t, 1, len(actualIPv6))
	assert.Contains(t, expectedIPv6, actualIPv6[0])
}
