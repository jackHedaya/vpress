package main

import (
	"fmt"
	"io/ioutil"
	"pbytearray"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dat, err := ioutil.ReadFile("../test_file")
	check(err)
	fmt.Println(dat)

	var i pByteArray.pByteArray
}
