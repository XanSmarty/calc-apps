package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	calc "github.com/XanSmarty/xan-calc-lib"
)

type Handler struct {
	stdout     io.Writer
	calculator calc.Calculator
}

func NewHandler(stdout io.Writer, calculator calc.Calculator) *Handler {
	return &Handler{stdout, calculator}
}

func (this *Handler) Handle(args []string) error {
	if this.calculator == nil {
		return errUnsupportedOperation
	}
	if len(args) != 2 {
		return errWrongArgCount
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[0])
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[1])
	}
	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintf(this.stdout, "%d", result)
	if err != nil {
		return fmt.Errorf("%w: %w", errOutputWriter, err)
	}
	return nil
}

var (
	errWrongArgCount        = errors.New("two args required")
	errInvalidArgument      = errors.New("invalid argument")
	errOutputWriter         = errors.New("output failure")
	errUnsupportedOperation = errors.New("unsupported operation")
)
