package evmasm

import (
	"unicode/utf8"
)

// hexadecimal parses the UTF-8 text in a []byte, trying to interpret it as hexadecimal digits, and tries to turn those hexadecimal digits into a number that fits into a byte.
//
// If what it finds are not all hexadecimal digits, then it returns a errNotNumeral error.
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
//	hexadecimal( []byte("0x01") )  -> 0, errNotNumeral
//
//	hexadecimal( []byte("apple") ) -> 0, errNotNumeral
//
//	hexadecimal( []byte("") )      -> 0, errNotNumeral
//
//
//	hexadecimal( []byte("100") )   -> 0, errOverflow
//
//	hexadecimal( []byte("fff") )   -> 0, errOverflow
//
//	hexadecimal( []byte("1000") )  -> 0, errOverflow
func hexadecimal(text []byte) (byte, error) {
	if nil == text {
		return 0, errNotNumeral
	}

	if 0 >= len(text) {
		return 0, errNotNumeral
	}

	var num byte
	{
		var p []byte = text

		for  {
			r, size := utf8.DecodeRune(p)
			if utf8.RuneError == r {
				return 0, errNotNumeral
			}
			p = p[size:]

			var n byte
			switch r {
			case '0','1','2','3','4','5','6','7','8','9':
				n = byte(r - '0')

			case 'a','b','c','d','e','f':
				n = byte(r - 'a' + 10)

			case 'A','B','C','D','E','F':
				n = byte(r - 'A' + 10)

			default:
				return 0, errNotNumeral
			}

			old := num

			num *= 16
			if old > num {

				return 0, errOverflow
			}

			num += n
			if old > num {
				return 0, errOverflow
			}

			if 0 >= len(p) {
				break
			}
		}
	}

	return num, nil
}
