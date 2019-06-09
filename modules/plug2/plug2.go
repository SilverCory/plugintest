package main

import (
	"fmt"
	"time"

	"github.com/SilverCory/plugintest/api"
)

type Plug2 struct {
	owner api.Owner
}

func New() api.Plugin {
	return new(Plug2)
}

func (p *Plug2) Init(owner api.Owner) error {
	p.owner = owner
	return nil
}

func (p *Plug2) Name() string {
	return "Plug2"
}

func (p *Plug2) Enable() error {
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Printf("%s says: %q\n", p.Name(), p.owner.GetThing())
	}()
	return nil
}

func (p *Plug2) Disable() error {
	return nil
}
