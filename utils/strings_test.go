package utils_test

import (
	"testing"

	"github.com/eric38383/csstools/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetCharactersInBetweenString(t *testing.T) {
	var testResultOne = utils.GetStringBetweenChars(
		"(max-width: 50px) @media",
		"(",
		")",
	)
	var expectedResultOne = []string{"max-width: 50px"}
	assert.Equal(t, testResultOne, expectedResultOne)

	var testResultTwo = utils.GetStringBetweenChars(
		"@media (max-width: 50px) and (min-width: 100px) (min-width: 100px)",
		"(",
		")",
	)
	var expectedResultTwo = []string{"max-width: 50px", "min-width: 100px", "min-width: 100px"}
	assert.Equal(t, testResultTwo, expectedResultTwo)

	var testResultThree = utils.GetStringBetweenChars(
		"@media (min-width: 50px  ) (max-width: 100px ( max-width: 20px ) hello )",
		"(",
		")",
	)
	var expectedResultThree = []string{"min-width: 50px  ", "max-width: 100px ( max-width: 20px "}
	assert.Equal(t, testResultThree, expectedResultThree)

	var testResultFour = utils.GetStringBetweenChars(
		"hello#world#hello there is a #world#",
		"#",
		"#",
	)
	var expectedResultFour = []string{"world", "world"}
	assert.Equal(t, testResultFour, expectedResultFour)

	var testResultFive = utils.GetStringBetweenChars(
		"hello#world#hello there is a #world#",
		"there",
		"world",
	)
	var expectedResultFive = []string{" is a #"}
	assert.Equal(t, testResultFive, expectedResultFive)

	var testResultSix = utils.GetStringBetweenChars(
		"hello#world#hello there is a #world#",
		"",
		"world",
	)
	var expectedResultSix = []string(nil)
	assert.Equal(t, testResultSix, expectedResultSix)
}
