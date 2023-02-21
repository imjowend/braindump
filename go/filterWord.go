package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	a := "palabra (word)"
	delimitator := " "
	asd := strings.Split(a, delimitator)

	//Solucion hecha por mi
	fmt.Println(strings.TrimFunc(asd[1], func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	//Solucion Google
	fmt.Println(strings.Contains("palabra (word)", "word"))
}
