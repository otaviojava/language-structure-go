package main

import (
	"fmt"
	"language-structure/dictionary/dto"
	"log"
)


func main() {
	c, err := dto.ReadConf("languages/java.yaml")
	if err != nil {
		log.Fatal(err)
	}
	manager := c.ToRuleManager()
	fmt.Printf("%v", c)
	fmt.Println("The sample code: ")
	fmt.Println("The rule manager is", manager)
}