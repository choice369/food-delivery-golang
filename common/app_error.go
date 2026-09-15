package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	RootErr    error  `json:"root_error"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorRes(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorRes(root error, statusCode int, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorRes(root, msg, root.Error(), key)
	}

	return NewErrorRes(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewErrorRes(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorRes(err, "invalid req", err.Error(), "ErrInvalidReq")
}

func ErrInternal(err error) *AppError {
	return NewErrorRes(err, "internal error", err.Error(), "InternalErr")
}

func ErrCannotListEntity(err error) *AppError {
	return NewErrorRes(err, "cannot list entity", err.Error(), "CannotListEntityErr")
}

func ErrCannotGetEntity(err error) *AppError {
	return NewErrorRes(err, "cannot get entity", err.Error(), "CannotGetEntityErr")
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewErrorRes(err, "cannot create entity", err.Error(), "CannotCreateEntityErr")
}

func ErrCannotUpdateEntity(err error) *AppError {
	return NewErrorRes(err, "cannot update entity", err.Error(), "CannotUpdateEntityErr")
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewErrorRes(err, "cannot delete entity", err.Error(), "CannotDeleteEntityErr")
}

func ErrNotPermission(err error) *AppError {
	return NewErrorRes(err, "not permission", err.Error(), "NotPermissionErr")
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewErrorRes(err, "entity not found", err.Error(), "EntityNotFoundErr")
}

var RecordNotFound = errors.New("record not found")
