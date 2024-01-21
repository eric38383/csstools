package parser

import (
	"strings"
)

// var selectorType = []string{
// 	"attribute",
// 	"class",
// 	"id",
// 	"nesting",
// 	"type",
// 	"universal",
// }

var regAtRules = []string{
	"@charset",
	"@import",
	"@namespace",
}

// These rules don't have nested styles within them
var unnestedRules = []string{
	"@color-profile",
	"@counter-style",
	"@font-face",
	"@font-values-palette",
	"@keyframes",
}

// this rules can be nested and also have nested styles
var nestedAtRules = []string{
	"@container",
	"@font-feature-values",
	"@layer",
	"@media",
	"@page",
	"@property",
	"@supports",
	//"@document", EXPERIEMENTAL
	//"@scope", EXPERIMENTAL
	//"@starting-style", EXPERIMENTAL
}

type Rule struct {
	Name         string
	Declarations []Declaration
	AtRule       bool
	Minwidth     string
	Maxwidth     string
}

type Declaration struct {
	Property string
	Value    string
}

func ParseRule(ruleblock []string, atRule bool, minwidth string, maxwidth string) []Rule {
	var rules []Rule
	var parentRule string
	for blockIndex, ruleStr := range ruleblock {
		var rulename string
		for index, char := range ruleStr {
			var s = string(char)
			if s == "{" {
				declarations := ParseDeclaration(ruleStr[index+1:])
				rulename = strings.TrimSpace(rulename)
				if blockIndex == 0 {
					parentRule = rulename
				}
				rulename = strings.Replace(rulename, "&", parentRule, -1)
				var rule = Rule{
					Name:         rulename,
					Declarations: declarations,
					AtRule:       atRule,
					Minwidth:     minwidth,
					Maxwidth:     maxwidth,
				}
				rules = append(rules, rule)
				break
			}
			rulename += s
		}
	}
	return rules
}

func ParseDeclaration(ruleStr string) []Declaration {
	declarationMap := make(map[string]string)
	var isKey bool = true
	var key string
	var value string

	reset := func() {
		//
		isKey = true
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		declarationMap[key] = value
		key = ""
		value = ""
	}

	for _, char := range ruleStr {
		var s = string(char)
		if s == ":" {
			isKey = false
			continue
		}
		if s == "}" {
			// A semicolon is not required in the last declaration of a rule

			// When key are value are empty, we end the loop because its the end
			// of the rule block and there is no final rule
			if key == "" || value == "" {
				break
			}
			reset()
			break
		}
		if s == ";" {
			reset()
			continue
		}
		if isKey {
			key += s
		} else {
			value += s
		}
	}
	var declarations []Declaration
	for k, v := range declarationMap {
		d := Declaration{Property: k, Value: v}
		declarations = append(declarations, d)
	}
	return declarations
}
