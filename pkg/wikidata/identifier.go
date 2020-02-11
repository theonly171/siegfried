package wikidata

import (
	"fmt"
	"strings"

	"github.com/richardlehane/siegfried/internal/identifier"
	"github.com/richardlehane/siegfried/pkg/config"
	"github.com/richardlehane/siegfried/pkg/core"

)

// Comment...
type Identifier struct {
	infos map[string]formatInfo
	*identifier.Base
}

// Comment...
func New(opts ...config.Option) (core.Identifier, error) {
	fmt.Println("Congratulations: doing something with the Wikidata identifier package!")

	wikidata, _ := newWikidata()

	return &Identifier{
		infos: infos(wikidata.Infos()),
		Base:  identifier.New(wikidata, "", ""),
	}, nil
}

// Comment...
func (i *Identifier) Fields() []string {
	return []string{
		"namespace", "id", "format", "full", "mime", "basis", "warning"}
}

// Comment...
func (r *Recorder) Record(m core.MatcherType, res core.Result) bool {
	return false
}

// Comment...
func (i *Identifier) Recorder() core.Recorder {
	return &Recorder{
		Identifier: i,
		ids:        make(pids, 0, 1),
	}
}

// Comment...
func (r *Recorder) Report() []core.Identification {
	return []core.Identification{Identification{
		Namespace: r.Name(),
		ID:        "UNKNOWN",
		Warning:   "no match",
	}}
}

// Comment...
func (r *Recorder) Satisfied(mt core.MatcherType) (bool, core.Hint) {
	return false, core.Hint{}
}

// Comment...
type Recorder struct {
	*Identifier
	ids        pids
	cscore     int
	satisfied  bool
	extActive  bool
	mimeActive bool
	textActive bool
}

// Comment...
func (r *Recorder) Active(m core.MatcherType) {
	// implement
}

// Comment...
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

// Comment...
func (id Identification) String() string {
	return id.ID
}

// Comment...
func (id Identification) Archive() config.Archive {
	return id.archive
}

// Comment...
func (id Identification) Known() bool {
	return id.ID != "UNKNOWN"
}

// Comment...
func (id Identification) Warn() string {
	return id.Warning
}

// Comment...
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

type pids []Identification
