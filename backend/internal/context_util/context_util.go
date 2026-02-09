package contextutil

import (
	"context"
	"errors"
)

func IsContextErr(err error) bool {
	return errors.Is(err, context.Canceled) ||
		errors.Is(err, context.DeadlineExceeded)
}
