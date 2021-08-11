package data

import "language-structure/dictionary/engine"

type Languages struct {
	data map[string]*engine.RuleManager
}

func (data Languages) Contains(key string) bool {
	_, ok := data.data[key]
	return ok
}

func (data Languages) Size() int  {
	return len(data.data)
}

func (data Languages) Add(ruleManager *engine.RuleManager)  {
	data.data[ruleManager.Name] = ruleManager
}