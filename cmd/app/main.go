package main

import (
	"fmt"
	"plugin"

	"github.com/SilverCory/plugintest/api"
)

func main() {
	app := &App{thing: "Hello World"}
	pl, err := plugin.Open("plugin.dll")
	if err != nil {
		panic(err)
	}

	sym, err := pl.Lookup("New")
	if err != nil {
		panic(err)
	}

	plug, ok := sym.(api.Plugin)
	if !ok {
		fmt.Println("type not plugin, ", plug)
		panic("type not plug")
	}

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
