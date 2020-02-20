package wikidata

import (
	"fmt"
	"strings"
	"time"

	"github.com/richardlehane/siegfried/internal/identifier"
	"github.com/richardlehane/siegfried/pkg/config"
	"github.com/richardlehane/siegfried/pkg/core"
)

// I can't remember what the definition of Golang init is...
func init() {
	core.RegisterIdentifier(core.Wikidata, Load)
}

// Identifier comment...
type Identifier struct {
	infos map[string]formatInfo
	*identifier.Base
}

// New is the entry point for an Identifier when it is compiled by the Roy tool
// to a brand new signature file.
//
// New will read a Wikidata report, and parse its information into structures
// suitable for compilation by Roy.
//
// New will also update its identification information with provenance-like
// info. It will enable signature extensions to be added by the utility, and
// enables configuration to be applied as well.
//
// Examples of extensions include: ...
//
// Examples of configuration include: ...
//
func New(opts ...config.Option) (core.Identifier, error) {

	fmt.Println("Congratulations: doing something with the Wikidata identifier package!")

	wikidata, _ := newWikidata()

	// TODO: What does date versioning look like from the Wikidata sources?
	updatedDate := time.Now().Format("2006-01-02")

	// TODO: Add extensions.
	// TODO: Apply configuration.

	return &Identifier{
		infos: infos(wikidata.Infos()),
		Base:  identifier.New(wikidata, "", updatedDate),
	}, nil
}

// Fields comment...
func (i *Identifier) Fields() []string {
	return []string{
		"namespace", "id", "format", "full", "mime", "basis", "warning"}
}

// Recorder comment that belongs to an identifier...
func (i *Identifier) Recorder() core.Recorder {
	return &Recorder{
		Identifier: i,
		ids:        make(pids, 0, 1),
	}
}

// Identification comment...
type Identification struct {
	Namespace  string
	ID         string
	Name       string
	LongName   string
	MIME       string
	Basis      []string
	Warning    string
	archive    config.Archive
	confidence int
}

// String creates a human readable representation of an identifier for output
// by fmt-like functions.
func (id Identification) String() string {
	return id.ID
}

// Archive comment... [Mandatory]
func (id Identification) Archive() config.Archive {
	return id.archive
}

// Known comment... [Mandatory]
func (id Identification) Known() bool {
	return id.ID != "UNKNOWN"
}

// Warn comment... [Mandatory]
func (id Identification) Warn() string {
	return id.Warning
}

// Values comment... [Mandatory]
func (id Identification) Values() []string {
	var basis string
	if len(id.Basis) > 0 {
		basis = strings.Join(id.Basis, "; ")
	}
	return []string{
		id.Namespace,
		id.ID,
		id.Name,
		id.LongName,
		id.MIME,
		basis,
		id.Warning,
	}
}

//-------------------------------------------------------------
// TODO how much of this is needed? ---------------------------
//-------------------------------------------------------------

func add(p pids, id string, f string, info formatInfo, basis string, c int) pids {
	for i, v := range p {
		if v.ID == f {
			p[i].confidence += c
			p[i].Basis = append(p[i].Basis, basis)
			return p
		}
	}
	return append(p, Identification{id, f, info.name, info.uri, info.puid, []string{basis}, "", config.IsArchive(f), c})
}
