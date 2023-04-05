package contentAnalysis

import (
	"errors"
)

var (
	// ErrInvalidAnswer is returned when the answer is not valid.
	ErrInvalidAnswer = errors.New("invalid answer")

	// ErrGeneratorCannotBeNil is returned when the generator is nil.
	ErrorGeneratorCannotBeNil = errors.New("generator cannot be nil")
)
