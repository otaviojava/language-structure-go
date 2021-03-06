package dto

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"language-structure/dictionary/engine"
	"regexp"
)

type Language struct {
	Name       string
	Extensions []string
	Rules      []rule
}

type rule struct {
	Type        string
	Metadata    engine.Metadata
	Expressions []string
}

func (rule rule) toExpressions() []*regexp.Regexp {
	var expressions []*regexp.Regexp
	for _, expression := range rule.Expressions {
		expressions = append(expressions, regexp.MustCompile(expression))
	}
	return expressions
}

func (rule rule) toMatchType() engine.MatchType {
	switch rule.Type {
	case "NotMatch":
		return engine.NotMatch
	case "OrMatch":
		return engine.OrMatch
	case "AndMatch":
		return engine.AndMatch
	default:
		return engine.Regular
	}
}

func (rule rule) toTextRule() engine.TextRule {
	return engine.TextRule{rule.Metadata, rule.toMatchType(), rule.toExpressions()}
}

func (language Language) ToRuleManager() *engine.RuleManager {
	var rules []engine.Rule
	for _, rule := range language.Rules {
		rules = append(rules, rule.toTextRule())
	}
	return engine.NewRuleManager(language.Name, rules, language.Extensions)
}

func ReadConf(filename string) (*Language, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	l := &Language{}
	err = yaml.Unmarshal(buf, l)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return l, nil
}
