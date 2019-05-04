package main

import (
	"os"

	"github.com/yusukemisa/bingo/src/bingo"
)

func main() {
	cli := &bingo.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
