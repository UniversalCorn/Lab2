package lab2

import (
	"fmt"
	"regexp"
	"strings"
)

type validator struct {
	Operator string
	Operand  string
}

func (v *validator) ValidSrting(input string) bool {
	validator := fmt.Sprintf(`^((%s|%s)\s){2,}(%s\s){0,}%s$`, v.Operand, v.Operator, v.Operator, v.Operator)
	isValid, _ := regexp.MatchString(validator, input)
	return isValid
}

func (v *validator) CheckArgsAmount(args []string) error {
	operators, operands := 0, 0
	macthOperator := fmt.Sprintf(`^%s$`, v.Operator)
	macthOperand := fmt.Sprintf(`^%s$`, v.Operand)
	for _, arg := range args {
		if isOperator, _ := regexp.MatchString(macthOperator, arg); isOperator {
			operators++
		} else if isOperand, _ := regexp.MatchString(macthOperand, arg); isOperand {
			operands++
		}
	}
	switch {
	case operators+operands != len(args):
		return fmt.Errorf("invalid expression argument(s)")
	case operands > operators+1:
		return fmt.Errorf("too many operands")
	case operators > operands-1:
		return fmt.Errorf("too many operators")
	default:
		return nil
	}
}

func (v *validator) IncludesOperator(str string) bool {
	includes, _ := regexp.MatchString(v.Operator, str)
	return includes
}

func (v *validator) IsOperator(str string) bool {
	macthOperator := fmt.Sprintf(`^%s$`, v.Operator)
	includes, _ := regexp.MatchString(macthOperator, str)
	return includes
}

func PostfixToInfix(postfixStr string) (infixStr string, err error) {
	v := validator{Operator: `[-\+\*\^\/]`, Operand: `(\d+|(\d+[,\.]\d+))`} //проверить на валидацию входящие символы, операторы и операнды с точкой или без
	if !v.ValidSrting(postfixStr) {
		err = fmt.Errorf("invalid input expression")
		return
	}
	var operatorsStack []string
	var infixHeap []string
	operatorsPriorities := map[string]uint8{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
	}
	postfixArgs := strings.Split(postfixStr, " ")
	if agrsErr := v.CheckArgsAmount(postfixArgs); agrsErr != nil {
		err = agrsErr
		return
	}

	for _, arg := range postfixArgs {
		if !v.IsOperator(arg) {
			infixHeap = append(infixHeap, arg)
			continue
		}
		operator := arg
		operatorsStack = append(operatorsStack, operator)
		slicedEnd := len(infixHeap) - 2
		sliced := infixHeap[slicedEnd:]
		infixHeap = infixHeap[:slicedEnd]
		operand1, operand2 := sliced[0], sliced[1]

		if len(operatorsStack) > 1 {
			prevOperatorIndex := len(operatorsStack) - 2
			prevOperator := operatorsStack[prevOperatorIndex]

			isPowerOperators := operatorsPriorities[operator] == 3 && operatorsPriorities[prevOperator] == 3
			higherPriority := operatorsPriorities[operator] > operatorsPriorities[prevOperator]
			if higherPriority || isPowerOperators {
				if v.IncludesOperator(operand2) {
					operand2 = "(" + operand2 + ")"
				} else {
					operand1 = "(" + operand1 + ")"
				}
			}
		}

		operand := fmt.Sprintf("%s %s %s", operand1, operator, operand2)
		infixHeap = append(infixHeap, operand)
	}

	infixStr = infixHeap[0]
	return
}
