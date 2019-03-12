package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please input some strings")
		os.Exit(1)
	}
	args := os.Args
	//var sum = 0.0
	for i := 1; i < len(args); i++ {

		fmt.Println(args[i])
		// if args[i]=='STOP' {
		// 	os.ErrClosed("this will stop")
		// }
	}

}
