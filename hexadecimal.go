package evmasm

import (
	"unicode/utf8"
)

// hexadecimal parses the UTF-8 text in a []byte, trying to interpret it as hexadecimal digits (i.e., base 16 digits), and tries to turn those hexadecimal digits into a number that fits into a byte.
//
// If what it finds are not all hexadecimal digits, then it returns a NotNumeral error.
//
// If what it finds are hexadecimal digits but represents a number larger than can fit in a byte (i.e., larger than 255), then it returns a errOverflow error.
//
// Examples:
//
//	hexadecimal( []byte("0") ) -> 0, nil
//
//	hexadecimal( []byte("1") ) -> 1, nil
//
//	hexadecimal( []byte("2") ) -> 2, nil
//
//	hexadecimal( []byte("3") ) -> 3, nil
//
//
//	hexadecimal( []byte("a") ) -> 10, nil
//
//	hexadecimal( []byte("b") ) -> 11, nil
//
//	hexadecimal( []byte("c") ) -> 12, nil
//
//
//	hexadecimal( []byte("A") ) -> 10, nil
//
//	hexadecimal( []byte("B") ) -> 11, nil
//
//	hexadecimal( []byte("C") ) -> 12, nil
//
//
//	hexadecimal( []byte("00") ) -> 0, nil
//
//	hexadecimal( []byte("01") ) -> 1, nil
//
//	hexadecimal( []byte("02") ) -> 2, nil
//
//	hexadecimal( []byte("03") ) -> 3, nil
//
//
//	hexadecimal( []byte("0a") ) -> 10, nil
//
//	hexadecimal( []byte("0b") ) -> 11, nil
//
//	hexadecimal( []byte("0c") ) -> 12, nil
//
//
//	hexadecimal( []byte("0A") ) -> 10, nil
//
//	hexadecimal( []byte("0B") ) -> 11, nil
//
//	hexadecimal( []byte("0C") ) -> 12, nil
//
// Error Examples:
//
//	hexadecimal( []byte("0x01") )  -> 0, NotNumeral
//
//	hexadecimal( []byte("apple") ) -> 0, NotNumeral
//
//	hexadecimal( []byte("") )      -> 0, NotNumeral
//
//
//	hexadecimal( []byte("100") )   -> 0, errOverflow
//
//	hexadecimal( []byte("fff") )   -> 0, errOverflow
//
//	hexadecimal( []byte("1000") )  -> 0, errOverflow
func hexadecimal(text []byte) (byte, error) {
	if nil == text {
		return 0, internalNotNumeral{string(text)}
	}

	if 0 >= len(text) {
		return 0, internalNotNumeral{string(text)}
	}

	var num uint16
	{
		var p []byte = text

		for  {
			r, size := utf8.DecodeRune(p)
			if utf8.RuneError == r {
				return 0, internalNotNumeral{string(text)}
			}
			p = p[size:]

			var n uint16
			switch r {
			case '0','1','2','3','4','5','6','7','8','9':
				n = uint16(r - '0')

			case 'a','b','c','d','e','f':
				n = uint16(r - 'a' + 10)

			case 'A','B','C','D','E','F':
				n = uint16(r - 'A' + 10)

			default:
				return 0, internalNotNumeral{string(text)}
			}

			num *= 16
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
