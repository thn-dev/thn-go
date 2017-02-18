package main

import (
	"fmt"
	"net"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	common_ip "github.com/thn-dev/thn-go/commons/net/ip"
)

var (
	app = kingpin.New("iptools", "IP Tools")

	ip2hex   = app.Command("ip2hex", "Convert an IP address to its hex representation")
	ip2hexIP = ip2hex.Arg("ip", "IP address").Required().IP()

	hex2ip      = app.Command("hex2ip", "Convert a hex value to an IP address")
	hex2ipValue = hex2ip.Arg("hex", "Hex Representation of an IP").Required().HexBytes()
)

func main() {
	inputCommand := kingpin.MustParse(app.Parse(os.Args[1:]))
	switch inputCommand {

	case ip2hex.FullCommand():
		convertIPToHexString(*ip2hexIP)
		break

	case hex2ip.FullCommand():
		convertHexToIP(*hex2ipValue)
		break
	}
}

func convertIPToHexString(ip net.IP) {
	var IPHexString string

	if ip.To4() != nil {
		IPHexString = common_ip.ByteArrayToHexString(ip.To4())
	} else {
		IPHexString = common_ip.ByteArrayToHexString(ip.To16())
	}

	fmt.Printf("IP(%s) -> HEX(%s)\n", ip.String(), IPHexString)
}

func convertHexToIP(hex []byte) {
	var IP net.IP
	IP = hex

	fmt.Printf("HEX(%s) -> IP(%s)\n", common_ip.ByteArrayToHexString(hex), IP.String())
}
