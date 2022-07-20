package dtxt

import (
	"github.com/reiver/go-fck"
)

const (
	errNilBytes = fck.Error("nil bytes")
)

type BytesTextMarshaler struct {
	Bytes []byte
}

func (receiver BytesTextMarshaler) MarshalText() ([]byte, error) {
	if nil == receiver.Bytes {
		return nil, errNilBytes
	}

	return append([]byte(nil), receiver.Bytes...), nil
}

