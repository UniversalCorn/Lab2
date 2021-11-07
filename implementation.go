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
	if operators+operands != len(args) {
		return fmt.Errorf("invalid expression argument")
	} else if operands > operators+1 {
		return fmt.Errorf("too many operands")
	} else if operators > operands-1 {
		return fmt.Errorf("too many operators")
	} else {
		return nil
	}
}

func (v *validator) IncludesOperator(str string) bool {
	includes, _ := regexp.MatchString(v.Operator, str)
	return includes
}

func (v *validator) CheckOperator(str string) bool {
	macthOperator := fmt.Sprintf(`^%s$`, v.Operator)
	includes, _ := regexp.MatchString(macthOperator, str)
	return includes
}

func PostfixToInfix(postfixStr string) (infixStr string, err error) {
	v := validator{Operator: `[-\+\*\^\/]`, Operand: `(\d+|(\d+[,\.]\d+))`}
	if !v.ValidSrting(postfixStr) {
		err = fmt.Errorf("invalid expression")
		return
	}
	var operatorsStack []string
	var infixHeap []string
	prop := map[string]uint8{
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

	for _, operator := range postfixArgs {
		if !v.CheckOperator(operator) {
			infixHeap = append(infixHeap, operator)
			continue
		}

		operatorsStack = append(operatorsStack, operator)
		sliced := infixHeap[(len(infixHeap) - 2):]
		infixHeap = infixHeap[:(len(infixHeap) - 2)]

		if len(operatorsStack) > 1 {
			prevOperator := operatorsStack[(len(operatorsStack) - 2)]

			if prop[operator] > prop[prevOperator] || prop[operator] == 3 && prop[prevOperator] == 3 {
				if v.IncludesOperator(sliced[1]) {
					sliced[1] = "(" + sliced[1] + ")"
				} else {
					sliced[0] = "(" + sliced[0] + ")"
				}
			}
		}

		operand := fmt.Sprintf("%s %s %s", sliced[0], operator, sliced[1])
		infixHeap = append(infixHeap, operand)
	}

	infixStr = infixHeap[0]
	return
}
