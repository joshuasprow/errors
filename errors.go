package errors

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

// TODO: add real comments

// Op is..,
type Op string

// Kind is...
type Kind int

func (k Kind) String() string {
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
	// KindUnexpected is...
	KindUnexpected Kind = iota + 1
	// KindUser is...
	KindUser
)

// Error is...
type Error struct {
	Err   error
	Op    Op
	Kind  Kind
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

// Ops is...
func Ops(e *Error) []Op {
	ops := []Op{e.Op}

	subErr, ok := e.Err.(*Error)
	if !ok {
		return ops
	}

	ops = append(ops, Ops(subErr)...)

	return ops
}

// New is...
func New(err error) Kind {
	e, ok := err.(*Error)
	if !ok {
		return KindUnexpected
	}

	if e.Kind != 0 {
		return e.Kind
	}

	return New(e.Err)
}

// E is...
func E(args ...interface{}) error {
	e := &Error{}

	for _, arg := range args {
		switch t := arg.(type) {
		case error:
			e.Err = t
		case Kind:
			e.Kind = t
		case zerolog.Level:
			e.Level = t
		case Op:
			e.Op = t
		default:
			fmt.Printf("")
		}
	}

	return e
}
