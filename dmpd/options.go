package dmpd

import (
	"time"
)

type Options struct {
	TCPAddress string

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewOpts() *Options {
	return &Options{
		TCPAddress:   "0.0.0.0:4150",
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}
}
