package errors

import "errors"

// Domain-specific errors untuk mudah debugging
var (
	ErrNotFound    = errors.New("data tidak ditemukan")
	ErrInvalidData = errors.New("data tidak valid")
	ErrDatabase    = errors.New("terjadi kesalahan di database")
	ErrInternal    = errors.New("terjadi kesalahan internal")
)

// HTTP Response error untuk handler
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Untuk membedakan error dari mana (Repository, Service, Handler)
type AppError struct {
	Layer   string // "repository", "service", "handler"
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func NewRepositoryError(message string, err error) *AppError {
	return &AppError{
		Layer:   "repository",
		Message: message,
		Err:     err,
	}
}

func NewServiceError(message string, err error) *AppError {
	return &AppError{
		Layer:   "service",
		Message: message,
		Err:     err,
	}
}

func NewHandlerError(message string, err error) *AppError {
	return &AppError{
		Layer:   "handler",
		Message: message,
		Err:     err,
	}
}
