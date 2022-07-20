package dtxt

import (
	"bytes"
	"io"

	"testing"
)

func TestEscaper_EncodeRunes(t *testing.T) {

	tests := []struct {
		Runes []rune
		Expected []byte
	}{
		{
			Runes:        []rune{0},
			Expected: []byte("\x00"),
		},
		{
			Runes:        []rune{1},
			Expected: []byte("\x01"),
		},
		{
			Runes:        []rune{2},
			Expected: []byte("\x02"),
		},
		{
			Runes:        []rune{3},
			Expected: []byte("\x03"),
		},
		{
			Runes:        []rune{4},
			Expected: []byte("\x04"),
		},
		{
			Runes:        []rune{5},
			Expected: []byte("\x05"),
		},
		{
			Runes:        []rune{6},
			Expected: []byte("\x06"),
		},
		{
			Runes:        []rune{7},
			Expected: []byte("\x07"),
		},
		{
			Runes:        []rune{8},
			Expected: []byte("\x08"),
		},
		{
			Runes:        []rune{9},
			Expected: []byte("\x09"),
		},
		{
			Runes:       []rune{10},
			Expected: []byte("\x0a"),
		},
		{
			Runes:       []rune{11},
			Expected: []byte("\x0b"),
		},
		{
			Runes:       []rune{12},
			Expected: []byte("\x0c"),
		},
		{
			Runes:       []rune{13},
			Expected: []byte("\x0d"),
		},
		{
			Runes:       []rune{14},
			Expected: []byte("\x0e"),
		},
		{
			Runes:       []rune{15},
			Expected: []byte("\x0f"),
		},
		{
			Runes:       []rune{16},
			Expected: []byte("\x10"),
		},
		{
			Runes:       []rune{17},
			Expected: []byte("\x11"),
		},
		{
			Runes:       []rune{18},
			Expected: []byte("\x12"),
		},
		{
			Runes:       []rune{19},
			Expected: []byte("\x13"),
		},
		{
			Runes:       []rune{20},
			Expected: []byte("\x14"),
		},
		{
			Runes:       []rune{21},
			Expected: []byte("\x15"),
		},
		{
			Runes:       []rune{22},
			Expected: []byte("\x16"),
		},
		{
			Runes:       []rune{23},
			Expected: []byte("\x17"),
		},
		{
			Runes:       []rune{24},
			Expected: []byte("\x18"),
		},
		{
			Runes:       []rune{25},
			Expected: []byte("\x19"),
		},
		{
			Runes:       []rune{26},
			Expected: []byte("\x1a"),
		},



		{
			Runes:           []rune{27},  //     ESC
			Expected: []byte("\x1b\x1b"), // ESC ESC
		},
		{
			Runes:           []rune{28},  //     FS
			Expected: []byte("\x1b\x1c"), // ESC FS
		},
		{
			Runes:           []rune{29},  //     GS
			Expected: []byte("\x1b\x1d"), // ESC GS
		},
		{
			Runes:           []rune{30},  //     RS
			Expected: []byte("\x1b\x1e"), // ESC RS
		},
		{
			Runes:           []rune{31},  //     US
			Expected: []byte("\x1b\x1f"), // ESC US
		},



		{
			Runes:       []rune{32},
			Expected: []byte("\x20"),
		},
		{
			Runes:       []rune{33},
			Expected: []byte("\x21"),
		},



		{
			//                 NUL SOH STX ETX EOT ENQ ACK BEL  BS  HT  LF  VT  FF  CR  SO  SI DLE DC1 DC2 DC3 DC4 NAK SYN ETB CAN  EM SUB     ESC      FS      GS      RS      FS  SP
			Runes:    []rune{    0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,     27,     28,     29,     30,     31, 32, 33, 34, 35},
			Expected: []byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1b\x1b\x1c\x1b\x1d\x1b\x1e\x1b\x1f\x20\x21\x22\x23"),
		},



		{
			Runes:    []rune{'1'},
			Expected: []byte("1"),
		},
		{
			Runes:    []rune{'2'},
			Expected: []byte("2"),
		},
		{
			Runes:    []rune{'3'},
			Expected: []byte("3"),
		},
		{
			Runes:    []rune{'4'},
			Expected: []byte("4"),
		},
		{
			Runes:    []rune{'5'},
			Expected: []byte("5"),
		},



		{
			Runes:    []rune{'1','2','3','4','5'},
			Expected: []byte("12345"),
		},



		{
			Runes:    []rune{'A'},
			Expected: []byte("A"),
		},
		{
			Runes:    []rune{'B'},
			Expected: []byte("B"),
		},
		{
			Runes:    []rune{'C'},
			Expected: []byte("C"),
		},
		{
			Runes:    []rune{'D'},
			Expected: []byte("D"),
		},
		{
			Runes:    []rune{'E'},
			Expected: []byte("E"),
		},



		{
			Runes:    []rune{'A','B','C','D','E'},
			Expected: []byte("ABCDE"),
		},



		{
			Runes:    []rune{'a'},
			Expected: []byte("a"),
		},
		{
			Runes:    []rune{'b'},
			Expected: []byte("b"),
		},
		{
			Runes:    []rune{'c'},
			Expected: []byte("c"),
		},
		{
			Runes:    []rune{'d'},
			Expected: []byte("d"),
		},
		{
			Runes:    []rune{'e'},
			Expected: []byte("e"),
		},



		{
			Runes:    []rune{'a','b','c','d','e'},
			Expected: []byte("abcde"),
		},



		{
			Runes:    []rune{'۱'},
			Expected: []byte("۱"),
		},
		{
			Runes:    []rune{'۲'},
			Expected: []byte("۲"),
		},
		{
			Runes:    []rune{'۳'},
			Expected: []byte("۳"),
		},
		{
			Runes:    []rune{'۴'},
			Expected: []byte("۴"),
		},
		{
			Runes:    []rune{'۵'},
			Expected: []byte("۵"),
		},



		{
			Runes:    []rune{'۱','۲','۳','۴','۵'},
			Expected: []byte("۱۲۳۴۵"),
		},



		{
			Runes:    []rune{'≡'},
			Expected: []byte("≡"),
		},



		{
			Runes:    []rune{'🙂'},
			Expected: []byte("🙂"),
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer
		var writer io.Writer = &buffer

		var esc escaper = escaperWrap(writer)

		err := esc.EncodeRunes(test.Runes)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("RUNES: %q", string(test.Runes))
			continue
		}

		{
			expected := len(test.Expected)
			actual   := buffer.Len()

			if expected != actual {
				t.Errorf("For test #%d, the actual number of bytes written is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			expected := append([]byte(nil), buffer.Bytes()...)
			actual   := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual bytes written is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", string(expected))
				t.Logf("ACTUAL:   %q", string(actual))
				t.Logf("RUNES: %q", string(test.Runes))
				continue
			}
		}
	}
}
