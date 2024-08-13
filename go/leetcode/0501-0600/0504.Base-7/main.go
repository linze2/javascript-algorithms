package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {
	return strconv.FormatInt(int64(num), 7)
}

func main() {
	fmt.Println(convertToBase7(100)) // Output: 202
}