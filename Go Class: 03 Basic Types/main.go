package main

//import (
//	"fmt"
//)

//func main() {
//a := 2
//b := 3.1

// https://www.w3schools.com/go/go_formatting_verbs.php
//fmt.Printf("a: %10T %[1]v\n", a)
//fmt.Printf("b: %10T %[1]v\n", b)
//
//a = int(b)
//fmt.Printf("a: %10T %[1]v\n", a)
//
//b = float64(a)
//fmt.Printf("b: %10T %[1]v\n", b)
//}

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var n int

	for {
		var val float64

		_, err := fmt.Fscanln(os.Stdin, &val)

		if err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "No values provided!")
		os.Exit(-1)
	}

	fmt.Println("The average is", sum/float64(n))

}
