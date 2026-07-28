package main

import (
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	modPath := os.Args[1]
	version := os.Args[2]
	srcDir := os.Args[3]
	outPath := os.Args[4]

	f, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, srcDir); err != nil {
		panic(err)
	}
}
