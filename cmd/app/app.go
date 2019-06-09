package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/SilverCory/plugintest/api"
	"golang.org/x/xerrors"
)

type App struct {
	thing   string
	plugins map[string]api.Plugin
}

func (a *App) startup() {
	pathS, err := os.Getwd()
	if err != nil {
		panic(xerrors.Errorf("unable to get cwd %w", err))
	}

	err = filepath.Walk(filepath.Join(pathS, "plugins"), func(path string, _ os.FileInfo, _ error) error {
		if filepath.Ext(path) == ".so" {
			if err := a.loadPlugin(path); err != nil {
				fmt.Printf("Error loading plugin at path %q\n\t%s", path, err)
			}
		}
		return nil
	})
	if err != nil {
		panic(xerrors.Errorf("unable to walk file path %w", err))
	}
}

func (a *App) enable() {
	for k, v := range a.plugins {
		if err := v.Enable(); err != nil {
			fmt.Println(xerrors.Errorf("unable to enable plugin %q: %w", k, err))
			continue
		}
		fmt.Printf("Enabled plugin %q\n", k)
	}
}

func (a *App) disable() {
	for k, v := range a.plugins {
		if err := v.Enable(); err != nil {
			fmt.Println(xerrors.Errorf("unable to disable plugin %q: %w", k, err))
			continue
		}
		fmt.Printf("disabled plugin %q\n", k)
	}
}

func (a *App) loadPlugin(path string) error {
	pl, err := plugin.Open(path)
	if err != nil {
		return xerrors.Errorf("open file %w", err)
	}

	sym, err := pl.Lookup("New")
	if err != nil {
		return xerrors.Errorf("lookup new %w", err)
	}

	instance, ok := sym.(func() api.Plugin)
	if !ok {
		return xerrors.Errorf("new instance not type Initializer, instead %#v", sym)
	}

	plug := instance()
	if plug2 := a.plugins[plug.Name()]; plug2 != nil {
		return xerrors.Errorf("plugin overrides same name %q", plug.Name())
	}

	if err := plug.Init(a); err != nil {
		return xerrors.Errorf("call Init(api.Owner) on plugin %w", err)
	}

	a.plugins[plug.Name()] = plug
	return nil
}

func (a *App) GetThing() string {
	return a.thing
}
