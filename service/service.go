// +build windows

package service

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/FireFart/reboot/server"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

var elog debug.Log

type myservice struct{}

// Execute is a method from the Handler interface. It's the main run method of a windows service
func (m *myservice) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	elog.Info(1, "reboot service started")
	p, err := servicePath()
	if err != nil {
		elog.Error(1, fmt.Sprintf("could not get service path: %v", err))
		changes <- svc.Status{State: svc.StopPending}
		errno = 1
		return
	}
	pwbyte, err := ioutil.ReadFile(fmt.Sprintf("%s\\password.conf", p))
	if err != nil {
		elog.Error(1, fmt.Sprintf("could not read password.conf: %v", err))
		changes <- svc.Status{State: svc.StopPending}
		errno = 1
		return
	}
	password := strings.TrimSpace(string(pwbyte))
	go server.Listen(1234, elog, password)
loop:
	for {
		select {
		case c := <-r:
			switch c.Cmd {
			// status update
			case svc.Interrogate:
				changes <- c.CurrentStatus
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			// stop service
			case svc.Stop, svc.Shutdown:
				break loop
			default:
				elog.Error(1, fmt.Sprintf("unexpected control request #%d", c))
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

// RunService runs a service. If debug mode is enabled all output will be sent to the terminal
func RunService(name string, isDebug bool) {
	var err error
	if isDebug {
		elog = debug.New(name)
	} else {
		elog, err = eventlog.Open(name)
		if err != nil {
			return
		}
	}
	defer elog.Close()

	elog.Info(1, fmt.Sprintf("starting %s service", name))
	run := svc.Run
	if isDebug {
		run = debug.Run
	}
	err = run(name, &myservice{})
	if err != nil {
		elog.Error(1, fmt.Sprintf("%s service failed: %v", name, err))
		return
	}
	elog.Info(1, fmt.Sprintf("%s service stopped", name))
}
