package dtxt

import (
	"bytes"
	"encoding"
	"io"

	"testing"
)

func TestUnitescaper_EncodeTextMarshaler(t *testing.T) {

	tests := []struct {
		Bytes []byte
		Expected []byte
	}{
		{
			Bytes:        []byte{0},      // NUL
			Expected: []byte("\x00\x1f"), // NUL US
		},
		{
			Bytes:        []byte{1},      // SOH
			Expected: []byte("\x01\x1f"), // SOH US
		},
		{
			Bytes:        []byte{2},      // STX
			Expected: []byte("\x02\x1f"), // STX US
		},
		{
			Bytes:        []byte{3},      // ETX
			Expected: []byte("\x03\x1f"), // ETX US
		},
		{
			Bytes:        []byte{4},      // EOT
			Expected: []byte("\x04\x1f"), // EOT US
		},
		{
			Bytes:        []byte{5},      // ENQ
			Expected: []byte("\x05\x1f"), // ENQ US
		},
		{
			Bytes:        []byte{6},      // ACK
			Expected: []byte("\x06\x1f"), // ACK US
		},
		{
			Bytes:        []byte{7},      // BEL
			Expected: []byte("\x07\x1f"), // BEL US
		},
		{
			Bytes:        []byte{8},      // BS
			Expected: []byte("\x08\x1f"), // BS US
		},
		{
			Bytes:        []byte{9},      // HT
			Expected: []byte("\x09\x1f"), // HT US
		},
		{
			Bytes:       []byte{10},      // LF
			Expected: []byte("\x0a\x1f"), // LF US
		},
		{
			Bytes:       []byte{11},      // VT
			Expected: []byte("\x0b\x1f"), // VT US
		},
		{
			Bytes:       []byte{12},      // FF
			Expected: []byte("\x0c\x1f"), // FF US
		},
		{
			Bytes:       []byte{13},      // CR
			Expected: []byte("\x0d\x1f"), // CR US
		},
		{
			Bytes:       []byte{14},      // SO
			Expected: []byte("\x0e\x1f"), // SO US
		},
		{
			Bytes:       []byte{15},      // SI
			Expected: []byte("\x0f\x1f"), // SI US
		},
		{
			Bytes:       []byte{16},      // DLE
			Expected: []byte("\x10\x1f"), // DLE US
		},
		{
			Bytes:       []byte{17},      // DC1
			Expected: []byte("\x11\x1f"), // DC1 US
		},
		{
			Bytes:       []byte{18},      // DC2
			Expected: []byte("\x12\x1f"), // DC2 US
		},
		{
			Bytes:       []byte{19},      // DC3
			Expected: []byte("\x13\x1f"), // DC3 US
		},
		{
			Bytes:       []byte{20},      // DC4
			Expected: []byte("\x14\x1f"), // DC4 US
		},
		{
			Bytes:       []byte{21},      // NAK
			Expected: []byte("\x15\x1f"), // NAK US
		},
		{
			Bytes:       []byte{22},      // SYN
			Expected: []byte("\x16\x1f"), // SYN US
		},
		{
			Bytes:       []byte{23},      // ETB
			Expected: []byte("\x17\x1f"), // ETB US
		},
		{
			Bytes:       []byte{24},      // CAN
			Expected: []byte("\x18\x1f"), // CAN US
		},
		{
			Bytes:       []byte{25},      // EM
			Expected: []byte("\x19\x1f"), // EM US
		},
		{
			Bytes:       []byte{26},      // SUB
			Expected: []byte("\x1a\x1f"), // SUB US
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
			Bytes:       []byte{32},      // SP
			Expected: []byte("\x20\x1f"), // SP US
		},
		{
			Bytes:       []byte{33},      // '!'
			Expected: []byte("\x21\x1f"), // '!' US
		},



		{
			//                 NUL SOH STX ETX EOT ENQ ACK BEL  BS  HT  LF  VT  FF  CR  SO  SI DLE DC1 DC2 DC3 DC4 NAK SYN ETB CAN  EM SUB     ESC      FS      GS      RS      FS  SP
			Bytes:    []byte{    0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,     27,     28,     29,     30,     31, 32, 33, 34, 35},
			Expected: []byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1b\x1b\x1c\x1b\x1d\x1b\x1e\x1b\x1f\x20\x21\x22\x23\x1f"),
			//                                                                                                                                                                                      US
		},



		{
			Bytes:    []byte{'1'},     // '1'
			Expected: []byte("1\x1f"), // '1' US
		},
		{
			Bytes:    []byte{'2'},     // '2'
			Expected: []byte("2\x1f"), // '2' US
		},
		{
			Bytes:    []byte{'3'},     // '3'
			Expected: []byte("3\x1f"), // '3' US
		},
		{
			Bytes:    []byte{'4'},     // '4'
			Expected: []byte("4\x1f"), // '4' US
		},
		{
			Bytes:    []byte{'5'},     // '5'
			Expected: []byte("5\x1f"), // '5' US
		},



		{
			Bytes:    []byte{'1','2','3','4','5'}, // '1' '2' '3' '4' '5'
			Expected: []byte("12345\x1f"),         // '1' '2' '3' '4' '5' US
		},



		{
			Bytes:    []byte{'A'},     // 'A'
			Expected: []byte("A\x1f"), // 'A' US
		},
		{
			Bytes:    []byte{'B'},     // 'B'
			Expected: []byte("B\x1f"), // 'B' US
		},
		{
			Bytes:    []byte{'C'},     // 'C'
			Expected: []byte("C\x1f"), // 'C' US
		},
		{
			Bytes:    []byte{'D'},     // 'D'
			Expected: []byte("D\x1f"), // 'D' US
		},
		{
			Bytes:    []byte{'E'},     // 'E'
			Expected: []byte("E\x1f"), // 'E' US
		},



		{
			Bytes:    []byte{'A','B','C','D','E'}, // 'A' 'B' 'C' 'D' 'E'
			Expected: []byte("ABCDE\x1f"),         // 'A' 'B' 'C' 'D' 'E' US
		},



		{
			Bytes:    []byte{'a'},     // 'a'
			Expected: []byte("a\x1f"), // 'a' US
		},
		{
			Bytes:    []byte{'b'},     // 'b'
			Expected: []byte("b\x1f"), // 'b' US
		},
		{
			Bytes:    []byte{'c'},     // 'c'
			Expected: []byte("c\x1f"), // 'c' US
		},
		{
			Bytes:    []byte{'d'},     // 'd'
			Expected: []byte("d\x1f"), // 'd' US
		},
		{
			Bytes:    []byte{'e'},     // 'e'
			Expected: []byte("e\x1f"), // 'e' US
		},



		{
			Bytes:    []byte{'a','b','c','d','e'}, // 'a' 'b' 'c' 'd' 'e'
			Expected: []byte("abcde\x1f"),         // 'a' 'b' 'c' 'd' 'e' US
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

		var textmarshaler encoding.TextMarshaler = BytesTextMarshaler{Bytes:test.Bytes}

		err := esc.EncodeTextMarshaler(textmarshaler)

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
				t.Logf("INPUT BYTES:   %q", string(test.Bytes))
				t.Logf("WRITTEN BYTES: %q", buffer.String())
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
