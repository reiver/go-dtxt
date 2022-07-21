package dtxt

import (
	"github.com/reiver/go-fck"
	"github.com/reiver/go-utf8"

	"bytes"
	"encoding"
	"io"
	"strings"
)

type unescaper struct {
	reader io.Reader
}

func unescaperWrap(reader io.Reader) unescaper {
	return unescaper{
		reader:reader,
	}
}

func (receiver unescaper) Decode(dst any) error {

	switch casted := dst.(type) {
	case encoding.TextUnmarshaler:
		return receiver.DecodeToTextUnmarshaler(casted)

	case io.Writer:
		return receiver.DecodeToWriter(casted)
	case *[]rune:
		return receiver.DecodeToRunes(casted)
	case *string:
		return receiver.DecodeToString(casted)
	case *[]byte:
		return receiver.DecodeToBytes(casted)
	default:
		return fck.Errorf("cannot descape-decode into value of type %T", dst)
	}
}

func (receiver unescaper) DecodeToBytes(dst *[]byte) error {

	if nil == dst {
		return errNilDestination
	}

	var buffer bytes.Buffer
	{
		var writer io.Writer = &buffer

		err := receiver.DecodeToWriter(writer)
		if nil != err {
			return err
		}
	}

	*dst = append([]byte(nil), buffer.Bytes()...)

	return nil
}

func (receiver unescaper) DecodeToRunes(dst *[]rune) error {

	if nil == dst {
		return errNilDestination
	}

	var p []byte
	{
		err := receiver.DecodeToBytes(&p)
		if nil != err {
			return err
		}
	}

	*dst = bytes.Runes(p)

	return nil
}

func (receiver unescaper) DecodeToString(dst *string) error {

	if nil == dst {
		return errNilDestination
	}

	var buffer strings.Builder
	{
		var writer io.Writer = &buffer

		err := receiver.DecodeToWriter(writer)
		if nil != err {
			return err
		}
	}

	*dst = buffer.String()

	return nil
}

func (receiver unescaper) DecodeToTextUnmarshaler(dst encoding.TextUnmarshaler) error {

	if nil == dst {
		return errNilDestination
	}

	var buffer bytes.Buffer
	{
		var writer io.Writer = &buffer

		err := receiver.DecodeToWriter(writer)
		if nil != err {
			return err
		}
	}

	{
		err := dst.UnmarshalText(buffer.Bytes())
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver unescaper) DecodeToWriter(writer io.Writer) error {

	if nil == writer {
		return errNilWriter
	}

	var reader io.Reader
	{
		reader = receiver.reader

		if nil == reader {
			return errNilReader
		}
	}

	for {
		var r rune
		{
			var err error

			r, _, err = readRune(reader)
			if nil != err {
				return err
			}
		}

		switch r {
		case esc:
			{
				var err error

				r, _, err = readRune(reader)
				if nil != err {
					return err
				}
			}
		case us:
			return nil
		}

		{
			_, err := utf8.WriteRune(writer, r)
			if nil != err {
				return err
			}
		}
	}
}
