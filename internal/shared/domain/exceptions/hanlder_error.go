package exceptions

import domainerrors "github.com/purplesvage/moneka-ai/internal/shared/domain/errors"

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	appErr, ok := err.(*domainerrors.AppError)
	return ok && appErr.Status == 404
}