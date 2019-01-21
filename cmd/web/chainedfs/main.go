// A program to show how to combine go-bindata embedded assets with a live filesystem
// where bindata provides hard-coded assets but live filesystem can override them
// simply by creating a file in the same path.
// For example, if you have "<bindata>/index.html" in bindata, and then create
// a "<live>/index.html", the second one will mask the first.
package main

import (
	"flag"
	"log"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/jkdoshi/go-experiments/chainedfs"
)

func main() {
	addr := flag.String("addr", ":8080", "address to listen on ([host]:port)")
	staticDir := flag.String("static", "static", "directory to serve static files from")
	flag.Parse() // must call this before flags are actually usable
	liveFS := http.Dir(*staticDir)
	bindataFS := &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "bindata",
	}
	chainedFS := chainedfs.ChainedFileSystem([]http.FileSystem{
		liveFS,
		bindataFS,
	})
	http.Handle("/", http.FileServer(chainedFS))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
