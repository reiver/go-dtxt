package dtxt

import (
	"bytes"
	"io"

	"testing"
)

func TestUnitescaper_EncodeBytes(t *testing.T) {

	tests := []struct {
		Bytes []byte
		Expected []byte
	}{
		{
			Bytes:        []byte{0},
			Expected: []byte("\x00\x1f"),
		},
		{
			Bytes:        []byte{1},
			Expected: []byte("\x01\x1f"),
		},
		{
			Bytes:        []byte{2},
			Expected: []byte("\x02\x1f"),
		},
		{
			Bytes:        []byte{3},
			Expected: []byte("\x03\x1f"),
		},
		{
			Bytes:        []byte{4},
			Expected: []byte("\x04\x1f"),
		},
		{
			Bytes:        []byte{5},
			Expected: []byte("\x05\x1f"),
		},
		{
			Bytes:        []byte{6},
			Expected: []byte("\x06\x1f"),
		},
		{
			Bytes:        []byte{7},
			Expected: []byte("\x07\x1f"),
		},
		{
			Bytes:        []byte{8},
			Expected: []byte("\x08\x1f"),
		},
		{
			Bytes:        []byte{9},
			Expected: []byte("\x09\x1f"),
		},
		{
			Bytes:       []byte{10},
			Expected: []byte("\x0a\x1f"),
		},
		{
			Bytes:       []byte{11},
			Expected: []byte("\x0b\x1f"),
		},
		{
			Bytes:       []byte{12},
			Expected: []byte("\x0c\x1f"),
		},
		{
			Bytes:       []byte{13},
			Expected: []byte("\x0d\x1f"),
		},
		{
			Bytes:       []byte{14},
			Expected: []byte("\x0e\x1f"),
		},
		{
			Bytes:       []byte{15},
			Expected: []byte("\x0f\x1f"),
		},
		{
			Bytes:       []byte{16},
			Expected: []byte("\x10\x1f"),
		},
		{
			Bytes:       []byte{17},
			Expected: []byte("\x11\x1f"),
		},
		{
			Bytes:       []byte{18},
			Expected: []byte("\x12\x1f"),
		},
		{
			Bytes:       []byte{19},
			Expected: []byte("\x13\x1f"),
		},
		{
			Bytes:       []byte{20},
			Expected: []byte("\x14\x1f"),
		},
		{
			Bytes:       []byte{21},
			Expected: []byte("\x15\x1f"),
		},
		{
			Bytes:       []byte{22},
			Expected: []byte("\x16\x1f"),
		},
		{
			Bytes:       []byte{23},
			Expected: []byte("\x17\x1f"),
		},
		{
			Bytes:       []byte{24},
			Expected: []byte("\x18\x1f"),
		},
		{
			Bytes:       []byte{25},
			Expected: []byte("\x19\x1f"),
		},
		{
			Bytes:       []byte{26},
			Expected: []byte("\x1a\x1f"),
		},



		{
			Bytes:           []byte{27},      //     ESC
			Expected: []byte("\x1b\x1b\x1f"), // ESC ESC US
		},
		{
			Bytes:           []byte{28},      //     FS
			Expected: []byte("\x1b\x1c\x1f"), // ESC FS US
		},
		{
			Bytes:           []byte{29},      //     GS
			Expected: []byte("\x1b\x1d\x1f"), // ESC GS US
		},
		{
			Bytes:           []byte{30},      //     RS
			Expected: []byte("\x1b\x1e\x1f"), // ESC RS US
		},
		{
			Bytes:           []byte{31},      //     US
			Expected: []byte("\x1b\x1f\x1f"), // ESC US US
		},



		{
			Bytes:       []byte{32},
			Expected: []byte("\x20\x1f"),
		},
		{
			Bytes:       []byte{33},
			Expected: []byte("\x21\x1f"),
		},



		{
			//                 NUL SOH STX ETX EOT ENQ ACK BEL  BS  HT  LF  VT  FF  CR  SO  SI DLE DC1 DC2 DC3 DC4 NAK SYN ETB CAN  EM SUB     ESC      FS      GS      RS      FS  SP
			Bytes:    []byte{    0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,     27,     28,     29,     30,     31, 32, 33, 34, 35},
			Expected: []byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1b\x1b\x1c\x1b\x1d\x1b\x1e\x1b\x1f\x20\x21\x22\x23\x1f"),
		},



		{
			Bytes:    []byte{'1'},
			Expected: []byte("1\x1f"),
		},
		{
			Bytes:    []byte{'2'},
			Expected: []byte("2\x1f"),
		},
		{
			Bytes:    []byte{'3'},
			Expected: []byte("3\x1f"),
		},
		{
			Bytes:    []byte{'4'},
			Expected: []byte("4\x1f"),
		},
		{
			Bytes:    []byte{'5'},
			Expected: []byte("5\x1f"),
		},



		{
			Bytes:    []byte{'1','2','3','4','5'},
			Expected: []byte("12345\x1f"),
		},



		{
			Bytes:    []byte{'A'},
			Expected: []byte("A\x1f"),
		},
		{
			Bytes:    []byte{'B'},
			Expected: []byte("B\x1f"),
		},
		{
			Bytes:    []byte{'C'},
			Expected: []byte("C\x1f"),
		},
		{
			Bytes:    []byte{'D'},
			Expected: []byte("D\x1f"),
		},
		{
			Bytes:    []byte{'E'},
			Expected: []byte("E\x1f"),
		},



		{
			Bytes:    []byte{'A','B','C','D','E'},
			Expected: []byte("ABCDE\x1f"),
		},



		{
			Bytes:    []byte{'a'},
			Expected: []byte("a\x1f"),
		},
		{
			Bytes:    []byte{'b'},
			Expected: []byte("b\x1f"),
		},
		{
			Bytes:    []byte{'c'},
			Expected: []byte("c\x1f"),
		},
		{
			Bytes:    []byte{'d'},
			Expected: []byte("d\x1f"),
		},
		{
			Bytes:    []byte{'e'},
			Expected: []byte("e\x1f"),
		},



		{
			Bytes:    []byte{'a','b','c','d','e'},
			Expected: []byte("abcde\x1f"),
		},



		{
			Bytes:    []byte("۱"),
			Expected: []byte("۱\x1f"),
		},
		{
			Bytes:    []byte("۲"),
			Expected: []byte("۲\x1f"),
		},
		{
			Bytes:    []byte("۳"),
			Expected: []byte("۳\x1f"),
		},
		{
			Bytes:    []byte("۴"),
			Expected: []byte("۴\x1f"),
		},
		{
			Bytes:    []byte("۵"),
			Expected: []byte("۵\x1f"),
		},



		{
			Bytes:    []byte("۱۲۳۴۵"),
			Expected: []byte("۱۲۳۴۵\x1f"),
		},



		{
			Bytes:    []byte("≡"),
			Expected: []byte("≡\x1f"),
		},



		{
			Bytes:    []byte("🙂"),
			Expected: []byte("🙂\x1f"),
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer
		var writer io.Writer = &buffer

		var esc unitescaper = unitescaperWrap(writer)

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
