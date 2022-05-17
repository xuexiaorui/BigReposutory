package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := 11112
	b := strconv.FormatInt(int64(a), 10)
	fmt.Printf("b: %v\n", b)

}
