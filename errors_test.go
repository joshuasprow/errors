package errors

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/rs/zerolog"
)

func TestKind_String(t *testing.T) {
	tests := []struct {
		name string
		k    Kind
		want string
	}{
		{
			name: "unexpected",
			k:    KindUnexpected,
			want: "unexpected",
		},
		{
			name: "(actually) unexpected",
			k:    -1,
			want: "unexpected",
		},
		{
			name: "unmarshal",
			k:    KindUnmarshal,
			want: "unmarshal",
		},
		{
			name: "user",
			k:    KindUser,
			want: "user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("Kind.String() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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
