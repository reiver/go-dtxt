package dtxt

	import (
	"io"
)

type Encoder struct {
	writer io.Writer
	begun bool
	ended bool
}

func EncoderWrap(writer io.Writer) Encoder {
	return Encoder{
		writer:writer,
	}
}

func (receiver *Encoder) Begin() error {
	if nil == receiver {
		return errNilReceiver
	}

	if receiver.begun {
		return errBegun
	}
	if receiver.ended {
		return errEnded
	}

	receiver.begun = true

	return nil
}


func (receiver Encoder) EncodeRow(values ...any) error {

	if !receiver.begun {
		return errNotBegun
	}
	if receiver.ended {
		return errEnded
	}

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

	if !receiver.begun {
		return errNotBegun
	}
	if receiver.ended {
		return errEnded
	}

	var writer io.Writer
	{
		writer = receiver.writer


		if nil == writer {
			return errNilWriter
		}
	}

	{
		var esc unitescaper = unitescaperWrap(writer)

		err := esc.Encode(value)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver *Encoder) End() error {
	if nil == receiver {
		return errNilReceiver
	}

	if !receiver.begun {
		return errNotBegun
	}
	if receiver.ended {
		return errEnded
	}

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
