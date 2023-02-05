package utils

type DomainErrorType int

const (
	ResourceFoundError = iota
	InvalidInputError
)

func NewError(message string, typeError DomainErrorType) *DomainError {
	return &DomainError{
		message:   message,
		typeError: typeError,
	}
}

type DomainError struct {
	message   string
	typeError DomainErrorType
}

func (e DomainError) Error() string {
	return e.message
}
