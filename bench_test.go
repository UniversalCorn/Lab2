package lab2

import (
	"fmt"
	"strings"
	"testing"
)

var cntRes string
var err error

func BenchmarkPostfixToInfix(b *testing.B) {
	const baseLen = 1000

	for i := 0; i < 20; i++ {
		l := baseLen * (i + 1)

		input := strings.Repeat("123 22.828 - 54 * ", l)

		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			cntRes, err = PostfixToInfix(input)
		})
	}
}
