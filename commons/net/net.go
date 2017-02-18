package net

import (
	"net"
	"strings"
)

// FormatServer - formats server with a default port number if no port
// number is included
func FormatServer(host string, defaultPort string) string {
	server := host

	// add default port number to the server
	if strings.ContainsAny(server, ":") != true {
		server = net.JoinHostPort(server, defaultPort)
	}

	return server
}
