package wikidata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/richardlehane/siegfried/internal/identifier"
	"github.com/richardlehane/siegfried/pkg/config"
	"github.com/richardlehane/siegfried/pkg/pronom"
	"github.com/richardlehane/siegfried/pkg/wikidata/internal/mappings"

	"github.com/ross-spencer/spargo/pkg/spargo"
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

func getID(wikidataURI string) string {
	splitURI := strings.Split(wikidataURI, "/")
	return splitURI[len(splitURI)-1]
}

func openWikidata() {
	path := filepath.Join(config.WikidataHome(), config.Wikidata())
	fmt.Println("roy: opening Wikidata...", path)

	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Errorf("roy: cannot open Wikidata file: %s", err)
	}

	var sparqlResponse spargo.SPARQLResult
	err = json.Unmarshal(jsonFile, &sparqlResponse)
	if err != nil {
		fmt.Errorf("roy: cannot open Wikidata file: %s", err)
	}

	/*
		type item struct {
			Lang     string `json:"xml:lang"` // Populated if requested in query.
			Type     string // Can be "uri", "literal"
			Value    string
			DataType string
		}

		type binding struct {
			Bindings []map[string]item
		}

		// SPARQLResult packages a SPARQL response from an endpoint.
		type SPARQLResult struct {
			Head    map[string]interface{}
			Results binding
			Human   string
		}
	*/

	/*
		map[
			extension:{Lang: Type:literal Value:pp5 DataType:}
			format:{Lang: Type:uri Value:http://www.wikidata.org/entity/Q73019451 DataType:}
			formatLabel:{Lang:en Type:literal Value:Picture Publisher Bitmap, version 5.0 DataType:}
			puid:{Lang: Type:literal Value:x-fmt/85 DataType:}
		]
	*/

	for _, wdRecord := range sparqlResponse.Results.Bindings {
		wd := mappings.Wikidata{getID(wdRecord["format"].Value), wdRecord["formatLabel"].Value, wdRecord["format"].Value, wdRecord["puid"].Value, wdRecord["extension"].Value}
		// fmt.Printf("%+v\n\n", wd)
		reportMappings = append(reportMappings, wd)
	}
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
	openWikidata()
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
	/*
		{"Q12345", "PNG", "http://wikidata.org/q12345", "fmt/11", "png"},
		{"Q23456", "FLAC", "http://wikidata.org/q23456", "fmt/279", "flac"},
		{"Q34567", "ICO", "http://wikidata.org/q34567", "x-fmt/418", "ico"},
		{"Q45678", "SIARD", "http://wikidata.org/q45678", "fmt/995", "siard"},
	*/
}
