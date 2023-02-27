package main

import "strings"

func reverseWords(s string) string {
	var result string
	sl := strings.Split(s, " ")
	for i := len(sl) - 1; i >= 0; i-- {
		if sl[i] == "" {
			continue
		}
		if result == "" {
			result += sl[i]
		} else {
			result += " " + sl[i]
		}
	}
	return result
}
