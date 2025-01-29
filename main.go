package main

import (
	"fmt"
	"log"
	"plugin"
	"text/template"
)

const symbolName = "Funcs"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	pl, err := plugin.Open("plugin/plugin_1.so")
	if err != nil {
		return err
	}
	symbol, err := pl.Lookup(symbolName)
	if err != nil {
		return err
	}
	funcs, isFuncs := symbol.(*template.FuncMap)
	if !isFuncs {
		return fmt.Errorf("symbol has to be of type `template.FuncMap`")
	}
	fmt.Println(funcs)
	return nil
}
