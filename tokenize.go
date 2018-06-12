package evmasm

import (
	"github.com/reiver/go-whitespace"

	"bufio"
	"io"
	"unicode/utf8"
)

type tokenize struct {
	reader     io.Reader
	bufreader *bufio.Reader
	eof        bool
}

func (receiver *tokenize) SetReader(reader io.Reader) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil != receiver.reader {
		return errReaderFound
	}

	if nil == reader {
		return errNilReader
	}
	receiver.reader = reader

	bufreader := bufio.NewReader(reader)
	if nil == bufreader {
		return errInternalError
	}
	receiver.bufreader = bufreader

	return nil
}

func (receiver *tokenize) Token(writer io.Writer) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == writer {
		return errNilWriter
	}

	bufreader := receiver.bufreader
	if nil == bufreader {
		return errNilReader
	}


	if receiver.eof {
		return io.EOF
	}


	var n uint64
	for {
		r, size, err := bufreader.ReadRune()
		if 0 > size {
			return errInternalError
		}
		if io.EOF == err {
			receiver.eof = true
			if 0 >= n {
				return io.EOF
			} else {
				return nil
			}
		}
		if nil != err {
			return err
		}

		if whitespace.IsWhitespace(r) {
			if 0 < n {
				return nil
			} else {
				continue
			}
		}

		var b [utf8.UTFMax]byte
		var p []byte
		{
			p = b[:]

			m := utf8.EncodeRune(p, r)

			p = p[:m]
		}

		{
			_, err := writer.Write(p)
			if nil != err {
				return err
			}
		}
		n += uint64(size)
	}
}
