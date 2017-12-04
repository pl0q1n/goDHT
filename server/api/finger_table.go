package api

import (
	pbNode "github.com/pl0q1n/goDHT/node_proto"
)

type Entry struct {
	Host string
	Hash uint64
}

type Pair struct {
	first  uint64
	second uint64
}

type FingerTable struct {
	Entries       [64]Entry
	PreviousEntry Entry
	SelfEntry     Entry
}

func diff(lhs uint64, rhs uint64) uint64 {
	if lhs > rhs {
		return lhs - rhs
	}
	return rhs - lhs
}

func (fingerTable *FingerTable) GetProtoFingerTable() *pbNode.FingerTable {
	protoFingerTable := &pbNode.FingerTable{}
	protoPrevious := &pbNode.FingerTable_Entry{}
	protoPrevious.Hash = fingerTable.PreviousEntry.Hash
	protoPrevious.Host = fingerTable.PreviousEntry.Host

	var entrySlice []*pbNode.FingerTable_Entry
	for _, elem := range fingerTable.Entries {
		protoEntry := &pbNode.FingerTable_Entry{
			Hash: elem.Hash,
			Host: elem.Host,
		}
		entrySlice = append(entrySlice, protoEntry)
	}
	protoFingerTable.Entry = entrySlice
	protoFingerTable.Previous = protoPrevious
	return protoFingerTable
}

func (fingerTable *FingerTable) add(entry *Entry) {
	for i := 0; i < len(fingerTable.Entries); i++ {
		target := fingerTable.SelfEntry.Hash + (2 << uint64(i))
		if entry.Hash >= target {
			if diff(fingerTable.Entries[i].Hash, target) > diff(target, entry.Hash) {
				fingerTable.Entries[i] = *entry
			}
		}
	}
	if fingerTable.Entries[0].Host == "" {
		for ind := range fingerTable.Entries {
			fingerTable.Entries[ind] = *entry
		}
	}
}

func (fingerTable *FingerTable) route(Hash uint64) string {
	if fingerTable.SelfEntry.Hash < fingerTable.PreviousEntry.Hash {
		if Hash >= fingerTable.PreviousEntry.Hash || Hash < fingerTable.SelfEntry.Hash {
			return fingerTable.SelfEntry.Host
		}
	}
	if Hash >= fingerTable.PreviousEntry.Hash && Hash < fingerTable.SelfEntry.Hash {
		return fingerTable.SelfEntry.Host
	}

	var min Pair = Pair{
		first:  diff(Hash, fingerTable.Entries[0].Hash),
		second: 0,
	}

	for ind, elem := range fingerTable.Entries {
		if diff(elem.Hash, Hash) < min.first && elem.Hash > Hash {
			min.first = diff(elem.Hash, Hash)
			min.second = uint64(ind)
		}
	}
	return fingerTable.Entries[min.second].Host
}
