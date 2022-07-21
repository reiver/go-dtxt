package dtxt

import (
	"io"
)

func readGS(reader io.Reader) error {

	var r rune
	{
		var err error

		r, _, err = readRune(reader)
		if nil != err {
			return err
		}
	}

	{
		const expected rune = rune(gs)
		actual := r

		if expected != actual {
			return errNotGS
		}
	}

	return nil
}
