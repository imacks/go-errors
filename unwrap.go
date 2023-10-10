package errors

// UnwrapAll unwraps the error chain wrapped by err recursively and returns the 
// first error that returns nil on calling Unwrap. Returns nil if err is nil or 
// does not implement Unwrap method.
func UnwrapAll(err error) error {
	for {
		if innerErr := Unwrap(err); innerErr != nil {
			err = innerErr
			continue
		}
		break
	}
	return err
}