// Lookup Windows event messages by their event id (up to Windows 10).
//
// Usage:
//
//	eventid [PROVIDER:]ID...
//
// The arguments are:
//
//	[provider:]id
//	    The event id with optional provider prefix (required).
package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"

	"go.foxforensics.eu/eventid/events"
)

var Usage = `© 2026 Fox Forensics. Licensed under MIT License.
Usage: eventid [PROVIDER:]ID...

Report bugs at: foxforensics.eu/issues`

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		_, _ = fmt.Fprintln(os.Stderr, Usage)
		os.Exit(2)
	}

	db, err := events.Get()

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	all := slices.Sorted(maps.Keys(db))

	for _, id := range os.Args[1:] {
		var part string
		var found bool

		if strings.Contains(id, ":") {
			t := strings.SplitN(id, ":", 2)
			part, id = t[0], t[1]
		}

		n, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		providers := all

		// reduce providers
		if len(part) > 0 {
			providers = providers[:0]
			for _, p := range all {
				if strings.Contains(p, part) {
					providers = append(providers, p)
				}
			}
		}

		// search remaining providers
		for _, p := range providers {
			if m, ok := db[p][n]; ok {
				_, _ = fmt.Printf("%s: %s: \"%s\"\n", id, p, m)
				found = true
			}
		}

		// nothing found
		if !found {
			_, _ = fmt.Fprintf(os.Stderr, "%s not found\n", id)
		}
	}
}
