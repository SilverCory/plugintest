package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/SilverCory/plugintest/api"
)

func main() {
	app := &App{
		thing:   "Hello World",
		plugins: map[string]api.Plugin{},
	}

	app.startup()
	fmt.Println("APP STARTED.")

	app.enable()
	fmt.Println("PLUGINS ENABLED; DONE")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("\nSHUTTING DOWN")
	app.disable()
}
