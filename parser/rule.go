package parser

import "strings"

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
	minwidth     string
	maxwidth     string
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
			rulename += s
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
					minwidth:     minwidth,
					maxwidth:     maxwidth,
				}
				rules = append(rules, rule)
				break
			}
		}
	}
	return rules
}

func ParseDeclaration(ruleStr string) []Declaration {
	declarationMap := make(map[string]string)
	var isKey bool
	var key string
	var value string
	for _, char := range ruleStr {
		var s = string(char)
		if s == ":" {
			isKey = false
		}
		if s == ";" {
			isKey = true
			key = strings.TrimSpace(key)
			value = strings.TrimSpace(value)
			declarationMap[key] = value
			key = ""
			value = ""
			continue
		}
		if isKey {
			key += s
		} else {
			value += s
		}
		if s == "}" {
			break
		}
	}
	declarationMap[key] = value
	var declarations []Declaration
	for k, v := range declarationMap {
		d := Declaration{Property: k, Value: v}
		declarations = append(declarations, d)
	}
	return declarations
}
