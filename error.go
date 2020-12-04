package errors

import (
	"encoding/json"
	"strings"

	"github.com/rs/zerolog"
)

// TODO: add real comments

// Error is...
type Error struct {
	Err   error
	Kind  Kind
	Level zerolog.Level
	Msg   string
	Op    Op
}

func (e *Error) Error() string {
	ops := []string{}

	for _, op := range GetOps(e) {
		if op == "" {
			continue
		}

		ops = append(ops, string(op))
	}

	return strings.Join(ops, ": ")
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
