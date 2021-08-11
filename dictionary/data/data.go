package data

import "language-structure/dictionary/engine"

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
		language.Extensions = append(language.Extensions, ruleManager.Extensions...)
		data.data[language.Name] = language
	} else {
		data.data[ruleManager.Name] = *ruleManager
	}
}

func (data *Languages) Get(key string) (engine.RuleManager, bool) {
	val, ok := data.data[key]
	return val, ok
}
