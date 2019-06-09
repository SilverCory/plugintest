package main

import (
	"fmt"
	"github.com/SilverCory/plugintest/api"
)

type Plug1 struct {
	owner api.Owner
}

func New() *Plug1 {
	return new(Plug1)
}

func (p *Plug1) Init(owner api.Owner) error {
	p.owner = owner
}

func (p *Plug1) Enable() error {
	fmt.Print(p.owner.GetThing(), "ENABLING")
	return nil
}

func (p *Plug1) Disable() error {
	fmt.Print(p.owner.GetThing(), "DISABLING")
}