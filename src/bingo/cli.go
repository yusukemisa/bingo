package bingo

import (
	"flag"
	"io"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var regular bool
	flags := flag.NewFlagSet("bingo", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.BoolVar(&regular, "regular", false, "Create regular type bingo card.")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	bingo := newBingo(regular, c.OutStream)
	bingo.run()

	return ExitCodeOK
}
