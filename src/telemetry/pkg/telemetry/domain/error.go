package domain

import "fmt"

type NotFoundError struct {
    Entity string
	Input string
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s with value %s not found", e.Entity, e.Input)
}

type ValidationError struct {
    Entity string
    Field  string
    Msg    string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error for %s: %s %s", e.Entity, e.Field, e.Msg)
}