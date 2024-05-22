package gerror

import (
	"errors"
	"testing"
)

func NewUserError() error {
	return errors.New("NewUserError")
}
func NewUserError2() *Error {
	if err := NewUserError(); err != nil {
		return GetError(err)
	}
	return nil
}

func TestGError(t *testing.T) {
	err := NewUserError()
	if err != nil {
		gerr := GetError(err)
		gerr.PrintConsole()
	}
	gerr := NewUserError2()
	if gerr != nil {
		gerr.PrintConsole()
	} else {
		t.Fatalf("%v", "Error TestGError")
	}

}
