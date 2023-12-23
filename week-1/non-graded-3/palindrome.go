package main

import "fmt"

func Palindrome(word string) {
	str := word
	arr := []rune(word)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	convSrting := string(arr)

	if str == convSrting {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
