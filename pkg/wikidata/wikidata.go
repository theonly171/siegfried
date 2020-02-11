package wikidata

import (
	"github.com/richardlehane/siegfried/internal/identifier"
)

type formatInfo struct {
	name     string
	longName string
	mimeType string
}

// turn generic FormatInfo into fdd formatInfo
func infos(m map[string]identifier.FormatInfo) map[string]formatInfo {
	i := make(map[string]formatInfo, len(m))
	for k, v := range m {
		i[k] = v.(formatInfo)
	}
	return i
}

func (f formatInfo) String() string {
	return f.name
}

func newWikidata() (identifier.Parseable, error) {
	/*
		We're going to load up a Wikidata "report" here. It's going to do a
		few things, but first, it's going to read the Wikidata IDs available.
		It's going to then read how those are mapped to PRONOM signatures.

		It's then going to load a PRONOM signature into memory (unless
	    noPRONOM is set) and then we'll return some form of mapping.

	    But first, let's load the mappings...
	*/
	var wikidata = identifier.Blank{}
	return wikidata, nil
}
