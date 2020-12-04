package errors

import (
	"fmt"

	"github.com/rs/zerolog"
)

func concatErrStrings(s1 string, s2 string) string {
	if s1 == "" {
		return s2
	}
	if s2 == "" {
		return s1
	}
	return s1 + ": " + s2
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
			concatErrStrings(msg, t)
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
