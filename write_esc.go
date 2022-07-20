package dtxt

import (
	"io"
)

func writeESC(writer io.Writer) error {
	return writeByte(writer, ESC)
}
