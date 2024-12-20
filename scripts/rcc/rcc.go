// original: https://github.com/mappu/miqt/blob/master/cmd/miqt-rcc/main.go

// miqt-rcc compiles a Qt resource XML file (*.qrc) to a binary resource file and
// creates a Go stub to load it.
package main

import (
	"flag"
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// Parse arguments

	input := flag.String("Input", "", "Path to .qrc input file")
	outputRcc := flag.String("OutputRcc", "", "(Optional) Path to .rcc output file. If omitted, inferred from the input file path")
	outputGo := flag.String("OutputGo", "", "(Optional) Path to .go output file. If omitted, interred from the input file path")
	packageName := flag.String("Package", "main", "Package to use in generated Go files")
	variableName := flag.String("VariableName", "_resourceRcc", "Temporary global variable name for loading embedded data")
	flag.Parse()

	// Check if input file exists

	if _, err := os.Stat(*input); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Input file '%s' not found\n", *input)
		os.Exit(1)
	}

	// Fill in default output names, if not specified

	if *outputRcc == "" {
		*outputRcc = strings.TrimSuffix(*input, `.qrc`) + `.rcc`
	}
	if *outputGo == "" {
		*outputGo = strings.TrimSuffix(*input, `.qrc`) + `.go`
	}

	// Compile qrc to binary resource file

	rccCmd := exec.Command(`rcc`, `--binary`, `-o`, *outputRcc, *input)
	rccCmd.Stderr = os.Stderr
	rccCmd.Stdout = os.Stdout
	err := rccCmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "rcc error: %s\n", err.Error())
		os.Exit(1)
	}

	//  Create Go file that loads the resource

	goSrcData := `
package ` + *packageName + `

//go:generate go run scripts/rcc/rcc.go "` + strings.Join(os.Args[1:], `" "`) + `"

import (
	_ "embed"
	
	qt "github.com/mappu/miqt/qt6"
)

//go:embed ` + *outputRcc + `
var ` + *variableName + ` []byte

func init() {
	qt.QResource_RegisterResourceWithRccData(& ` + *variableName + `[0])
}

`

	outputData, err := format.Source([]byte(goSrcData))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Format error: %s\n", err.Error())
		os.Exit(1)
	}

	err = os.WriteFile(*outputGo, outputData, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("RCC OK")
}
