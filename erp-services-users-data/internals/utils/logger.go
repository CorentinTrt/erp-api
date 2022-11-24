// Package utils regroup all utilities with a general purpose
package utils

import (
	"log"
	"os"
)

// Log function helps you to to print line to the default output with a time prefix
func Log(msg string, sig os.Signal) {
	l := log.New(os.Stdout, "service-users-data - ", log.LstdFlags)

	l.Println(msg, sig)
}
