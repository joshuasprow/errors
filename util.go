package errors

import (
	"fmt"

	"github.com/rs/zerolog"
)

func concatMessage(msg string, t string) string {
	return msg + ": " + t
}

// E is...
func E(args ...interface{}) error {
	e := &Error{}

	msg := ""

	for _, arg := range args {
		switch t := arg.(type) {
		case error:
			e.Err = t
		case Kind:
			e.Kind = t
		case zerolog.Level:
			e.Level = t
		case string:
			concatMessage(msg, t)
		case Op:
			e.Op = t
		default:
			fmt.Printf("errors.E: unhandled type: %v", t)
		}
	}

	e.Msg = msg

	return e
}

// GetKind is...
func GetKind(err error) Kind {
	e, ok := err.(*Error)
	if !ok {
		return KindUnexpected
	}

	if e.Kind != 0 {
		return e.Kind
	}

	return GetKind(e.Err)
}

// GetOps is...
func GetOps(e *Error) []Op {
	ops := []Op{e.Op}

	subErr, ok := e.Err.(*Error)
	if !ok {
		return ops
	}

	ops = append(ops, GetOps(subErr)...)

	return ops
}
