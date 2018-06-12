package evmasm

import (
	"testing"
)

func TestNotNumeral(t *testing.T) {

	var x NotNumeral = internalNotNumeral{} // THIS IS THAT PART THAT ACTUALLY MATTERS.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
