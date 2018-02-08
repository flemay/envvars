package envvars

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorAppender_AppendError(t *testing.T) {
	// given
	a := NewErrorAppender("")
	err1 := errors.New("err1")
	err2 := errors.New("err2")

	// when
	a.AppendError(err1)
	a.AppendError(err2)
	err := a.Error()

	// then
	assert.EqualError(t, err, "err1err2")
}

func TestErrorAppender_AppendError_toNotAppendNilError(t *testing.T) {
	// given
	a := NewErrorAppender("")

	// when
	a.AppendError(nil)
	err := a.Error()

	// then
	assert.NoError(t, err)
}

func TestErrorAppender_AppendString(t *testing.T) {
	// given
	a := NewErrorAppender("")

	// when
	a.AppendString("err1")
	a.AppendString("err2")
	err := a.Error()

	// then
	assert.EqualError(t, err, "err1err2")
}

func TestErrorAppender_AppendString_toNotAppendEmptyString(t *testing.T) {
	// given
	a := NewErrorAppender("")

	// when
	a.AppendString("")
	err := a.Error()

	// then
	assert.NoError(t, err)
}

func TestErrorAppender_Error_toReturnNilIfNoError(t *testing.T) {
	// given
	a := NewErrorAppender("")

	// when
	err := a.Error()

	// then
	assert.NoError(t, err)
}

func TestErrorAppender_Error_toJoinErrorsWithSeparator(t *testing.T) {
	// given
	a := NewErrorAppender("; ")
	err1 := errors.New("err1")
	err2 := errors.New("err2")

	// when
	a.AppendError(err1)
	a.AppendError(err2)
	err := a.Error()

	// then
	assert.EqualError(t, err, "err1; err2")
}

func TestErrorAppender_Wrap(t *testing.T) {
	// given
	a := NewErrorAppender("; ")
	err1 := errors.New("err1")
	err2 := errors.New("err2")

	// when
	a.AppendError(err1)
	a.AppendError(err2)
	err := a.Wrap("errors: ")

	// then
	assert.EqualError(t, err, "errors: err1; err2")
}

func TestErrorAppender_Wrap_toReturnNilIfNoError(t *testing.T) {
	// given
	a := NewErrorAppender("; ")

	// when
	err := a.Wrap("errors: ")

	// then
	assert.NoError(t, err)
}
