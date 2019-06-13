package ldap

import (
	"log"

	"gopkg.in/asn1-ber.v1"
)

// Debugger is the interface that wraps the debug output methods
type Debugger interface {
	Printf(format string, args ...interface{})
	PrintPacket(packet *ber.Packet)
}

// DefaultDebugger is the default implementation of the debug output methods
type DefaultDebugger struct{}

// Printf write debug output
func (DefaultDebugger) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// PrintPacket write debug packet output
func (DefaultDebugger) PrintPacket(packet *ber.Packet) {
	ber.PrintPacket(packet)
}
