package errors

// IsJoin returns true if err implements "Unwrap []error".
func IsJoin(err error) bool {
	_, ok := err.(interface {
		Unwrap() []error
	})
	return ok
}

// Split returns the result of calling the Unwrap method on err, if err's type 
// contains an Unwrap method returning []error. Otherwise, Split returns nil.
func Split(err error) []error {
	jerr, ok := err.(interface {
		Unwrap() []error
	})
	if !ok {
		return nil
	}
	return jerr.Unwrap()
}