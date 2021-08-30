package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s1 := strings.Fields(s)
	var res string
	for i := len(s1)-1; i >= 0; i-- {
		if i != len(s1)-1 {
			res += " "
		}
		res += s1[i]
	}
	return res
}

//func reverseWords(s string) string {
//	str := strings.Fields(s)
//	var res string
//	for i, v := range str {
//		if i != 0 {
//			res += " "
//		}
//		res += reverse(string(v))
//	}
//	return res
//}
//
//func reverse(s string) string {
//	b := []byte(s)
//	l, r := 0, len(b)-1
//	for l < r {
//		b[l], b[r] = b[r], b[l]
//		l++
//		r--
//	}
//	return string(b)
//}

func main() {
	fmt.Println(reverseWords("   hello world!  "))
}