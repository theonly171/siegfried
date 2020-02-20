package wikidata

import (
	"fmt"

	"github.com/richardlehane/siegfried/pkg/config"
	"github.com/richardlehane/siegfried/pkg/core"
)

const (
	extScore = 1 << iota
	mimeScore
	textScore
	incScore
)

type pids []Identification

// Recorder comment...
type Recorder struct {
	*Identifier
	ids        pids
	cscore     int
	satisfied  bool
	extActive  bool
	mimeActive bool
	textActive bool
}

// Active comment...
func (r *Recorder) Active(m core.MatcherType) {
	// implement
	fmt.Println("TODO Wikidata: Active isn't implemented....")
}

// Record comment...
func (r *Recorder) Record(m core.MatcherType, res core.Result) bool {
	switch m {
	default:
		return false
	case core.ByteMatcher:
		if hit, id := r.Hit(m, res.Index()); hit {
			if r.satisfied {
				return true
			}
			r.cscore += incScore
			basis := res.Basis()
			p, t := r.Place(core.ByteMatcher, res.Index())
			if t > 1 {
				basis = basis + fmt.Sprintf(" (signature %d/%d)", p, t)
			}
			r.ids = add(r.ids, r.Name(), id, r.infos[id], basis, r.cscore)
			return true
		}
		return false
	}
}

// Satisfied comment...
func (r *Recorder) Satisfied(mt core.MatcherType) (bool, core.Hint) {
	// TODO: Not sure what Satisfied does in the context of a Signature
	// matcher.
	fmt.Println("TODO Wikidata: Satisfied isn't implemented...")
	return false, core.Hint{}
}

// Report comment...
func (r *Recorder) Report() []core.Identification {
	// Happy path for zero results...
	if len(r.ids) == 0 {
		return []core.Identification{Identification{
			Namespace: r.Name(),
			ID:        "UNKNOWN",
			Warning:   "no match",
		}}
	}

	confidence := r.ids[0].confidence

	//--------------------------------------------------------
	// TODO: conf <= textScore ?? Why??
	//--------------------------------------------------------

	//--------------------------------------------------------
	// TODO: what is config.Single?
	// handle single result only
	//--------------------------------------------------------

	ret := make([]core.Identification, len(r.ids))
	for i, v := range r.ids {
		if i > 0 {
			switch r.Multi() {
			case config.Single:
				return ret[:i]
			case config.Conclusive:
				if v.confidence < confidence {
					return ret[:i]
				}
			default:
				if v.confidence < incScore {
					return ret[:i]
				}
			}
		}

		//--------------------------------------------------------
		// TODO: updateWarning uses the scoring to say something about the
		// match returned.
		//--------------------------------------------------------

		ret[i] = v
	}

	fmt.Printf("TODO Wikidata: returning identification %+v\n", ret)

	return ret
}
