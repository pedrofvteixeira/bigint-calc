package main

import (
	"flag"
	"fmt"
	"math/big"
	"strings"
)

type Params struct {
	Operand string
	Arg1    string
	Arg2    string
}

func main() {

	var p Params

	flag.StringVar(&p.Operand, "operand", "", "arithmethic operand")
	flag.StringVar(&p.Arg1, "a1", "", "arg 1")
	flag.StringVar(&p.Arg2, "a2", "", "arg 2")

	flag.Parse()

	if IsAnyEmpty(p.Arg1, p.Operand, p.Arg2) {
		panic(fmt.Sprintf("Missing one or more required params (received a1:'%v', operand:'%v', a2:'%v')", p.Arg1, p.Operand, p.Arg2))
	}

	a1 := MakeBigInt(p.Arg1)
	a2 := MakeBigInt(p.Arg2)

	var result *big.Int

	switch p.Operand {
	case "+":
		result = new(big.Int).Add(a1, a2)
	case "-":
		result = new(big.Int).Sub(a1, a2)
	case "x":
		result = new(big.Int).Mul(a1, a2)
	case "/":
		result = new(big.Int).Div(a1, a2)
	default:
		panic(fmt.Sprintf("Unknown operand '%v'. Must be one of +,-,x,/", p.Operand))
	}

	fmt.Printf("%v\n", result)
}

// converts a string representation of a number into a Big.Int
// @see math/big
func MakeBigInt(val string) *big.Int {
	i, success := new(big.Int).SetString(val, 10)

	if !success {
		panic("Failed to create BigInt from string")
	}

	return i
}

// predicate that checks if the given trimmed string is empty
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// predicate that checks if the given trimmed string is empty
func IsAnyEmpty(s ...string) bool {
	for _, val := range s {
		if IsEmpty(val) {
			return true
		}
	}
	return false
}
