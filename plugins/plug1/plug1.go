package main

import (
	"fmt"
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

func (p *Plug1) Enable() error {
	fmt.Println(p.owner.GetThing(), "ENABLING")
	return nil
}

func (p *Plug1) Disable() error {
	fmt.Println(p.owner.GetThing(), "DISABLING")
	return nil
}
