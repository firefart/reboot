//go:build !windows

package commands

import "context"

// Reboot is a dummy function
func Reboot(_ context.Context, _ string) []byte {
	return nil
}
