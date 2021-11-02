package main

import (
	"flag"
	"fmt"
	"strings"
	"io"
	"os"
	lab2 "https://github.com/IP94-rocketBunny-architecture/Lab2"
)

func getFlagsValues() (inputExpression, fileIn, fileOut *string) {
	defer flag.Parse()
	
	inputExpression = flag.String("e", "", "Expression to compute")
	filein = flag.String("i", "", "input file")
	fileout = flag.String("o", "", "output file")
	return
}


func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	res, _ := lab2.PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}
