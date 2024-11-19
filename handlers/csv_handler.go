package handlers

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	calc "github.com/XanSmarty/xan-calc-lib"
)

type CSVHandler struct {
	logger      *log.Logger
	input       *csv.Reader
	output      *csv.Writer
	calculators map[string]calc.Calculator
}

func NewCSVHandler(logger *log.Logger, input io.Reader, output io.Writer, calculators map[string]calc.Calculator) *CSVHandler {
	return &CSVHandler{
		logger:      logger,
		input:       csv.NewReader(input),
		output:      csv.NewWriter(output),
		calculators: calculators,
	}
}

func (this *CSVHandler) Handle() (err error) {
	defer func() {
		this.output.Flush()
		if err == nil {
			err = this.output.Error()
		}
	}()
	for {
		record, err := this.input.Read()
		if err == io.EOF {
			break
		}
		if len(record) != 3 {
			this.logger.Println("must provide exactly 3 fields")
			continue
		}
		if err != nil {
			return err
		}
		a, err := strconv.Atoi(record[0])
		if err != nil {
			this.logger.Println("invalid arg:", record[0])
			continue
		}
		op := record[1]
		b, err := strconv.Atoi(record[2])
		if err != nil {
			this.logger.Println("invalid arg:", record[2])
			continue
		}
		calculator, ok := this.calculators[op]
		if !ok {
			this.logger.Println("unsupported operator:", op)
			continue
		}
		c := calculator.Calculate(a, b)
		_ = this.output.Write(append(record, strconv.Itoa(c)))
	}
	return this.output.Error()
}
