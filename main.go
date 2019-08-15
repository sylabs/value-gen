package main

import (
	"fmt"
	"io"
	"os"

	"github.com/sylabs/value-gen/values"
)

func usage() {
	fmt.Println("Usage:\nvalue-gen <values.yaml>")
}

func Run(out io.Writer) error {
	var val values.Values
	if err := val.Configure(); err != nil {
		return err
	}
	return val.Render(out)
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	outFileName := os.Args[1]
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Error opening file %v: %v\n", outFile, err)
		os.Exit(1)
	}
	err = Run(outFile)
	if err != nil {
		fmt.Printf("Error rendering values: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Successfully rendered yaml to %s\n", outFileName)
	}
}
