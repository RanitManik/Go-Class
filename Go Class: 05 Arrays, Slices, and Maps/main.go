package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*func main() {
	t := []byte("string")

	fmt.Println(len(t), t)
	fmt.Println(t[2])
	fmt.Println(t[:2])
	fmt.Println(t[2:])
	fmt.Println(t[3:5], len(t[3:5]))
}
*/

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words found")

	type kv struct {
		key string
		val int
	}

	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, kv := range ss {
		fmt.Println(kv.key, "appears", kv.val, "times")
	}
}
