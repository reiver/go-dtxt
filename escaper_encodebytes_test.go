package dtxt

import (
	"bytes"
	"io"

	"testing"
)

func TestEscaper_EncodeBytes(t *testing.T) {

	tests := []struct {
		Bytes []byte
		Expected []byte
	}{
		{
			Bytes:        []byte{0},
			Expected: []byte("\x00"),
		},
		{
			Bytes:        []byte{1},
			Expected: []byte("\x01"),
		},
		{
			Bytes:        []byte{2},
			Expected: []byte("\x02"),
		},
		{
			Bytes:        []byte{3},
			Expected: []byte("\x03"),
		},
		{
			Bytes:        []byte{4},
			Expected: []byte("\x04"),
		},
		{
			Bytes:        []byte{5},
			Expected: []byte("\x05"),
		},
		{
			Bytes:        []byte{6},
			Expected: []byte("\x06"),
		},
		{
			Bytes:        []byte{7},
			Expected: []byte("\x07"),
		},
		{
			Bytes:        []byte{8},
			Expected: []byte("\x08"),
		},
		{
			Bytes:        []byte{9},
			Expected: []byte("\x09"),
		},
		{
			Bytes:       []byte{10},
			Expected: []byte("\x0a"),
		},
		{
			Bytes:       []byte{11},
			Expected: []byte("\x0b"),
		},
		{
			Bytes:       []byte{12},
			Expected: []byte("\x0c"),
		},
		{
			Bytes:       []byte{13},
			Expected: []byte("\x0d"),
		},
		{
			Bytes:       []byte{14},
			Expected: []byte("\x0e"),
		},
		{
			Bytes:       []byte{15},
			Expected: []byte("\x0f"),
		},
		{
			Bytes:       []byte{16},
			Expected: []byte("\x10"),
		},
		{
			Bytes:       []byte{17},
			Expected: []byte("\x11"),
		},
		{
			Bytes:       []byte{18},
			Expected: []byte("\x12"),
		},
		{
			Bytes:       []byte{19},
			Expected: []byte("\x13"),
		},
		{
			Bytes:       []byte{20},
			Expected: []byte("\x14"),
		},
		{
			Bytes:       []byte{21},
			Expected: []byte("\x15"),
		},
		{
			Bytes:       []byte{22},
			Expected: []byte("\x16"),
		},
		{
			Bytes:       []byte{23},
			Expected: []byte("\x17"),
		},
		{
			Bytes:       []byte{24},
			Expected: []byte("\x18"),
		},
		{
			Bytes:       []byte{25},
			Expected: []byte("\x19"),
		},
		{
			Bytes:       []byte{26},
			Expected: []byte("\x1a"),
		},



		{
			Bytes:           []byte{27},  //     ESC
			Expected: []byte("\x1b\x1b"), // ESC ESC
		},
		{
			Bytes:           []byte{28},  //     FS
			Expected: []byte("\x1b\x1c"), // ESC FS
		},
		{
			Bytes:           []byte{29},  //     GS
			Expected: []byte("\x1b\x1d"), // ESC GS
		},
		{
			Bytes:           []byte{30},  //     RS
			Expected: []byte("\x1b\x1e"), // ESC RS
		},
		{
			Bytes:           []byte{31},  //     US
			Expected: []byte("\x1b\x1f"), // ESC US
		},



		{
			Bytes:       []byte{32},
			Expected: []byte("\x20"),
		},
		{
			Bytes:       []byte{33},
			Expected: []byte("\x21"),
		},



		{
			//                 NUL SOH STX ETX EOT ENQ ACK BEL  BS  HT  LF  VT  FF  CR  SO  SI DLE DC1 DC2 DC3 DC4 NAK SYN ETB CAN  EM SUB     ESC      FS      GS      RS      FS  SP
			Bytes:    []byte{    0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,     27,     28,     29,     30,     31, 32, 33, 34, 35},
			Expected: []byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1b\x1b\x1c\x1b\x1d\x1b\x1e\x1b\x1f\x20\x21\x22\x23"),
		},



		{
			Bytes:    []byte{'1'},
			Expected: []byte("1"),
		},
		{
			Bytes:    []byte{'2'},
			Expected: []byte("2"),
		},
		{
			Bytes:    []byte{'3'},
			Expected: []byte("3"),
		},
		{
			Bytes:    []byte{'4'},
			Expected: []byte("4"),
		},
		{
			Bytes:    []byte{'5'},
			Expected: []byte("5"),
		},



		{
			Bytes:    []byte{'1','2','3','4','5'},
			Expected: []byte("12345"),
		},



		{
			Bytes:    []byte{'A'},
			Expected: []byte("A"),
		},
		{
			Bytes:    []byte{'B'},
			Expected: []byte("B"),
		},
		{
			Bytes:    []byte{'C'},
			Expected: []byte("C"),
		},
		{
			Bytes:    []byte{'D'},
			Expected: []byte("D"),
		},
		{
			Bytes:    []byte{'E'},
			Expected: []byte("E"),
		},



		{
			Bytes:    []byte{'A','B','C','D','E'},
			Expected: []byte("ABCDE"),
		},



		{
			Bytes:    []byte{'a'},
			Expected: []byte("a"),
		},
		{
			Bytes:    []byte{'b'},
			Expected: []byte("b"),
		},
		{
			Bytes:    []byte{'c'},
			Expected: []byte("c"),
		},
		{
			Bytes:    []byte{'d'},
			Expected: []byte("d"),
		},
		{
			Bytes:    []byte{'e'},
			Expected: []byte("e"),
		},



		{
			Bytes:    []byte{'a','b','c','d','e'},
			Expected: []byte("abcde"),
		},



		{
			Bytes:    []byte("۱"),
			Expected: []byte("۱"),
		},
		{
			Bytes:    []byte("۲"),
			Expected: []byte("۲"),
		},
		{
			Bytes:    []byte("۳"),
			Expected: []byte("۳"),
		},
		{
			Bytes:    []byte("۴"),
			Expected: []byte("۴"),
		},
		{
			Bytes:    []byte("۵"),
			Expected: []byte("۵"),
		},



		{
			Bytes:    []byte("۱۲۳۴۵"),
			Expected: []byte("۱۲۳۴۵"),
		},



		{
			Bytes:    []byte("≡"),
			Expected: []byte("≡"),
		},



		{
			Bytes:    []byte("🙂"),
			Expected: []byte("🙂"),
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer
		var writer io.Writer = &buffer

		var esc escaper = escaperWrap(writer)

		err := esc.EncodeBytes(test.Bytes)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("BYTES: %q", string(test.Bytes))
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
				t.Logf("BYTES: %q", string(test.Bytes))
				continue
			}
		}
	}
}
