package errors

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_concatErrStrings(t *testing.T) {
	t.Run("called once", func(t *testing.T) {
		s1 := "s1"
		s2 := "s2"

		want := "s1: s2"

		got := concatErrStrings(s1, s2)
		if got != want {
			t.Errorf("concatErrStrings() got = %v, want %v", got, want)
		}
	})

	t.Run("missing first string", func(t *testing.T) {
		s1 := ""
		s2 := "s2"

		want := "s2"

		got := concatErrStrings(s1, s2)
		if got != want {
			t.Errorf("concatErrStrings() got = %v, want %v", got, want)
		}
	})

	t.Run("missing second string", func(t *testing.T) {
		s1 := "s1"
		s2 := ""

		want := "s1"

		got := concatErrStrings(s1, s2)
		if got != want {
			t.Errorf("concatErrStrings() got = %v, want %v", got, want)
		}
	})

	t.Run("called repeatedly", func(t *testing.T) {
		ss := []string{"s1", "s2", "s3", "s4"}

		got := ""
		want := "s1: s2: s3: s4"

		for _, s := range ss {
			got = concatErrStrings(got, s)
		}

		if got != want {
			t.Errorf("concatErrStrings() got = %v, want %v", got, want)
		}
	})
}

func TestE(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want Error
	}{
		{
			name: "correctly formats message",
			args: []interface{}{"msg1", "msg2"},
			want: Error{
				Msg: "msg1: msg2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := E(tt.args...).(*Error)
			got := *err
			if !cmp.Equal(got, tt.want) {
				t.Errorf("E() diff = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}
