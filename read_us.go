package dtxt

import (
	"io"
)

func readUS(reader io.Reader) error {

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
			return errNotUS
		}
	}

	return nil
}
