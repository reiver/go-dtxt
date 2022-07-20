package dtxt

import (
	"bytes"
	"io"

	"testing"
)

func TestWriteGS(t *testing.T) {

	var buffer bytes.Buffer
	var writer io.Writer = &buffer

	err := writeGS(writer)

	if nil != err {
		t.Errorf("Did not expect an error but actually got one.")
		t.Logf("ERROR: (T%) %s", err, err)
		return
	}

	{
		expected := 1
		actual   := buffer.Len()

		if expected != actual {
			t.Errorf("The actual number of bytes written is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	{
		var expected byte = gs
		var actual   byte = buffer.Bytes()[0]

		if expected != actual {
			t.Errorf("The actual byte written is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}
}
