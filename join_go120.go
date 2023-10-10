// +build go1.20

package errors

import (
	goerrors "errors"
)

// Join is an alias to standard library errors.Join.
var Join = goerrors.Join
