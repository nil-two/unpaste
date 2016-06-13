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
	cmdName    = "unpaste"
	cmdVersion = "0.1.0"

	flagset    = pflag.NewFlagSet(cmdName, pflag.ContinueOnError)
	delimiters = flagset.StringP("delimiters", "d", "\t", "")
	isSerial   = flagset.BoolP("serial", "s", false, "")
	isHelp     = flagset.BoolP("help", "h", false, "")
	isVersion  = flagset.BoolP("version", "v", false, "")
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [OPTION]... [FILE]...
Separate corresponding or subsequent lines of files.

With no FILE, or when FILE is -, output to standard output.

Options:
  -d, --delimiters=LIST   reuse characters from LIST instead of TABs
  -s, --serial            unpaste one file at a time instead of in parallel
      --help              display this help and exit
      --version           display version information and exit
`[1:], cmdName)
}

func printVersion() {
	fmt.Fprintf(os.Stderr, "%s\n", cmdVersion)
}

func printErr(err interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", cmdName, err)
}

func guideToHelp() {
	fmt.Fprintf(os.Stderr, "Try '%s --help' for more information.\n", cmdName)
}

func do(ws []io.Writer, s *Separator, isSerial bool) error {
	b := bufio.NewScanner(os.Stdin)
	if isSerial {
		for i := 0; i < len(ws) && b.Scan(); i++ {
			a := s.Separate(b.Text())
			for _, line := range a {
				fmt.Fprintln(ws[i], line)
			}
		}
	} else {
		for b.Scan() {
			a := s.Separate(b.Text())
			for i := 0; i < len(ws); i++ {
				if i < len(a) {
					fmt.Fprintln(ws[i], a[i])
				} else {
					fmt.Fprintln(ws[i], "")
				}
			}
		}
	}
	return b.Err()
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

	var ws []io.Writer
	if flagset.NArg() < 1 {
		ws = append(ws, os.Stdout)
	} else {
		for _, name := range flagset.Args() {
			if name == "-" {
				ws = append(ws, os.Stdout)
				continue
			}

			f, err := os.Create(name)
			if err != nil {
				printErr(err)
				guideToHelp()
				return 2
			}
			defer f.Close()

			ws = append(ws, f)
		}
	}

	s := NewSeparator(*delimiters)
	if err := do(ws, s, *isSerial); err != nil {
		printErr(err)
		return 1
	}
	return 0
}

func main() {
	e := _main()
	os.Exit(e)
}
