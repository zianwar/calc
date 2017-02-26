package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/fatih/color"
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
	if l := len(os.Args[1:]); l < 2 {
		usage()
	}

	op, rest := os.Args[1], os.Args[2:]

	if _, ok := operations[strings.Replace(op, " ", "", -1)]; !ok {
		fmt.Print("\"" + op + "\" operation not recognized.")
	}

	expression := strings.Replace(strings.Join(rest, " "), " ", "", -1)
	URL := "https://newton.now.sh/" + op + "/" + url.QueryEscape(expression)

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var result map[string]string
	json.Unmarshal([]byte(body), &result)
	fmt.Println(result["result"])
}

func usage() {
	bold := color.New(color.Bold)
	green := color.New(color.FgGreen)

	fmt.Printf("\nUsage: %s [operation] [expression]\n\n", bold.Sprint(green.Sprint("calc")))
	fmt.Print("  • Where operation is defined as:\n\n")
	for cmd, desc := range operations {
		fmt.Printf("\t %s:  %s\n", cmd, desc)
	}
	fmt.Print("\n  • And expression could be number or any mathematical equation.\n\n")
	os.Exit(1)
}
