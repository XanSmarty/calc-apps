package main

import (
	"log"
	"os"

	"github.com/XanSmarty/calc-apps/handlers"
	calc "github.com/XanSmarty/xan-calc-lib"
)

func main() {
	handler := handlers.NewHandler(os.Stdout, &calc.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
