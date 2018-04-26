
// tgpl-01-cmd-line

package main

import (
	"os"
	"fmt"
	"strings"
)

// ======== Command line arguments ========

// One way of passing input to a program is through the command line
// They are given to the program through the "os" package.

// os.Args is a slice of strings, where the element at index 0 will be
// the name of the program, and the subsequent elements the arguments.

// Notes on slices : the indexes are 0-based. One can also access
// contiguous elements with the following notation s[m:n] meaning
// that we're going to egt the elements from m to n-1.
// All indexing in go is said to be "half-open", meaning that the
// last index is excluded.


// Variable declaration / initialization - these are all equivalent
// s := "" (only inside a function) - preferred initial value is important
// var s string (initialized with zero value - preferred when initial value isn't important
// var s = "" (rarely used)
// var s string = ""


func main() {
	fmt.Println(formatArguments_basic(os.Args))
	fmt.Println(formatArguments_range(os.Args))
	fmt.Println(formatArguments_join(os.Args))

	// Usefull for printing slices
	fmt.Println(os.Args[1:])
}

// Simple version, it uses the full form of the for loop

//for initialization; condition; post {
//	...
//}

func formatArguments_basic(args []string) (string) {

	// Declare 2 variables, both strings
	var s, sep string

	// loop over the arguments
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	return s;
}


// Another form of the for loop using the range keyword
func formatArguments_range(args []string) (string) {

	// Declare 2 variables, both strings
	var s, sep string

	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}

	return s;
}

// The += on s produces a new string corresponding to the concatenation
// of s and sep. This can be costly, so a better version would be
// using strings.Join
func formatArguments_join(args []string) (string) {
	return strings.Join(args[1:], " ")
}