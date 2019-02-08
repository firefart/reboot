// +build !windows

package main

import (
	"github.com/FireFart/reboot/logger"
	"github.com/FireFart/reboot/server"
)

func main() {
	var l logger.CLILogger
	l.Warning(1, "This program only supports windows")
	l.Warning(1, "Entering testing mode")
	server.Listen(1234, l, "password")
}
