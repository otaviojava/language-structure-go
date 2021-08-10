package dto

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"language-structure/dictionary/engine"
)

type Language struct {
	Name        string
	Extensions  []string
	Rules []Rule
}

type Rule struct {
	TYPE string
	Metadata Metadata
	Expressions []string
}

type Metadata struct {
	Id string
	Name string
	Description string
	Severity string
	Confidence string
}

func (language Language) ToRuleManager() *engine.RuleManager {
	rules := []engine.Rule{}
	return engine.NewRuleManager(language.Name, rules, language.Extensions)
}

func ReadConf(filename string) (*Language, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Language{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}