package schema

import (
	"errors"
	"fmt"
)

type Error struct {
	Msg string
	Pos Position
	Err error // Wrapped error, if any
}

func (e Error) Error() string { return fmt.Sprintf("%s: %s", e.Pos, e.Msg) }
func (e Error) Unwrap() error { return e.Err }

func Errorf(pos Position, format string, args ...any) Error {
	return Error{Msg: fmt.Sprintf(format, args...), Pos: pos}
}

func Wrapf(pos Position, err error, format string, args ...any) Error {
	if format == "" {
		format = "%s"
	} else {
		format += ": %s"
	}
	// Propagate existing error position if available
	var newPos Position
	if perr := (Error{}); errors.As(err, &perr) {
		newPos = perr.Pos
		args = append(args, perr.Msg)
	} else {
		newPos = pos
		args = append(args, err)
	}
	return Error{Msg: fmt.Sprintf(format, args...), Pos: newPos, Err: err}
}
