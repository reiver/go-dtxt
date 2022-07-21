package dtxt

import (
	"io"
)

func readRS(reader io.Reader) error {

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
			return errNotRS
		}
	}

	return nil
}
