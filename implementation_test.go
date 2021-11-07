package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func TestImplementation(t *testing.T) { TestingT(t) }

func (s *TestSuite) TestPostfixToInfix(c *C) {
	examples := map[string]string{
		"123 22.828 - 54 *":          "(123 - 22.828) * 54",
		"20211202 11 +":              "20211202 + 11",
		"1 2 3 -":                    "too many argument",
		"3.14 11 ^ 44 *":             "3.14 ^ 11 * 44",
		"10 9 8 7 6 5 4 - / ^ - * +": "10 + 9 * (8 - 7 ^ (6 / (5 - 4)))",
		"Random text....":            "invalid input, there can be only operators and numbers",
		"993 711.2021 - 11 + - -":    "too many operators",
		"":                           "invalid input",
	}

	for postfix, expected := range examples {
		res, err := PostfixToInfix(postfix)
		if err != nil {
			c.Assert(err, ErrorMatches, expected)
		} else {
			c.Assert(res, Equals, expected)
		}
	}
}

func ExamplePostfixToInfix() {
	res, err := PostfixToInfix("1000 7 - 0 ^")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}

	// Output:
	// (1000 - 7) ^ 0
}
