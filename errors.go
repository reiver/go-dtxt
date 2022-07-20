package dtxt

import (
	"github.com/reiver/go-fck"
)

const (
	errBegun            = fck.Error("begun")
	errEnded            = fck.Error("ended")
	errNilReceiver      = fck.Error("nil receiver")
	errNilTextMarshaler = fck.Error("nil encoding.TextMarshaler")
	errNilWriter        = fck.Error("nil writer")
	errNotBegun         = fck.Error("not begun")
)
