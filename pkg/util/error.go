package util

import "fmt"

// Pe Panic when condition is not satisfied.
func Pe(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(msg, v...))
	}
}
