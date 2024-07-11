package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.ErrUnsupported

	err = fmt.Errorf("something failed: %w", err)
	fmt.Println(errors.Is(err, errors.ErrUnsupported))

	err = fmt.Errorf("something else failed: %v", err)
	fmt.Println(errors.Is(err, errors.ErrUnsupported))
}
