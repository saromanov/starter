package cmd

import "errors"

var (
	errNoName   = errors.New("name is not defined")
	errNoAuthor = errors.New("author is not defined")
)
