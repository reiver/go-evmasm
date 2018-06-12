package evmasm

import (
	"bytes"
	"io"
	"strings"

	"testing"
)

func TestTokenize(t *testing.T) {

	tests := []struct{
		Code     string
		Expected []string
	}{
		{
			Code: "",
			Expected: []string{
			},
		},



		{
			Code: "\u0009", // horizontal tab
			Expected: []string{
			},
		},
		{
			Code: "\u000A", // line feed
			Expected: []string{
			},
		},
		{
			Code: "\u000B", // vertical tab
			Expected: []string{
			},
		},
		{
			Code: "\u000C", // form feed
			Expected: []string{
			},
		},
		{
			Code: "\u000D", // carriage return
			Expected: []string{
			},
		},
		{
			Code: "\u0020", // space
			Expected: []string{
			},
		},
		{
			Code: "\u0085", // next line
			Expected: []string{
			},
		},
		{
			Code: "\u00A0", // no-break space
			Expected: []string{
			},
		},
		{
			Code: "\u1680", // ogham space mark
			Expected: []string{
			},
		},
		{
			Code: "\u180E", // mongolian vowel separator
			Expected: []string{
			},
		},
		{
			Code: "\u2000", // en quad
			Expected: []string{
			},
		},
		{
			Code: "\u2001", // em quad
			Expected: []string{
			},
		},
		{
			Code: "\u2002", // en space
			Expected: []string{
			},
		},
		{
			Code: "\u2003", // em space
			Expected: []string{
			},
		},
		{
			Code: "\u2004", // three-per-em space
			Expected: []string{
			},
		},
		{
			Code: "\u2005", // four-per-em space
			Expected: []string{
			},
		},
		{
			Code: "\u2006", // six-per-em space
			Expected: []string{
			},
		},
		{
			Code: "\u2007", // figure space
			Expected: []string{
			},
		},
		{
			Code: "\u2008", // punctuation space
			Expected: []string{
			},
		},
		{
			Code: "\u2009", // thin space
			Expected: []string{
			},
		},
		{
			Code: "\u200A", // hair space
			Expected: []string{
			},
		},
		{
			Code: "\u2028", // line separator
			Expected: []string{
			},
		},
		{
			Code: "\u2029", // paragraph separator
			Expected: []string{
			},
		},
		{
			Code: "\u202F", // narrow no-break space
			Expected: []string{
			},
		},
		{
			Code: "\u205F", // medium mathematical space
			Expected: []string{
			},
		},
		{
			Code: "\u3000", // ideographic space
			Expected: []string{
			},
		},



		{
			Code: "push1 0x02 push1 0x01 add",
			Expected: []string{
				"push1",
				"0x02",
				"push1",
				"0x01",
				"add",
			},
		},
		{
			Code: "push1  0x02  push1  0x01  add",
			Expected: []string{
				"push1",
				"0x02",
				"push1",
				"0x01",
				"add",
			},
		},
		{
			Code: "push1 0x02 push1 0x01 add ",
			Expected: []string{
				"push1",
				"0x02",
				"push1",
				"0x01",
				"add",
			},
		},
		{
			Code: "push1  0x02  push1  0x01  add  ",
			Expected: []string{
				"push1",
				"0x02",
				"push1",
				"0x01",
				"add",
			},
		},
		{
			Code: " push1 0x02 push1 0x01 add",
			Expected: []string{
				"push1",
				"0x02",
				"push1",
				"0x01",
				"add",
			},
		},
		{
			Code: "  push1  0x02  push1  0x01  add",
			Expected: []string{
				"push1",
				"0x02",
				"push1",
				"0x01",
				"add",
			},
		},
		{
			Code: " push1 0x02 push1 0x01 add ",
			Expected: []string{
				"push1",
				"0x02",
				"push1",
				"0x01",
				"add",
			},
		},
		{
			Code: "  push1  0x02  push1  0x01  add  ",
			Expected: []string{
				"push1",
				"0x02",
				"push1",
				"0x01",
				"add",
			},
		},
		{
			Code: "\u1680 \u2000push1\u000A0x02\u000B\u000Bpush1\u00A00x01             add\u000D\u000A",
			Expected: []string{
				"push1",
				"0x02",
				"push1",
				"0x01",
				"add",
			},
		},



		{
			Code: "    one  2  three 0x4   FIVE     ",
			Expected: []string{
				"one",
				"2",
				"three",
				"0x4",
				"FIVE",
			},
		},



		{
			Code: "push1  0x01",
			Expected: []string{
				"push1",
				"0x01",
			},
		},
		{
			Code: "push2  0x01  0x02",
			Expected: []string{
				"push2",
				"0x01",
				"0x02",
			},
		},
		{
			Code: "push3  0x01  0x02  0x03",
			Expected: []string{
				"push3",
				"0x01",
				"0x02",
				"0x03",
			},
		},
		{
			Code: "push4  0x01  0x02  0x03  0x04",
			Expected: []string{
				"push4",
				"0x01",
				"0x02",
				"0x03",
				"0x04",
			},
		},
		{
			Code: "push5  0x01  0x02  0x03  0x04  0x05",
			Expected: []string{
				"push5",
				"0x01",
				"0x02",
				"0x03",
				"0x04",
				"0x05",
			},
		},
		{
			Code: "push6  0x01  0x02  0x03  0x04  0x05  0x06",
			Expected: []string{
				"push6",
				"0x01",
				"0x02",
				"0x03",
				"0x04",
				"0x05",
				"0x06",
			},
		},
		{
			Code: "push7  0x01  0x02  0x03  0x04  0x05  0x06  0x07",
			Expected: []string{
				"push7",
				"0x01",
				"0x02",
				"0x03",
				"0x04",
				"0x05",
				"0x06",
				"0x07",
			},
		},
	}


	TestLoop: for testNumber, test := range tests {

		var actualTokens []string
		{
			var tokenizer tokenize

			if err := tokenizer.SetReader( strings.NewReader(test.Code) ); nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
				continue TestLoop
			}

			TokenLoop: for tokenNumber:=0; true; tokenNumber++ {

				var buffer bytes.Buffer

				err := tokenizer.Token(&buffer)
				if io.EOF == err {
					break TokenLoop
				}
				if nil != err {
					t.Errorf("For test #%d and token #%d, did not expect an error, but actually got one: (%T) %q", testNumber, tokenNumber, err, err)
					continue TestLoop
				}

				actualTokens = append(actualTokens, buffer.String())
			}
		}

		if expected, actual := len(test.Expected), len(actualTokens); expected != actual {
			t.Errorf("For test #%d, expected %d tokens, but actually got %d tokens.", testNumber, expected, actual)
			t.Errorf("EXPECTED: %#v", test.Expected)
			t.Errorf("ACTUAL:   %#v", actualTokens)
			t.Errorf("CODE: %q", test.Code)
			t.Errorf("")
			continue TestLoop
		}

		for tokenNumber:=0; tokenNumber<len(test.Expected); tokenNumber++ {
			expected := test.Expected[tokenNumber]
			actual   := actualTokens[tokenNumber]

			if expected != actual {
				t.Errorf("For test #%d and token #%d, expected %q, but actually got %q.", testNumber, tokenNumber, expected, actual)
				t.Errorf("EXPECTED: %#v", test.Expected)
				t.Errorf("ACTUAL:   %#v", actualTokens)
				t.Errorf("CODE: %q", test.Code)
				t.Errorf("")
				continue
			}
		}
	}
}
