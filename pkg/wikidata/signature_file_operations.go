package wikidata

import (
	"github.com/richardlehane/siegfried/internal/identifier"
	"github.com/richardlehane/siegfried/internal/persist"
	"github.com/richardlehane/siegfried/pkg/core"
)

// Save writes a Wikidata identifier to the Siegfried signature file.
func (i *Identifier) Save(ls *persist.LoadSaver) {
	// Save the Wikidata magic enum from core.
	ls.SaveByte(core.Wikidata)

	// Save the no. formatInfo entries to read.
	ls.SaveSmallInt(len(i.infos))

	// Save the information in the formatInfo records.
	for key, value := range i.infos {
		ls.SaveString(key)
		ls.SaveString(value.name)
		ls.SaveString(value.uri)
		ls.SaveString(value.puid)
	}
	i.Base.Save(ls)
}

// Load reads a Wikidata identifier from the Siegfried signature file.
func Load(ls *persist.LoadSaver) core.Identifier {
	i := &Identifier{}
	le := ls.LoadSmallInt()
	i.infos = make(map[string]formatInfo)
	for j := 0; j < le; j++ {
		i.infos[ls.LoadString()] = formatInfo{
			ls.LoadString(),
			ls.LoadString(),
			ls.LoadString(),
		}
	}
	i.Base = identifier.Load(ls)
	return i
}

// TODO: I think this can hold pretty much anything, so what else is important
// and that characterizes a Wikidata signature?
type formatInfo struct {
	name string
	uri  string
	puid string
}

// turn generic FormatInfo into fdd formatInfo
func infos(m map[string]identifier.FormatInfo) map[string]formatInfo {
	i := make(map[string]formatInfo, len(m))
	for k, v := range m {
		i[k] = v.(formatInfo)
	}
	return i
}

// String belongs to formatInfo and outputs the formatInfo struct in a human
// readable format.
func (f formatInfo) String() string {
	return f.name
}

// Infos arranges summary information about formats within an Identifier into a
// structure suitable for output in a Siegfried signature file.
func (wdd wikidataFDDs) Infos() map[string]identifier.FormatInfo {
	formatInfoMap := make(map[string]identifier.FormatInfo, len(wdd.formats))
	for _, value := range wdd.formats {
		fi := formatInfo{
			name: value.Name,
			uri:  value.URI,
			puid: value.PRONOM,
		}
		formatInfoMap[value.ID] = fi
	}
	return formatInfoMap
}
