package dtxt

import (
	"testing"

	"io"
	"strings"
)

func TestReadGS(t *testing.T) {

	const GS string = string(rune(gs))

	var input string = GS + "apple banana cherry"
	var reader io.Reader = strings.NewReader(input)

	{
		err := readGS(reader)
		if nil != err {
			t.Errorf("Did not expect an error but actually got one.")
			t.Logf("ERROR: (%T) %q", err, err)
			return
		}
	}
}

func TestReadGS_error(t *testing.T) {

	const GS string = string(rune(gs))

	var input string = "apple banana cherry" + GS
	var reader io.Reader = strings.NewReader(input)

	{
		err := readGS(reader)
		if nil == err {
			t.Errorf("Expect an error but did not actually get one.")
			t.Logf("ERROR: (%T) %q", err, err)
			return
		}

		if expected, actual := errNotGS, err; expected != actual {
			t.Errorf("Actual error is not what was expected.")
			t.Logf("EXPECTED ERROR: (%T) %s", expected, expected)
			t.Logf("ATUAL    ERROR: (%T) %s", actual, actual)
			return
		}
	}
}
