package dtxt

import (
	"bytes"
	"io"

	"testing"
)

func TestEscaper_EncodeString(t *testing.T) {

	tests := []struct {
		String string
		Expected []byte
	}{
		{
			String:          "\x00",
			Expected: []byte("\x00"),
		},
		{
			String:          "\x01",
			Expected: []byte("\x01"),
		},
		{
			String:          "\x02",
			Expected: []byte("\x02"),
		},
		{
			String:          "\x03",
			Expected: []byte("\x03"),
		},
		{
			String:          "\x04",
			Expected: []byte("\x04"),
		},
		{
			String:          "\x05",
			Expected: []byte("\x05"),
		},
		{
			String:          "\x06",
			Expected: []byte("\x06"),
		},
		{
			String:          "\x07",
			Expected: []byte("\x07"),
		},
		{
			String:          "\x08",
			Expected: []byte("\x08"),
		},
		{
			String:          "\x09",
			Expected: []byte("\x09"),
		},
		{
			String:          "\x0a",
			Expected: []byte("\x0a"),
		},
		{
			String:          "\x0b",
			Expected: []byte("\x0b"),
		},
		{
			String:          "\x0c",
			Expected: []byte("\x0c"),
		},
		{
			String:          "\x0d",
			Expected: []byte("\x0d"),
		},
		{
			String:          "\x0e",
			Expected: []byte("\x0e"),
		},
		{
			String:          "\x0f",
			Expected: []byte("\x0f"),
		},
		{
			String:          "\x10",
			Expected: []byte("\x10"),
		},
		{
			String:          "\x11",
			Expected: []byte("\x11"),
		},
		{
			String:          "\x12",
			Expected: []byte("\x12"),
		},
		{
			String:          "\x13",
			Expected: []byte("\x13"),
		},
		{
			String:          "\x14",
			Expected: []byte("\x14"),
		},
		{
			String:          "\x15",
			Expected: []byte("\x15"),
		},
		{
			String:          "\x16",
			Expected: []byte("\x16"),
		},
		{
			String:          "\x17",
			Expected: []byte("\x17"),
		},
		{
			String:          "\x18",
			Expected: []byte("\x18"),
		},
		{
			String:          "\x19",
			Expected: []byte("\x19"),
		},
		{
			String:          "\x1a",
			Expected: []byte("\x1a"),
		},



		{
			String:              "\x1b",  //     ESC
			Expected: []byte("\x1b\x1b"), // ESC ESC
		},
		{
			String:              "\x1c",  //     FS
			Expected: []byte("\x1b\x1c"), // ESC FS
		},
		{
			String:              "\x1d",  //     GS
			Expected: []byte("\x1b\x1d"), // ESC GS
		},
		{
			String:              "\x1e",  //     RS
			Expected: []byte("\x1b\x1e"), // ESC RS
		},
		{
			String:              "\x1f",  //     US
			Expected: []byte("\x1b\x1f"), // ESC US
		},



		{
			String:          "\x20",
			Expected: []byte("\x20"),
		},
		{
			String:          "\x21",
			Expected: []byte("\x21"),
		},



		{
			//                        NUL SOH STX ETX EOT ENQ ACK BEL  BS  HT  LF  VT  FF  CR  SO  SI DLE DC1 DC2 DC3 DC4 NAK SYN ETB CAN  EM SUB     ESC      FS      GS      RS      FS  SP
			String:   string([]rune{    0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,     27,     28,     29,     30,     31, 32, 33, 34, 35}),
			Expected:        []byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1b\x1b\x1c\x1b\x1d\x1b\x1e\x1b\x1f\x20\x21\x22\x23"),
		},



		{
			String:          "1",
			Expected: []byte("1"),
		},
		{
			String:          "2",
			Expected: []byte("2"),
		},
		{
			String:          "3",
			Expected: []byte("3"),
		},
		{
			String:          "4",
			Expected: []byte("4"),
		},
		{
			String:          "5",
			Expected: []byte("5"),
		},



		{
			String:          "12345",
			Expected: []byte("12345"),
		},



		{
			String:          "A",
			Expected: []byte("A"),
		},
		{
			String:          "B",
			Expected: []byte("B"),
		},
		{
			String:          "C",
			Expected: []byte("C"),
		},
		{
			String:          "D",
			Expected: []byte("D"),
		},
		{
			String:          "E",
			Expected: []byte("E"),
		},



		{
			String:          "ABCDE",
			Expected: []byte("ABCDE"),
		},



		{
			String:          "a",
			Expected: []byte("a"),
		},
		{
			String:          "b",
			Expected: []byte("b"),
		},
		{
			String:          "c",
			Expected: []byte("c"),
		},
		{
			String:          "d",
			Expected: []byte("d"),
		},
		{
			String:          "e",
			Expected: []byte("e"),
		},



		{
			String:          "abcde",
			Expected: []byte("abcde"),
		},



		{
			String:          "۱",
			Expected: []byte("۱"),
		},
		{
			String:          "۲",
			Expected: []byte("۲"),
		},
		{
			String:          "۳",
			Expected: []byte("۳"),
		},
		{
			String:          "۴",
			Expected: []byte("۴"),
		},
		{
			String:          "۵",
			Expected: []byte("۵"),
		},



		{
			String:          "۱۲۳۴۵",
			Expected: []byte("۱۲۳۴۵"),
		},



		{
			String:          "≡",
			Expected: []byte("≡"),
		},



		{
			String:          "🙂",
			Expected: []byte("🙂"),
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer
		var writer io.Writer = &buffer

		var esc escaper = escaperWrap(writer)

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
