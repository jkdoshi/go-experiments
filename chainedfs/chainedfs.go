// Package chainedfs allows programs to embed web assets in the binary (using go-bindata)
// but then override them from outside the binary (in live filesystem) so that
// bindata provides fallback, while live filesystem allows for changes/customization.
package chainedfs

import (
	"errors"
	"net/http"
)

// ChainedFileSystem chains an array of filesystems that can try to Open a given file
// until one of them succeeds. Fundamentally, ChainedFileSystem is just an array/slice
// of FileSystem instances.
//
// Usage example:
//   liveFS := http.Dir(*staticDir)
//   bindataFS := &assetfs.AssetFS{
//   	Asset:     Asset,
//  	AssetDir:  AssetDir,
//  	AssetInfo: AssetInfo,
//  	Prefix:    "bindata",
//   }
//   // it's just a slice of FileSystems!
//   chainedFS := chainedfs.ChainedFileSystem([]http.FileSystem{
//  	liveFS,
//  	bindataFS,
//   })
type ChainedFileSystem []http.FileSystem

// Open tries to open a given file in the chain of FileSystems until one succeeds
func (chain ChainedFileSystem) Open(name string) (http.File, error) {
	// iterate over the FileSystems in reverse order
	for i := len(chain) - 1; i >= 0; i-- {
		fs := chain[i]
		if fp, err := fs.Open(name); err == nil {
			return fp, err
		}
	}
	// For some reason, "err" is nil here. So we have to create a new one.
	return nil, errors.New("does not exist")
}
