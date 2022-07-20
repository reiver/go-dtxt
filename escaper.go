package dtxt

import (
	"github.com/reiver/go-fck"
	"github.com/reiver/go-utf8"

	"encoding"
	"io"
	"unsafe"
)

type escaper struct {
	writer io.Writer
}

func escaperWrap(writer io.Writer) escaper {
	return escaper{
		writer:writer,
	}
}

func (receiver escaper) Encode(value any) error {

	switch casted := value.(type) {
	case encoding.TextMarshaler:
		return receiver.EncodeTextMarshaler(casted)
	case []rune:
		return receiver.EncodeRunes(casted)
	case string:
		return receiver.EncodeString(casted)
	case []byte:
		return receiver.EncodeBytes(casted)
	case rune:
		return receiver.EncodeRune(casted)
	default:
		return fck.Errorf("cannot escape escape-encode value of type %T", value)
	}
}

func (receiver escaper) EncodeByte(b byte) error {

	var writer io.Writer
	{
		writer = receiver.writer

		if nil == writer {
			return errNilWriter
		}
	}

	{
		switch b {
		case esc,fs,gs,rs,us:
			err := writeESC(writer)
			if nil != err {
				return err
			}
		}
	}

	{
		err := writeByte(writer, b)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver escaper) EncodeBytes(value []byte) error {

	var writer io.Writer
	{
		writer = receiver.writer

		if nil == writer {
			return errNilWriter
		}
	}

	for _, b := range value {
		err := receiver.EncodeByte(b)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver escaper) EncodeRune(r rune) error {

	var writer io.Writer
	{
		writer = receiver.writer

		if nil == writer {
			return errNilWriter
		}
	}

	{
		switch r {
		case esc,fs,gs,rs,us:
			err := writeESC(writer)
			if nil != err {
				return err
			}
		}
	}

	{
		_, err := utf8.WriteRune(writer, r)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver escaper) EncodeRunes(value []rune) error {

	var writer io.Writer
	{
		writer = receiver.writer

		if nil == writer {
			return errNilWriter
		}
	}

	for _, r := range value {
		err := receiver.EncodeRune(r)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver escaper) EncodeString(value string) error {

	var p []byte
	{
		p = *(*[]byte)(unsafe.Pointer(&value))
	}

	return receiver.EncodeBytes(p)

}

func (receiver escaper) EncodeTextMarshaler(value encoding.TextMarshaler) error {

	if nil == value {
		return errNilTextMarshaler
	}

	p, err := value.MarshalText()
	if nil != err {
		return err
	}

	return receiver.EncodeBytes(p)
}
