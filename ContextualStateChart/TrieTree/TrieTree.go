package TrieTree

type NamesTrie struct {
	NamePartTree map[string]int `json:"NextNamePartTree"`
	StateID      int            `json:"StateID",omitempty`
}
type InsertNameParameters struct {
	Name    []string
	StateID int
}

type TrieTree []NamesTrie

func (trieTree TrieTree) InsertName(input InsertNameParameters) int {
	name := input.Name
	namesTrie := trieTree
	stateID := input.StateID

	if len(name) == 0 {
		return -1
	}
	if len(namesTrie) == 0 {
		namePart := name[0]

		namesTrie = append(namesTrie, NamesTrie{NamePartTree: map[string]int{namePart: 1}, StateID: stateID})
		namesTrie = append(namesTrie, NamesTrie{StateID: -1})

		return stateID
	}

	namesTracker := 0

	for i := 0; i < len(name); i++ {
		namePart := name[i]
		nextNameID, ok := namesTrie[namesTracker].NamePartTree[namePart]

		if !ok {
			namesTrie = append(namesTrie, NamesTrie{StateID: -1})
			nextNameID = len(namesTrie) - 1

			if namesTrie[namesTracker].NamePartTree == nil {
				namesTrie[namesTracker].NamePartTree = map[string]int{namePart: nextNameID}

			} else {
				namesTrie[namesTracker].NamePartTree[namePart] = nextNameID
			}
		}
		namesTracker = nextNameID
	}

	// if item is new
	if namesTrie[namesTracker].StateID == -1 {
		namesTrie[namesTracker].StateID = stateID
	}
	return stateID

}
