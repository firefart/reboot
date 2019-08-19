package server

import (
	"fmt"
	"net"

	"github.com/FireFart/reboot/logger"
)

// Listen listens for incoming TCP connections
func Listen(port int, logger logger.Log, password string) {
	addr := fmt.Sprintf(":%d", port)
	// nolint: errcheck
	logger.Info(1, fmt.Sprintf("listening on %s\n", addr))
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		// nolint: errcheck
		logger.Error(1, err.Error())
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			// nolint: errcheck
			logger.Error(1, err.Error())
			continue
		}
		go handleConnection(c, logger, password)
	}
}
