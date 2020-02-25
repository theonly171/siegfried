package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/richardlehane/siegfried/pkg/config"
	"github.com/ross-spencer/spargo/pkg/spargo"
)

func harvestWikidata() {
	log.Println("roy harvesting Wikidata definitions")
	err := os.MkdirAll(config.WikidataHome(), os.ModePerm)
	if err != nil {
		fmt.Errorf("roy: error harvesting Wikidata definitions: %s", err)
	}
	sparqlMe := spargo.SPARQLClient{}
	sparqlMe.ClientInit(config.WikidataEndpoint(), config.WikidataSPARQL())
	sparqlMe.SetUserAgent(config.UserAgent())
	res := sparqlMe.SPARQLGo()
	path := filepath.Join(config.WikidataHome(), config.Wikidata())
	err = ioutil.WriteFile(path, []byte(res.Human), config.WikidataFileMode())
	if err != nil {
		fmt.Printf("roy: error harvesting Wikidata: %s", err)
	}
	log.Printf("roy harvesting Wikidata definitions '%s' complete\n", path)
}
