package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type myData struct {
	Conf struct {
		Hits      int64
		Time      int64
		CamelCase string `yaml:"camelCase"`
	}
}

func readConf(filename string) (*myData, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &myData{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}

func main() {
	c, err := readConf("conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", c)
}