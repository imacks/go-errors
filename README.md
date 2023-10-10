go-errors
=========
Backport of errors.Join in Go standard lib.

This package is for people who cannot bump their module to Go 1.20, but still want to use the `errors.Join` function to create multi-errors. It works exactly like the standard library.

Additional functions I have added because they make my life a bit easier:

  - `UnwrapAll(err error) error`: recursively calls `Unwrap` and returns the first error that does not implement `Unwrap() error`.
  - `IsJoin(err error) bool`: returns true if err implements `Unwrap() []error`.
  - `Split(err error) []error`: if err implements `Unwrap() []error`, returns the result of calling `Unwrap`. Otherwise returns nil.
