package dtxt

import (
	"bytes"
	"io"

	"testing"
)

func TestEscaper_EncodeRune(t *testing.T) {

	tests := []struct {
		Rune rune
		Expected []byte
	}{
		{
			Rune:                0,
			Expected: []byte("\x00"),
		},
		{
			Rune:                1,
			Expected: []byte("\x01"),
		},
		{
			Rune:                2,
			Expected: []byte("\x02"),
		},
		{
			Rune:                3,
			Expected: []byte("\x03"),
		},
		{
			Rune:                4,
			Expected: []byte("\x04"),
		},
		{
			Rune:                5,
			Expected: []byte("\x05"),
		},
		{
			Rune:                6,
			Expected: []byte("\x06"),
		},
		{
			Rune:                7,
			Expected: []byte("\x07"),
		},
		{
			Rune:                8,
			Expected: []byte("\x08"),
		},
		{
			Rune:                9,
			Expected: []byte("\x09"),
		},
		{
			Rune:               10,
			Expected: []byte("\x0a"),
		},
		{
			Rune:               11,
			Expected: []byte("\x0b"),
		},
		{
			Rune:               12,
			Expected: []byte("\x0c"),
		},
		{
			Rune:               13,
			Expected: []byte("\x0d"),
		},
		{
			Rune:               14,
			Expected: []byte("\x0e"),
		},
		{
			Rune:               15,
			Expected: []byte("\x0f"),
		},
		{
			Rune:               16,
			Expected: []byte("\x10"),
		},
		{
			Rune:               17,
			Expected: []byte("\x11"),
		},
		{
			Rune:               18,
			Expected: []byte("\x12"),
		},
		{
			Rune:               19,
			Expected: []byte("\x13"),
		},
		{
			Rune:               20,
			Expected: []byte("\x14"),
		},
		{
			Rune:               21,
			Expected: []byte("\x15"),
		},
		{
			Rune:               22,
			Expected: []byte("\x16"),
		},
		{
			Rune:               23,
			Expected: []byte("\x17"),
		},
		{
			Rune:               24,
			Expected: []byte("\x18"),
		},
		{
			Rune:               25,
			Expected: []byte("\x19"),
		},
		{
			Rune:               26,
			Expected: []byte("\x1a"),
		},



		{
			Rune:                   27,   //     ESC
			Expected: []byte("\x1b\x1b"), // ESC ESC
		},
		{
			Rune:                   28,   //     FS
			Expected: []byte("\x1b\x1c"), // ESC FS
		},
		{
			Rune:                   29,   //     GS
			Expected: []byte("\x1b\x1d"), // ESC GS
		},
		{
			Rune:                   30,   //     RS
			Expected: []byte("\x1b\x1e"), // ESC RS
		},
		{
			Rune:                   31,   //     US
			Expected: []byte("\x1b\x1f"), // ESC US
		},



		{
			Rune:               32,
			Expected: []byte("\x20"),
		},
		{
			Rune:               33,
			Expected: []byte("\x21"),
		},



		{
			Rune:            '1',
			Expected: []byte("1"),
		},
		{
			Rune:            '2',
			Expected: []byte("2"),
		},
		{
			Rune:            '3',
			Expected: []byte("3"),
		},
		{
			Rune:            '4',
			Expected: []byte("4"),
		},
		{
			Rune:            '5',
			Expected: []byte("5"),
		},



		{
			Rune:            'A',
			Expected: []byte("A"),
		},
		{
			Rune:            'B',
			Expected: []byte("B"),
		},
		{
			Rune:            'C',
			Expected: []byte("C"),
		},
		{
			Rune:            'D',
			Expected: []byte("D"),
		},
		{
			Rune:            'E',
			Expected: []byte("E"),
		},



		{
			Rune:            'a',
			Expected: []byte("a"),
		},
		{
			Rune:            'b',
			Expected: []byte("b"),
		},
		{
			Rune:            'c',
			Expected: []byte("c"),
		},
		{
			Rune:            'd',
			Expected: []byte("d"),
		},
		{
			Rune:            'e',
			Expected: []byte("e"),
		},



		{
			Rune:            '۱',
			Expected: []byte("۱"),
		},
		{
			Rune:            '۲',
			Expected: []byte("۲"),
		},
		{
			Rune:            '۳',
			Expected: []byte("۳"),
		},
		{
			Rune:            '۴',
			Expected: []byte("۴"),
		},
		{
			Rune:            '۵',
			Expected: []byte("۵"),
		},



		{
			Rune:            '≡',
			Expected: []byte("≡"),
		},



		{
			Rune:            '🙂',
			Expected: []byte("🙂"),
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer
		var writer io.Writer = &buffer

		var esc escaper = escaperWrap(writer)

		err := esc.EncodeRune(test.Rune)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("RUNE: %q", string(test.Rune))
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
				t.Logf("RUNE: %q", string(test.Rune))
				continue
			}
		}
	}
}
