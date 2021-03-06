// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package syslog provides a simple interface to the system log service. It can
// send messages to the syslog daemon using UNIX domain sockets, UDP or TCP.
//
// Only one call to Dial is necessary. On write failures, the syslog client will
// attempt to reconnect to the server and write again.
//
// Package syslog provides a simple interface to the system log service.
//
// Package syslog provides a simple interface to the system log service.
package syslog

// NewLogger creates a log.Logger whose output is written to the system log service
// with the specified priority. The logFlag argument is the flag set passed through
// to log.New to create the Logger.
func NewLogger(p Priority, logFlag int) (*log.Logger, error)

// The Priority is a combination of the syslog facility and severity. For example,
// LOG_ALERT | LOG_FTP sends an alert severity message from the FTP facility. The
// default severity is LOG_EMERG; the default facility is LOG_KERN.
type Priority int

const (

	// From /usr/include/sys/syslog.h.
	// These are the same on Linux, BSD, and OS X.
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

const (

	// From /usr/include/sys/syslog.h.
	// These are the same up to LOG_FTP on Linux, BSD, and OS X.
	LOG_KERN Priority = iota << 3
	LOG_USER
	LOG_MAIL
	LOG_DAEMON
	LOG_AUTH
	LOG_SYSLOG
	LOG_LPR
	LOG_NEWS
	LOG_UUCP
	LOG_CRON
	LOG_AUTHPRIV
	LOG_FTP
	_ // unused
	_ // unused
	_ // unused
	_ // unused
	LOG_LOCAL0
	LOG_LOCAL1
	LOG_LOCAL2
	LOG_LOCAL3
	LOG_LOCAL4
	LOG_LOCAL5
	LOG_LOCAL6
	LOG_LOCAL7
)

// A Writer is a connection to a syslog server.
type Writer struct {
	// contains filtered or unexported fields
}

// Dial establishes a connection to a log daemon by connecting to address raddr on
// the specified network. Each write to the returned writer sends a log message
// with the given facility, severity and tag. If network is empty, Dial will
// connect to the local syslog server.
func Dial(network, raddr string, priority Priority, tag string) (*Writer, error)

// New establishes a new connection to the system log daemon. Each write to the
// returned writer sends a log message with the given priority and prefix.
func New(priority Priority, tag string) (w *Writer, err error)

// Alert logs a message with severity LOG_ALERT, ignoring the severity passed to
// New.
func (w *Writer) Alert(m string) (err error)

// Close closes a connection to the syslog daemon.
func (w *Writer) Close() error

// Crit logs a message with severity LOG_CRIT, ignoring the severity passed to New.
func (w *Writer) Crit(m string) (err error)

// Debug logs a message with severity LOG_DEBUG, ignoring the severity passed to
// New.
func (w *Writer) Debug(m string) (err error)

// Emerg logs a message with severity LOG_EMERG, ignoring the severity passed to
// New.
func (w *Writer) Emerg(m string) (err error)

// Err logs a message with severity LOG_ERR, ignoring the severity passed to New.
func (w *Writer) Err(m string) (err error)

// Info logs a message with severity LOG_INFO, ignoring the severity passed to New.
func (w *Writer) Info(m string) (err error)

// Notice logs a message with severity LOG_NOTICE, ignoring the severity passed to
// New.
func (w *Writer) Notice(m string) (err error)

// Warning logs a message with severity LOG_WARNING, ignoring the severity passed
// to New.
func (w *Writer) Warning(m string) (err error)

// Write sends a log message to the syslog daemon.
func (w *Writer) Write(b []byte) (int, error)
