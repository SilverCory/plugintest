package api

type Owner interface {
	GetThing() string
}

type Plugin interface {
	Init(owner Owner) error
	Enable() error
	Disable() error
}
