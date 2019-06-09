package api

// Plugin is an interface for plugins that want to be loaded enabled and disabled.
type Plugin interface {
	Init(owner Owner) error
	Name() string
	Enable() error
	Disable() error
}
