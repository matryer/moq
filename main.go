package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/sudo-suhas/moqit/pkg/moq"
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
		outFile  = flag.String("out", "", "output file (default stdout)")
		pkgName  = flag.String("pkg", "", "package name (default will infer)")
		mockName = flag.String("mock", "", "mock name (default will infer)")
	)
	flag.Usage = func() {
		fmt.Println(`moq [flags] destination interface`)
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		err = errors.New("not enough arguments")
		return
	}
	destination, iface := args[0], args[1]
	args = args[1:]
	var buf bytes.Buffer
	var out io.Writer
	out = os.Stdout
	if len(*outFile) > 0 {
		out = &buf
	}
	m, err := moq.New(destination, *pkgName)
	if err != nil {
		return
	}
	if err = m.Mock(out, iface, *mockName); err != nil {
		return
	}
	// create the file
	if len(*outFile) > 0 {
		err = ioutil.WriteFile(*outFile, buf.Bytes(), 0644)
	}
}
