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
	postfix, err := GetPostfixExpr(infixExpr)
	assert.Equal(t, "2 2 +", postfix)
	assert.Equal(t, nil, err)

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

func TestCalculateMediumExpressions(t *testing.T) {
	t.Parallel()
	tSuite := []testSuite{
		{expr: "1 + 2) * 3", answer: "", err: errors.New("ExpressionError: Brackets must be paired")},
		{expr: "((1 + 2) * 3", answer: "", err: errors.New("ExpressionError: Brackets must be paired")},
		{expr: "((1 + 2 * 3", answer: "", err: errors.New("ExpressionError: Brackets must be paired")},

		{expr: " 20 - 57 * 12 - (  58 + 84 * 32 / 27  ) ", answer: "-821.555540", err: nil},
		{expr: " 77 + 79 / 25 * (  64 * 63 - 89 * 14  ) * 49 ", answer: "431461.240000", err: nil},

		{expr: " 100 - 60 / 38 + (  19 / 88 * 97 / 82 / 94  ) * 92 ", answer: "98.671017", err: nil},
		{expr: " (  97 / 48 + 86 + 56 * 94  ) / 43 + 57 ", answer: "181.465601", err: nil},
		{expr: " (  68 - 85 / 75 * 64  ) / 15 + 73 ", answer: "72.697779", err: nil},

		{expr: " 91 + 18 / (  42 + 62 + 84 * 95  ) + 30 ", answer: "121.002227", err: nil},
		{expr: " 49 * 31 * (  20 - 83 / 63 / 46 * 29  ) / 68 ", answer: "428.212176", err: nil},
		{expr: " 35 - 45 / 37 + 84 + (  41 + 86 / 18 / 41 * 73  ) ", answer: "167.290547", err: nil},

		{expr: " 44 * 13 / (  26 + 24 * 70 + 89 * 7  ) + 81 ", answer: "81.245608", err: nil},
		{expr: " 53 - 88 + 7 + (  34 / 54 + 15 / 23 / 6  ) * 73 ", answer: "25.897798", err: nil},
		{expr: " 57 - 71 + (  14 + 3 - 24 * 100 / 23  ) / 53 ", answer: "-15.648072", err: nil},

		{expr: " (  41 * 76 * 79 - 61  ) / 60 + 83 ", answer: "4184.716667", err: nil},
		{expr: " (  73 + 85 + 64 / 17  ) * 17 + 31 / 60 ", answer: "2750.516669", err: nil},
		{expr: " 74 * 96 + 62 / (  25 / 33 + 96 + 87 + 78  ) ", answer: "7104.236860", err: nil},

		{expr: " 33 - 96 + (  95 - 76 * 98 / 11  ) * 15 ", answer: "-8794.363740", err: nil},
		{expr: " 72 / 75 + 4 * (  14 * 2 / 57 * 21  ) / 15 ", answer: "3.710900", err: nil},
		{expr: " 72 * 95 + 53 + (  2 + 76 - 52 / 1 - 47  ) ", answer: "6872.000000", err: nil},

		{expr: " 85 * 97 / (  89 / 11 - 18 * 96  ) - 61 ", answer: "-65.793830", err: nil},
		{expr: " 29 + 24 / 91 - (  14 * 71 * 18 / 20 * 100  ) + 63 ", answer: "-89367.736264", err: nil},
		{expr: " 52 * 62 * (  61 + 12 - 14 * 79  ) + 39 ", answer: "-3330353.000000", err: nil},
	}

	for _, c := range tSuite {
		result, err := CalculateExpression(c.expr)
		assert.Equal(t, c.err, err)
		assert.Equal(t, c.answer, result)
	}
}
