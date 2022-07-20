package dtxt

import (
	"bytes"
	"io"

	"testing"
)

func TestEscaper_EncodeByte(t *testing.T) {

	tests := []struct {
		Byte byte
		Expected []byte
	}{
		{
			Byte:                0,
			Expected: []byte("\x00"),
		},
		{
			Byte:                1,
			Expected: []byte("\x01"),
		},
		{
			Byte:                2,
			Expected: []byte("\x02"),
		},
		{
			Byte:                3,
			Expected: []byte("\x03"),
		},
		{
			Byte:                4,
			Expected: []byte("\x04"),
		},
		{
			Byte:                5,
			Expected: []byte("\x05"),
		},
		{
			Byte:                6,
			Expected: []byte("\x06"),
		},
		{
			Byte:                7,
			Expected: []byte("\x07"),
		},
		{
			Byte:                8,
			Expected: []byte("\x08"),
		},
		{
			Byte:                9,
			Expected: []byte("\x09"),
		},
		{
			Byte:               10,
			Expected: []byte("\x0a"),
		},
		{
			Byte:               11,
			Expected: []byte("\x0b"),
		},
		{
			Byte:               12,
			Expected: []byte("\x0c"),
		},
		{
			Byte:               13,
			Expected: []byte("\x0d"),
		},
		{
			Byte:               14,
			Expected: []byte("\x0e"),
		},
		{
			Byte:               15,
			Expected: []byte("\x0f"),
		},
		{
			Byte:               16,
			Expected: []byte("\x10"),
		},
		{
			Byte:               17,
			Expected: []byte("\x11"),
		},
		{
			Byte:               18,
			Expected: []byte("\x12"),
		},
		{
			Byte:               19,
			Expected: []byte("\x13"),
		},
		{
			Byte:               20,
			Expected: []byte("\x14"),
		},
		{
			Byte:               21,
			Expected: []byte("\x15"),
		},
		{
			Byte:               22,
			Expected: []byte("\x16"),
		},
		{
			Byte:               23,
			Expected: []byte("\x17"),
		},
		{
			Byte:               24,
			Expected: []byte("\x18"),
		},
		{
			Byte:               25,
			Expected: []byte("\x19"),
		},
		{
			Byte:               26,
			Expected: []byte("\x1a"),
		},



		{
			Byte:                   27,   //     ESC
			Expected: []byte("\x1b\x1b"), // ESC ESC
		},
		{
			Byte:                   28,   //     FS
			Expected: []byte("\x1b\x1c"), // ESC FS
		},
		{
			Byte:                   29,   //     GS
			Expected: []byte("\x1b\x1d"), // ESC GS
		},
		{
			Byte:                   30,   //     RS
			Expected: []byte("\x1b\x1e"), // ESC RS
		},
		{
			Byte:                    31,  //     US
			Expected: []byte("\x1b\x1f"), // ESC US
		},



		{
			Byte:               32,
			Expected: []byte("\x20"),
		},
		{
			Byte:               33,
			Expected: []byte("\x21"),
		},



		{
			Byte:            '1',
			Expected: []byte("1"),
		},
		{
			Byte:            '2',
			Expected: []byte("2"),
		},
		{
			Byte:            '3',
			Expected: []byte("3"),
		},
		{
			Byte:            '4',
			Expected: []byte("4"),
		},
		{
			Byte:            '5',
			Expected: []byte("5"),
		},



		{
			Byte:            'A',
			Expected: []byte("A"),
		},
		{
			Byte:            'B',
			Expected: []byte("B"),
		},
		{
			Byte:            'C',
			Expected: []byte("C"),
		},
		{
			Byte:            'D',
			Expected: []byte("D"),
		},
		{
			Byte:            'E',
			Expected: []byte("E"),
		},



		{
			Byte:            'a',
			Expected: []byte("a"),
		},
		{
			Byte:            'b',
			Expected: []byte("b"),
		},
		{
			Byte:            'c',
			Expected: []byte("c"),
		},
		{
			Byte:            'd',
			Expected: []byte("d"),
		},
		{
			Byte:            'e',
			Expected: []byte("e"),
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer
		var writer io.Writer = &buffer

		var esc escaper = escaperWrap(writer)

		err := esc.EncodeByte(test.Byte)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("BYTE: %q", string(test.Byte))
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
				t.Logf("BYTE: %q", string(test.Byte))
				continue
			}
		}
	}
}
