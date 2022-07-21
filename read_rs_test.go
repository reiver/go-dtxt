package dtxt

import (
	"testing"

	"io"
	"strings"
)

func TestReadRS(t *testing.T) {

	const RS string = string(rune(rs))

	var input string = RS + "apple banana cherry"
	var reader io.Reader = strings.NewReader(input)

	{
		err := readRS(reader)
		if nil != err {
			t.Errorf("Did not expect an error but actually got one.")
			t.Logf("ERROR: (%T) %q", err, err)
			return
		}
	}
}

func TestReadRS_error(t *testing.T) {

	const RS string = string(rune(rs))

	var input string = "apple banana cherry" + RS
	var reader io.Reader = strings.NewReader(input)

	{
		err := readRS(reader)
		if nil == err {
			t.Errorf("Expect an error but did not actually get one.")
			t.Logf("ERROR: (%T) %q", err, err)
			return
		}

		if expected, actual := errNotRS, err; expected != actual {
			t.Errorf("Actual error is not what was expected.")
			t.Logf("EXPECTED ERROR: (%T) %s", expected, expected)
			t.Logf("ATUAL    ERROR: (%T) %s", actual, actual)
			return
		}
	}
}
