package common

import "fmt"

// Pe common error processing
func Pe(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(msg, v...))
	}
}
