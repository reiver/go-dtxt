package dtxt

import (
	"github.com/reiver/go-fck"
	"github.com/reiver/go-utf8"

	"encoding"
	"io"
	"unsafe"
)

type unitescaper struct {
	writer io.Writer
}

func unitescaperWrap(writer io.Writer) unitescaper {
	return unitescaper{
		writer:writer,
	}
}

// Encode encodes a unit (i.e., a field) that is gives as a `encoding.TextMarshaler:`, `[]rune`, `string`, or `[]byte`.
func (receiver unitescaper) Encode(value any) error {

	switch casted := value.(type) {
	case encoding.TextMarshaler:
		return receiver.EncodeTextMarshaler(casted)
	case []rune:
		return receiver.EncodeRunes(casted)
	case string:
		return receiver.EncodeString(casted)
	case []byte:
		return receiver.EncodeBytes(casted)
	default:
		return fck.Errorf("cannot escape-encode value of type %T", value)
	}
}

// encodeByte does NOT  encode a unit (i.e., a field).
//
// encodeByte is a helper method used by other methods to encode a single byte.
func (receiver unitescaper) encodeByte(b byte) error {

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

// EncodeBytes encodes a unit (i.e., a field) that is gives as a `[]byte`.
func (receiver unitescaper) EncodeBytes(value []byte) error {

	var writer io.Writer
	{
		writer = receiver.writer

		if nil == writer {
			return errNilWriter
		}
	}

	for _, b := range value {
		err := receiver.encodeByte(b)
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

// encodeRune does NOT  encode a unit (i.e., a field).
//
// encodeRune is a helper method used by other methods to encode a single rune.
func (receiver unitescaper) encodeRune(r rune) error {

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

// EncodeRunes encodes a unit (i.e., a field) that is gives as a `[]rune`.
func (receiver unitescaper) EncodeRunes(value []rune) error {

	var writer io.Writer
	{
		writer = receiver.writer

		if nil == writer {
			return errNilWriter
		}
	}

	for _, r := range value {
		err := receiver.encodeRune(r)
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

// EncodeString encodes a unit (i.e., a field) that is gives as a `string`.
func (receiver unitescaper) EncodeString(value string) error {

	var p []byte
	{
		p = *(*[]byte)(unsafe.Pointer(&value))
	}

	return receiver.EncodeBytes(p)

}

// EncodeTextMarshaler encodes a unit (i.e., a field) that is gives as a `encoding.TextMarshaler`.
func (receiver unitescaper) EncodeTextMarshaler(value encoding.TextMarshaler) error {

	if nil == value {
		return errNilTextMarshaler
	}

	p, err := value.MarshalText()
	if nil != err {
		return err
	}

	return receiver.EncodeBytes(p)
}
