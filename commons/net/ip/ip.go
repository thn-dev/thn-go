package ip

import (
	"encoding/hex"
	"net"

	"github.com/thn-dev/thn-go/commons"
)

// ByteArrayToString - converts from a byte array to a human-readable string
func ByteArrayToString(baValue []byte) string {
	if len(baValue) == net.IPv6len || len(baValue) == net.IPv4len {
		var IP net.IP
		IP = baValue

		if IPv4 := IP.To4(); IPv4 != nil {
			return IPv4.String()
		}

		if IPv6 := IP.To16(); IPv6 != nil {
			return IPv6.String()
		}
	}

	return commons.Blank
}

// ByteArrayToHexString - converts a byte array to a hex representation string
func ByteArrayToHexString(baValue []byte) string {
	if len(baValue) == net.IPv6len || len(baValue) == net.IPv4len {
		return hex.EncodeToString(baValue)
	}

	return commons.Blank
}

// StringToByteArray - converts a human-readable IP string to a byte array
func StringToByteArray(strValue string) []byte {
	if IP := net.ParseIP(strValue); IP != nil {
		if IPv4 := IP.To4(); IPv4 != nil {
			return IPv4
		}

		if IPv6 := IP.To16(); IPv6 != nil {
			return IPv6
		}
	}

	return []byte(commons.Blank)
}

// StringToHexString - converts a human-readable IP string to a hex representation string
func StringToHexString(strValue string) string {
	return ByteArrayToHexString(StringToByteArray(strValue))
}
