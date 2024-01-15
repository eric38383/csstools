package parser

import (
	"fmt"
	"slices"
	"strings"

	"github.com/eric38383/csstools/utils"
	"github.com/gorilla/css/scanner"
)

type Parser struct {
	scanner *scanner.Scanner
}

func New(str string) *Parser {
	return &Parser{scanner: scanner.New(str)}
}

// Iterates though all tokens by called the Next() method
// Separates rules when reaching and AT Keyword or regular rule
// returns a list of css style rules
func (parser *Parser) Stylesheet() []Rule {
	var rules []Rule
	var maxwidth string = ""
	var minwidth string = ""
	for {
		token := parser.scanner.Next()
		if parser.SkipToken(*token) {
			continue
		}
		if parser.EndToken(*token) {
			break
		}
		if token.Type == scanner.TokenAtKeyword {
			if slices.Contains(regAtRules, token.Value) {
				var value = token.Value + " " + parser.SkipTo(";")
				var rule = Rule{Name: value, AtRule: true}
				rules = append(rules, rule)
			} else if slices.Contains(unnestedRules, token.Value) {
				var ruleBlocks = parser.RuleBlock(token.Value)
				parsedRules := ParseRule(ruleBlocks, true, minwidth, maxwidth)
				rules = append(rules, parsedRules...)
			} else if token.Value == "@media" {
				//there are other nested rules but we'll handle it later
				var mediaStr = parser.SkipTo("{")
				var widths = utils.GetBetweenTwoChars(mediaStr, "(", ")")
				for _, s := range widths {
					var splitRule = strings.Split(s, ":")
					var key = strings.TrimSpace(splitRule[0])
					var val = strings.TrimSpace(splitRule[1])
					if key == "min-width" {
						minwidth = val
					} else if key == "max-width" {
						maxwidth = val
					}
				}

			}
			continue
		}
		if token.Type == scanner.TokenChar || token.Type == scanner.TokenIdent {
			//this will skip the last "}" when we have a nested atrule
			// and reset any min or max width values
			if token.Value == "}" {
				minwidth = ""
				maxwidth = ""
				continue
			}
			var ruleBlocks = parser.RuleBlock(token.Value)
			parsedRules := ParseRule(ruleBlocks, false, minwidth, maxwidth)
			rules = append(rules, parsedRules...)
		}
	}
	return rules
}

// Skips to the next "something" token and stops at that value
// Returns all tokens as string from start to end
func (parser *Parser) SkipTo(stop string) string {
	var value string
	var index = 0
	for {
		token := parser.scanner.Next()

		if token.Value == stop {
			break
		}
		value += token.Value
		if index == 25 {
			break
		}
		index += 1
	}
	return value
}

// Skips the token if its a specific token type.
// These tokens are not important to us
func (parser *Parser) SkipToken(token scanner.Token) bool {
	return token.Type == scanner.TokenS || token.Type == scanner.TokenComment || token.Type == scanner.TokenCDC || token.Type == scanner.TokenCDO
}

// Returns true if there is a token error or token is the end of the file
func (parser *Parser) EndToken(token scanner.Token) bool {
	if token.Type == scanner.TokenError {
		fmt.Printf("There is an error in your file: %s", scanner.TokenError)
		return true
	}
	if token.Type == scanner.TokenEOF {
		fmt.Println("End of File!")
		return true
	}
	return false
}

// Separates parent and nested rules into Rule structs. Updates the ampersand with the parent rule.
// Returns all rules when iterating though the tokens reaches the final closing tag "}"
func (parser *Parser) RuleBlock(currentToken string) []string {
	var level int = 0
	var rules = []string{currentToken}
	for {
		token := parser.scanner.Next()
		if token.Value == "&" {
			level += 1
			rules = append(rules, "")
		}
		rules[level] += token.Value
		if level == 0 && token.Value == "}" {
			return rules
		}
		if token.Value == "}" {
			level -= 1
		}
	}
}
