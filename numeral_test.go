package evmasm

import (
	"testing"
)

func TestNumeral(t *testing.T) {

	tests := []struct{
		Text   []byte
		Expected byte
	}{
		{
			Text: []byte("0x00"),
			Expected:     0x00,
		},
		{
			Text: []byte("0x01"),
			Expected:     0x01,
		},
		{
			Text: []byte("0x02"),
			Expected:     0x02,
		},
		{
			Text: []byte("0x03"),
			Expected:     0x03,
		},
		{
			Text: []byte("0x04"),
			Expected:     0x04,
		},
		{
			Text: []byte("0x05"),
			Expected:     0x05,
		},
		{
			Text: []byte("0x06"),
			Expected:     0x06,
		},
		{
			Text: []byte("0x07"),
			Expected:     0x07,
		},
		{
			Text: []byte("0x08"),
			Expected:     0x08,
		},
		{
			Text: []byte("0x09"),
			Expected:     0x09,
		},



		{
			Text: []byte("0x0a"),
			Expected:     0x0a,
		},
		{
			Text: []byte("0x0A"),
			Expected:     0x0a,
		},



		{
			Text: []byte("0x0b"),
			Expected:     0x0b,
		},
		{
			Text: []byte("0x0B"),
			Expected:     0x0b,
		},



		{
			Text: []byte("0x0c"),
			Expected:     0x0c,
		},
		{
			Text: []byte("0x0C"),
			Expected:     0x0c,
		},



		{
			Text: []byte("0x0d"),
			Expected:     0x0d,
		},
		{
			Text: []byte("0x0D"),
			Expected:     0x0d,
		},



		{
			Text: []byte("0x0e"),
			Expected:     0x0e,
		},
		{
			Text: []byte("0x0E"),
			Expected:     0x0e,
		},



		{
			Text: []byte("0x0f"),
			Expected:     0x0f,
		},
		{
			Text: []byte("0x0F"),
			Expected:     0x0f,
		},



		{
			Text: []byte("0x10"),
			Expected:     0x10,
		},



		{
			Text: []byte("0x11"),
			Expected:     0x11,
		},



		{
			Text: []byte("0x12"),
			Expected:     0x12,
		},



		{
			Text: []byte("0x13"),
			Expected:     0x13,
		},



		{
			Text: []byte("0x14"),
			Expected:     0x14,
		},



		{
			Text: []byte("0x15"),
			Expected:     0x15,
		},



		{
			Text: []byte("0x16"),
			Expected:     0x16,
		},



		{
			Text: []byte("0x17"),
			Expected:     0x17,
		},



		{
			Text: []byte("0x18"),
			Expected:     0x18,
		},



		{
			Text: []byte("0x19"),
			Expected:     0x19,
		},



		{
			Text: []byte("0x1a"),
			Expected:     0x1a,
		},
		{
			Text: []byte("0x1A"),
			Expected:     0x1a,
		},



		{
			Text: []byte("0x22"),
			Expected:     0x22,
		},
		{
			Text: []byte("0x33"),
			Expected:     0x33,
		},
		{
			Text: []byte("0x44"),
			Expected:     0x44,
		},



		{
			Text: []byte("0"),
			Expected:     0,
		},
		{
			Text: []byte("1"),
			Expected:     1,
		},
		{
			Text: []byte("2"),
			Expected:     2,
		},
		{
			Text: []byte("3"),
			Expected:     3,
		},
		{
			Text: []byte("4"),
			Expected:     4,
		},
		{
			Text: []byte("5"),
			Expected:     5,
		},
		{
			Text: []byte("6"),
			Expected:     6,
		},
		{
			Text: []byte("7"),
			Expected:     7,
		},
		{
			Text: []byte("8"),
			Expected:     8,
		},
		{
			Text: []byte("9"),
			Expected:     9,
		},
		{
			Text: []byte("10"),
			Expected:     10,
		},
		{
			Text: []byte("11"),
			Expected:     11,
		},
		{
			Text: []byte("12"),
			Expected:     12,
		},
		{
			Text: []byte("13"),
			Expected:     13,
		},
		{
			Text: []byte("14"),
			Expected:     14,
		},
		{
			Text: []byte("15"),
			Expected:     15,
		},
		{
			Text: []byte("16"),
			Expected:     16,
		},
		{
			Text: []byte("17"),
			Expected:     17,
		},
		{
			Text: []byte("18"),
			Expected:     18,
		},
		{
			Text: []byte("19"),
			Expected:     19,
		},
		{
			Text: []byte("20"),
			Expected:     20,
		},
		{
			Text: []byte("21"),
			Expected:     21,
		},
		{
			Text: []byte("22"),
			Expected:     22,
		},
		{
			Text: []byte("23"),
			Expected:     23,
		},
		{
			Text: []byte("24"),
			Expected:     24,
		},
		{
			Text: []byte("25"),
			Expected:     25,
		},
		{
			Text: []byte("26"),
			Expected:     26,
		},
		{
			Text: []byte("27"),
			Expected:     27,
		},
		{
			Text: []byte("28"),
			Expected:     28,
		},
		{
			Text: []byte("29"),
			Expected:     29,
		},
		{
			Text: []byte("30"),
			Expected:     30,
		},
		{
			Text: []byte("31"),
			Expected:     31,
		},
		{
			Text: []byte("32"),
			Expected:     32,
		},
		{
			Text: []byte("33"),
			Expected:     33,
		},
		{
			Text: []byte("34"),
			Expected:     34,
		},
		{
			Text: []byte("35"),
			Expected:     35,
		},
		{
			Text: []byte("36"),
			Expected:     36,
		},
		{
			Text: []byte("37"),
			Expected:     37,
		},
		{
			Text: []byte("38"),
			Expected:     38,
		},
		{
			Text: []byte("39"),
			Expected:     39,
		},
		{
			Text: []byte("40"),
			Expected:     40,
		},
		{
			Text: []byte("41"),
			Expected:     41,
		},
		{
			Text: []byte("42"),
			Expected:     42,
		},
		{
			Text: []byte("43"),
			Expected:     43,
		},
		{
			Text: []byte("44"),
			Expected:     44,
		},
		{
			Text: []byte("45"),
			Expected:     45,
		},
		{
			Text: []byte("46"),
			Expected:     46,
		},
		{
			Text: []byte("47"),
			Expected:     47,
		},
		{
			Text: []byte("48"),
			Expected:     48,
		},
		{
			Text: []byte("49"),
			Expected:     49,
		},
		{
			Text: []byte("50"),
			Expected:     50,
		},
		{
			Text: []byte("51"),
			Expected:     51,
		},
		{
			Text: []byte("52"),
			Expected:     52,
		},
		{
			Text: []byte("53"),
			Expected:     53,
		},
		{
			Text: []byte("54"),
			Expected:     54,
		},
		{
			Text: []byte("55"),
			Expected:     55,
		},
		{
			Text: []byte("56"),
			Expected:     56,
		},
		{
			Text: []byte("57"),
			Expected:     57,
		},
		{
			Text: []byte("58"),
			Expected:     58,
		},
		{
			Text: []byte("59"),
			Expected:     59,
		},



		{
			Text: []byte("111"),
			Expected:     111,
		},
		{
			Text: []byte("222"),
			Expected:     222,
		},



		{
			Text: []byte("250"),
			Expected:     250,
		},
		{
			Text: []byte("251"),
			Expected:     251,
		},
		{
			Text: []byte("252"),
			Expected:     252,
		},
		{
			Text: []byte("253"),
			Expected:     253,
		},
		{
			Text: []byte("254"),
			Expected:     254,
		},
		{
			Text: []byte("255"),
			Expected:     255,
		},
	}


	for testNumber, test := range tests {

		actual, err := numeral(test.Text)
		if nil != err {
			t.Errorf("For test #%d, (from %q) did not expect an error, but actually got one: (%T) %q", testNumber, test.Text, err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, (from %q) expected %d, but actually got %d.", testNumber, test.Text, expected, actual)
			continue
		}
	}
}
