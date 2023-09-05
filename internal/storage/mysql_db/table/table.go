package table

import "errors"

var (
	ErrNameTaskNotFound = errors.New("Name task not found")
	ErrNameTaskExists   = errors.New("Name task exists")
)
