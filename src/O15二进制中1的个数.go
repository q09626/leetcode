package main

import "fmt"

func hammingWeight(num uint32) int {
	var ones int
	for i := 0; i < 32; i++ {
		if 1 << i & num > 0 {
			ones++
		}
	}
	return ones
}

func main() {
	fmt.Println(hammingWeight(11))
	fmt.Println(hammingWeight(128))
	fmt.Println(hammingWeight(4294967293))
}