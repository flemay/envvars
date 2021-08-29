package errorappender

import (
	"errors"
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
	got := a.Error()

	// then
	want := "err1err2"
	if got.Error() != want {
		t.Errorf("want %q got %q", want, got.Error())
	}
}

func TestErrorAppender_AppendError_toNotAppendNilError(t *testing.T) {
	// given
	a := NewErrorAppender("")

	// when
	a.AppendError(nil)
	got := a.Error()

	// then
	if got != nil {
		t.Errorf("want no error, got %q", got.Error())
	}
}

func TestErrorAppender_AppendString(t *testing.T) {
	// given
	a := NewErrorAppender("")

	// when
	a.AppendString("err1")
	a.AppendString("err2")
	got := a.Error()

	// then
	want := "err1err2"
	if got.Error() != want {
		t.Errorf("want %q, got %q", want, got.Error())
	}
}

func TestErrorAppender_AppendString_toNotAppendEmptyString(t *testing.T) {
	// given
	a := NewErrorAppender("")

	// when
	a.AppendString("")
	got := a.Error()

	// then
	if got != nil {
		t.Errorf("want no error, got %q", got.Error())
	}
}

func TestErrorAppender_Error_toReturnNilIfNoError(t *testing.T) {
	// given
	a := NewErrorAppender("")

	// when
	got := a.Error()

	// then
	if got != nil {
		t.Errorf("want no error, got %q", got.Error())
	}
}

func TestErrorAppender_Error_toJoinErrorsWithSeparator(t *testing.T) {
	// given
	a := NewErrorAppender("; ")
	err1 := errors.New("err1")
	err2 := errors.New("err2")

	// when
	a.AppendError(err1)
	a.AppendError(err2)
	got := a.Error()

	// then
	want := "err1; err2"
	if got.Error() != want {
		t.Errorf("want %q got %q", want, got.Error())
	}
}

func TestErrorAppender_Wrap(t *testing.T) {
	// given
	a := NewErrorAppender("; ")
	err1 := errors.New("err1")
	err2 := errors.New("err2")

	// when
	a.AppendError(err1)
	a.AppendError(err2)
	got := a.Wrap("errors: ")

	// then
	want := "errors: err1; err2"
	if got.Error() != want {
		t.Errorf("want %q got %q", want, got.Error())
	}
}

func TestErrorAppender_Wrap_toReturnNilIfNoError(t *testing.T) {
	// given
	a := NewErrorAppender("; ")

	// when
	got := a.Wrap("errors: ")

	// then
	if got != nil {
		t.Errorf("want no error, got %q", got.Error())
	}
}
