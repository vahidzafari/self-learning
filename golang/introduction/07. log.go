package main

import (
	"log"
	"log/syslog"
)

// Generally speaking, using a log file to write some information used to be considered
// a better practice than writing the same output on screen for two reasons: firstly,
// because the output does not get lost as it is stored on a file, and secondly, because
// you can search and process log files using UNIX tools, such as grep(1), awk(1), and
// sed(1), which cannot be done when messages are printed on a terminal window.

// The UNIX logging service has support for two properties named logging level and
// logging facility. The logging level is a value that specifies the severity of the log
// entry. There are various logging levels, including debug, info, notice, warning, err,
// crit, alert, and emerg, in reverse order of severity. The logging facility is like
// a category used for logging information. The value of the logging facility part can be
// one of auth, authpriv, cron, daemon, kern, lpr, mail, mark, news, syslog, user, UUCP,
// local0, local1, local2, local3, local4, local5, local6, or local7 and is defined
// inside /etc/syslog.conf, /etc/rsyslog.conf, or another appropriate file depending
// on the server process used for system logging on your UNIX machine.

// The log package sends log messages to standard error. Part of the log package is the
// log/syslog package, which allows you to send log messages to the syslog server
// of your machine. Although by default log writes to standard error, the use of log.
// SetOutput() modifies that behavior. The list of functions for sending logging data
// includes log.Printf(), log.Print(), log.Println(), log.Fatalf(), log.Fatalln(),
// log.Panic(), log.Panicln() and log.Panicf().

func main() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")
	if err != nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(sysLog)
		log.Print("Everything is fine!")
	}
}

// if you execute journalctl -xe on a Linux machine, you can see the
