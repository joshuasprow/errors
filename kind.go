package errors

// Kind is...
type Kind int

func (k Kind) String() string {
	unexpected := "unexpected"

	switch k {
	case KindUnexpected:
		return unexpected
	case KindUnmarshal:
		return "unmarshal"
	case KindUser:
		return "user"
	default:
		return unexpected
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
