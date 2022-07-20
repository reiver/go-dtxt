package dtxt

import (
	"github.com/reiver/go-fck"

	"io"
)

func writeByte(writer io.Writer, b byte) error {

	if nil == writer {
		return errNilWriter
	}

	{
		var buffer [1]byte
		var p []byte = buffer[:]

		buffer[0] = b

		n, err := writer.Write(p)

		if nil != err {
			return err
		}
		if expected, actual := 1, n; expected != actual {
			return fck.Errorf("the actual number of bytes writter (%d) is not what was expected (%d)", actual, expected)
		}
	}

	return nil
}
