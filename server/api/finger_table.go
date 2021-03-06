package api

import (
	pbNode "github.com/pl0q1n/goDHT/node_proto"
)

type Entry struct {
	Host string
	Hash uint64
}

type Pair struct {
	Diff  uint64
	Index uint64
}

type Update struct {
	updates [64]string
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

func (fingerTable *FingerTable) chosePrevious(possiblePrevious *Entry) {
	if fingerTable.PreviousEntry.Host == "" {
		fingerTable.PreviousEntry.Host = possiblePrevious.Host
		fingerTable.PreviousEntry.Hash = possiblePrevious.Hash
	}

	if fingerTable.SelfEntry.Host == "" {
		fingerTable.SelfEntry.Host = possiblePrevious.Host
		fingerTable.SelfEntry.Hash = possiblePrevious.Hash
		return
	}
	selfDiff := diff(fingerTable.SelfEntry.Hash, fingerTable.PreviousEntry.Hash)
	targetDiff := diff(possiblePrevious.Hash, fingerTable.PreviousEntry.Hash)
	if selfDiff > targetDiff && possiblePrevious.Hash < fingerTable.SelfEntry.Hash {
		fingerTable.PreviousEntry.Hash = possiblePrevious.Hash
		fingerTable.PreviousEntry.Host = possiblePrevious.Host
	}
}

func GetFingerTableFromProto(protoTable *pbNode.FingerTable) *FingerTable {
	fingerTable := &FingerTable{
		PreviousEntry: Entry{
			Hash: protoTable.Previous.Hash,
			Host: protoTable.Previous.Host,
		},
		SelfEntry: Entry{
			Hash: protoTable.SelfEntry.Hash,
			Host: protoTable.SelfEntry.Host,
		},
	}
	for ind, elem := range protoTable.Entry {
		fingerTable.Entries[ind].Hash = elem.Hash
		fingerTable.Entries[ind].Host = elem.Host
	}

	return fingerTable
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
	protoFingerTable.SelfEntry = &pbNode.FingerTable_Entry{
		Hash: fingerTable.SelfEntry.Hash,
		Host: fingerTable.SelfEntry.Host,
	}
	return protoFingerTable
}

func (fingerTable *FingerTable) Add(entry *Entry) Update {
	update := Update{}
	if fingerTable.SelfEntry.Host == entry.Host {
		return update
	}
	fingerTable.chosePrevious(entry)
	if fingerTable.Entries[0].Host == "" {
		for ind := range fingerTable.Entries {
			fingerTable.Entries[ind] = *entry
			update.updates[ind] = fingerTable.Entries[ind].Host
		}
		return update
	}
	for i := 0; i < len(fingerTable.Entries); i++ {
		target := fingerTable.SelfEntry.Hash + (1 << uint64(i))
		if entry.Hash >= target {
			if diff(fingerTable.Entries[i].Hash, target) > diff(target, entry.Hash) {
				fingerTable.Entries[i] = *entry
				update.updates[i] = entry.Host
			}
		}
	}
	return update

}

func (fingerTable *FingerTable) Route(Hash uint64) (string, int) {
	// Over zero trip check
	if fingerTable.SelfEntry.Hash < fingerTable.PreviousEntry.Hash {
		if Hash >= fingerTable.PreviousEntry.Hash || Hash < fingerTable.SelfEntry.Hash {
			return fingerTable.SelfEntry.Host, 64
		}
	}

	size := uint64(0)

	// fingerTable size check (TODO: add size as member of fingerTable)
	for _, elem := range fingerTable.Entries {
		if elem.Host != "" {
			size++
		}
	}

	// self check
	if Hash >= fingerTable.PreviousEntry.Hash && Hash < fingerTable.SelfEntry.Hash || size == 0 {
		return fingerTable.SelfEntry.Host, 64
	}

	var HashHolder uint64
	index := uint64(len(fingerTable.Entries))

	for ind, elem := range fingerTable.Entries {
		if elem.Hash < uint64(Hash) && elem.Hash > HashHolder && elem.Hash > fingerTable.SelfEntry.Hash {
			index = uint64(ind)
			HashHolder = elem.Hash
		}
	}

	if index == uint64(len(fingerTable.Entries)) {
		return fingerTable.Entries[0].Host, 0
	}
	return fingerTable.Entries[index].Host, int(index)
}
