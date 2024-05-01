package main

import (
	"errors"
	"fmt"
)

func divison(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("erro divison by por 0")
	}

	return a / b, nil
}

func main() {

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    for i:= 0, i < len(a); i++ {

		v, err := divison(a[i], b[i])
		if err != nil {
			fmt.Println(err)
		    continue 
	}
	fmt.Println(a[i], "/", b[i], "=", v)
	}
}
