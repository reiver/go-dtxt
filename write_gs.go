package dtxt

import (
	"io"
)

func writeGS(writer io.Writer) error {
	return writeByte(writer, GS)
}
