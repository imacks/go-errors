// +build go1.20

package errors

import (
	goerrors "errors"
)

var (
	// As is an alias to standard library errors.As.
	As = goerrors.As
	// Is is an alias to standard library errors.Is.
	Is = goerrors.Is
)
