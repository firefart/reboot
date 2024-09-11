//go:build windows

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/firefart/reboot/service"
	"golang.org/x/sys/windows/svc"
)

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command>\n"+
			"       where <command> is one of\n"+
			"       install, remove, debug, start, stop, pause or continue.\n",
		errmsg, os.Args[0])
	os.Exit(1)
}

func main() {
	const svcName = "reboot"

	isService, err := svc.IsWindowsService()
	if err != nil {
		log.Fatalf("failed to determine if we are running as a service: %v", err)
	}
	if !isService {
		service.RunService(svcName, false)
		return
	}

	if len(os.Args) < 2 {
		usage("no command specified")
	}

	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "debug":
		service.RunService(svcName, true)
		return
	case "install":
		err = service.InstallService(svcName, "reboot service listens for incoming network reboot requests")
	case "remove":
		err = service.RemoveService(svcName)
	case "start":
		err = service.StartService(svcName)
	case "stop":
		err = service.ControlService(svcName, svc.Stop, svc.Stopped)
	default:
		usage(fmt.Sprintf("invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, svcName, err)
	}
}
