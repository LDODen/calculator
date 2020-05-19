package exprcalc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LDODen/exprcalc/stack"
)

// GetPostfixExpr returns expr in postfix form
func GetPostfixExpr(infixExpr string) string {

	infixExpr = strings.Replace(infixExpr, " ", "", -1)
	infixExpr = strings.Replace(infixExpr, "+", ",+,", -1)
	infixExpr = strings.Replace(infixExpr, "-", ",-,", -1)
	infixExpr = strings.Replace(infixExpr, "*", ",*,", -1)
	infixExpr = strings.Replace(infixExpr, "/", ",/,", -1)
	infixExpr = strings.Replace(infixExpr, "(", ",(,", -1)
	infixExpr = strings.Replace(infixExpr, ")", ",),", -1)

	tempInfix := strings.Split(infixExpr, ",")

	st := stack.NewStack()
	postfixExpr := []string{}
	for _, str := range tempInfix {
		switch string(str) {
		case "(":
			st.Push(stack.NewStackElement(string(str)))
			break
		case ")":
			for {
				if st.Length() == 0 {
					break
				}
				s := st.Pop()
				if s.Value == "(" {
					break
				}
				postfixExpr = append(postfixExpr, s.Value)
			}
			break
		case "+", "-", "*", "/":
			for {
				if st.Length() == 0 || st.Head.Value == "(" {
					break
				}
				if string(str) == "/" && st.Head.Value == "*" {
					break
				}
				if string(str) == "*" && (st.Head.Value == "+" || st.Head.Value == "-") {
					break
				}
				s := st.Pop()

				postfixExpr = append(postfixExpr, s.Value)
			}
			st.Push(stack.NewStackElement(string(str)))
			break
		default:
			postfixExpr = append(postfixExpr, string(str))
		}
	}
	for st.Length() > 0 {
		s := st.Pop()
		postfixExpr = append(postfixExpr, s.Value)
	}
	return strings.Join(postfixExpr, " ")
}

//CalculateExpression calculates given expression, returns solution or error
func CalculateExpression(expression string) (string, error) {

	postfixExpr := strings.Split(GetPostfixExpr(expression), " ")

	st := stack.NewStack()
	for _, el := range postfixExpr {
		if el == "+" || el == "-" || el == "*" || el == "/" {
			el1, _ := strconv.ParseFloat(st.Pop().Value, 64)
			el2, _ := strconv.ParseFloat(st.Pop().Value, 64)
			var result float64
			switch el {
			case "+":
				result = el1 + el2
			case "-":
				result = el2 - el1
			case "*":
				result = el1 * el2
			case "/":
				result = el2 / el1
			}
			st.Push(stack.NewStackElement(fmt.Sprintf("%f", result)))
		} else {
			st.Push(stack.NewStackElement(el))
		}
	}
	return st.Pop().Value, nil
}
