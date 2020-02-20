package wikidata

import (
	"fmt"

	"github.com/richardlehane/siegfried/internal/identifier"
	"github.com/richardlehane/siegfried/pkg/config"
	"github.com/richardlehane/siegfried/pkg/pronom"
	"github.com/richardlehane/siegfried/pkg/wikidata/internal/mappings"
)

// TODO: figure out a better name for Parseable, it is described as;
//
// 		Parseable is something we can parse to derive filename,
//		MIME, XML and byte signatures.
//
type wikidataFDDs struct {
	formats   []mappings.Wikidata
	parseable identifier.Parseable
	identifier.Blank
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

	for _, x := range reportMappings {
		fmt.Println(x)
	}

	var pronomID identifier.Parseable = identifier.Blank{}
	var err error
	if !config.NoPRONOM() {
		pronomID, err = pronom.NewPronom()
		if err != nil {
			return nil, err
		}
	}

	return wikidataFDDs{reportMappings, pronomID, identifier.Blank{}}, nil
}

// Basic mapping to load into newWikidata that we will then map to and return
// when PRONOM identifies the files...
var reportMappings = []mappings.Wikidata{
	{"Q12345", "PNG", "http://wikidata.org/q12345", "fmt/11"},
	{"Q23456", "FLAC", "http://wikidata.org/q23456", "fmt/279"},
	{"Q34567", "ICO", "http://wikidata.org/q34567", "x-fmt/418"},
	{"Q45678", "SIARD", "http://wikidata.org/q45678", "fmt/995"},
}
