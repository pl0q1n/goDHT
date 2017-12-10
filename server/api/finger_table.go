package api

import (
	"log"
	"math"

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
	if fingerTable.SelfEntry.Host == "" {
		fingerTable.SelfEntry.Host = possiblePrevious.Host
		fingerTable.SelfEntry.Hash = possiblePrevious.Hash
		return
	}
	selfDiff := diff(fingerTable.SelfEntry.Hash, fingerTable.SelfEntry.Hash)
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
		log.Printf("First branch of add with: %s", entry.Host)
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
				log.Printf("Second branch of add with old: %s and new %s and index: %d", fingerTable.Entries[i].Host, entry.Host, i)
				fingerTable.Entries[i] = *entry
				update.updates[i] = entry.Host
			}
		}
	}
	return update

}

func (fingerTable *FingerTable) Route(Hash uint64) string {
	if fingerTable.SelfEntry.Hash < fingerTable.PreviousEntry.Hash {
		if Hash >= fingerTable.PreviousEntry.Hash || Hash < fingerTable.SelfEntry.Hash {
			return fingerTable.SelfEntry.Host
		}
	}

	if Hash >= fingerTable.PreviousEntry.Hash && Hash < fingerTable.SelfEntry.Hash {
		return fingerTable.SelfEntry.Host
	}

	var min Pair = Pair{
		first:  math.MaxUint64,
		second: uint64(len(fingerTable.Entries)),
	}

	for ind, elem := range fingerTable.Entries {
		if diff(elem.Hash, Hash) < min.first && elem.Hash > Hash {
			min.first = diff(elem.Hash, Hash)
			min.second = uint64(ind)
		}
	}

	if min.second == uint64(len(fingerTable.Entries)) {
		return fingerTable.Entries[0].Host
	}
	return fingerTable.Entries[min.second].Host
}
