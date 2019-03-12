package main

import (
	"fmt"
	"os"
	"strconv"
)

func main1() {
	if len(os.Args) == 1 {
		fmt.Println("Please input some decimals")
		os.Exit(1)
	}
	args := os.Args
	var sum = 0.0
	for i := 1; i < len(args); i++ {
		n, _ := strconv.ParseFloat(args[i], 64)
		sum = sum + n
	}
	fmt.Println("Sum:", sum)
	fmt.Println("Mean:", sum/float64(len(args)))
}
