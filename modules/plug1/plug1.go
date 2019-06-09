package main

import (
	"fmt"
	"time"

	"github.com/SilverCory/plugintest/api"
)

type Plug1 struct {
	owner api.Owner
}

func New() api.Plugin {
	return new(Plug1)
}

func (p *Plug1) Init(owner api.Owner) error {
	p.owner = owner
	return nil
}

func (p *Plug1) Name() string {
	return "Plug1"
}

func (p *Plug1) Enable() error {
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Printf("%s says: %q\n", p.Name(), p.owner.GetThing())
	}()
	return nil
}

func (p *Plug1) Disable() error {
	return nil
}
