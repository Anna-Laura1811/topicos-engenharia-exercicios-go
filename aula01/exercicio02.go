package main

import (
	"errors"
	"fmt"
)

func isString(s string) (int, error) {

	if s == "" {
		return 0, errors.New("a string está vazia")
	}

	return len(s), nil
}

func main() {

	v, err := isString("Anna Laura")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}
