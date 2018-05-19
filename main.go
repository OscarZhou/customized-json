package main

import (
	"fmt"
)

func main() {
	cp, err := New()
	if err != nil {
		panic(err)
	}

	cp.Add("Method", "GET")
	fmt.Println(cp)

	err = cp.CreateFile("test")
	if err != nil {
		panic(err)
	}
}
