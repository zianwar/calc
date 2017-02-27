package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var operations = map[string]string{
	"simplify":  ` Simplifies expression.     eg:  2^2 + 2(2)   =>                  8`,
	"factor":    `   Factorise expression.      eg:    x^2 + 2x   =>           x(x + 2)`,
	"derive":    `   Derives expression.        eg:    x^2 + 2x   =>             2x + 2`,
	"integrate": `integrates expression.     eg:    x^2 + 2x   =>  1/3 x^3 + x^2 + C`,
	"zeroes":    `   Finds expression zeros.    eg:    x^2 + 2x   =>            [-2, 0]`,
	"cos":       `      Cosine of number.          eg:          pi   =>                 -1`,
	"sin":       `      Sine of number.            eg:           0   =>                  0`,
	"tan":       `      Tangent of number.         eg:           0   =>                  0`,
	"arccos":    `   Inverse Cosine of number.  eg:           1   =>                  0`,
	"arcsin":    `   Inverse Sine of number.    eg:           0   =>                  0`,
	"arctan":    `   Inverse Tangent of number. eg:           0   =>                  0`,
	"abs":       `      Absolute Value of number.  eg:          -1   =>                  1`,
}

func main() {
	m, err := validateArgs(os.Args)
	if err != nil {
		fmt.Print(err)
		usage()
	}
	url := fmt.Sprintf("https://newton.now.sh/%s/%s", m["operation"], m["expression"])

	r, err := get(url)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(r["result"])
}

func validateArgs(args []string) (map[string]string, error) {
	var valid map[string]string

	if l := len(args[1:]); l < 2 {
		usage()
		os.Exit(1)
	}

	op, exp := args[1], args[2:]
	operation := strings.Replace(op, " ", "", -1)

	if _, ok := operations[op]; !ok {
		return valid, errors.New("\"" + op + "\" operation not recognized.")
	}

	expression := url.QueryEscape(strings.Replace(strings.Join(exp, " "), " ", "", -1))
	valid = map[string]string{
		"operation":  operation,
		"expression": expression,
	}

	return valid, nil
}

func get(url string) (result map[string]string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	json.Unmarshal([]byte(body), &result)

	if _, ok := result["error"]; ok {
		return nil, fmt.Errorf("Error: %s", result["error"])
	}

	return result, nil
}

func usage() {
	fmt.Printf("\nUsage: calc [operation] [expression]\n\n")
	fmt.Print("  • Where operation is defined as:\n\n")
	for cmd, desc := range operations {
		fmt.Printf("\t %s:  %s\n", cmd, desc)
	}
	fmt.Print("\n  • And expression could be number or any mathematical equation.\n\n")
}
