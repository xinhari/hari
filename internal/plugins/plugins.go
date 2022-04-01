// Package plugins includes the plugins we want to load
package plugins

import (
	"github.com/ebelanja/go-micro/config/cmd"

	// import specific plugins
	ckStore "github.com/ebelanja/go-micro/store/cockroach"
	fileStore "github.com/ebelanja/go-micro/store/file"
	memStore "github.com/ebelanja/go-micro/store/memory"

	// we only use CF internally for certs
	cfStore "github.com/ebelanja/hari/internal/plugins/store/cloudflare"
)

func init() {
	// TODO: make it so we only have to import them
	cmd.DefaultStores["cloudflare"] = cfStore.NewStore
	cmd.DefaultStores["cockroach"] = ckStore.NewStore
	cmd.DefaultStores["file"] = fileStore.NewStore
	cmd.DefaultStores["memory"] = memStore.NewStore
}
