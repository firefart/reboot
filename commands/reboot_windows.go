//go:build windows

package commands

import (
	"context"
	"fmt"
	"os/exec"
)

// Reboot executes a reboot on a windows based system
func Reboot(ctx context.Context, initiator string) []byte {
	path := "C:\\WINDOWS\\system32\\shutdown.exe"
	args := []string{
		"/g",
		"/t", "5",
		"/f",
		"/c", fmt.Sprintf("reboot via reboot service by %s", initiator),
	}

	out, err := exec.CommandContext(ctx, path, args...).CombinedOutput()
	if err != nil {
		return []byte(fmt.Sprintf("error on running reboot: %v Output: %q", err, out))
	}
	return out
}
