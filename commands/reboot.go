//go:build !windows

package commands

import "context"

// Reboot is a dummy function
func Reboot(ctx context.Context, initiator string) []byte {
	return nil
}
