package main

import (
	"encoding/pem"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()

	for i := range flag.Args() {
		err := splitPEM(flag.Arg(i))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(i + 1)
		}
	}
}

func splitPEM(n string) error {

	b, err := os.ReadFile(n)
	if err != nil {
		return err
	}

	block, rest := pem.Decode(b)
	i := 0
	base := filepath.Base(n)
	base = strings.TrimSuffix(base, filepath.Ext(base))

	for block != nil {
		err := writePEM(block, fmt.Sprintf("%s.%d.pem", base, i))
		if err != nil {
			return err
		}
		block, rest = pem.Decode(rest)
		i++
	}

	return nil
}

func writePEM(block *pem.Block, n string) error {
	f, err := os.Create(n)
	if err != nil {
		return err
	}
	defer f.Close()
	err = pem.Encode(f, block)
	return err
}
