package main

import "errors"

func NewUserError() error {
	return errors.New("NewUserError")
}
func NewUserError2() *Error {
	if err := NewUserError(); err != nil {
		return GetError(err)
	}
	return nil
}

func main() {
	err := NewUserError()
	if err != nil {
		gerr := GetError(err)
		gerr.PrintConsole()
	}
	gerr := NewUserError2()
	if gerr != nil {
		gerr.PrintConsole()
	}
}
