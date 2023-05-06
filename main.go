package main

import (
	"fmt"
	"os"

	"chipcom/lib"
)

func main() {
	if err := lib.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
