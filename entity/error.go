package entity

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("Not found")

// ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("Invalid entity")

// ErrCannotBeUpdated cannot be updated
var ErrCannotBeUpdated = errors.New("Can't be updated")

// ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("Can't be deleted")
