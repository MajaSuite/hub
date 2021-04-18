package app

import "errors"

var (
	ErrNotConnected = errors.New("not connected to network")
	ErrNotParsed = errors.New("can't parse data")
)
