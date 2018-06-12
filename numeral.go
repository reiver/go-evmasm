package evmasm

import (
	"unicode/utf8"
)

// numeral parses the UTF-8 text in a []byte, trying to interpret it as numeral, and tries to turn those numerals into a number that fits into a byte.
//
// If what it finds are not all numerals, then it returns a NotNumeral error.
//
// If what it finds are numerals but represents a number larger than can fit in a byte (i.e., larger than 255), then it returns a errOverflow error.
//
// Example:
//
//	numeral("0") -> 0, nil
//
//	numeral("1") -> 1, nil
//
//	numeral("10") -> 10, nil
//
//	numeral("0x10") -> 16, nil
func numeral(text []byte) (byte, error) {

	p := text

	var r0 rune
	{
		var size int

		r0, size = utf8.DecodeRune(p)
		if utf8.RuneError == r0 {
			return 0, internalNotNumeral{string(text)}
		}
		p = p[size:]

		if 0 >= len(p) {
			switch r0 {
			case '0','1','2','3','4','5','6','7','8','9':
				x := byte(r0) - '0'
				return x, nil

			default:
				return 0, internalNotNumeral{string(text)}
			}
		}
	}

	var r1 rune
	{
		var size int

		r1, size = utf8.DecodeRune(p)
		if utf8.RuneError == r1 {
			return 0, internalNotNumeral{string(text)}
		}
		p = p[size:]

		if '0' == r0 && 'x' != r1 {
			return 0, internalNotNumeral{string(text)}
		}

		if '0' == r0 && 'x' == r1 && 0 >= len(p) {
			return 0, internalNotNumeral{string(text)}
		}

		if 0 >= len(p) {
			switch r1 {
			case '0','1','2','3','4','5','6','7','8','9':
				x := (10*(byte(r0)-'0')) + (byte(r1) - '0')

				return x, nil
			default:
				return 0, internalNotNumeral{string(text)}
			}
		}
	}


	switch {
	case '0' == r0 && 'x' == r1:
		return hexadecimal(p)
	case '0' <= r0 && r0 <= '9' && '0' <= r1 && r1 <= '9':
		return decimal(text)
	default:
		return 0, internalNotNumeral{string(text)}
	}
}
