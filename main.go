package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/ogier/pflag"
)

var (
	name    = "unpaste"
	version = "0.1.0"

	flagset    = pflag.NewFlagSet(name, pflag.ContinueOnError)
	delimiters = flagset.StringP("delimiters", "d", "\t", "")
	isSerial   = flagset.BoolP("serial", "s", false, "")
	isHelp     = flagset.BoolP("help", "h", false, "")
	isVersion  = flagset.BoolP("version", "v", false, "")
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [OPTION]... [FILE]...
Separate corresponding or subsequent lines of files.

Options:
  -d, --delimiters=LIST   reuse characters from LIST instead of TABs
  -s, --serial            unpaste one file at a time instead of in parallel
      --help              display this help and exit
      --version           display version information and exit
`[1:], name)
}

func printVersion() {
	fmt.Fprintf(os.Stderr, "%s\n", version)
}

func printErr(err interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
}

func guideToHelp() {
	fmt.Fprintf(os.Stderr, "Try '%s --help' for more information.\n", name)
}

func _main() int {
	flagset.SetOutput(ioutil.Discard)
	if err := flagset.Parse(os.Args[1:]); err != nil {
		printErr(err)
		guideToHelp()
		return 2
	}
	if *isHelp {
		printUsage()
		return 0
	}
	if *isVersion {
		printVersion()
		return 0
	}

	var writers []io.Writer
	if flagset.NArg() < 1 {
		writers = append(writers, os.Stdout)
	} else {
		for _, name := range flagset.Args() {
			if name == "-" {
				writers = append(writers, os.Stdout)
				continue
			}

			f, err := os.Create(name)
			if err != nil {
				printErr(err)
				guideToHelp()
				return 2
			}
			defer f.Close()

			writers = append(writers, f)
		}
	}

	s := NewSeparator(*delimiters)
	b := bufio.NewScanner(os.Stdin)
	if *isSerial {
		for i := 0; i < len(writers) && b.Scan(); i++ {
			a := s.Separate(b.Text())
			for _, line := range a {
				fmt.Fprintln(writers[i], line)
			}
		}
	} else {
		for b.Scan() {
			a := s.Separate(b.Text())
			for i := 0; i < len(writers); i++ {
				if i < len(a) {
					fmt.Fprintln(writers[i], a[i])
				} else {
					fmt.Fprintln(writers[i], "")
				}
			}
		}
	}
	if err := b.Err(); err != nil {
		printErr(err)
		return 1
	}
	return 0
}

func main() {
	e := _main()
	os.Exit(e)
}
