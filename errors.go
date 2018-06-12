package evmasm

import (
	"errors"
)

var (
	errInternalError = errors.New("Internal Error")
	errNilReader     = errors.New("Nil Reader")
	errNilReceiver   = errors.New("Nil Receiver")
	errNilWriter     = errors.New("Nil Writer")
	errOverflow      = errors.New("Overflow")
	errReaderFound   = errors.New("Reader Found")
)
