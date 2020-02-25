package config

import (
	"os"
	"path/filepath"
)

var wikidata = struct {
	definitions string
	name        string
	nopronom    bool
	filemode    os.FileMode

	endpoint string
	sparql   string
}{
	// TODO: identify something helpful for versioning Wikidata definitions file.
	definitions: "wikidata-definitions-0.0.1-beta",

	name:     "wikidata",
	filemode: 0644,

	endpoint: "https://query.wikidata.org/sparql",
	sparql: `
		SELECT ?format ?puid ?ldd ?formatLabel ?extension ?mimetype ?sig WHERE
		{
		  ?format wdt:P2748 ?puid.
		  OPTIONAL { ?format wdt:P3266 ?ldd }
		  OPTIONAL { ?format wdt:P1195 ?extension }
		  OPTIONAL { ?format wdt:P1163 ?mimetype }
		  OPTIONAL { ?format p:P4152 ?sig }
		  SERVICE wikibase:label { bd:serviceParam wikibase:language "[AUTO_LANGUAGE],en". }
		}`,
}

// Wikidata ...
func Wikidata() string {
	return wikidata.definitions
}

// SetWikidata ...
func SetWikidata() func() private {
	return func() private {
		return private{}
	}
}

// WikidataHome ...
func WikidataHome() string {
	return filepath.Join(siegfried.home, wikidata.name)
}

// WikidataFileMode ...
func WikidataFileMode() os.FileMode {
	return wikidata.filemode
}

// WikidataEndpoint ...
func WikidataEndpoint() string {
	return wikidata.endpoint
}

// WikidataSPARQL ...
func WikidataSPARQL() string {
	return wikidata.sparql
}
