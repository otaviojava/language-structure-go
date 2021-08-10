package main

import (
	"fmt"
	"language-structure/dictionary/engine"
	"language-structure/dictionary/java"
)

func main()  {
	dictionary := make(map[string]*engine.RuleManager)
	dictionary["java"] = java.NewRules()
	for language := range dictionary {
		fmt.Println(language)
	}
	fmt.Println("Running code ", len(dictionary))
}
