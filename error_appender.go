package envvars

import (
	"errors"
	"strings"
)

type ErrorAppender struct {
	errs []string
	sep  string
}

func NewErrorAppender(sep string) *ErrorAppender {
	return &ErrorAppender{
		errs: make([]string, 0),
		sep:  sep,
	}
}

func (a *ErrorAppender) AppendError(e error) {
	if e != nil {
		a.errs = append(a.errs, e.Error())
	}
}

func (a *ErrorAppender) AppendString(e string) {
	if e != "" {
		a.errs = append(a.errs, e)
	}
}

func (a *ErrorAppender) Error() error {
	if len(a.errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(a.errs, a.sep))
}

// Wrap prepends a string to the error if any
func (a *ErrorAppender) Wrap(s string) error {
	e := a.Error()
	if e == nil {
		return nil
	}
	return errors.New(s + e.Error())
}
