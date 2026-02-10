package main

import (
	"fmt"
	"os"

	"github.com/dhth/commits/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		followUp, ok := cmd.GetErrorFollowUp(err)
		if ok {
			fmt.Fprint(os.Stderr, followUp)
		}
		os.Exit(1)
	}
}
