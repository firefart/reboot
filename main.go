//go:build !windows

package main

import (
	"github.com/firefart/reboot/logger"
	"github.com/firefart/reboot/server"
)

func main() {
	var l logger.CLILogger
	// nolint: errcheck
	l.Warning(1, "This program only supports windows")
	// nolint: errcheck
	l.Warning(1, "Entering testing mode")
	server.Listen(1234, l, "password")
}
