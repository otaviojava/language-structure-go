package main

import (
	"fmt"
	"language-structure/dictionary/dto"
	"log"
)

func main() {
	l, err := dto.ReadConf("languages/java.yaml")
	if err != nil {
		log.Fatal(err)
	}
	manager := l.ToRuleManager()
	fmt.Printf("%v", l)
	fmt.Println("The sample code ")
	fmt.Println("The rule manager is", manager)
}