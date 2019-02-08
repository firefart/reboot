package logger

import "log"

// CLILogger is a Log implementation for testing purposes under non windows OS
type CLILogger struct{}

// Close does nothing
func (CLILogger) Close() error {
	return nil
}

// Info prints a info message
func (CLILogger) Info(eid uint32, msg string) error {
	log.Printf("[INFO] %s", msg)
	return nil
}

// Warning prints a warning message
func (CLILogger) Warning(eid uint32, msg string) error {
	log.Printf("[WARNING] %s", msg)
	return nil
}

// Error prints an error message
func (CLILogger) Error(eid uint32, msg string) error {
	log.Printf("[ERROR] %s", msg)
	return nil
}
