package evmasm

import (
	"fmt"
)

type NotNumeral interface {
	error
	NotNumeral()
}

type internalNotNumeral struct {
	value string
}

func (receiver internalNotNumeral) Error() string {
	return fmt.Sprintf("%q is not a numeral", receiver.value)
}

func (internalNotNumeral) NotNumeral() {
	// Nothing here.
}
