package dto

import "language-structure/dictionary/engine"

type Language struct {
	Name        string
	Extensions  []string
	Expressions []string
}

func (language Language) ToRuleManager() *engine.RuleManager {
	rules := []engine.Rule{}
	return engine.NewRuleManager(language.Name, rules, language.Extensions)
}
