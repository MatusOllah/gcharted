package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	inFile      = flag.String("InFile", "", "Input .ui file")
	outFile     = flag.String("OutFile", "-", "Output .go file, or - for stdout")
	packageName = flag.String("Package", "main", "Custom package name")
)

func main() {
	flag.Parse()

	if *inFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	gosrc, err := generate()
	if err != nil {
		panic(err)
	}

	gosrc = update(gosrc)

	if *outFile == "-" {
		fmt.Println(string(gosrc))

	} else {
		err = os.WriteFile(*outFile, gosrc, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func generate() ([]byte, error) {
	cmd := exec.Command("miqt-uic", "-InFile", *inFile, "-OutFile", "-", "-Package", *packageName)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return stdout.Bytes(), nil
}

func update(code []byte) []byte {
	replacements := map[string]string{
		`"github.com/mappu/miqt/qt"`: `"github.com/MatusOllah/gcharted/internal/i18n"
	qt "github.com/mappu/miqt/qt6"`,
		`SetObjectName\("([^"]+)"\)`: `SetObjectName(*qt.NewQAnyStringView3("$1"))`,
		`SetGeometry\(qt.NewQRect\((?<numbers>[0-9]+), (?<numbers>[0-9]+), (?<numbers>[0-9]+), (?<numbers>[0-9]+)\)\)`: `SetGeometry($1, $2, $3, $4)`,
		`qt.*_Tr\("(.*)"\)`:                      `i18n.L("$1")`,
		`SetOrientation\(qt.Orientation__(.*)\)`: `SetOrientation(qt.$1)`,
		`(?m)^//go:generate\s+.*`:                `//go:generate go run ../../scripts/uic/uic.go "` + strings.Join(os.Args[1:], `" "`) + `"`,
	}

	updatedCode := code

	// Perform replacements
	for old, new := range replacements {
		re := regexp.MustCompile(old)
		updatedCode = re.ReplaceAll(updatedCode, []byte(new))
	}

	return updatedCode
}
