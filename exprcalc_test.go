package exprcalc

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSuite struct {
	expr   string
	answer string
	err    error
}

func TestGetPostfixExpr(t *testing.T) {
	t.Parallel()
	infixExpr := "2+2"
	postfix := GetPostfixExpr(infixExpr)
	assert.Equal(t, "2 2 +", postfix)
}

func TestCalculateEasyExpressions(t *testing.T) {
	t.Parallel()
	tSuite := []testSuite{
		{expr: "2 + 2", answer: "4.000000", err: nil},
		{expr: "2-2", answer: "0.000000", err: nil},
		{expr: "2*3", answer: "6.000000", err: nil},
		{expr: "1/2", answer: "0.500000", err: nil},
		{expr: "1 / 0", answer: "", err: errors.New("error occured due to division by zero")},
		{expr: " 49 * 63 / 58 * 36 ", answer: "1916.069148", err: nil},
		{expr: " 84 + 62 / 33 * 10 + 15 ", answer: "117.787880", err: nil},
		{expr: " 48 + 59 * 86 * 92 * 23 ", answer: "10736632.000000", err: nil},
		{expr: " 16 + 25 - 92 + 54 / 66 ", answer: "-50.181818", err: nil},
		{expr: " 64 + 19 - 77 - 93 ", answer: "-87.000000", err: nil},
		{expr: " 88 - 72 + 55 * 57", answer: "3151.000000", err: nil},
		{expr: " 99 * 55 / 30 + 50 ", answer: "231.499967", err: nil},
		{expr: " 11 - 88 + 84 - 48 ", answer: "-41.000000", err: nil},
		{expr: " 68 * 60 / 87 / 53 + 17 ", answer: "17.884816", err: nil},
		{expr: " 63 - 69 - 46 + 57 ", answer: "5.000000", err: nil},
		{expr: " 60 + 29 / 57 - 85 ", answer: "-24.491228", err: nil},
		{expr: " 34 * 18 * 55 - 50 ", answer: "33610.000000", err: nil},
		{expr: " 12 * 3 - 18 + 34 - 84 ", answer: "-32.000000", err: nil},
		{expr: " 70 / 42 - 52 - 64 / 35 ", answer: "-52.161904", err: nil},
		{expr: " 39 / 41 + 100 + 45 ", answer: "145.951220", err: nil},
	}

	for _, c := range tSuite {
		result, err := CalculateExpression(c.expr)
		assert.Equal(t, c.err, err)
		assert.Equal(t, c.answer, result)
	}
}
