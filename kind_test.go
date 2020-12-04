package errors

import "testing"

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
