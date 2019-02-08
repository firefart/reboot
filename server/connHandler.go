package server

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/FireFart/reboot/commands"
	"github.com/FireFart/reboot/logger"
)

func handleConnection(c net.Conn, logger logger.Log, password string) {
	defer c.Close()
	initiator := c.RemoteAddr().String()
	logger.Info(1, fmt.Sprintf("Serving %s\n", initiator))

	reader := bufio.NewReader(c)
	command, err := reader.ReadString('\n')
	if err != nil {
		logger.Error(1, err.Error())
		return
	}

	temp := strings.TrimSpace(command)
	if temp == "REBOOT" {
		_, err = c.Write([]byte("Please enter password: "))
		if err != nil {
			logger.Error(1, err.Error())
			return
		}
		pw, err := reader.ReadString('\n')
		if err != nil {
			logger.Error(1, err.Error())
			return
		}
		pw = strings.TrimSpace(pw)
		if pw != password {
			return
		}
		_, err = c.Write([]byte("Initiating reboot...\n"))
		if err != nil {
			logger.Error(1, err.Error())
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		out := commands.Reboot(ctx, initiator)
		text := out
		if len(out) == 0 {
			text = []byte("Done\n")
		}
		logger.Info(1, fmt.Sprintf("%s", text))
		_, err = c.Write(text)
		if err != nil {
			logger.Error(1, err.Error())
			return
		}
	}
}
