package errors

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

type Op string

type kind int

func (k kind) String() string {
	switch k {
	case KindUnexpected:
		return "unexpected"
	case KindUser:
		return "user"
	default:
		panic(k)
	}
}

const (
	KindUnexpected kind = iota + 1
	KindUser
)

type Error struct {
	Err   error
	Op    Op
	Kind  kind
	Level zerolog.Level
}

func (e *Error) Error() string {
	strs := []string{}

	for _, op := range Ops(e) {
		if op == "" {
			continue
		}

		strs = append(strs, string(op))
	}

	return strings.Join(strs, ": ")
}

func Ops(e *Error) []Op {
	ops := []Op{e.Op}

	subErr, ok := e.Err.(*Error)
	if !ok {
		return ops
	}

	ops = append(ops, Ops(subErr)...)

	return ops
}

func Kind(err error) kind {
	e, ok := err.(*Error)
	if !ok {
		return KindUnexpected
	}

	if e.Kind != 0 {
		return e.Kind
	}

	return Kind(e.Err)
}

func E(args ...interface{}) error {
	e := &Error{}

	for _, arg := range args {
		switch t := arg.(type) {
		case error:
			e.Err = t
		case kind:
			e.Kind = t
		case zerolog.Level:
			e.Level = t
		case Op:
			e.Op = t
		default:
			panic(fmt.Sprintf("failed to find arg type: %q", t))
		}
	}

	return e
}
