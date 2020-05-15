package exprcalc

import (
	"fmt"
	"strconv"
	"strings"
)

//CalculateExpression calculates given expression, returns solution or error
func CalculateExpression(expression string) (string, error) {
	expression = strings.Replace(expression, " ", "", -1)
	expression = strings.Replace(expression, "+", ",+,", -1)
	expression = strings.Replace(expression, "-", ",-,", -1)
	expression = strings.Replace(expression, "*", ",*,", -1)
	expression = strings.Replace(expression, "/", ",/,", -1)
	expression = strings.Replace(expression, "(", ",(,", -1)
	expression = strings.Replace(expression, ")", ",),", -1)

	infixExpr := strings.Split(expression, ",")

	stack := NewStack()
	postfixExpr := []string{}
	for _, st := range infixExpr {
		switch string(st) {
		case "(":
			stack.Push(NewStackElement(string(st)))
			break
		case ")":
			for {
				if stack.Length() == 0 {
					break
				}
				s := stack.Pop()
				if s.Value == "(" {
					break
				}
				postfixExpr = append(postfixExpr, s.Value)
			}
			break
		case "+", "-", "*", "/":
			for {
				if stack.Length() == 0 || stack.Head.Value == "(" {
					break
				}
				if string(st) == "/" && stack.Head.Value == "*" {
					break
				}
				if string(st) == "*" && (stack.Head.Value == "+" || stack.Head.Value == "-") {
					break
				}
				s := stack.Pop()

				postfixExpr = append(postfixExpr, s.Value)
			}
			stack.Push(NewStackElement(string(st)))
			break
		default:
			postfixExpr = append(postfixExpr, string(st))
		}
	}

	for {
		if stack.Length() == 0 {
			break
		}
		s := stack.Pop()
		postfixExpr = append(postfixExpr, s.Value)
	}
	// fmt.Println(postfixExpr)

	stack = NewStack()
	for _, el := range postfixExpr {
		if el == "+" || el == "-" || el == "*" || el == "/" {
			el1, _ := strconv.ParseFloat(stack.Pop().Value, 64)
			el2, _ := strconv.ParseFloat(stack.Pop().Value, 64)
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
			stack.Push(NewStackElement(fmt.Sprintf("%f", result)))
		} else {
			stack.Push(NewStackElement(el))
		}
	}
	// fmt.Println(stack.Pop().Value)
	return stack.Pop().Value, nil
}
