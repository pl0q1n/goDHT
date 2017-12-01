package api

type Entry struct {
	host string
	hash uint64
}

type Pair struct {
	first  uint64
	second uint64
}

type FingerTable struct {
	entries   [64]Entry
	start     uint64
	selfEntry Entry
}

func diff(lhs uint64, rhs uint64) uint64 {
	if lhs > rhs {
		return lhs - rhs
	}
	return rhs - lhs
}

func (fingerTable *FingerTable) add(entry *Entry) {
	for i := 0; i < len(fingerTable.entries); i++ {
		target := fingerTable.selfEntry.hash + (2 << uint64(i))
		if entry.hash >= target {
			if diff(fingerTable.entries[i].hash, target) > diff(target, entry.hash) {
				fingerTable.entries[i] = *entry
			}
		}
	}
	if len(fingerTable.entries) == 0 {
		fingerTable.entries[0] = *entry
	}
}

func (fingerTable *FingerTable) route(hash uint64) string {
	if fingerTable.selfEntry.hash < fingerTable.start {
		if hash >= fingerTable.start || hash < fingerTable.selfEntry.hash {
			return fingerTable.selfEntry.host
		}
	}
	if hash >= fingerTable.start && hash < fingerTable.selfEntry.hash {
		return fingerTable.selfEntry.host
	}

	var min Pair = Pair{
		first:  diff(hash, fingerTable.entries[0].hash),
		second: 0,
	}

	for ind, elem := range fingerTable.entries {
		if diff(elem.hash, hash) < min.first && elem.hash > hash {
			min.first = diff(elem.hash, hash)
			min.second = uint64(ind)
		}
	}
	return fingerTable.entries[min.second].host
}
