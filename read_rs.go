package dtxt

import (
	"github.com/reiver/go-fck"

	"io"
)

func readerRS(reader io.Reader) error {

	var r rune
	{
		var err error

		r, _, err = readRune(reader)
		if nil != err {
			return err
		}
	}

	{
		const expected rune = rune(rs)
		actual := r

		if expected != actual {
			return fck.Errorf("expected to read a %d control code character (i.e., a Record Separator (RS) control code character) but actually got %d", expected, actual)
		}
	}

	return nil
}
