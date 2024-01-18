package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var skipCheck bool
	flag.BoolVar(&skipCheck, "skip-check", false, "skip the check phase")
	flag.Parse()

	env, err := cel.NewEnv()
	if err != nil {
		return fmt.Errorf("create env: %v", err)
	}
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("read: %v", err)
	}
	source := common.NewTextSource(string(in))
	ast, iss := env.ParseSource(source)
	if iss.Err() != nil {
		return fmt.Errorf("compile: %v", iss.Err())
	}
	if !skipCheck {
		ast, iss = env.Check(ast)
		if iss.Err() != nil {
			return fmt.Errorf("check: %v", iss.Err())
		}
	}
	p, err := env.Program(ast)
	if err != nil {
		return err
	}
	result, _, err := p.Eval(map[string]any{})
	if err != nil {
		return err
	}
	fmt.Println(result.Value())
	return nil
}
