package generic_error

type GenericError struct {
	Message string `json:"message"`
}

func NewGenericError(err error) GenericError {
	return GenericError{
		Message: err.Error(),
	}
}
