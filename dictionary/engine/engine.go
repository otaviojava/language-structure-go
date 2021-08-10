package engine

import "regexp"


// UnitType defines which type of content, and therefore, which kind of rule
// is needed in order to extract information about the program we are analyzing
type UnitType int
type Severity string
type MatchType int

const (
	Regular MatchType = iota
	NotMatch
	OrMatch
	AndMatch
)
const (
	Critical Severity = "CRITICAL"
	High     Severity = "HIGH"
	Medium   Severity = "MEDIUM"
	Low      Severity = "LOW"
	Unknown  Severity = "UNKNOWN"
	Info     Severity = "INFO"
)

const (
	ProgramTextUnit UnitType = iota
	StructuredDataUnit
)

// RuleManager is a generic implementation of formatters.RuleManager
// that can be reused between all engines to load rules
type RuleManager struct {
	Name string
	Rules      []Rule
	Extensions []string
}

// Rule defines a generic rule for any kind of analysis the engine have to execute
type Rule interface {
	IsFor(UnitType) bool // Indicates which kind of program unit this rules can be ran on
}

// Metadata holds information for the rule to match a useful advisory
type Metadata struct {
	ID          string
	Name        string
	CodeSample  string
	Description string

	// Metadata levels
	Severity   string
	Confidence string
}

type TextRule struct {
	Metadata
	Type        MatchType
	Expressions []*regexp.Regexp
}

// nolint // create pointer is not necessary for now
func (rule TextRule) IsFor(unitType UnitType) bool {
	return ProgramTextUnit == unitType
}
func (s Severity) ToString() string {
	return string(s)
}


func NewRuleManager(name string, rules []Rule, extensions []string) *RuleManager {
	return &RuleManager{
		Name: name,
		Rules:      rules,
		Extensions: extensions,
	}
}
