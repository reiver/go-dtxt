package dtxt

import (
	"github.com/reiver/go-fck"

	"io"
)

func readerUS(reader io.Reader) error {

	var r rune
	{
		var err error

		r, _, err = readRune(reader)
		if nil != err {
			return err
		}
	}

	{
		const expected rune = rune(us)
		actual := r

		if expected != actual {
			return fck.Errorf("expected to read a %d control code character (i.e., a Unit Separator (US) control code character) but actually got %d", expected, actual)
		}
	}

	return nil
}
