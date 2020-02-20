package config

import "path/filepath"

var wikidata = struct {
	definitions string
	name        string
	nopronom    bool
}{
	definitions: "wikidata_definitions-{version-info-tbd}",
	name:        "wikidata",
}

func Wikidata() string {
	return filepath.Join("", wikidata.definitions)
}

func SetWikidata() func() private {
	return func() private {
		return private{}
	}
}
