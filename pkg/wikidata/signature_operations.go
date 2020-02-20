package wikidata

import (
	"github.com/richardlehane/siegfried/internal/bytematcher/frames"
	"github.com/richardlehane/siegfried/internal/identifier"
)

func (wdd wikidataFDDs) Signatures() ([]frames.Signature, []string, error) {
	// TODO Siegfried uses an errors slice to return cumulative errors here...
	// e.g. var errs []error, we can do the same...
	var errs []error
	var puidsIDs map[string][]string
	if len(wdd.parseable.IDs()) > 0 {
		puidsIDs = make(map[string][]string)
	}
	sigs, ids := make([]frames.Signature, 0, len(wdd.formats)), make([]string, 0, len(wdd.formats))
	for _, v := range wdd.formats {
		// Append Wikidata identifiers to signatures...
		ids = append(ids, v.ID)
		if puidsIDs != nil {
			for _, puid := range v.PUIDs() {
				puidsIDs[puid] = append(puidsIDs[puid], v.ID)
			}
		}
	}
	// TODO Can we combine with puidsIDs above?
	if puidsIDs != nil {
		puids := make([]string, 0, len(puidsIDs))
		for p := range puidsIDs {
			puids = append(puids, p)
		}
		np := identifier.Filter(puids, wdd.parseable)
		// ns == BOF?
		// ps == EOF?
		ns, ps, err := np.Signatures()
		if err != nil {
			errs = append(errs, err)
		}
		for i, v := range ps {
			for _, id := range puidsIDs[v] {
				sigs = append(sigs, ns[i])
				ids = append(ids, id)
			}
		}

	}
	return sigs, ids, nil
}
