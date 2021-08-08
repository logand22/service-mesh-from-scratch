package main

import (
	"fmt"
	"os"

	"github.com/logand22/service-mesh-from-scratch/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
