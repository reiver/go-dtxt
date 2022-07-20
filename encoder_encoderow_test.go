package dtxt

import (
	"bytes"
	"io"

	"testing"
)

func TestEncoder_EncodeRow(t *testing.T) {

	tests := []struct{
		Row []any
		Expected []byte
	}{
		{
			Row: []any{
				BytesTextMarshaler{Bytes:[]byte("apple")},
				BytesTextMarshaler{Bytes:[]byte("banana")},
				BytesTextMarshaler{Bytes:[]byte("cherry")},
			},
			Expected:
				[]byte(
					"apple"  +string(rune(US))+
					"banana" +string(rune(US))+
					"cherry" +string(rune(US))+
					          string(rune(RS)),
				),
		},



		{
			Row: []any{
				"ONE",
				[]byte("TWO"),
				[]rune("THREE"),
				rune('4'),
				"\x00\x01\x02\x03\x04\x05\x06\x07",
				"\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f",
				"\x10\x11\x12\x13\x14\x15\x16\x17",
				"\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f",
				"\x20\x21\x22\x23\x24\x25\x26\x27",
				"\x28\x29\x2a\x2b\x2c\x2d\x2e\x2f",
				"\x30\x31\x32\x33\x34\x35\x36\x37",
				"\x38\x39\x3a\x3b\x3c\x3d\x3e\x3f",
				"\x40\x41\x42\x43\x44\x45\x46\x47",
				"\x48\x49\x4a\x4b\x4c\x4d\x4e\x4f",
				"\x50\x51\x52\x53\x54\x55\x56\x57",
				"\x58\x59\x5a\x5b\x5c\x5d\x5e\x5f",
				"\x60\x61\x62\x63\x64\x65\x66\x67",
				"\x68\x69\x6a\x6b\x6c\x6d\x6e\x6f",
				"\x70\x71\x72\x73\x74\x75\x76\x77",
				"\x78\x79\x7a\x7b\x7c\x7d\x7e\x7f",
				"\x80\x81\x82\x83\x84\x85\x86\x87",
				"\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f",
			},
			Expected:
				[]byte(
					"ONE"                                                  +string(rune(US))+
					"TWO"                                                  +string(rune(US))+
					"THREE"                                                +string(rune(US))+
					"4"                                                    +string(rune(US))+
					"\x00\x01\x02\x03\x04\x05\x06\x07"                     +string(rune(US))+
					"\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f"                     +string(rune(US))+
					"\x10\x11\x12\x13\x14\x15\x16\x17"                     +string(rune(US))+
					"\x18\x19\x1a\x1b\x1b\x1b\x1c\x1b\x1d\x1b\x1e\x1b\x1f" +string(rune(US))+
					"\x20\x21\x22\x23\x24\x25\x26\x27"                     +string(rune(US))+
					"\x28\x29\x2a\x2b\x2c\x2d\x2e\x2f"                     +string(rune(US))+
					"\x30\x31\x32\x33\x34\x35\x36\x37"                     +string(rune(US))+
					"\x38\x39\x3a\x3b\x3c\x3d\x3e\x3f"                     +string(rune(US))+
					"\x40\x41\x42\x43\x44\x45\x46\x47"                     +string(rune(US))+
					"\x48\x49\x4a\x4b\x4c\x4d\x4e\x4f"                     +string(rune(US))+
					"\x50\x51\x52\x53\x54\x55\x56\x57"                     +string(rune(US))+
					"\x58\x59\x5a\x5b\x5c\x5d\x5e\x5f"                     +string(rune(US))+
					"\x60\x61\x62\x63\x64\x65\x66\x67"                     +string(rune(US))+
					"\x68\x69\x6a\x6b\x6c\x6d\x6e\x6f"                     +string(rune(US))+
					"\x70\x71\x72\x73\x74\x75\x76\x77"                     +string(rune(US))+
					"\x78\x79\x7a\x7b\x7c\x7d\x7e\x7f"                     +string(rune(US))+
					"\x80\x81\x82\x83\x84\x85\x86\x87"                     +string(rune(US))+
					"\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f"                     +string(rune(US))+
					                                                        string(rune(RS)),
				),
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer
		var writer io.Writer = &buffer

		var enc Encoder = EncoderWrap(writer)

		err := enc.EncodeRow(test.Row...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.Expected
			actual   := append([]byte(nil), buffer.Bytes()...)

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("EXPECTED: %q", string(expected))
				t.Logf("ACTUAL:   %q", string(actual))
				continue
			}
		}
	}
}
