// Package events provides an event message database.
package events

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"io"
	"sync"

	"github.com/ulikunitz/xz"
)

//go:embed database.xz
var database []byte

var cache Providers

// Providers mapping of event ids and messages.
type Providers map[string]map[int64]string

// Get returns the decompressed embedded providers.
func Get() (Providers, error) {
	var err error

	sync.OnceFunc(func() {
		err = decompress()
	})()

	return cache, err
}

func decompress() error {
	r, err := xz.NewReader(bytes.NewReader(database))

	if err != nil {
		return err
	}

	b, err := io.ReadAll(r)

	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &cache)

	return err
}
