package dmpd

import (
	"testing"
)

func TestTCPServer(t *testing.T) {
	opts := NewOpts()
	d := New(opts)
	d.Main()
}
