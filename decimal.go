package evmasm

import (
	"unicode/utf8"
)

// decimal parses the UTF-8 text in a []byte, trying to interpret it as decimal digits (i.e., base 10 digits), and tries to turn those decimal digits into a number that fits into a byte.
//
// If what it finds are not all decimal digits, then it returns a errNotNumeral error.
//
// If what it finds are decimal digits but represents a number larger than can fit in a byte (i.e., larger than 255), then it returns a errOverflow error.
//
// Examples:
//
//	decimal( []byte("0") ) -> 0, nil
//
//	decimal( []byte("1") ) -> 1, nil
//
//	decimal( []byte("2") ) -> 2, nil
//
//	decimal( []byte("3") ) -> 3, nil
//
//
//	decimal( []byte("00") ) -> 0, nil
//
//	decimal( []byte("01") ) -> 1, nil
//
//	decimal( []byte("02") ) -> 2, nil
//
//	decimal( []byte("03") ) -> 3, nil
//
//
//	decimal( []byte("253") ) -> 253, nil
//
//	decimal( []byte("254") ) -> 254, nil
//
//	decimal( []byte("255") ) -> 255, nil
//
// Error Examples:
//
//	decimal( []byte("ff") )   -> 0, errNotNumeral
//
//	decimal( []byte("0x01") )  -> 0, errNotNumeral
//
//	decimal( []byte("apple") ) -> 0, errNotNumeral
//
//	decimal( []byte("") )      -> 0, errNotNumeral
//
//
//	decimal( []byte("256") )   -> 0, errOverflow
//
//	decimal( []byte("1000") )  -> 0, errOverflow
func decimal(text []byte) (byte, error) {
	if nil == text {
		return 0, errNotNumeral
	}

	if 0 >= len(text) {
		return 0, errNotNumeral
	}

	var num uint16
	{
		var p []byte = text

		for  {
			r, size := utf8.DecodeRune(p)
			if utf8.RuneError == r {
				return 0, errNotNumeral
			}
			p = p[size:]

			var n uint16
			switch r {
			case '0','1','2','3','4','5','6','7','8','9':
				n = uint16(r - '0')

			default:
				return 0, errNotNumeral
			}

			num *= 10
			if num > 255 {
				return 0, errOverflow
			}

			num += n
			if num > 255 {
				return 0, errOverflow
			}

			if 0 >= len(p) {
				break
			}
		}
	}
	if num > 255 {
		return 0, errOverflow
	}

	return byte(num), nil
}
