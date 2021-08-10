package dictionary

import "regexp"

func NewRules() *RuleManager {
	return NewRuleManager(rules(), extensions())
}

func extensions() []string {
	return []string{".java"}
}

func rules() []Rule {
	var java []Rule
	{
		NewJavaAndMessageDigestIsCustom()
	}
	return java
}

func NewJavaAndMessageDigestIsCustom() TextRule {
	return TextRule{
		Metadata: Metadata{
			ID:          "d34c6b79-4051-4f73-bf8e-37db9becc896",
			Name:        "Message digest is custom",
			Description: "Implementing a custom MessageDigest is error-prone. NIST recommends the use of SHA-224, SHA-256, SHA-384, SHA-512, SHA-512/224, or SHA-512/256. For more information checkout the CWE-327 (https://cwe.mitre.org/data/definitions/327.html) advisory.",
			Severity:    High.ToString(),
			Confidence:  Medium.ToString(),
		},
		Type: AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`extends\sMessageDigest`),
			regexp.MustCompile(`@Override`),
			regexp.MustCompile(`protected\sbyte\[\]\sengineDigest\(\)`),
		},
	}
}