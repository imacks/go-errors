// Package errors is a drop-in replacement for the standard library errors 
// package. It contains backport of errors.Join pre-Go 1.20, and the extended 
// IsJoin and Split utility functions for working with multi-error unwraps.
package errors

import (
	goerrors "errors"
)

var (
	// New is an alias to standard library errors.New.
	New = goerrors.New
	// Unwrap is an alias to standard library errors.Unwrap.
	Unwrap = goerrors.Unwrap
)
