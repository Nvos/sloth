package main

import (
	"errors"
	"fmt"
	"sloth"
)

var notFound = errors.New("not found")

func main() {
	err := fmt.Errorf("wrapper 1: %w", notFound)
	err = fmt.Errorf("wrapper 2: %w", err)
	err = sloth.Wrap(err).WithOp("main")
	err = fmt.Errorf("wrapper 3: %w", err)
	err = fmt.Errorf("wrapper 4: %w", err)
	err = sloth.Wrap(err).WithOp("main")
	err = sloth.Wrap(err).WithOp("main")
	err = sloth.Wrap(err).WithOp("main")
	err = sloth.Wrap(err).WithOp("main")
	err = sloth.Wrap(err).WithOp("main")
	err = sloth.Wrap(err).WithOp("main")
	err = sloth.Wrap(err).WithOp("main")
	err = sloth.Wrap(err).WithOp("main").WithInternalMessage(fmt.Sprintf("%v", "user"))



	if errors.Is(err, notFound) {
		println("it is actually not found!")
		println(err.Error())
	}
}
