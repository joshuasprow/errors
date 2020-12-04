package errors

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
