package main

import (
	"errors"
	"fmt"
	"github.com/pborges/errs"
)

func main() {
	errs.Push(errors.New("rawr"))
	test := errs.Push(errors.New("Test"))

	fmt.Println(test)
}
