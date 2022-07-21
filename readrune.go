package dtxt

import (
	"github.com/reiver/go-utf8"

	"io"
)

func readRune(reader io.Reader) (rune, int, error) {
	return utf8.ReadRune(reader)
}
