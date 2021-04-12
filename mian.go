package main

import (
	"fmt"

	interface_abstract "github.com/promacanthus/design-patterns/interface-abstract"
)

func main() {
	fileLogger, err := interface_abstract.NewFileLogger("file", true, interface_abstract.Level{Value: 1}, "")
	if err != nil {
		fmt.Println(err)
	}
	fileLogger.Log(interface_abstract.Level{Value: 2}, "test log message")
}
