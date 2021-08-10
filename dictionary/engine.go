package dictionary

// RuleManager is a generic implementation of formatters.RuleManager
// that can be reused between all engines to load rules
type RuleManager struct {
	rules      []Rule
	extensions []string
}

// Rule defines a generic rule for any kind of analysis the engine have to execute
type Rule interface {
	IsFor(UnitType) bool // Indicates which kind of program unit this rules can be ran on
}