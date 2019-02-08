package logger

// Log is a copy of https://godoc.org/golang.org/x/sys/windows/svc/debug#Log
// this hack is needed to make this also build under non windows
// because the whole package has a build constraint
type Log interface {
	Close() error
	Info(eid uint32, msg string) error
	Warning(eid uint32, msg string) error
	Error(eid uint32, msg string) error
}
