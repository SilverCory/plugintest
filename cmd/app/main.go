package main

import (
	"fmt"
	"plugin"

	"github.com/SilverCory/plugintest/api"
)

func main() {
	app := &App{thing: "Hello World"}
	pl, err := plugin.Open("plug1.so")
	if err != nil {
		panic(err)
	}

	sym, err := pl.Lookup("New")
	if err != nil {
		panic(err)
	}

	new, ok := sym.(func() api.Plugin)
	if !ok {
		fmt.Printf("type not new plugin, %#v\n", sym)
		panic("type not plug")
	}

	plug := new()

	if err := plug.Init(app); err != nil {
		panic(err)
	}

	if err := plug.Enable(); err != nil {
		panic(err)
	}

	if err := plug.Disable(); err != nil {
		panic(err)
	}
}

type App struct {
	thing string
}

func (a *App) GetThing() string {
	return a.thing
}
