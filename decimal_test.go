package evmasm

import (
	"testing"
)

func TestDecimal(t *testing.T) {

	tests := []struct{
		Text   []byte
		Expected byte
	}{
		{
			Text: []byte(  "0"),
			Expected:       0,
		},
		{
			Text: []byte( "00"),
			Expected:       0,
		},
		{
			Text: []byte("000"),
			Expected:       0,
		},



		{
			Text: []byte(  "1"),
			Expected:       1,
		},
		{
			Text: []byte( "01"),
			Expected:       1,
		},
		{
			Text: []byte("001"),
			Expected:       1,
		},



		{
			Text: []byte(  "2"),
			Expected:       2,
		},
		{
			Text: []byte( "02"),
			Expected:       2,
		},
		{
			Text: []byte("002"),
			Expected:       2,
		},



		{
			Text: []byte(  "3"),
			Expected:       3,
		},
		{
			Text: []byte( "03"),
			Expected:       3,
		},
		{
			Text: []byte("003"),
			Expected:       3,
		},



		{
			Text: []byte(  "4"),
			Expected:       4,
		},
		{
			Text: []byte( "04"),
			Expected:       4,
		},
		{
			Text: []byte("004"),
			Expected:       4,
		},



		{
			Text: []byte(  "5"),
			Expected:       5,
		},
		{
			Text: []byte( "05"),
			Expected:       5,
		},
		{
			Text: []byte("005"),
			Expected:       5,
		},



		{
			Text: []byte(  "6"),
			Expected:       6,
		},
		{
			Text: []byte( "06"),
			Expected:       6,
		},
		{
			Text: []byte("006"),
			Expected:       6,
		},



		{
			Text: []byte(  "7"),
			Expected:       7,
		},
		{
			Text: []byte( "07"),
			Expected:       7,
		},
		{
			Text: []byte("007"),
			Expected:       7,
		},



		{
			Text: []byte(  "8"),
			Expected:       8,
		},
		{
			Text: []byte( "08"),
			Expected:       8,
		},
		{
			Text: []byte("008"),
			Expected:       8,
		},



		{
			Text: []byte(  "9"),
			Expected:       9,
		},
		{
			Text: []byte( "09"),
			Expected:       9,
		},
		{
			Text: []byte("009"),
			Expected:       9,
		},



		{
			Text: []byte( "10"),
			Expected:      10,
		},
		{
			Text: []byte("010"),
			Expected:      10,
		},



		{
			Text: []byte( "11"),
			Expected:      11,
		},
		{
			Text: []byte("011"),
			Expected:      11,
		},



		{
			Text: []byte( "12"),
			Expected:      12,
		},
		{
			Text: []byte("012"),
			Expected:      12,
		},



		{
			Text: []byte( "13"),
			Expected:      13,
		},
		{
			Text: []byte("013"),
			Expected:      13,
		},



		{
			Text: []byte( "14"),
			Expected:      14,
		},
		{
			Text: []byte("014"),
			Expected:      14,
		},



		{
			Text: []byte( "15"),
			Expected:      15,
		},
		{
			Text: []byte("015"),
			Expected:      15,
		},



		{
			Text: []byte( "16"),
			Expected:      16,
		},
		{
			Text: []byte("016"),
			Expected:      16,
		},



		{
			Text: []byte( "17"),
			Expected:      17,
		},
		{
			Text: []byte("017"),
			Expected:      17,
		},



		{
			Text: []byte( "18"),
			Expected:      18,
		},
		{
			Text: []byte("018"),
			Expected:      18,
		},



		{
			Text: []byte( "19"),
			Expected:      19,
		},
		{
			Text: []byte("019"),
			Expected:      19,
		},



		{
			Text: []byte( "20"),
			Expected:      20,
		},
		{
			Text: []byte("020"),
			Expected:      20,
		},



		{
			Text: []byte( "21"),
			Expected:      21,
		},
		{
			Text: []byte("021"),
			Expected:      21,
		},



		{
			Text: []byte( "22"),
			Expected:      22,
		},
		{
			Text: []byte("022"),
			Expected:      22,
		},



		{
			Text: []byte( "23"),
			Expected:      23,
		},
		{
			Text: []byte("023"),
			Expected:      23,
		},



		{
			Text: []byte( "24"),
			Expected:      24,
		},
		{
			Text: []byte("024"),
			Expected:      24,
		},



		{
			Text: []byte( "25"),
			Expected:      25,
		},
		{
			Text: []byte("025"),
			Expected:      25,
		},



		{
			Text: []byte( "26"),
			Expected:      26,
		},
		{
			Text: []byte("026"),
			Expected:      26,
		},



		{
			Text: []byte( "27"),
			Expected:      27,
		},
		{
			Text: []byte("027"),
			Expected:      27,
		},



		{
			Text: []byte( "28"),
			Expected:      28,
		},
		{
			Text: []byte("028"),
			Expected:      28,
		},



		{
			Text: []byte( "29"),
			Expected:      29,
		},
		{
			Text: []byte("029"),
			Expected:      29,
		},



		{
			Text: []byte( "30"),
			Expected:      30,
		},
		{
			Text: []byte("030"),
			Expected:      30,
		},



		{
			Text: []byte( "31"),
			Expected:      31,
		},
		{
			Text: []byte("031"),
			Expected:      31,
		},



		{
			Text: []byte( "32"),
			Expected:      32,
		},
		{
			Text: []byte("032"),
			Expected:      32,
		},



		{
			Text: []byte( "33"),
			Expected:      33,
		},
		{
			Text: []byte("033"),
			Expected:      33,
		},



		{
			Text: []byte( "34"),
			Expected:      34,
		},
		{
			Text: []byte("034"),
			Expected:      34,
		},



		{
			Text: []byte( "35"),
			Expected:      35,
		},
		{
			Text: []byte("035"),
			Expected:      35,
		},



		{
			Text: []byte( "36"),
			Expected:      36,
		},
		{
			Text: []byte("036"),
			Expected:      36,
		},



		{
			Text: []byte( "37"),
			Expected:      37,
		},
		{
			Text: []byte("037"),
			Expected:      37,
		},



		{
			Text: []byte( "38"),
			Expected:      38,
		},
		{
			Text: []byte("038"),
			Expected:      38,
		},



		{
			Text: []byte( "39"),
			Expected:      39,
		},
		{
			Text: []byte("039"),
			Expected:      39,
		},



		{
			Text: []byte( "40"),
			Expected:      40,
		},
		{
			Text: []byte("040"),
			Expected:      40,
		},



		{
			Text: []byte( "41"),
			Expected:      41,
		},
		{
			Text: []byte("041"),
			Expected:      41,
		},



		{
			Text: []byte( "42"),
			Expected:      42,
		},
		{
			Text: []byte("042"),
			Expected:      42,
		},



		{
			Text: []byte( "43"),
			Expected:      43,
		},
		{
			Text: []byte("043"),
			Expected:      43,
		},



		{
			Text: []byte( "44"),
			Expected:      44,
		},
		{
			Text: []byte("044"),
			Expected:      44,
		},



		{
			Text: []byte( "45"),
			Expected:      45,
		},
		{
			Text: []byte("045"),
			Expected:      45,
		},



		{
			Text: []byte( "46"),
			Expected:      46,
		},
		{
			Text: []byte("046"),
			Expected:      46,
		},



		{
			Text: []byte( "47"),
			Expected:      47,
		},
		{
			Text: []byte("047"),
			Expected:      47,
		},



		{
			Text: []byte( "48"),
			Expected:      48,
		},
		{
			Text: []byte("048"),
			Expected:      48,
		},



		{
			Text: []byte( "49"),
			Expected:      49,
		},
		{
			Text: []byte("049"),
			Expected:      49,
		},



		{
			Text: []byte( "50"),
			Expected:      50,
		},
		{
			Text: []byte("050"),
			Expected:      50,
		},



		{
			Text: []byte( "51"),
			Expected:      51,
		},
		{
			Text: []byte("051"),
			Expected:      51,
		},



		{
			Text: []byte( "52"),
			Expected:      52,
		},
		{
			Text: []byte("052"),
			Expected:      52,
		},



		{
			Text: []byte( "53"),
			Expected:      53,
		},
		{
			Text: []byte("053"),
			Expected:      53,
		},



		{
			Text: []byte( "54"),
			Expected:      54,
		},
		{
			Text: []byte("054"),
			Expected:      54,
		},



		{
			Text: []byte( "55"),
			Expected:      55,
		},
		{
			Text: []byte("055"),
			Expected:      55,
		},



		{
			Text: []byte( "56"),
			Expected:      56,
		},
		{
			Text: []byte("056"),
			Expected:      56,
		},



		{
			Text: []byte( "57"),
			Expected:      57,
		},
		{
			Text: []byte("057"),
			Expected:      57,
		},



		{
			Text: []byte( "58"),
			Expected:      58,
		},
		{
			Text: []byte("058"),
			Expected:      58,
		},



		{
			Text: []byte( "59"),
			Expected:      59,
		},
		{
			Text: []byte("059"),
			Expected:      59,
		},



		{
			Text: []byte( "60"),
			Expected:      60,
		},
		{
			Text: []byte("060"),
			Expected:      60,
		},



		{
			Text: []byte( "61"),
			Expected:      61,
		},
		{
			Text: []byte("061"),
			Expected:      61,
		},



		{
			Text: []byte( "62"),
			Expected:      62,
		},
		{
			Text: []byte("062"),
			Expected:      62,
		},



		{
			Text: []byte( "63"),
			Expected:      63,
		},
		{
			Text: []byte("063"),
			Expected:      63,
		},



		{
			Text: []byte( "64"),
			Expected:      64,
		},
		{
			Text: []byte("064"),
			Expected:      64,
		},



		{
			Text: []byte( "65"),
			Expected:      65,
		},
		{
			Text: []byte("065"),
			Expected:      65,
		},



		{
			Text: []byte( "66"),
			Expected:      66,
		},
		{
			Text: []byte("066"),
			Expected:      66,
		},



		{
			Text: []byte( "67"),
			Expected:      67,
		},
		{
			Text: []byte("067"),
			Expected:      67,
		},



		{
			Text: []byte( "68"),
			Expected:      68,
		},
		{
			Text: []byte("068"),
			Expected:      68,
		},



		{
			Text: []byte( "69"),
			Expected:      69,
		},
		{
			Text: []byte("069"),
			Expected:      69,
		},



		{
			Text: []byte( "70"),
			Expected:      70,
		},
		{
			Text: []byte("070"),
			Expected:      70,
		},



		{
			Text: []byte( "71"),
			Expected:      71,
		},
		{
			Text: []byte("071"),
			Expected:      71,
		},



		{
			Text: []byte( "72"),
			Expected:      72,
		},
		{
			Text: []byte("072"),
			Expected:      72,
		},



		{
			Text: []byte( "73"),
			Expected:      73,
		},
		{
			Text: []byte("073"),
			Expected:      73,
		},



		{
			Text: []byte( "74"),
			Expected:      74,
		},
		{
			Text: []byte("074"),
			Expected:      74,
		},



		{
			Text: []byte( "75"),
			Expected:      75,
		},
		{
			Text: []byte("075"),
			Expected:      75,
		},



		{
			Text: []byte( "76"),
			Expected:      76,
		},
		{
			Text: []byte("076"),
			Expected:      76,
		},



		{
			Text: []byte( "77"),
			Expected:      77,
		},
		{
			Text: []byte("077"),
			Expected:      77,
		},



		{
			Text: []byte( "78"),
			Expected:      78,
		},
		{
			Text: []byte("078"),
			Expected:      78,
		},



		{
			Text: []byte( "79"),
			Expected:      79,
		},
		{
			Text: []byte("079"),
			Expected:      79,
		},



		{
			Text: []byte( "80"),
			Expected:      80,
		},
		{
			Text: []byte("080"),
			Expected:      80,
		},



		{
			Text: []byte( "81"),
			Expected:      81,
		},
		{
			Text: []byte("081"),
			Expected:      81,
		},



		{
			Text: []byte( "82"),
			Expected:      82,
		},
		{
			Text: []byte("082"),
			Expected:      82,
		},



		{
			Text: []byte( "83"),
			Expected:      83,
		},
		{
			Text: []byte("083"),
			Expected:      83,
		},



		{
			Text: []byte( "84"),
			Expected:      84,
		},
		{
			Text: []byte("084"),
			Expected:      84,
		},



		{
			Text: []byte( "85"),
			Expected:      85,
		},
		{
			Text: []byte("085"),
			Expected:      85,
		},



		{
			Text: []byte( "86"),
			Expected:      86,
		},
		{
			Text: []byte("086"),
			Expected:      86,
		},



		{
			Text: []byte( "87"),
			Expected:      87,
		},
		{
			Text: []byte("087"),
			Expected:      87,
		},



		{
			Text: []byte( "88"),
			Expected:      88,
		},
		{
			Text: []byte("088"),
			Expected:      88,
		},



		{
			Text: []byte( "89"),
			Expected:      89,
		},
		{
			Text: []byte("089"),
			Expected:      89,
		},



		{
			Text: []byte( "90"),
			Expected:      90,
		},
		{
			Text: []byte("090"),
			Expected:      90,
		},



		{
			Text: []byte( "91"),
			Expected:      91,
		},
		{
			Text: []byte("091"),
			Expected:      91,
		},



		{
			Text: []byte( "92"),
			Expected:      92,
		},
		{
			Text: []byte("092"),
			Expected:      92,
		},



		{
			Text: []byte( "93"),
			Expected:      93,
		},
		{
			Text: []byte("093"),
			Expected:      93,
		},



		{
			Text: []byte( "94"),
			Expected:      94,
		},
		{
			Text: []byte("094"),
			Expected:      94,
		},



		{
			Text: []byte( "95"),
			Expected:      95,
		},
		{
			Text: []byte("095"),
			Expected:      95,
		},



		{
			Text: []byte( "96"),
			Expected:      96,
		},
		{
			Text: []byte("096"),
			Expected:      96,
		},



		{
			Text: []byte( "97"),
			Expected:      97,
		},
		{
			Text: []byte("097"),
			Expected:      97,
		},



		{
			Text: []byte( "98"),
			Expected:      98,
		},
		{
			Text: []byte("098"),
			Expected:      98,
		},



		{
			Text: []byte( "99"),
			Expected:      99,
		},
		{
			Text: []byte("099"),
			Expected:      99,
		},



		{
			Text: []byte("100"),
			Expected:     100,
		},



		{
			Text: []byte("101"),
			Expected:     101,
		},



		{
			Text: []byte("102"),
			Expected:     102,
		},



		{
			Text: []byte("103"),
			Expected:     103,
		},



		{
			Text: []byte("104"),
			Expected:     104,
		},



		{
			Text: []byte("105"),
			Expected:     105,
		},



		{
			Text: []byte("106"),
			Expected:     106,
		},



		{
			Text: []byte("107"),
			Expected:     107,
		},



		{
			Text: []byte("108"),
			Expected:     108,
		},



		{
			Text: []byte("109"),
			Expected:     109,
		},



		{
			Text: []byte("110"),
			Expected:     110,
		},



		{
			Text: []byte("111"),
			Expected:     111,
		},



		{
			Text: []byte("112"),
			Expected:     112,
		},



		{
			Text: []byte("113"),
			Expected:     113,
		},



		{
			Text: []byte("114"),
			Expected:     114,
		},



		{
			Text: []byte("115"),
			Expected:     115,
		},



		{
			Text: []byte("116"),
			Expected:     116,
		},



		{
			Text: []byte("117"),
			Expected:     117,
		},



		{
			Text: []byte("118"),
			Expected:     118,
		},



		{
			Text: []byte("119"),
			Expected:     119,
		},



		{
			Text: []byte("120"),
			Expected:     120,
		},



		{
			Text: []byte("121"),
			Expected:     121,
		},



		{
			Text: []byte("122"),
			Expected:     122,
		},



		{
			Text: []byte("123"),
			Expected:     123,
		},



		{
			Text: []byte("124"),
			Expected:     124,
		},



		{
			Text: []byte("125"),
			Expected:     125,
		},



		{
			Text: []byte("126"),
			Expected:     126,
		},



		{
			Text: []byte("127"),
			Expected:     127,
		},



		{
			Text: []byte("128"),
			Expected:     128,
		},



		{
			Text: []byte("129"),
			Expected:     129,
		},



		{
			Text: []byte("130"),
			Expected:     130,
		},



		{
			Text: []byte("131"),
			Expected:     131,
		},



		{
			Text: []byte("132"),
			Expected:     132,
		},



		{
			Text: []byte("133"),
			Expected:     133,
		},



		{
			Text: []byte("134"),
			Expected:     134,
		},



		{
			Text: []byte("135"),
			Expected:     135,
		},



		{
			Text: []byte("136"),
			Expected:     136,
		},



		{
			Text: []byte("137"),
			Expected:     137,
		},



		{
			Text: []byte("138"),
			Expected:     138,
		},



		{
			Text: []byte("139"),
			Expected:     139,
		},



		{
			Text: []byte("140"),
			Expected:     140,
		},









		{
			Text: []byte("245"),
			Expected:     245,
		},



		{
			Text: []byte("246"),
			Expected:     246,
		},



		{
			Text: []byte("247"),
			Expected:     247,
		},



		{
			Text: []byte("248"),
			Expected:     248,
		},



		{
			Text: []byte("249"),
			Expected:     249,
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

		actual, err := decimal(test.Text)
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

func TestDecimalError(t *testing.T) {

	tests := []struct{
		Text []byte
	}{
		{
			Text: []byte(""),
		},



		{
			Text: []byte("ff"),
		},



		{
			Text: []byte("256"),
		},
		{
			Text: []byte("257"),
		},
		{
			Text: []byte("258"),
		},
		{
			Text: []byte("259"),
		},
		{
			Text: []byte("260"),
		},



		{
			Text: []byte("270"),
		},



		{
			Text: []byte("280"),
		},
		{
			Text: []byte("281"),
		},
		{
			Text: []byte("282"),
		},
		{
			Text: []byte("283"),
		},
		{
			Text: []byte("284"),
		},
		{
			Text: []byte("285"),
		},
		{
			Text: []byte("286"),
		},
		{
			Text: []byte("287"),
		},
		{
			Text: []byte("288"),
		},
		{
			Text: []byte("289"),
		},
		{
			Text: []byte("290"),
		},
		{
			Text: []byte("291"),
		},
		{
			Text: []byte("292"),
		},
		{
			Text: []byte("293"),
		},
		{
			Text: []byte("294"),
		},
		{
			Text: []byte("295"),
		},
		{
			Text: []byte("296"),
		},
		{
			Text: []byte("297"),
		},
		{
			Text: []byte("298"),
		},
		{
			Text: []byte("299"),
		},
		{
			Text: []byte("300"),
		},



		{
			Text: []byte("400"),
		},



		{
			Text: []byte("500"),
		},



		{
			Text: []byte("600"),
		},



		{
			Text: []byte("700"),
		},



		{
			Text: []byte("800"),
		},



		{
			Text: []byte("900"),
		},



		{
			Text: []byte("987"),
		},
		{
			Text: []byte("988"),
		},
		{
			Text: []byte("989"),
		},
		{
			Text: []byte("990"),
		},
		{
			Text: []byte("991"),
		},
		{
			Text: []byte("992"),
		},
		{
			Text: []byte("993"),
		},
		{
			Text: []byte("994"),
		},
		{
			Text: []byte("995"),
		},
		{
			Text: []byte("996"),
		},
		{
			Text: []byte("997"),
		},
		{
			Text: []byte("998"),
		},
		{
			Text: []byte("999"),
		},
		{
			Text: []byte("1000"),
		},
		{
			Text: []byte("1001"),
		},
		{
			Text: []byte("1002"),
		},
		{
			Text: []byte("1003"),
		},
		{
			Text: []byte("1004"),
		},
		{
			Text: []byte("1005"),
		},
		{
			Text: []byte("1006"),
		},
		{
			Text: []byte("1007"),
		},
		{
			Text: []byte("1008"),
		},
		{
			Text: []byte("1009"),
		},
		{
			Text: []byte("1010"),
		},



		{
			Text: []byte("1020"),
		},
		{
			Text: []byte("1021"),
		},
		{
			Text: []byte("1022"),
		},
		{
			Text: []byte("1023"),
		},
		{
			Text: []byte("1024"),
		},
		{
			Text: []byte("1025"),
		},
		{
			Text: []byte("1026"),
		},
		{
			Text: []byte("1027"),
		},
		{
			Text: []byte("1028"),
		},
		{
			Text: []byte("1029"),
		},
		{
			Text: []byte("1030"),
		},



		{
			Text: []byte("1040"),
		},



		{
			Text: []byte("1050"),
		},



		{
			Text: []byte("1060"),
		},



		{
			Text: []byte("1070"),
		},



		{
			Text: []byte("1080"),
		},



		{
			Text: []byte("1090"),
		},



		{
			Text: []byte("10000"),
		},



		{
			Text: []byte("100000"),
		},



		{
			Text: []byte("1000000"),
		},



		{
			Text: []byte("a"),
		},
		{
			Text: []byte("A"),
		},



		{
			Text: []byte("b"),
		},
		{
			Text: []byte("B"),
		},



		{
			Text: []byte("c"),
		},
		{
			Text: []byte("C"),
		},



		{
			Text: []byte("d"),
		},
		{
			Text: []byte("D"),
		},



		{
			Text: []byte("e"),
		},
		{
			Text: []byte("E"),
		},



		{
			Text: []byte("f"),
		},
		{
			Text: []byte("F"),
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
			Text: []byte("2a"),
		},
		{
			Text: []byte("2A"),
		},



		{
			Text: []byte("2b"),
		},
		{
			Text: []byte("2B"),
		},



		{
			Text: []byte("2c"),
		},
		{
			Text: []byte("2C"),
		},



		{
			Text: []byte("2d"),
		},
		{
			Text: []byte("2D"),
		},



		{
			Text: []byte("2e"),
		},
		{
			Text: []byte("2E"),
		},



		{
			Text: []byte("2f"),
		},
		{
			Text: []byte("2F"),
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

		_, err := decimal(test.Text)
		if nil == err {
			t.Errorf("For test #%d, (for %q) expected an error, but did not actually get one: %#v", testNumber, test.Text, err)
			continue
		}
	}
}
