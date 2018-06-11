package evmasm

import (
	"errors"
)

var (
	errInternalError = errors.New("Internal Error")
	errNilReader     = errors.New("Nil Reader")
	errNilReceiver   = errors.New("Nil Receiver")
	errNilWriter     = errors.New("Nil Writer")
	errNotNumeral    = errors.New("Not Numeral")
	errOverflow      = errors.New("Overflow")
	errReaderFound   = errors.New("Reader Found")
)
