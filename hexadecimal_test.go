package evmasm

import (
	"testing"
)

func TestHexadecimal(t *testing.T) {

	tests := []struct{
		Text   []byte
		Expected byte
	}{
		{
			Text: []byte( "0"),
			Expected:   0x00,
		},
		{
			Text: []byte("00"),
			Expected:   0x00,
		},



		{
			Text: []byte( "1"),
			Expected:   0x01,
		},
		{
			Text: []byte("01"),
			Expected:   0x01,
		},



		{
			Text: []byte( "2"),
			Expected:   0x02,
		},
		{
			Text: []byte("02"),
			Expected:   0x02,
		},



		{
			Text: []byte( "3"),
			Expected:   0x03,
		},
		{
			Text: []byte("03"),
			Expected:   0x03,
		},



		{
			Text: []byte( "4"),
			Expected:   0x04,
		},
		{
			Text: []byte("04"),
			Expected:   0x04,
		},



		{
			Text: []byte( "5"),
			Expected:   0x05,
		},
		{
			Text: []byte("05"),
			Expected:   0x05,
		},



		{
			Text: []byte( "6"),
			Expected:   0x06,
		},
		{
			Text: []byte("06"),
			Expected:   0x06,
		},



		{
			Text: []byte( "7"),
			Expected:   0x07,
		},
		{
			Text: []byte("07"),
			Expected:   0x07,
		},



		{
			Text: []byte( "8"),
			Expected:   0x08,
		},
		{
			Text: []byte("08"),
			Expected:   0x08,
		},



		{
			Text: []byte( "9"),
			Expected:   0x09,
		},
		{
			Text: []byte("09"),
			Expected:   0x09,
		},



		{
			Text: []byte( "a"),
			Expected:   0x0a,
		},
		{
			Text: []byte( "A"),
			Expected:   0x0a,
		},
		{
			Text: []byte("0a"),
			Expected:   0x0a,
		},
		{
			Text: []byte("0A"),
			Expected:   0x0a,
		},



		{
			Text: []byte( "b"),
			Expected:   0x0b,
		},
		{
			Text: []byte( "B"),
			Expected:   0x0b,
		},
		{
			Text: []byte("0b"),
			Expected:   0x0b,
		},
		{
			Text: []byte("0B"),
			Expected:   0x0b,
		},



		{
			Text: []byte( "c"),
			Expected:   0x0c,
		},
		{
			Text: []byte( "C"),
			Expected:   0x0c,
		},
		{
			Text: []byte("0c"),
			Expected:   0x0c,
		},
		{
			Text: []byte("0C"),
			Expected:   0x0c,
		},



		{
			Text: []byte( "d"),
			Expected:   0x0d,
		},
		{
			Text: []byte( "D"),
			Expected:   0x0d,
		},
		{
			Text: []byte("0d"),
			Expected:   0x0d,
		},
		{
			Text: []byte("0D"),
			Expected:   0x0d,
		},



		{
			Text: []byte( "e"),
			Expected:   0x0e,
		},
		{
			Text: []byte( "E"),
			Expected:   0x0e,
		},
		{
			Text: []byte("0e"),
			Expected:   0x0e,
		},
		{
			Text: []byte("0E"),
			Expected:   0x0e,
		},



		{
			Text: []byte( "f"),
			Expected:   0x0f,
		},
		{
			Text: []byte( "F"),
			Expected:   0x0f,
		},
		{
			Text: []byte("0f"),
			Expected:   0x0f,
		},
		{
			Text: []byte("0F"),
			Expected:   0x0f,
		},



		{
			Text: []byte("10"),
			Expected:   0x10,
		},



		{
			Text: []byte("11"),
			Expected:   0x11,
		},



		{
			Text: []byte("12"),
			Expected:   0x12,
		},



		{
			Text: []byte("13"),
			Expected:   0x13,
		},



		{
			Text: []byte("14"),
			Expected:   0x14,
		},



		{
			Text: []byte("15"),
			Expected:   0x15,
		},



		{
			Text: []byte("16"),
			Expected:   0x16,
		},



		{
			Text: []byte("17"),
			Expected:   0x17,
		},



		{
			Text: []byte("18"),
			Expected:   0x18,
		},



		{
			Text: []byte("19"),
			Expected:   0x19,
		},



		{
			Text: []byte("1a"),
			Expected:   0x1a,
		},
		{
			Text: []byte("1A"),
			Expected:   0x1a,
		},



		{
			Text: []byte("1b"),
			Expected:   0x1b,
		},
		{
			Text: []byte("1B"),
			Expected:   0x1b,
		},



		{
			Text: []byte("1c"),
			Expected:   0x1c,
		},
		{
			Text: []byte("1C"),
			Expected:   0x1c,
		},



		{
			Text: []byte("1d"),
			Expected:   0x1d,
		},
		{
			Text: []byte("1D"),
			Expected:   0x1d,
		},



		{
			Text: []byte("1e"),
			Expected:   0x1e,
		},
		{
			Text: []byte("1E"),
			Expected:   0x1e,
		},



		{
			Text: []byte("1f"),
			Expected:   0x1f,
		},
		{
			Text: []byte("1F"),
			Expected:   0x1f,
		},



		{
			Text: []byte("20"),
			Expected:   0x20,
		},



		{
			Text: []byte("21"),
			Expected:   0x21,
		},



		{
			Text: []byte("22"),
			Expected:   0x22,
		},



		{
			Text: []byte("23"),
			Expected:   0x23,
		},



		{
			Text: []byte("24"),
			Expected:   0x24,
		},



		{
			Text: []byte("25"),
			Expected:   0x25,
		},



		{
			Text: []byte("26"),
			Expected:   0x26,
		},



		{
			Text: []byte("27"),
			Expected:   0x27,
		},



		{
			Text: []byte("28"),
			Expected:   0x28,
		},



		{
			Text: []byte("29"),
			Expected:   0x29,
		},



		{
			Text: []byte("2a"),
			Expected:   0x2a,
		},
		{
			Text: []byte("2A"),
			Expected:   0x2a,
		},



		{
			Text: []byte("2b"),
			Expected:   0x2b,
		},
		{
			Text: []byte("2B"),
			Expected:   0x2b,
		},



		{
			Text: []byte("2c"),
			Expected:   0x2c,
		},
		{
			Text: []byte("2C"),
			Expected:   0x2c,
		},



		{
			Text: []byte("2d"),
			Expected:   0x2d,
		},
		{
			Text: []byte("2D"),
			Expected:   0x2d,
		},



		{
			Text: []byte("2e"),
			Expected:   0x2e,
		},
		{
			Text: []byte("2E"),
			Expected:   0x2e,
		},



		{
			Text: []byte("2f"),
			Expected:   0x2f,
		},
		{
			Text: []byte("2F"),
			Expected:   0x2f,
		},



		{
			Text: []byte("30"),
			Expected:   0x30,
		},



		{
			Text: []byte("31"),
			Expected:   0x31,
		},



		{
			Text: []byte("32"),
			Expected:   0x32,
		},



		{
			Text: []byte("33"),
			Expected:   0x33,
		},



		{
			Text: []byte("34"),
			Expected:   0x34,
		},



		{
			Text: []byte("35"),
			Expected:   0x35,
		},



		{
			Text: []byte("36"),
			Expected:   0x36,
		},



		{
			Text: []byte("37"),
			Expected:   0x37,
		},



		{
			Text: []byte("38"),
			Expected:   0x38,
		},



		{
			Text: []byte("39"),
			Expected:   0x39,
		},



		{
			Text: []byte("3a"),
			Expected:   0x3a,
		},
		{
			Text: []byte("3A"),
			Expected:   0x3a,
		},



		{
			Text: []byte("3b"),
			Expected:   0x3b,
		},
		{
			Text: []byte("3B"),
			Expected:   0x3b,
		},



		{
			Text: []byte("3c"),
			Expected:   0x3c,
		},
		{
			Text: []byte("3C"),
			Expected:   0x3c,
		},



		{
			Text: []byte("3d"),
			Expected:   0x3d,
		},
		{
			Text: []byte("3D"),
			Expected:   0x3d,
		},



		{
			Text: []byte("3e"),
			Expected:   0x3e,
		},
		{
			Text: []byte("3E"),
			Expected:   0x3e,
		},



		{
			Text: []byte("3f"),
			Expected:   0x3f,
		},
		{
			Text: []byte("3F"),
			Expected:   0x3f,
		},



		{
			Text: []byte("40"),
			Expected:   0x40,
		},



		{
			Text: []byte("41"),
			Expected:   0x41,
		},



		{
			Text: []byte("42"),
			Expected:   0x42,
		},



		{
			Text: []byte("43"),
			Expected:   0x43,
		},



		{
			Text: []byte("44"),
			Expected:   0x44,
		},



		{
			Text: []byte("45"),
			Expected:   0x45,
		},



		{
			Text: []byte("46"),
			Expected:   0x46,
		},



		{
			Text: []byte("47"),
			Expected:   0x47,
		},



		{
			Text: []byte("48"),
			Expected:   0x48,
		},



		{
			Text: []byte("49"),
			Expected:   0x49,
		},



		{
			Text: []byte("4a"),
			Expected:   0x4a,
		},
		{
			Text: []byte("4A"),
			Expected:   0x4a,
		},



		{
			Text: []byte("4b"),
			Expected:   0x4b,
		},
		{
			Text: []byte("4B"),
			Expected:   0x4b,
		},



		{
			Text: []byte("4c"),
			Expected:   0x4c,
		},
		{
			Text: []byte("4C"),
			Expected:   0x4c,
		},



		{
			Text: []byte("4d"),
			Expected:   0x4d,
		},
		{
			Text: []byte("4D"),
			Expected:   0x4d,
		},



		{
			Text: []byte("4e"),
			Expected:   0x4e,
		},
		{
			Text: []byte("4E"),
			Expected:   0x4e,
		},



		{
			Text: []byte("4f"),
			Expected:   0x4f,
		},
		{
			Text: []byte("4F"),
			Expected:   0x4f,
		},



		{
			Text: []byte("50"),
			Expected:   0x50,
		},



		{
			Text: []byte("51"),
			Expected:   0x51,
		},



		{
			Text: []byte("52"),
			Expected:   0x52,
		},



		{
			Text: []byte("53"),
			Expected:   0x53,
		},



		{
			Text: []byte("54"),
			Expected:   0x54,
		},



		{
			Text: []byte("55"),
			Expected:   0x55,
		},



		{
			Text: []byte("56"),
			Expected:   0x56,
		},



		{
			Text: []byte("57"),
			Expected:   0x57,
		},



		{
			Text: []byte("58"),
			Expected:   0x58,
		},



		{
			Text: []byte("59"),
			Expected:   0x59,
		},



		{
			Text: []byte("5a"),
			Expected:   0x5a,
		},
		{
			Text: []byte("5A"),
			Expected:   0x5a,
		},



		{
			Text: []byte("5b"),
			Expected:   0x5b,
		},
		{
			Text: []byte("5B"),
			Expected:   0x5b,
		},



		{
			Text: []byte("5c"),
			Expected:   0x5c,
		},
		{
			Text: []byte("5C"),
			Expected:   0x5c,
		},



		{
			Text: []byte("5d"),
			Expected:   0x5d,
		},
		{
			Text: []byte("5D"),
			Expected:   0x5d,
		},



		{
			Text: []byte("5e"),
			Expected:   0x5e,
		},
		{
			Text: []byte("5E"),
			Expected:   0x5e,
		},



		{
			Text: []byte("5f"),
			Expected:   0x5f,
		},
		{
			Text: []byte("5F"),
			Expected:   0x5f,
		},



		{
			Text: []byte("60"),
			Expected:   0x60,
		},



		{
			Text: []byte("61"),
			Expected:   0x61,
		},



		{
			Text: []byte("62"),
			Expected:   0x62,
		},



		{
			Text: []byte("63"),
			Expected:   0x63,
		},



		{
			Text: []byte("64"),
			Expected:   0x64,
		},



		{
			Text: []byte("65"),
			Expected:   0x65,
		},



		{
			Text: []byte("66"),
			Expected:   0x66,
		},



		{
			Text: []byte("67"),
			Expected:   0x67,
		},



		{
			Text: []byte("68"),
			Expected:   0x68,
		},



		{
			Text: []byte("69"),
			Expected:   0x69,
		},



		{
			Text: []byte("6a"),
			Expected:   0x6a,
		},
		{
			Text: []byte("6A"),
			Expected:   0x6a,
		},



		{
			Text: []byte("6b"),
			Expected:   0x6b,
		},
		{
			Text: []byte("6B"),
			Expected:   0x6b,
		},



		{
			Text: []byte("6c"),
			Expected:   0x6c,
		},
		{
			Text: []byte("6C"),
			Expected:   0x6c,
		},



		{
			Text: []byte("6d"),
			Expected:   0x6d,
		},
		{
			Text: []byte("6D"),
			Expected:   0x6d,
		},



		{
			Text: []byte("6e"),
			Expected:   0x6e,
		},
		{
			Text: []byte("6E"),
			Expected:   0x6e,
		},



		{
			Text: []byte("6f"),
			Expected:   0x6f,
		},
		{
			Text: []byte("6F"),
			Expected:   0x6f,
		},



		{
			Text: []byte("70"),
			Expected:   0x70,
		},



		{
			Text: []byte("71"),
			Expected:   0x71,
		},



		{
			Text: []byte("72"),
			Expected:   0x72,
		},



		{
			Text: []byte("73"),
			Expected:   0x73,
		},



		{
			Text: []byte("74"),
			Expected:   0x74,
		},



		{
			Text: []byte("75"),
			Expected:   0x75,
		},



		{
			Text: []byte("76"),
			Expected:   0x76,
		},



		{
			Text: []byte("77"),
			Expected:   0x77,
		},



		{
			Text: []byte("78"),
			Expected:   0x78,
		},



		{
			Text: []byte("79"),
			Expected:   0x79,
		},



		{
			Text: []byte("7a"),
			Expected:   0x7a,
		},
		{
			Text: []byte("7A"),
			Expected:   0x7a,
		},



		{
			Text: []byte("7b"),
			Expected:   0x7b,
		},
		{
			Text: []byte("7B"),
			Expected:   0x7b,
		},



		{
			Text: []byte("7c"),
			Expected:   0x7c,
		},
		{
			Text: []byte("7C"),
			Expected:   0x7c,
		},



		{
			Text: []byte("7d"),
			Expected:   0x7d,
		},
		{
			Text: []byte("7D"),
			Expected:   0x7d,
		},



		{
			Text: []byte("7e"),
			Expected:   0x7e,
		},
		{
			Text: []byte("7E"),
			Expected:   0x7e,
		},



		{
			Text: []byte("7f"),
			Expected:   0x7f,
		},
		{
			Text: []byte("7F"),
			Expected:   0x7f,
		},



		{
			Text: []byte("80"),
			Expected:   0x80,
		},



		{
			Text: []byte("81"),
			Expected:   0x81,
		},



		{
			Text: []byte("82"),
			Expected:   0x82,
		},



		{
			Text: []byte("83"),
			Expected:   0x83,
		},



		{
			Text: []byte("84"),
			Expected:   0x84,
		},



		{
			Text: []byte("85"),
			Expected:   0x85,
		},



		{
			Text: []byte("86"),
			Expected:   0x86,
		},



		{
			Text: []byte("87"),
			Expected:   0x87,
		},



		{
			Text: []byte("88"),
			Expected:   0x88,
		},



		{
			Text: []byte("89"),
			Expected:   0x89,
		},



		{
			Text: []byte("8a"),
			Expected:   0x8a,
		},
		{
			Text: []byte("8A"),
			Expected:   0x8a,
		},



		{
			Text: []byte("8b"),
			Expected:   0x8b,
		},
		{
			Text: []byte("8B"),
			Expected:   0x8b,
		},



		{
			Text: []byte("8c"),
			Expected:   0x8c,
		},
		{
			Text: []byte("8C"),
			Expected:   0x8c,
		},



		{
			Text: []byte("8d"),
			Expected:   0x8d,
		},
		{
			Text: []byte("8D"),
			Expected:   0x8d,
		},



		{
			Text: []byte("8e"),
			Expected:   0x8e,
		},
		{
			Text: []byte("8E"),
			Expected:   0x8e,
		},



		{
			Text: []byte("8f"),
			Expected:   0x8f,
		},
		{
			Text: []byte("8F"),
			Expected:   0x8f,
		},



		{
			Text: []byte("90"),
			Expected:   0x90,
		},



		{
			Text: []byte("91"),
			Expected:   0x91,
		},



		{
			Text: []byte("92"),
			Expected:   0x92,
		},



		{
			Text: []byte("93"),
			Expected:   0x93,
		},



		{
			Text: []byte("94"),
			Expected:   0x94,
		},



		{
			Text: []byte("95"),
			Expected:   0x95,
		},



		{
			Text: []byte("96"),
			Expected:   0x96,
		},



		{
			Text: []byte("97"),
			Expected:   0x97,
		},



		{
			Text: []byte("98"),
			Expected:   0x98,
		},



		{
			Text: []byte("99"),
			Expected:   0x99,
		},



		{
			Text: []byte("9a"),
			Expected:   0x9a,
		},
		{
			Text: []byte("9A"),
			Expected:   0x9a,
		},



		{
			Text: []byte("9b"),
			Expected:   0x9b,
		},
		{
			Text: []byte("9B"),
			Expected:   0x9b,
		},



		{
			Text: []byte("9c"),
			Expected:   0x9c,
		},
		{
			Text: []byte("9C"),
			Expected:   0x9c,
		},



		{
			Text: []byte("9d"),
			Expected:   0x9d,
		},
		{
			Text: []byte("9D"),
			Expected:   0x9d,
		},



		{
			Text: []byte("9e"),
			Expected:   0x9e,
		},
		{
			Text: []byte("9E"),
			Expected:   0x9e,
		},



		{
			Text: []byte("9f"),
			Expected:   0x9f,
		},
		{
			Text: []byte("9F"),
			Expected:   0x9f,
		},



		{
			Text: []byte("a0"),
			Expected:   0xa0,
		},
		{
			Text: []byte("A0"),
			Expected:   0xa0,
		},



		{
			Text: []byte("a1"),
			Expected:   0xa1,
		},
		{
			Text: []byte("A1"),
			Expected:   0xa1,
		},



		{
			Text: []byte("a2"),
			Expected:   0xa2,
		},
		{
			Text: []byte("A2"),
			Expected:   0xa2,
		},



		{
			Text: []byte("a3"),
			Expected:   0xa3,
		},
		{
			Text: []byte("A3"),
			Expected:   0xa3,
		},



		{
			Text: []byte("a4"),
			Expected:   0xa4,
		},
		{
			Text: []byte("A4"),
			Expected:   0xa4,
		},



		{
			Text: []byte("a5"),
			Expected:   0xa5,
		},
		{
			Text: []byte("A5"),
			Expected:   0xa5,
		},



		{
			Text: []byte("a6"),
			Expected:   0xa6,
		},
		{
			Text: []byte("A6"),
			Expected:   0xa6,
		},



		{
			Text: []byte("a7"),
			Expected:   0xa7,
		},
		{
			Text: []byte("A7"),
			Expected:   0xa7,
		},



		{
			Text: []byte("a8"),
			Expected:   0xa8,
		},
		{
			Text: []byte("A8"),
			Expected:   0xa8,
		},



		{
			Text: []byte("a9"),
			Expected:   0xa9,
		},
		{
			Text: []byte("A9"),
			Expected:   0xa9,
		},



		{
			Text: []byte("aa"),
			Expected:   0xaa,
		},
		{
			Text: []byte("Aa"),
			Expected:   0xaa,
		},
		{
			Text: []byte("aA"),
			Expected:   0xaa,
		},
		{
			Text: []byte("AA"),
			Expected:   0xaa,
		},



		{
			Text: []byte("ab"),
			Expected:   0xab,
		},
		{
			Text: []byte("Ab"),
			Expected:   0xab,
		},
		{
			Text: []byte("aB"),
			Expected:   0xab,
		},
		{
			Text: []byte("AB"),
			Expected:   0xab,
		},



		{
			Text: []byte("ac"),
			Expected:   0xac,
		},
		{
			Text: []byte("Ac"),
			Expected:   0xac,
		},
		{
			Text: []byte("aC"),
			Expected:   0xac,
		},
		{
			Text: []byte("AC"),
			Expected:   0xac,
		},



		{
			Text: []byte("ad"),
			Expected:   0xad,
		},
		{
			Text: []byte("Ad"),
			Expected:   0xad,
		},
		{
			Text: []byte("aD"),
			Expected:   0xad,
		},
		{
			Text: []byte("AD"),
			Expected:   0xad,
		},



		{
			Text: []byte("ae"),
			Expected:   0xae,
		},
		{
			Text: []byte("Ae"),
			Expected:   0xae,
		},
		{
			Text: []byte("aE"),
			Expected:   0xae,
		},
		{
			Text: []byte("AE"),
			Expected:   0xae,
		},



		{
			Text: []byte("af"),
			Expected:   0xaf,
		},
		{
			Text: []byte("Af"),
			Expected:   0xaf,
		},
		{
			Text: []byte("aF"),
			Expected:   0xaf,
		},
		{
			Text: []byte("AF"),
			Expected:   0xaf,
		},



		{
			Text: []byte("b0"),
			Expected:   0xb0,
		},
		{
			Text: []byte("B0"),
			Expected:   0xb0,
		},



		{
			Text: []byte("b1"),
			Expected:   0xb1,
		},
		{
			Text: []byte("B1"),
			Expected:   0xb1,
		},



		{
			Text: []byte("b2"),
			Expected:   0xb2,
		},
		{
			Text: []byte("B2"),
			Expected:   0xb2,
		},



		{
			Text: []byte("b3"),
			Expected:   0xb3,
		},
		{
			Text: []byte("B3"),
			Expected:   0xb3,
		},



		{
			Text: []byte("b4"),
			Expected:   0xb4,
		},
		{
			Text: []byte("B4"),
			Expected:   0xb4,
		},



		{
			Text: []byte("b5"),
			Expected:   0xb5,
		},
		{
			Text: []byte("B5"),
			Expected:   0xb5,
		},



		{
			Text: []byte("b6"),
			Expected:   0xb6,
		},
		{
			Text: []byte("B6"),
			Expected:   0xb6,
		},



		{
			Text: []byte("b7"),
			Expected:   0xb7,
		},
		{
			Text: []byte("B7"),
			Expected:   0xb7,
		},



		{
			Text: []byte("b8"),
			Expected:   0xb8,
		},
		{
			Text: []byte("B8"),
			Expected:   0xb8,
		},



		{
			Text: []byte("b9"),
			Expected:   0xb9,
		},
		{
			Text: []byte("B9"),
			Expected:   0xb9,
		},



		{
			Text: []byte("ba"),
			Expected:   0xba,
		},
		{
			Text: []byte("Ba"),
			Expected:   0xba,
		},
		{
			Text: []byte("bA"),
			Expected:   0xba,
		},
		{
			Text: []byte("BA"),
			Expected:   0xba,
		},



		{
			Text: []byte("bb"),
			Expected:   0xbb,
		},
		{
			Text: []byte("Bb"),
			Expected:   0xbb,
		},
		{
			Text: []byte("bB"),
			Expected:   0xbb,
		},
		{
			Text: []byte("BB"),
			Expected:   0xbb,
		},



		{
			Text: []byte("bc"),
			Expected:   0xbc,
		},
		{
			Text: []byte("Bc"),
			Expected:   0xbc,
		},
		{
			Text: []byte("bC"),
			Expected:   0xbc,
		},
		{
			Text: []byte("BC"),
			Expected:   0xbc,
		},



		{
			Text: []byte("bd"),
			Expected:   0xbd,
		},
		{
			Text: []byte("Bd"),
			Expected:   0xbd,
		},
		{
			Text: []byte("bD"),
			Expected:   0xbd,
		},
		{
			Text: []byte("BD"),
			Expected:   0xbd,
		},



		{
			Text: []byte("be"),
			Expected:   0xbe,
		},
		{
			Text: []byte("Be"),
			Expected:   0xbe,
		},
		{
			Text: []byte("bE"),
			Expected:   0xbe,
		},
		{
			Text: []byte("BE"),
			Expected:   0xbe,
		},



		{
			Text: []byte("bf"),
			Expected:   0xbf,
		},
		{
			Text: []byte("Bf"),
			Expected:   0xbf,
		},
		{
			Text: []byte("bF"),
			Expected:   0xbf,
		},
		{
			Text: []byte("BF"),
			Expected:   0xbf,
		},



		{
			Text: []byte("c0"),
			Expected:   0xc0,
		},
		{
			Text: []byte("C0"),
			Expected:   0xc0,
		},



		{
			Text: []byte("c1"),
			Expected:   0xc1,
		},
		{
			Text: []byte("C1"),
			Expected:   0xc1,
		},



		{
			Text: []byte("c2"),
			Expected:   0xc2,
		},
		{
			Text: []byte("C2"),
			Expected:   0xc2,
		},



		{
			Text: []byte("c3"),
			Expected:   0xc3,
		},
		{
			Text: []byte("C3"),
			Expected:   0xc3,
		},



		{
			Text: []byte("c4"),
			Expected:   0xc4,
		},
		{
			Text: []byte("C4"),
			Expected:   0xc4,
		},



		{
			Text: []byte("c5"),
			Expected:   0xc5,
		},
		{
			Text: []byte("C5"),
			Expected:   0xc5,
		},



		{
			Text: []byte("c6"),
			Expected:   0xc6,
		},
		{
			Text: []byte("C6"),
			Expected:   0xc6,
		},



		{
			Text: []byte("c7"),
			Expected:   0xc7,
		},
		{
			Text: []byte("C7"),
			Expected:   0xc7,
		},



		{
			Text: []byte("c8"),
			Expected:   0xc8,
		},
		{
			Text: []byte("C8"),
			Expected:   0xc8,
		},



		{
			Text: []byte("c9"),
			Expected:   0xc9,
		},
		{
			Text: []byte("C9"),
			Expected:   0xc9,
		},



		{
			Text: []byte("ca"),
			Expected:   0xca,
		},
		{
			Text: []byte("Ca"),
			Expected:   0xca,
		},
		{
			Text: []byte("cA"),
			Expected:   0xca,
		},
		{
			Text: []byte("CA"),
			Expected:   0xca,
		},



		{
			Text: []byte("cb"),
			Expected:   0xcb,
		},
		{
			Text: []byte("Cb"),
			Expected:   0xcb,
		},
		{
			Text: []byte("cB"),
			Expected:   0xcb,
		},
		{
			Text: []byte("CB"),
			Expected:   0xcb,
		},



		{
			Text: []byte("cc"),
			Expected:   0xcc,
		},
		{
			Text: []byte("Cc"),
			Expected:   0xcc,
		},
		{
			Text: []byte("cC"),
			Expected:   0xcc,
		},
		{
			Text: []byte("CC"),
			Expected:   0xcc,
		},



		{
			Text: []byte("cd"),
			Expected:   0xcd,
		},
		{
			Text: []byte("Cd"),
			Expected:   0xcd,
		},
		{
			Text: []byte("cD"),
			Expected:   0xcd,
		},
		{
			Text: []byte("CD"),
			Expected:   0xcd,
		},



		{
			Text: []byte("ce"),
			Expected:   0xce,
		},
		{
			Text: []byte("Ce"),
			Expected:   0xce,
		},
		{
			Text: []byte("cE"),
			Expected:   0xce,
		},
		{
			Text: []byte("CE"),
			Expected:   0xce,
		},



		{
			Text: []byte("cf"),
			Expected:   0xcf,
		},
		{
			Text: []byte("Cf"),
			Expected:   0xcf,
		},
		{
			Text: []byte("cF"),
			Expected:   0xcf,
		},
		{
			Text: []byte("CF"),
			Expected:   0xcf,
		},



		{
			Text: []byte("d0"),
			Expected:   0xd0,
		},
		{
			Text: []byte("D0"),
			Expected:   0xd0,
		},



		{
			Text: []byte("d1"),
			Expected:   0xd1,
		},
		{
			Text: []byte("D1"),
			Expected:   0xd1,
		},



		{
			Text: []byte("d2"),
			Expected:   0xd2,
		},
		{
			Text: []byte("D2"),
			Expected:   0xd2,
		},



		{
			Text: []byte("d3"),
			Expected:   0xd3,
		},
		{
			Text: []byte("D3"),
			Expected:   0xd3,
		},



		{
			Text: []byte("d4"),
			Expected:   0xd4,
		},
		{
			Text: []byte("D4"),
			Expected:   0xd4,
		},



		{
			Text: []byte("d5"),
			Expected:   0xd5,
		},
		{
			Text: []byte("D5"),
			Expected:   0xd5,
		},



		{
			Text: []byte("d6"),
			Expected:   0xd6,
		},
		{
			Text: []byte("D6"),
			Expected:   0xd6,
		},



		{
			Text: []byte("d7"),
			Expected:   0xd7,
		},
		{
			Text: []byte("D7"),
			Expected:   0xd7,
		},



		{
			Text: []byte("d8"),
			Expected:   0xd8,
		},
		{
			Text: []byte("D8"),
			Expected:   0xd8,
		},



		{
			Text: []byte("d9"),
			Expected:   0xd9,
		},
		{
			Text: []byte("D9"),
			Expected:   0xd9,
		},



		{
			Text: []byte("da"),
			Expected:   0xda,
		},
		{
			Text: []byte("Da"),
			Expected:   0xda,
		},
		{
			Text: []byte("dA"),
			Expected:   0xda,
		},
		{
			Text: []byte("DA"),
			Expected:   0xda,
		},



		{
			Text: []byte("db"),
			Expected:   0xdb,
		},
		{
			Text: []byte("Db"),
			Expected:   0xdb,
		},
		{
			Text: []byte("dB"),
			Expected:   0xdb,
		},
		{
			Text: []byte("DB"),
			Expected:   0xdb,
		},



		{
			Text: []byte("dc"),
			Expected:   0xdc,
		},
		{
			Text: []byte("Dc"),
			Expected:   0xdc,
		},
		{
			Text: []byte("dC"),
			Expected:   0xdc,
		},
		{
			Text: []byte("DC"),
			Expected:   0xdc,
		},



		{
			Text: []byte("dd"),
			Expected:   0xdd,
		},
		{
			Text: []byte("Dd"),
			Expected:   0xdd,
		},
		{
			Text: []byte("dD"),
			Expected:   0xdd,
		},
		{
			Text: []byte("DD"),
			Expected:   0xdd,
		},



		{
			Text: []byte("de"),
			Expected:   0xde,
		},
		{
			Text: []byte("De"),
			Expected:   0xde,
		},
		{
			Text: []byte("dE"),
			Expected:   0xde,
		},
		{
			Text: []byte("DE"),
			Expected:   0xde,
		},



		{
			Text: []byte("df"),
			Expected:   0xdf,
		},
		{
			Text: []byte("Df"),
			Expected:   0xdf,
		},
		{
			Text: []byte("dF"),
			Expected:   0xdf,
		},
		{
			Text: []byte("DF"),
			Expected:   0xdf,
		},



		{
			Text: []byte("e0"),
			Expected:   0xe0,
		},
		{
			Text: []byte("E0"),
			Expected:   0xe0,
		},



		{
			Text: []byte("e1"),
			Expected:   0xe1,
		},
		{
			Text: []byte("E1"),
			Expected:   0xe1,
		},



		{
			Text: []byte("e2"),
			Expected:   0xe2,
		},
		{
			Text: []byte("E2"),
			Expected:   0xe2,
		},



		{
			Text: []byte("e3"),
			Expected:   0xe3,
		},
		{
			Text: []byte("E3"),
			Expected:   0xe3,
		},



		{
			Text: []byte("e4"),
			Expected:   0xe4,
		},
		{
			Text: []byte("E4"),
			Expected:   0xe4,
		},



		{
			Text: []byte("e5"),
			Expected:   0xe5,
		},
		{
			Text: []byte("E5"),
			Expected:   0xe5,
		},



		{
			Text: []byte("e6"),
			Expected:   0xe6,
		},
		{
			Text: []byte("E6"),
			Expected:   0xe6,
		},



		{
			Text: []byte("e7"),
			Expected:   0xe7,
		},
		{
			Text: []byte("E7"),
			Expected:   0xe7,
		},



		{
			Text: []byte("e8"),
			Expected:   0xe8,
		},
		{
			Text: []byte("E8"),
			Expected:   0xe8,
		},



		{
			Text: []byte("e9"),
			Expected:   0xe9,
		},
		{
			Text: []byte("E9"),
			Expected:   0xe9,
		},



		{
			Text: []byte("ea"),
			Expected:   0xea,
		},
		{
			Text: []byte("Ea"),
			Expected:   0xea,
		},
		{
			Text: []byte("eA"),
			Expected:   0xea,
		},
		{
			Text: []byte("EA"),
			Expected:   0xea,
		},



		{
			Text: []byte("eb"),
			Expected:   0xeb,
		},
		{
			Text: []byte("Eb"),
			Expected:   0xeb,
		},
		{
			Text: []byte("eB"),
			Expected:   0xeb,
		},
		{
			Text: []byte("EB"),
			Expected:   0xeb,
		},



		{
			Text: []byte("ec"),
			Expected:   0xec,
		},
		{
			Text: []byte("Ec"),
			Expected:   0xec,
		},
		{
			Text: []byte("eC"),
			Expected:   0xec,
		},
		{
			Text: []byte("EC"),
			Expected:   0xec,
		},



		{
			Text: []byte("ed"),
			Expected:   0xed,
		},
		{
			Text: []byte("Ed"),
			Expected:   0xed,
		},
		{
			Text: []byte("eD"),
			Expected:   0xed,
		},
		{
			Text: []byte("ED"),
			Expected:   0xed,
		},



		{
			Text: []byte("ee"),
			Expected:   0xee,
		},
		{
			Text: []byte("Ee"),
			Expected:   0xee,
		},
		{
			Text: []byte("eE"),
			Expected:   0xee,
		},
		{
			Text: []byte("EE"),
			Expected:   0xee,
		},



		{
			Text: []byte("ef"),
			Expected:   0xef,
		},
		{
			Text: []byte("Ef"),
			Expected:   0xef,
		},
		{
			Text: []byte("eF"),
			Expected:   0xef,
		},
		{
			Text: []byte("EF"),
			Expected:   0xef,
		},



		{
			Text: []byte("f0"),
			Expected:   0xf0,
		},
		{
			Text: []byte("F0"),
			Expected:   0xf0,
		},



		{
			Text: []byte("f1"),
			Expected:   0xf1,
		},
		{
			Text: []byte("F1"),
			Expected:   0xf1,
		},



		{
			Text: []byte("f2"),
			Expected:   0xf2,
		},
		{
			Text: []byte("F2"),
			Expected:   0xf2,
		},



		{
			Text: []byte("f3"),
			Expected:   0xf3,
		},
		{
			Text: []byte("F3"),
			Expected:   0xf3,
		},



		{
			Text: []byte("f4"),
			Expected:   0xf4,
		},
		{
			Text: []byte("F4"),
			Expected:   0xf4,
		},



		{
			Text: []byte("f5"),
			Expected:   0xf5,
		},
		{
			Text: []byte("F5"),
			Expected:   0xf5,
		},



		{
			Text: []byte("f6"),
			Expected:   0xf6,
		},
		{
			Text: []byte("F6"),
			Expected:   0xf6,
		},



		{
			Text: []byte("f7"),
			Expected:   0xf7,
		},
		{
			Text: []byte("F7"),
			Expected:   0xf7,
		},



		{
			Text: []byte("f8"),
			Expected:   0xf8,
		},
		{
			Text: []byte("F8"),
			Expected:   0xf8,
		},



		{
			Text: []byte("f9"),
			Expected:   0xf9,
		},
		{
			Text: []byte("F9"),
			Expected:   0xf9,
		},



		{
			Text: []byte("fa"),
			Expected:   0xfa,
		},
		{
			Text: []byte("Fa"),
			Expected:   0xfa,
		},
		{
			Text: []byte("fA"),
			Expected:   0xfa,
		},
		{
			Text: []byte("FA"),
			Expected:   0xfa,
		},



		{
			Text: []byte("fb"),
			Expected:   0xfb,
		},
		{
			Text: []byte("Fb"),
			Expected:   0xfb,
		},
		{
			Text: []byte("fB"),
			Expected:   0xfb,
		},
		{
			Text: []byte("FB"),
			Expected:   0xfb,
		},



		{
			Text: []byte("fc"),
			Expected:   0xfc,
		},
		{
			Text: []byte("Fc"),
			Expected:   0xfc,
		},
		{
			Text: []byte("fC"),
			Expected:   0xfc,
		},
		{
			Text: []byte("FC"),
			Expected:   0xfc,
		},



		{
			Text: []byte("fd"),
			Expected:   0xfd,
		},
		{
			Text: []byte("Fd"),
			Expected:   0xfd,
		},
		{
			Text: []byte("fD"),
			Expected:   0xfd,
		},
		{
			Text: []byte("FD"),
			Expected:   0xfd,
		},



		{
			Text: []byte("fe"),
			Expected:   0xfe,
		},
		{
			Text: []byte("Fe"),
			Expected:   0xfe,
		},
		{
			Text: []byte("fE"),
			Expected:   0xfe,
		},
		{
			Text: []byte("FE"),
			Expected:   0xfe,
		},



		{
			Text: []byte("ff"),
			Expected:   0xff,
		},
		{
			Text: []byte("Ff"),
			Expected:   0xff,
		},
		{
			Text: []byte("fF"),
			Expected:   0xff,
		},
		{
			Text: []byte("FF"),
			Expected:   0xff,
		},
	}


	for testNumber, test := range tests {

		actual, err := hexadecimal(test.Text)
		if nil != err {
			t.Errorf("For test #%d, (for %q) did not expect an error, but actually got one: (%T) %q", testNumber, string(test.Text), err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, (for %q) expected %d, but actually got %d.", testNumber, string(test.Text), expected, actual)
			continue
		}
	}
}

func TestHexadecimalError(t *testing.T) {

	tests := []struct{
		Text []byte
	}{
		{
			Text: []byte(""),
		},



		{
			Text: []byte("100"),
		},



		{
			Text: []byte("111"),
		},



		{
			Text: []byte("fff"),
		},



		{
			Text: []byte("1000"),
		},



		{
			Text: []byte("g"),
		},
		{
			Text: []byte("G"),
		},



		{
			Text: []byte("h"),
		},
		{
			Text: []byte("H"),
		},



		{
			Text: []byte("i"),
		},
		{
			Text: []byte("I"),
		},



		{
			Text: []byte("j"),
		},
		{
			Text: []byte("J"),
		},



		{
			Text: []byte("k"),
		},
		{
			Text: []byte("K"),
		},



		{
			Text: []byte("l"),
		},
		{
			Text: []byte("L"),
		},



		{
			Text: []byte("m"),
		},
		{
			Text: []byte("M"),
		},



		{
			Text: []byte("n"),
		},
		{
			Text: []byte("N"),
		},



		{
			Text: []byte("o"),
		},
		{
			Text: []byte("O"),
		},



		{
			Text: []byte("p"),
		},
		{
			Text: []byte("P"),
		},



		{
			Text: []byte("q"),
		},
		{
			Text: []byte("Q"),
		},



		{
			Text: []byte("r"),
		},
		{
			Text: []byte("R"),
		},



		{
			Text: []byte("s"),
		},
		{
			Text: []byte("S"),
		},



		{
			Text: []byte("t"),
		},
		{
			Text: []byte("T"),
		},



		{
			Text: []byte("u"),
		},
		{
			Text: []byte("U"),
		},



		{
			Text: []byte("v"),
		},
		{
			Text: []byte("V"),
		},



		{
			Text: []byte("w"),
		},
		{
			Text: []byte("W"),
		},



		{
			Text: []byte("x"),
		},
		{
			Text: []byte("X"),
		},



		{
			Text: []byte("y"),
		},
		{
			Text: []byte("Y"),
		},



		{
			Text: []byte("z"),
		},
		{
			Text: []byte("Z"),
		},



		{
			Text: []byte("2g"),
		},
		{
			Text: []byte("2G"),
		},



		{
			Text: []byte("2h"),
		},
		{
			Text: []byte("2H"),
		},



		{
			Text: []byte("2i"),
		},
		{
			Text: []byte("2I"),
		},



		{
			Text: []byte("2j"),
		},
		{
			Text: []byte("2J"),
		},



		{
			Text: []byte("2k"),
		},
		{
			Text: []byte("2K"),
		},



		{
			Text: []byte("2l"),
		},
		{
			Text: []byte("2L"),
		},



		{
			Text: []byte("2m"),
		},
		{
			Text: []byte("2M"),
		},



		{
			Text: []byte("2n"),
		},
		{
			Text: []byte("2N"),
		},



		{
			Text: []byte("2o"),
		},
		{
			Text: []byte("2O"),
		},



		{
			Text: []byte("2p"),
		},
		{
			Text: []byte("2P"),
		},



		{
			Text: []byte("2q"),
		},
		{
			Text: []byte("2Q"),
		},



		{
			Text: []byte("2r"),
		},
		{
			Text: []byte("2R"),
		},



		{
			Text: []byte("2s"),
		},
		{
			Text: []byte("2S"),
		},



		{
			Text: []byte("2t"),
		},
		{
			Text: []byte("2T"),
		},



		{
			Text: []byte("2u"),
		},
		{
			Text: []byte("2U"),
		},



		{
			Text: []byte("2v"),
		},
		{
			Text: []byte("2V"),
		},



		{
			Text: []byte("2w"),
		},
		{
			Text: []byte("2W"),
		},



		{
			Text: []byte("2x"),
		},
		{
			Text: []byte("2X"),
		},



		{
			Text: []byte("2y"),
		},
		{
			Text: []byte("2Y"),
		},



		{
			Text: []byte("2z"),
		},
		{
			Text: []byte("2Z"),
		},
	}


	for testNumber, test := range tests {

		_, err := hexadecimal(test.Text)
		if nil == err {
			t.Errorf("For test #%d, (for %q) expected an error, but did not actually get one: %#v", testNumber, test.Text, err)
			continue
		}
	}
}
