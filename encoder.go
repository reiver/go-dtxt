package dtxt

	import (
	"io"
)

type Encoder struct {
	writer io.Writer
}

func EncoderWrap(writer io.Writer) Encoder {
	return Encoder{
		writer:writer,
	}
}

func (receiver Encoder) EncodeRow(values ...any) error {

	var writer io.Writer
	{
		writer = receiver.writer


		if nil == writer {
			return errNilWriter
		}
	}

	for _, value := range values {

		err := receiver.encodeUnit(value)
		if nil != err {
			return err
		}
	}

	{
		err := writeRS(writer)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver Encoder) encodeUnit(value any) error {

	var writer io.Writer
	{
		writer = receiver.writer


		if nil == writer {
			return errNilWriter
		}
	}

	{
		var esc escaper = escaperWrap(writer)

		err := esc.Encode(value)
		if nil != err {
			return err
		}
	}

	{
		err := writeUS(writer)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver Encoder) Flush() error {

	var writer io.Writer
	{
		writer = receiver.writer


		if nil == writer {
			return errNilWriter
		}
	}

	{
		err := writeGS(writer)
		if nil != err {
			return err
		}
	}

	return nil
}

