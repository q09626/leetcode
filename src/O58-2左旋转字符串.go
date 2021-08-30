package main

func reverseLeftWords(s string, n int) string {
	s1 := s[:n]
	return s[n:]+s1
}