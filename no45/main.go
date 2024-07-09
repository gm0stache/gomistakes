package main

import (
	"errors"
	"fmt"
	"log"
)

type ErrNotEven struct {
	err error
}

func (ene *ErrNotEven) Add(err error) {
	ene.err = errors.Join(ene.err, err)
}

func (ene *ErrNotEven) Error() string {
	return ene.err.Error()
}

func assertEven(num *int) error {
	var err *ErrNotEven

	if num == nil {
		err = &ErrNotEven{}
		err.Add(errors.New("argument can not be null"))
	}

	if (*num % 2) != 0 {
		if err == nil {
			err = &ErrNotEven{}
		}
		err.Add(fmt.Errorf("number '%d' is not even", *num))
	}

	return err
}

func main() {
	num := 2
	if err := assertEven(&num); err != nil {
		log.Fatalln(err)
	}
}
