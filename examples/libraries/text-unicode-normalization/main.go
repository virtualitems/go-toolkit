package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	str1 := "\u00F1"  // "ñ" (precomposed)
	str2 := "n\u0303" // "n" + "~" (decomposed)

	fmt.Println("not normalized:", str1 == str2)                               // false
	fmt.Println("normalized:", norm.NFC.String(str1) == norm.NFC.String(str2)) // true
}
