package exprcalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPostfixExpr(t *testing.T) {
	t.Parallel()
	infixExpr := "2+2"

	postfix := GetPostfixExpr(infixExpr)

	assert.Equal(t, "2 2 +", postfix)
}

func TestCalculateExpression(t *testing.T) {
	t.Parallel()
	infixExpr := "2*3-(2+2)"

	result, _ := CalculateExpression(infixExpr)

	assert.Equal(t, "2.000000", result)
}
