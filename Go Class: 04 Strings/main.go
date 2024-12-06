package main

//import "fmt"

//func main() {
//	s := "élite"
//
//	fmt.Printf("%8T %[1]v %d\n", s, len(s))
//	fmt.Printf("%8T %[1]v\n", []rune(s))
//
//	b := []byte(s)
//	fmt.Printf("%8T %[1]v %d\n", b, len(b))

/*
   The length of a string in Go is:
       - The number of bytes required to represent it in UTF-8 encoding.
       - NOT the number of Unicode characters.
*/

//}

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("not enough arguments!")
		os.Exit(-1)
	}
	old, current := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, current)

		fmt.Println(t)
	}
}
