package errors

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

type QueryError struct {
	Reason string
}

type DatabaseError struct {
	Reason string
}

type ServiceError struct {
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Error while validating: field '%s': %s.", e.Field, e.Message)
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("Query error: %s.", e.Reason)
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Database error: %s.", e.Reason)
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("Service error: %s.", e.Reason)
}

func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{Field: field, Message: message}
}

func NewQueryError(reason string) *QueryError {
	return &QueryError{Reason: reason}
}

func NewDatabaseError(reason string) *DatabaseError {
	return &DatabaseError{Reason: reason}
}

func NewServiceError(reason string) *ServiceError {
	return &ServiceError{Reason: reason}
}

func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

func IsQueryError(err error) bool {
	_, ok := err.(*QueryError)
	return ok
}

func IsDatabaseError(err error) bool {
	_, ok := err.(*DatabaseError)
	return ok
}

func IsServiceError(err error) bool {
	_, ok := err.(*ServiceError)
	return ok
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
