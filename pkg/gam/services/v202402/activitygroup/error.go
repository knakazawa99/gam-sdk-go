// Code generated by gamwsdl/main.go. DO NOT EDIT.
package activitygroup

import (
	"errors"
)

func NewActivityError(reason string) error {
	return &ActivityError{
		Reason: reason,
	}
}

func (e *ActivityError) Error() string {
	return "ActivityError error: " + e.Reason
}

func (e *ActivityError) Is(target error) bool {
	var errType *ActivityError
	return errors.As(target, &errType)
}

func NewApiError(reason string) error {
	return &ApiError{
		Reason: reason,
	}
}

func (e *ApiError) Error() string {
	return "ApiError error: " + e.Reason
}

func (e *ApiError) Is(target error) bool {
	var errType *ApiError
	return errors.As(target, &errType)
}

func NewApiException(reason string) error {
	return &ApiException{
		Reason: reason,
	}
}

func (e *ApiException) Error() string {
	return "ApiException error: " + e.Reason
}

func (e *ApiException) Is(target error) bool {
	var errType *ApiException
	return errors.As(target, &errType)
}

func NewApiVersionError(reason string) error {
	return &ApiVersionError{
		Reason: reason,
	}
}

func (e *ApiVersionError) Error() string {
	return "ApiVersionError error: " + e.Reason
}

func (e *ApiVersionError) Is(target error) bool {
	var errType *ApiVersionError
	return errors.As(target, &errType)
}

func NewApplicationException(reason string) error {
	return &ApplicationException{
		Reason: reason,
	}
}

func (e *ApplicationException) Error() string {
	return "ApplicationException error: " + e.Reason
}

func (e *ApplicationException) Is(target error) bool {
	var errType *ApplicationException
	return errors.As(target, &errType)
}

func NewAuthenticationError(reason string) error {
	return &AuthenticationError{
		Reason: reason,
	}
}

func (e *AuthenticationError) Error() string {
	return "AuthenticationError error: " + e.Reason
}

func (e *AuthenticationError) Is(target error) bool {
	var errType *AuthenticationError
	return errors.As(target, &errType)
}

func NewCollectionSizeError(reason string) error {
	return &CollectionSizeError{
		Reason: reason,
	}
}

func (e *CollectionSizeError) Error() string {
	return "CollectionSizeError error: " + e.Reason
}

func (e *CollectionSizeError) Is(target error) bool {
	var errType *CollectionSizeError
	return errors.As(target, &errType)
}

func NewCommonError(reason string) error {
	return &CommonError{
		Reason: reason,
	}
}

func (e *CommonError) Error() string {
	return "CommonError error: " + e.Reason
}

func (e *CommonError) Is(target error) bool {
	var errType *CommonError
	return errors.As(target, &errType)
}

func NewFeatureError(reason string) error {
	return &FeatureError{
		Reason: reason,
	}
}

func (e *FeatureError) Error() string {
	return "FeatureError error: " + e.Reason
}

func (e *FeatureError) Is(target error) bool {
	var errType *FeatureError
	return errors.As(target, &errType)
}

func NewInternalApiError(reason string) error {
	return &InternalApiError{
		Reason: reason,
	}
}

func (e *InternalApiError) Error() string {
	return "InternalApiError error: " + e.Reason
}

func (e *InternalApiError) Is(target error) bool {
	var errType *InternalApiError
	return errors.As(target, &errType)
}

func NewNotNullError(reason string) error {
	return &NotNullError{
		Reason: reason,
	}
}

func (e *NotNullError) Error() string {
	return "NotNullError error: " + e.Reason
}

func (e *NotNullError) Is(target error) bool {
	var errType *NotNullError
	return errors.As(target, &errType)
}

func NewParseError(reason string) error {
	return &ParseError{
		Reason: reason,
	}
}

func (e *ParseError) Error() string {
	return "ParseError error: " + e.Reason
}

func (e *ParseError) Is(target error) bool {
	var errType *ParseError
	return errors.As(target, &errType)
}

func NewPermissionError(reason string) error {
	return &PermissionError{
		Reason: reason,
	}
}

func (e *PermissionError) Error() string {
	return "PermissionError error: " + e.Reason
}

func (e *PermissionError) Is(target error) bool {
	var errType *PermissionError
	return errors.As(target, &errType)
}

func NewPublisherQueryLanguageContextError(reason string) error {
	return &PublisherQueryLanguageContextError{
		Reason: reason,
	}
}

func (e *PublisherQueryLanguageContextError) Error() string {
	return "PublisherQueryLanguageContextError error: " + e.Reason
}

func (e *PublisherQueryLanguageContextError) Is(target error) bool {
	var errType *PublisherQueryLanguageContextError
	return errors.As(target, &errType)
}

func NewPublisherQueryLanguageSyntaxError(reason string) error {
	return &PublisherQueryLanguageSyntaxError{
		Reason: reason,
	}
}

func (e *PublisherQueryLanguageSyntaxError) Error() string {
	return "PublisherQueryLanguageSyntaxError error: " + e.Reason
}

func (e *PublisherQueryLanguageSyntaxError) Is(target error) bool {
	var errType *PublisherQueryLanguageSyntaxError
	return errors.As(target, &errType)
}

func NewQuotaError(reason string) error {
	return &QuotaError{
		Reason: reason,
	}
}

func (e *QuotaError) Error() string {
	return "QuotaError error: " + e.Reason
}

func (e *QuotaError) Is(target error) bool {
	var errType *QuotaError
	return errors.As(target, &errType)
}

func NewRangeError(reason string) error {
	return &RangeError{
		Reason: reason,
	}
}

func (e *RangeError) Error() string {
	return "RangeError error: " + e.Reason
}

func (e *RangeError) Is(target error) bool {
	var errType *RangeError
	return errors.As(target, &errType)
}

func NewRequiredCollectionError(reason string) error {
	return &RequiredCollectionError{
		Reason: reason,
	}
}

func (e *RequiredCollectionError) Error() string {
	return "RequiredCollectionError error: " + e.Reason
}

func (e *RequiredCollectionError) Is(target error) bool {
	var errType *RequiredCollectionError
	return errors.As(target, &errType)
}

func NewRequiredError(reason string) error {
	return &RequiredError{
		Reason: reason,
	}
}

func (e *RequiredError) Error() string {
	return "RequiredError error: " + e.Reason
}

func (e *RequiredError) Is(target error) bool {
	var errType *RequiredError
	return errors.As(target, &errType)
}

func NewRequiredNumberError(reason string) error {
	return &RequiredNumberError{
		Reason: reason,
	}
}

func (e *RequiredNumberError) Error() string {
	return "RequiredNumberError error: " + e.Reason
}

func (e *RequiredNumberError) Is(target error) bool {
	var errType *RequiredNumberError
	return errors.As(target, &errType)
}

func NewServerError(reason string) error {
	return &ServerError{
		Reason: reason,
	}
}

func (e *ServerError) Error() string {
	return "ServerError error: " + e.Reason
}

func (e *ServerError) Is(target error) bool {
	var errType *ServerError
	return errors.As(target, &errType)
}

func NewStatementError(reason string) error {
	return &StatementError{
		Reason: reason,
	}
}

func (e *StatementError) Error() string {
	return "StatementError error: " + e.Reason
}

func (e *StatementError) Is(target error) bool {
	var errType *StatementError
	return errors.As(target, &errType)
}

func NewStringFormatError(reason string) error {
	return &StringFormatError{
		Reason: reason,
	}
}

func (e *StringFormatError) Error() string {
	return "StringFormatError error: " + e.Reason
}

func (e *StringFormatError) Is(target error) bool {
	var errType *StringFormatError
	return errors.As(target, &errType)
}

func NewStringLengthError(reason string) error {
	return &StringLengthError{
		Reason: reason,
	}
}

func (e *StringLengthError) Error() string {
	return "StringLengthError error: " + e.Reason
}

func (e *StringLengthError) Is(target error) bool {
	var errType *StringLengthError
	return errors.As(target, &errType)
}

func NewUniqueError(reason string) error {
	return &UniqueError{
		Reason: reason,
	}
}

func (e *UniqueError) Error() string {
	return "UniqueError error: " + e.Reason
}

func (e *UniqueError) Is(target error) bool {
	var errType *UniqueError
	return errors.As(target, &errType)
}

// RaiseError raises an error based on the error type and reason.
func RaiseError(errorType, reason string) error {
	switch errorType {

	case "ActivityError":
		return NewActivityError(reason)

	case "ApiError":
		return NewApiError(reason)

	case "ApiException":
		return NewApiException(reason)

	case "ApiVersionError":
		return NewApiVersionError(reason)

	case "ApplicationException":
		return NewApplicationException(reason)

	case "AuthenticationError":
		return NewAuthenticationError(reason)

	case "CollectionSizeError":
		return NewCollectionSizeError(reason)

	case "CommonError":
		return NewCommonError(reason)

	case "FeatureError":
		return NewFeatureError(reason)

	case "InternalApiError":
		return NewInternalApiError(reason)

	case "NotNullError":
		return NewNotNullError(reason)

	case "ParseError":
		return NewParseError(reason)

	case "PermissionError":
		return NewPermissionError(reason)

	case "PublisherQueryLanguageContextError":
		return NewPublisherQueryLanguageContextError(reason)

	case "PublisherQueryLanguageSyntaxError":
		return NewPublisherQueryLanguageSyntaxError(reason)

	case "QuotaError":
		return NewQuotaError(reason)

	case "RangeError":
		return NewRangeError(reason)

	case "RequiredCollectionError":
		return NewRequiredCollectionError(reason)

	case "RequiredError":
		return NewRequiredError(reason)

	case "RequiredNumberError":
		return NewRequiredNumberError(reason)

	case "ServerError":
		return NewServerError(reason)

	case "StatementError":
		return NewStatementError(reason)

	case "StringFormatError":
		return NewStringFormatError(reason)

	case "StringLengthError":
		return NewStringLengthError(reason)

	case "UniqueError":
		return NewUniqueError(reason)

	default:
		return nil
	}
}
