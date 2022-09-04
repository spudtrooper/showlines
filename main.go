package main

import (
	"flag"

	"github.com/fatih/color"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/goutil/io"
	"github.com/spudtrooper/goutil/or"
	"github.com/thomaso-mirodin/intmath/intgr"
)

var (
	file    = flags.String("file", "File to show line from")
	line    = flags.Int("line", "Line to show")
	before  = flags.Int("before", "Number of lines to show before")
	after   = flags.Int("after", "Number of lines to show after")
	context = flag.Int("context", 10, "Number of context lines to show before and after. This takes precedence over --before and --after")
)

func showLine(file string, line int, before, after int) error {
	lines, err := io.ReadLines(file)
	if err != nil {
		return err
	}
	start, end := intgr.Max(0, line-before-1), intgr.Min(len(lines), line+after)
	for i := start; i < len(lines) && i < end; i++ {
		ln := lines[i]
		if i == line {
			color.New(color.FgHiYellow, color.Italic, color.Bold).Printf("%-7d", i)
			color.New(color.FgYellow, color.Italic).Printf("%s\n", ln)
		} else {
			color.New(color.FgHiWhite, color.Bold).Printf("%-7d", i)
			color.New(color.FgWhite).Printf("%s\n", ln)
		}
	}
	return nil
}

func main() {
	flag.Parse()
	flags.RequireString(file, "file")
	flags.RequireInt(line, "line")
	beforeLines, afterLines := or.Int(*before, *context), or.Int(*after, *context)
	check.Err(showLine(*file, *line, beforeLines, afterLines))
}
