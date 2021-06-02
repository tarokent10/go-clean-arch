package log

import (
	"fmt"
	"io"
	"log"
	"os"
	conf "picture-go-app/infrastructure/config"
)

type loglevel int

const (
	// DEBUG level
	ldebug loglevel = 0
	// INFO level
	linfo loglevel = 1
	// WARN level
	lwarn loglevel = 2
	// ERROR level
	lerror loglevel = 3
)

// Logger instance
type logger struct {
	level loglevel
	out   io.Writer
}

var singleton *logger = &logger{}

// TODO logrusに変更したい
func init() {
	c := conf.Get()
	singleton.level = loglevel(c.LOGLEVEL)
	if file := c.LOGFILE; len(file) == 0 {
		singleton.out = os.Stdout
	} else {
		//TODO 書き込むたびにOpenすべきかも（このままじゃファイルディスクリプタが枯渇する）
		if f, err := os.OpenFile(file, os.O_SYNC, 0666); err == nil {
			singleton.out = f
		} else {
			fmt.Printf("failed to open log file:%s\n", file)
		}
	}
}

// Debug log
func Debug(s string) error {
	var err error
	if singleton.level <= ldebug {
		_, err = fmt.Fprintf(singleton.out, "debug: %s\n", s)
	}
	return err
}

// Info log
func Info(s string) error {
	var err error
	if singleton.level <= linfo {
		_, err = fmt.Fprintf(singleton.out, "info: %s\n", s)
	}
	return err
}

// Warn warn log
func Warn(s string) error {
	var err error
	if singleton.level <= lwarn {
		_, err = fmt.Fprintf(singleton.out, "warn: %s\n", s)
	}
	return err
}

// Err Error log
func Err(s string) error {
	var err error
	if singleton.level <= lerror {
		_, err = fmt.Fprintf(singleton.out, "error: %s\n", s)
	}
	return err
}

// Err Error log
func Fatal(s string) error {
	var err error
	if singleton.level <= lerror {
		log.Fatal(fmt.Sprintf("fatal: %s\n", s))
	}
	return err
}
