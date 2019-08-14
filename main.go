package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/sylabs/value-gen/values"
)

func usage() {

}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	outFileName := os.Args[1]
	outFile, err := os.Open(outFileName)
	if err != nil {
		fmt.Printf("Error opening file %v: %v", outFile, err)
		os.Exit(1)
	}
	var val values.Values
	val.Configure()
	t, err := template.New("values").Parse(values.Template)
	if err != nil {
		panic(err) // bad hardcoded string, panic
	}
	err = t.Execute(outFile, val)
	if err != nil {
		fmt.Printf("Error rendering values: %v", err)
		os.Exit(1)
	}
}
