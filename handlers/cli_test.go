package handlers

import (
	"bytes"
	"errors"
	"testing"

	calc "github.com/XanSmarty/xan-calc-lib"
)

func assertErr(t *testing.T, actual error, targets ...error) {
	for _, target := range targets {
		if !errors.Is(actual, target) {
			t.Helper()
			t.Errorf("want: %v, got : %v", target, actual)
		}
	}
}

func TestHandler_TwoArgsRequired(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	assertErr(t, err, errWrongArgCount)
}

func TestHandler_FirstArgInvalid(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"INVALID", "42"})
	assertErr(t, err, errInvalidArgument)
}

func TestHandler_SecondArgInvalid(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"42", "INVALID"})
	assertErr(t, err, errInvalidArgument)
}

func TestHandler_ResultWrittenToOutput(t *testing.T) {
	buf := &bytes.Buffer{}

	handler := NewHandler(buf, &calc.Addition{})
	err := handler.Handle([]string{"2", "3"})
	assertErr(t, err, nil)
	if buf.String() != "5" {
		t.Errorf("want: %v, got %v", "5", buf.String())
	}
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}

func TestHandler_OutputWriterError(t *testing.T) {
	boink := errors.New("boink")
	errStdout := &ErringWriter{boink}
	handler := NewHandler(errStdout, &calc.Addition{})
	err := handler.Handle([]string{"2", "3"})
	assertErr(t, err, errOutputWriter, boink)
}

func TestHandler_NoCalculatorProvided(t *testing.T) {
	//handler := NewHandler(nil, nil)
	//err := handler.Handle([]string{})
}
