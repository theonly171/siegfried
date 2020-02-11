package config

import "path/filepath"

var wikidata = struct {
	definitions string
	name string
}{
	definitions: "wikidata_definitions",
	name: "wikidata",
}

func Wikidata() string {
	return filepath.Join("", wikidata.definitions)
}

func SetWikidata() func() private {
	return func() private {
		return private{}
	}
}
