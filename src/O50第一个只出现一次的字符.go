package main

func firstUniqChar(s string) byte {
	dic := make(map[byte]int)
	sc := []byte(s)
	for _, c := range sc {
		dic[c]++
	}
	for _, c := range sc {
		if dic[c] == 1 {
			return c
		}
	}
	return ' '
}