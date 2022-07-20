package dtxt

import (
	"io"
)

func writeRS(writer io.Writer) error {
	return writeByte(writer, RS)
}
