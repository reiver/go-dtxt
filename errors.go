package dtxt

import (
	"github.com/reiver/go-fck"
)

const (
	errNilTextMarshaler = fck.Error("nil encoding.TextMarshaler")
	errNilWriter        = fck.Error("nil writer")
)
