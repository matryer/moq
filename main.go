package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/matryer/moq/package/moq"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			flag.Usage()
			os.Exit(1)
		}
	}()
	var (
		outFile = flag.String("out", "", "output file (default stdout)")
		pkgName = flag.String("pkg", "", "package name (default will infer)")
	)
	flag.Usage = func() {
		fmt.Println(`moq [flags] destination interface [interface2 [interface3 [...]]]`)
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		err = errors.New("not enough arguments")
		return
	}
	destination := args[0]
	args = args[1:]
	out := os.Stdout
	if len(*outFile) > 0 {
		out, err = os.Create(*outFile)
		if err != nil {
			return
		}
		defer out.Close()
	}
	m, err := moq.New(destination, *pkgName)
	if err != nil {
		return
	}
	err = m.Mock(out, args...)
}
