package data

import (
	"language-structure/dictionary/engine"
	"strconv"
)

type Languages struct {
	data map[string]engine.RuleManager
}

func Create() Languages {
	l := Languages{make(map[string]engine.RuleManager)}
	return l
}

func (data *Languages) Contains(key string) bool {
	_, ok := data.data[key]
	return ok
}

func (data *Languages) Size() int {
	return len(data.data)
}

func (data *Languages) Add(ruleManager *engine.RuleManager) {
	language, ok := data.data[ruleManager.Name]
	if ok {
		language.Extensions = includeExtensions(language.Extensions, ruleManager.Extensions)
		language.Rules = includeRules(language.Rules, ruleManager.Rules)
		data.data[language.Name] = language
	} else {
		data.data[ruleManager.Name] = *ruleManager
	}
}

func includeRules(rules, newRules []engine.Rule) []engine.Rule {
	check := make(map[string]engine.Rule)
	allRules := append(rules, newRules...)
	rulesMerge := make([]engine.Rule, 0)
	id := 1
	for _, val := range allRules {
		var object interface{} = val
		r, ok := object.(engine.TextRule)
		if ok {
			check[r.ID] = r
		} else {
			check[strconv.Itoa(id)] = val
			id++
		}
	}
	for _, rule := range check {
		rulesMerge = append(rulesMerge, rule)
	}
	return rulesMerge
}
func includeExtensions(extensions, newExtensions []string) []string {
	check := make(map[string]int)
	allExtensions := append(extensions, newExtensions...)
	extensionsMerge := make([]string, 0)
	for _, val := range allExtensions {
		check[val] = 1
	}
	for ext, _ := range check {
		extensionsMerge = append(extensionsMerge, ext)
	}
	return extensionsMerge
}
func (data *Languages) Get(key string) (engine.RuleManager, bool) {
	val, ok := data.data[key]
	return val, ok
}
