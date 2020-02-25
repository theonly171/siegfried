package mappings

import (
	"fmt"
)

// Wikidata describes a complete record of a signature in the Wikidata
// database including its relationship to other identifier types such as
// PRONOM.

// Wikidata ..
// TODO: relationships, e.g. priorities of one signature over another.
// 		 PRONOM mapping, perhaps this needs to be a slice?
type Wikidata struct {
	ID        string // Wikidata short name, e.g. Q12345 can be appended to a URI to be dereferenced.
	Name      string // Name of the format as described in Wikidata.
	URI       string // URI is the absolute URL in Wikidata terms that can be dereferenced.
	PRONOM    string // 1:1 mapping to PRONOM wherever possible.
	Extension string // Extension returned by Wikidata
}

// String ...
// TODO rename f to something helpful (wdd)
func (f Wikidata) String() string {
	return fmt.Sprintf("ID: %s\nName: %s\nLong Name: %s\nPUIDs: %s\n",
		f.ID,
		f.Name,
		f.URI,
		f.PRONOM,
	)
}

// PUIDs ...
// TODO rename f to something helpful (wdd)
// TODO other characteristics of Wikidata which determine more than one PUID?
func (f Wikidata) PUIDs() []string {
	var puids []string
	puids = append(puids, f.PRONOM)
	return puids
}
