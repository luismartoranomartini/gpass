package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("gpass", flag.ContinueOnError)
	fs.SetOutput(stderr)

	tamanho := fs.Int("n", 16, "tamanho de cada senha, em `caracteres`")
	quantidade := fs.Int("c", 1, "quantidade da senha")
	comSimbolos := fs.Bool("s", false, "incluir símbolos")

	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "uso gpass [opções]\n\nOpções:\n")
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		return 2
	}

	if fs.NArg() > 0 {
		fmt.Fprintf(stderr, "gpass: argumento inesperado %q\n", fs.Arg(0))
		return 2
	}

	if *tamanho < 1 || *quantidade < 1 {
		fmt.Fprintf(stderr, "gpass: -n e -c deve ser maiores que zero")
		return 2
	}

	alfa := alfabeto(*comSimbolos)

	for i := 0; i < *quantidade; i++ {
		senha, err := gerar(*tamanho, alfa)
		if err != nil {
			fmt.Fprintln(stderr, "gpass:", err)
			return 1
		}
		fmt.Fprintln(stdout, senha)
	}
	return 0
}
