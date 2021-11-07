package main

import (
	"io"
	"os"
	"flag"
	"fmt"
	"strings"
	Lab2 "https://github.com/IP94-rocketBunny-architecture/Lab2"
)

func getFlagsValues() (inputExpression, fileIn, fileout *string) {
	defer flag.Parse()
	
	inputExpression = flag.String("e", "", "Expression to compute")
	filein = flag.String("i", "", "input file")
	fileout = flag.String("o", "", "output file")
	return
}

func main() {
	var in io.Reader
	var out io.Writer
	var inputExpression, fileIn, fileout *string = getFlagsValues()
	
	if *fileIn != "" && *inputExpression != "" {
		err := fmt.Errorf("more than one expr is not needed")
		panic(err)
	}
	
	if *inputExpression != "" {
		in = strings.NewReader(*inputExpression)
	} else if *fileIn != "" {
		in, _ = os.Open(*fileIn)
	}
	if *fileout != "" {
		out, _ = os.Create(*outputFilename)
	} else {
		out = os.Stdout
	}
	handler := Lab2.ComputeHandler{Input: in, Output: out}
	err := handler.Compute()
	if err != nil {
		panic(err)
	}
}
