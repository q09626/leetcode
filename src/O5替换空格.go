package main

import "fmt"

func replaceSpace(s string) string {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			n := len(str)
			str = append(str, str[n-2:]...)
			for k := n-1; k > i+2; k-- {
				str[k] = str[k-2]
			}
			str[i] = '%'
			str[i+1] = '2'
			str[i+2] = '0'
		}
	}
	return string(str)
}

func main() {
	s := "We are happy"
	fmt.Println(replaceSpace(s))
}