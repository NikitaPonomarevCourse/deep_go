package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	ErrorsArr []error
}

func (e *MultiError) Error() string {
	answ := fmt.Sprintf("%v errors occured:\n", len(e.ErrorsArr))
	for _, v := range e.ErrorsArr {
		answ += "\t* " + v.Error()
	}
	if len(e.ErrorsArr) == 0 {
		answ += "<nil>"
	}
	answ += "\n"
	return answ
}

func Append(err error, errs ...error) *MultiError {
	if err == nil {
		return &MultiError{
			ErrorsArr: append([]error{}, errs...),
		}
	} else if multiError, ok := err.(*MultiError); ok {
		multiError.ErrorsArr = append(multiError.ErrorsArr, errs...)
		return multiError
	}
	return &MultiError{
		ErrorsArr: append([]error{err}, errs...),
	}
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
