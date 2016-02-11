package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kisom/goutils/lib"
	"github.com/kisom/jsonfmt/jfmt"
)

const version = "1.0.0"

func usage() {
	progname := lib.ProgName()
	fmt.Printf("%s version %s\n", progname, version)
	fmt.Printf(`Usage: %s [-h] files...
	%s is used to lint and prettify (or compact) JSON files. The
	files will be printed to standard output, unless the -w flag is
	given. In this case, the files will be updated in place.

	Flags:
	-c	Compact files.
	-h	Print this help message.
	-w	write result to source file instead of stdout.
`, progname, progname)

}

func init() {
	flag.Usage = usage
}

func act(action func([]byte) ([]byte, error), file string, w bool) error {
	in, err := ioutil.ReadFile(file)
	if err != nil {
		lib.Warn(err, "ReadFile")
		return err
	}

	in, err = action(in)
	if err != nil {
		lib.Warn(err, "Compact")
		return err
	}

	if w {
		err = ioutil.WriteFile(file, in, 0644)
		if err != nil {
			lib.Warn(err, "WriteFile")
		}
	} else {
		fmt.Printf("%s", string(in))
	}

	return err
}

func main() {
	var shouldCompact, showUsage, writeFile bool
	flag.BoolVar(&shouldCompact, "c", false, "Compact files instead of prettifying.")
	flag.BoolVar(&showUsage, "h", false, "Print a usage message.")
	flag.BoolVar(&writeFile, "w", false, "Write result to source file instead of stdout.")
	flag.Parse()

	if showUsage {
		usage()
		os.Exit(0)
	}

	action := jfmt.Pretty
	if shouldCompact {
		action = jfmt.Compact
	}

	var errCount int
	for _, fileName := range flag.Args() {
		err := act(action, fileName, writeFile)
		if err != nil {
			errCount++
		}
	}

	if errCount > 0 {
		lib.Errx(lib.ExitFailure, "Not all files succeeded.")
	}
}
