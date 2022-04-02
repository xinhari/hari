// Package plugins includes the plugins we want to load
package plugins

import (
	"xinhari.com/xinhari/config/cmd"

	// import specific plugins
	ckStore "xinhari.com/xinhari/store/cockroach"
	fileStore "xinhari.com/xinhari/store/file"
	memStore "xinhari.com/xinhari/store/memory"

	// we only use CF internally for certs
	cfStore "xinhari.com/hari/internal/plugins/store/cloudflare"
)

func init() {
	// TODO: make it so we only have to import them
	cmd.DefaultStores["cloudflare"] = cfStore.NewStore
	cmd.DefaultStores["cockroach"] = ckStore.NewStore
	cmd.DefaultStores["file"] = fileStore.NewStore
	cmd.DefaultStores["memory"] = memStore.NewStore
}
