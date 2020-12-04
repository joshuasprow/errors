package errors

import (
	"encoding/json"
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
	case KindUnmarshal:
		return "unmarshal"
	case KindUser:
		return "user"
	default:
		panic(k)
	}
}

const (
	// KindUnexpected is...
	KindUnexpected Kind = iota + 1
	// KindUnmarshal is...
	KindUnmarshal
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

// MarshalJSON is...
func (e *Error) MarshalJSON() ([]byte, error) {
	type Alias Error
	return json.Marshal(&struct {
		Err   error
		Op    Op
		Kind  Kind
		Level zerolog.Level
	}{
		Err:   e.Err,
		Op:    e.Op,
		Kind:  e.Kind,
		Level: e.Level,
	})
}

// UnmarshalJSON is... based on (copied from) http://choly.ca/post/go-json-marshalling/
func (e *Error) UnmarshalJSON(data []byte) error {
	var op Op = "Error.UnmarshalJSON"

	type Alias Error

	aux := &struct {
		// LastSeen int64 `json:"lastSeen"`
		Err   error         `json:"err"`
		Op    Op            `json:"op"`
		Kind  Kind          `json:"kind"`
		Level zerolog.Level `json:"level"`
		*Alias
	}{
		Err:   e.Err,
		Op:    e.Op,
		Kind:  e.Kind,
		Level: e.Level,
		Alias: (*Alias)(e),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return E(op, KindUnmarshal, err, "json.Unmarshal")
	}

	// u.LastSeen = time.Unix(aux.LastSeen, 0)

	return nil
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
