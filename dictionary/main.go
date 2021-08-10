package main

import (
	"fmt"
	"language-structure/dictionary/java"
)

func main()  {
	javaLanguage := java.NewRules()
	fmt.Println(javaLanguage.Extensions)
	fmt.Println(javaLanguage.Rules)
	fmt.Println("Running code")
}
