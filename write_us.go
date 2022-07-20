package dtxt

import (
	"io"
)

func writeUS(writer io.Writer) error {
	return writeByte(writer, US)
}
