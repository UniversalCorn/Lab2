package lab2

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	if ch.Input == nil {
		return fmt.Errorf("input is undefined")
	}
	if ch.Output == nil {
		return fmt.Errorf("output is undefined")
	}
	buf, inputError := ioutil.ReadAll(ch.Input)
	if inputError != nil {
		return inputError
	}
	bufString = string(buf)
	inputString := strings.Trim(bufString, "\n")
	computed, computeErr := PostfixToInfix(inputString)
	if computeErr != nil {
		return computeErr
	}
	res := []byte(computed + "\n")
	_, outputError := ch.Output.Write(res)
	if outputError != nil {
		return outputError
	}
	return nil
}
