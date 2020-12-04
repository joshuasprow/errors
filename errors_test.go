package errors

import (
	"errors"
	"testing"

	"github.com/rs/zerolog"
	"k8s.io/apimachinery/pkg/util/json"
)

func TestError_UnmarshalJSON(t *testing.T) {
	t.Run("no bytes", func(t *testing.T) {
		e := &Error{
			Err:   errors.New("test error"),
			Op:    "testOp",
			Kind:  KindUnexpected,
			Level: zerolog.TraceLevel,
		}

		err := e.UnmarshalJSON([]byte{})
		if err == nil {
			t.Errorf("Error.UnmarshalJSON() error = %v, wantErr true", err)
		}
	})

	t.Run("empty fields", func(t *testing.T) {
		e := &Error{
			Err:   errors.New("test error"),
			Op:    "testOp",
			Kind:  KindUnexpected,
			Level: zerolog.TraceLevel,
		}

		p, err := json.Marshal(e)

		err = e.UnmarshalJSON(p)
		if err != nil {
			t.Errorf("Error.UnmarshalJSON() error = %v, wantErr false", err)
		}
	})
}
