package webhooks

import (
	"encoding/json"
	"errors"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
)

const (
	UnknownErrorType                   = "UnknownError"
	UnknownErrorMessage                = "An unknown error has occurred"
	ImmutableFieldErrorType            = "ImmutableFieldError"
	ImmutableFieldErrorMessageTemplate = "'%s' field is immutable"
)

type ValidationError struct {
	Type    string `json:"validationErrorType"`
	Message string `json:"message"`
}

func (v ValidationError) Error() string {
	return "ValidationError-" + v.Type + ": " + v.Message
}

func (v ValidationError) GetMessage() string {
	return v.Message
}

func (v ValidationError) ExportJSONError() error {
	bytes, err := json.Marshal(v)
	if err != nil { // This (probably) can't fail, untested
		return err
	}

	return errors.New(string(bytes))
}

func WebhookErrorToValidationError(err error) (ValidationError, bool) {
	statusErr := new(k8serrors.StatusError)
	if !errors.As(err, &statusErr) {
		return ValidationError{}, false
	}

	validationErr := new(ValidationError)
	if err := json.Unmarshal([]byte(statusErr.Status().Reason), validationErr); err != nil {
		return ValidationError{}, false
	}

	return *validationErr, true
}
