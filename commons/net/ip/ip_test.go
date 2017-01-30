package ip_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thn-dev/thn-go/commons"
	. "github.com/thn-dev/thn-go/commons/net/ip"
)

func TestStringToByteArray_IPv4(t *testing.T) {
	// invalid IP
	validateInvalidStringToByteArray(t, "mnpq")
	validateInvalidStringToByteArray(t, "1")

	// IPv4 (loopback address)
	validateValidStringToByteArray(t, "127.0.0.1", net.IPv4len)

	// IPv4 (normal address)
	validateValidStringToByteArray(t, "192.168.1.1", net.IPv4len)
}

func TestStringToByteArray_IPv6(t *testing.T) {
	// invalid IPv6
	validateInvalidStringToByteArray(t, "1:2:3:4")
	validateInvalidStringToByteArray(t, "f:f:f")

	// IPv6 (loopback address)
	validateValidStringToByteArray(t, "::1", net.IPv6len)

	// IPv6 (short address)
	validateValidStringToByteArray(t, "fe80::f65c:89ff:fec5:fc0b", net.IPv6len)

	// IPv6 (long address)
	validateValidStringToByteArray(t, "2001:638:902:1:201:2ff:fee2:7596", net.IPv6len)
}

func validateValidStringToByteArray(t *testing.T, ipValue string, expectedLength int) {
	baValue := StringToByteArray(ipValue)
	assert.Equal(t, expectedLength, len(baValue))
}

func validateInvalidStringToByteArray(t *testing.T, ipValue string) {
	baValue := StringToByteArray(ipValue)
	assert.NotEqual(t, net.IPv4len, len(baValue))
	assert.NotEqual(t, net.IPv6len, len(baValue))
}

func TestByteArrayToString_IPv4(t *testing.T) {
	expectedIP := "127.0.0.1"
	validByteArray := StringToByteArray(expectedIP)
	actualIP := ByteArrayToString(validByteArray)
	assert.Equal(t, expectedIP, actualIP)

	invalidByteArray := []byte{0x20, 0x21, 0x22}
	actual := ByteArrayToString(invalidByteArray)
	assert.Equal(t, commons.Blank, actual)
}

func TestByteArrayToString_IPv6(t *testing.T) {
	// loopback address
	expectedIP := "::1"
	actualIP := ByteArrayToString(StringToByteArray(expectedIP))
	assert.Equal(t, expectedIP, actualIP)

	// short address
	expectedShortIP := "ef23::a38d:ffff:bfa2:af7e"
	actualShortIP := ByteArrayToString(StringToByteArray(expectedShortIP))
	assert.Equal(t, expectedShortIP, actualShortIP)

	// long address
	expectedLongIP := "3210:789:654:0:432:2ff:abbd:9876"
	actualLongIP := ByteArrayToString(StringToByteArray(expectedLongIP))
	assert.Equal(t, expectedLongIP, actualLongIP)

	invalidByteArray := []byte{0x20, 0x21, 0x22, 0x33, 0x55}
	actual := ByteArrayToString(invalidByteArray)
	assert.Equal(t, commons.Blank, actual)
}

func TestByteArrayToHexString_IPv4(t *testing.T) {
	validByteArray := []byte{0x51, 0x83, 0x43, 0x83}
	expectedHexString := "51834383"
	actualHexString := ByteArrayToHexString(validByteArray)
	assert.Equal(t, expectedHexString, actualHexString)

	invalidByteArray := []byte{0x20, 0x21, 0x22}
	actual := ByteArrayToHexString(invalidByteArray)
	assert.Equal(t, commons.Blank, actual)
}

func TestByteArrayToHexString_IPv6(t *testing.T) {
	validByteArray := []byte{0x20, 0x01, 0x06, 0x38, 0x09, 0x02, 0x00, 0x01, 0x02, 0x01, 0x02, 0xff, 0xfe, 0xe2, 0x75, 0x96}
	expectedHexString := "2001063809020001020102fffee27596"
	actualHexString := ByteArrayToHexString(validByteArray)
	assert.Equal(t, expectedHexString, actualHexString)

	invalidByteArray := []byte{0x20, 0x01}
	actual := ByteArrayToHexString(invalidByteArray)
	assert.Equal(t, commons.Blank, actual)
}

func TestStringToHexString_IPv4(t *testing.T) {
	assert.Equal(t, commons.Blank, StringToHexString("1234"))
	assert.Equal(t, commons.Blank, StringToHexString("abcd"))

	assert.Equal(t, "c0a80338", StringToHexString("192.168.3.56"))
	assert.Equal(t, "51834383", StringToHexString("81.131.67.131"))
}

func TestStringToHexString_IPv6(t *testing.T) {
	assert.Equal(t, commons.Blank, StringToHexString("1:2:3:4"))
	assert.Equal(t, commons.Blank, StringToHexString("fe80::859f:b56e:6625:859h"))

	assert.Equal(t, "fe80000000000000859fb56e6625859d", StringToHexString("fe80::859f:b56e:6625:859d"))
	assert.Equal(t, "2001063809020001020102fffee27596", StringToHexString("2001:638:902:1:201:2ff:fee2:7596"))
}
