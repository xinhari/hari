package runtime

import (
	"fmt"

	"xinhari.com/hari/plugin"
)

var (
	defaultManager = plugin.NewManager()
)

// Plugins lists the runtime plugins
func Plugins() []plugin.Plugin {
	return defaultManager.Plugins()
}

// Register registers an runtime plugin
func Register(pl plugin.Plugin) error {
	if plugin.IsRegistered(pl) {
		return fmt.Errorf("%s registered globally", pl.String())
	}
	return defaultManager.Register(pl)
}
