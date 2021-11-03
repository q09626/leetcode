package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string转[]byte
	by := []byte("string")
	fmt.Println(by)
	// []byte转string
	st := string([]byte{'s', 't', 'r', 'i', 'n', 'g'})
	fmt.Println(st)

	// int转string
	fmt.Println(strconv.Itoa(99))
	// string转int
	s, _ := strconv.Atoi("-88")
	fmt.Println(s)

	// ParseXX函数

	// string转bool
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)

	// string转float
	f, _ := strconv.ParseFloat("12.34", 64)
	fmt.Println(f)

	// string转int
	i, _ := strconv.ParseInt("-45", 10, 64)
	fmt.Println(i)

	// string转uint
	u, _ := strconv.ParseUint("78", 10, 64)
	fmt.Println(u)

	// FormatXX函数

	s1 := strconv.FormatBool(true)
	fmt.Println(s1)

	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(s2)

	s3 := strconv.FormatInt(-42, 16)
	fmt.Println(s3)

	s4 := strconv.FormatUint(42, 16)
	fmt.Println(s4)
}
