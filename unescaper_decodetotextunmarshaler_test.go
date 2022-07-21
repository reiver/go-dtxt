package dtxt

import (
	"bytes"
	"encoding"
	"io"

	"testing"
)

func TestUnescaper_DecodeToTextUnmarshaler(t *testing.T) {

	tests := []struct{
		Input    []byte
		Expected []byte
	}{
		{
			Input:    []byte{'O','N','C','E',us},
			Expected: []byte{'O','N','C','E'},
		},
		{
			Input:    []byte{'T','W','I','C','E',us},
			Expected: []byte{'T','W','I','C','E'},
		},
		{
			Input:    []byte{'T','H','R','I','C','E',us},
			Expected: []byte{'T','H','R','I','C','E'},
		},
		{
			Input:    []byte{'F','O','U','R','C','E',us},
			Expected: []byte{'F','O','U','R','C','E'},
		},



		{
			Input:    []byte{'o','n','e',us,'t','w','o',us,'t','h','r','e','e',us},
			Expected: []byte{'o','n','e'},
		},



		{
			Input:    []byte{0x00,0x01,0x02,0x03,0x04,0x05,0x06,0x07,us},
			Expected: []byte{0x00,0x01,0x02,0x03,0x04,0x05,0x06,0x07},
		},
		{
			Input:    []byte{0x08,0x09,0x0a,0x0b,0x0c,0x0d,0x0e,0x0f,us},
			Expected: []byte{0x08,0x09,0x0a,0x0b,0x0c,0x0d,0x0e,0x0f},
		},
		{
			Input:    []byte{0x10,0x11,0x12,0x13,0x14,0x15,0x16,0x17,us},
			Expected: []byte{0x10,0x11,0x12,0x13,0x14,0x15,0x16,0x17},
		},
		{
			Input:    []byte{0x18,0x19,0x1a,0x1b,0x1b,0x1b,0x1c,0x1b,0x1d,0x1b,0x1e,0x1b,0x1f,us},
			Expected: []byte{0x18,0x19,0x1a,     0x1b,     0x1c,     0x1d,     0x1e,     0x1f},
		},
		{
			Input:    []byte{0x20,0x21,0x22,0x23,0x24,0x25,0x26,0x27,us},
			Expected: []byte{0x20,0x21,0x22,0x23,0x24,0x25,0x26,0x27},
		},
		{
			Input:    []byte{0x28,0x29,0x2a,0x2b,0x2c,0x2d,0x2e,0x2f,us},
			Expected: []byte{0x28,0x29,0x2a,0x2b,0x2c,0x2d,0x2e,0x2f},
		},



		{
			Input:    []byte{esc,fs,gs,rs,us,us},
			Expected: []byte{    fs,gs,rs},
		},
		{
			Input:    []byte{esc,esc,fs,gs,rs,us,us},
			Expected: []byte{    esc,fs,gs,rs},
		},
		{
			Input:    []byte{esc,esc,fs,gs,rs,esc,us,us},
			Expected: []byte{    esc,fs,gs,rs,    us},
		},
	}

	for testNumber, test := range tests {

		var unesc unescaper
		{
			var reader io.Reader = bytes.NewReader(test.Input)

			unesc = unescaperWrap(reader)
		}

		var bytestextunmarshaler BytesTextUnmarshaler
		var textunmarshaler encoding.TextUnmarshaler = &bytestextunmarshaler

		{
			err := unesc.DecodeToTextUnmarshaler(textunmarshaler)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err ,err)
				continue
			}
		}

		{
			expected := test.Expected
			actual   := bytestextunmarshaler.Bytes

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", string(expected))
				t.Logf("ACTUAL:   %q", string(actual))
				continue
			}
		}
	}
}
