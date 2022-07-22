package dtxt

import (
	"bytes"
	"io"

	"testing"
)

func TestUnitescaper_EncodeString(t *testing.T) {

	tests := []struct {
		String string
		Expected []byte
	}{
		{
			String:          "\x00",
			Expected: []byte("\x00\x1f"),
		},
		{
			String:          "\x01",
			Expected: []byte("\x01\x1f"),
		},
		{
			String:          "\x02",
			Expected: []byte("\x02\x1f"),
		},
		{
			String:          "\x03",
			Expected: []byte("\x03\x1f"),
		},
		{
			String:          "\x04",
			Expected: []byte("\x04\x1f"),
		},
		{
			String:          "\x05",
			Expected: []byte("\x05\x1f"),
		},
		{
			String:          "\x06",
			Expected: []byte("\x06\x1f"),
		},
		{
			String:          "\x07",
			Expected: []byte("\x07\x1f"),
		},
		{
			String:          "\x08",
			Expected: []byte("\x08\x1f"),
		},
		{
			String:          "\x09",
			Expected: []byte("\x09\x1f"),
		},
		{
			String:          "\x0a",
			Expected: []byte("\x0a\x1f"),
		},
		{
			String:          "\x0b",
			Expected: []byte("\x0b\x1f"),
		},
		{
			String:          "\x0c",
			Expected: []byte("\x0c\x1f"),
		},
		{
			String:          "\x0d",
			Expected: []byte("\x0d\x1f"),
		},
		{
			String:          "\x0e",
			Expected: []byte("\x0e\x1f"),
		},
		{
			String:          "\x0f",
			Expected: []byte("\x0f\x1f"),
		},
		{
			String:          "\x10",
			Expected: []byte("\x10\x1f"),
		},
		{
			String:          "\x11",
			Expected: []byte("\x11\x1f"),
		},
		{
			String:          "\x12",
			Expected: []byte("\x12\x1f"),
		},
		{
			String:          "\x13",
			Expected: []byte("\x13\x1f"),
		},
		{
			String:          "\x14",
			Expected: []byte("\x14\x1f"),
		},
		{
			String:          "\x15",
			Expected: []byte("\x15\x1f"),
		},
		{
			String:          "\x16",
			Expected: []byte("\x16\x1f"),
		},
		{
			String:          "\x17",
			Expected: []byte("\x17\x1f"),
		},
		{
			String:          "\x18",
			Expected: []byte("\x18\x1f"),
		},
		{
			String:          "\x19",
			Expected: []byte("\x19\x1f"),
		},
		{
			String:          "\x1a",
			Expected: []byte("\x1a\x1f"),
		},



		{
			String:              "\x1b",      //     ESC
			Expected: []byte("\x1b\x1b\x1f"), // ESC ESC US
		},
		{
			String:              "\x1c",      //     FS
			Expected: []byte("\x1b\x1c\x1f"), // ESC FS US
		},
		{
			String:              "\x1d",      //     GS
			Expected: []byte("\x1b\x1d\x1f"), // ESC GS US
		},
		{
			String:              "\x1e",      //     RS
			Expected: []byte("\x1b\x1e\x1f"), // ESC RS US
		},
		{
			String:              "\x1f",      //     US
			Expected: []byte("\x1b\x1f\x1f"), // ESC US US
		},



		{
			String:          "\x20",
			Expected: []byte("\x20\x1f"),
		},
		{
			String:          "\x21",
			Expected: []byte("\x21\x1f"),
		},



		{
			//                        NUL SOH STX ETX EOT ENQ ACK BEL  BS  HT  LF  VT  FF  CR  SO  SI DLE DC1 DC2 DC3 DC4 NAK SYN ETB CAN  EM SUB     ESC      FS      GS      RS      FS  SP
			String:   string([]rune{    0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,     27,     28,     29,     30,     31, 32, 33, 34, 35}),
			Expected:        []byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1b\x1b\x1c\x1b\x1d\x1b\x1e\x1b\x1f\x20\x21\x22\x23\x1f"),
			//                                                                                                                                                                                             US
		},



		{
			String:          "1",
			Expected: []byte("1\x1f"),
		},
		{
			String:          "2",
			Expected: []byte("2\x1f"),
		},
		{
			String:          "3",
			Expected: []byte("3\x1f"),
		},
		{
			String:          "4",
			Expected: []byte("4\x1f"),
		},
		{
			String:          "5",
			Expected: []byte("5\x1f"),
		},



		{
			String:          "12345",
			Expected: []byte("12345\x1f"),
		},



		{
			String:          "A",
			Expected: []byte("A\x1f"),
		},
		{
			String:          "B",
			Expected: []byte("B\x1f"),
		},
		{
			String:          "C",
			Expected: []byte("C\x1f"),
		},
		{
			String:          "D",
			Expected: []byte("D\x1f"),
		},
		{
			String:          "E",
			Expected: []byte("E\x1f"),
		},



		{
			String:          "ABCDE",
			Expected: []byte("ABCDE\x1f"),
		},



		{
			String:          "a",
			Expected: []byte("a\x1f"),
		},
		{
			String:          "b",
			Expected: []byte("b\x1f"),
		},
		{
			String:          "c",
			Expected: []byte("c\x1f"),
		},
		{
			String:          "d",
			Expected: []byte("d\x1f"),
		},
		{
			String:          "e",
			Expected: []byte("e\x1f"),
		},



		{
			String:          "abcde",
			Expected: []byte("abcde\x1f"),
		},



		{
			String:          "۱",
			Expected: []byte("۱\x1f"),
		},
		{
			String:          "۲",
			Expected: []byte("۲\x1f"),
		},
		{
			String:          "۳",
			Expected: []byte("۳\x1f"),
		},
		{
			String:          "۴",
			Expected: []byte("۴\x1f"),
		},
		{
			String:          "۵",
			Expected: []byte("۵\x1f"),
		},



		{
			String:          "۱۲۳۴۵",
			Expected: []byte("۱۲۳۴۵\x1f"),
		},



		{
			String:          "≡",
			Expected: []byte("≡\x1f"),
		},



		{
			String:          "🙂",
			Expected: []byte("🙂\x1f"),
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer
		var writer io.Writer = &buffer

		var esc unitescaper = unitescaperWrap(writer)

		err := esc.EncodeString(test.String)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRING: %q", string(test.String))
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
				t.Logf("STRING: %q", string(test.String))
				continue
			}
		}
	}
}
