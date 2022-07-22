package dtxt

import (
	"github.com/reiver/go-fck"
)

const (
	GS = fck.Error("GS") // ‘GS’ represents the end of a table. ‘GS’ means ‘Group Separator’. ‘GS’ also sometimes called a ‘Table Separator’.
)

const (
	errBegun            = fck.Error("begun")
	errEnded            = fck.Error("ended")
	errNilDestination   = fck.Error("nil destination")
	errNilReader        = fck.Error("nil reader")
	errNilReceiver      = fck.Error("nil receiver")
	errNilTextMarshaler = fck.Error("nil encoding.TextMarshaler")
	errNilWriter        = fck.Error("nil writer")
	errNotBegun         = fck.Error("not begun")
	errNotGS            = fck.Error("not group separator")
	errNotRS            = fck.Error("not row separator")
	errNotUS            = fck.Error("not unit separator")
)
