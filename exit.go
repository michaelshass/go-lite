package main

// ExitCode type of the code indicating the exit reason.
type ExitCode int

const (
	// ExitSuccess code 0 - program excited successfully without an error.
	ExitSuccess ExitCode = iota
	// ExitReadInputError code 1 - an error occured while reading the user input.
	ExitReadInputError
)
