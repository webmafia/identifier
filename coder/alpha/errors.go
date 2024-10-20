package alpha

import (
	"errors"
	"fmt"
)

const alphaMinChars = 3

var (
	ErrTooShort       = fmt.Errorf("alphabet must be at least %d characters", alphaMinChars)
	ErrDuplicateChars = errors.New("alphabet contains duplicated characters")
)
