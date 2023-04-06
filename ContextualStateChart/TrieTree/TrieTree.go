package TrieTree

type NamesTrie struct {
	NamePartTree map[string]int `json:"NextNamePartTree"`
	stateID      int            `json:"StateID",omitempty`
}
type InsertNameParameters struct {
	NamesTrie []NamesTrie
	Name      []string
	StateID   int
}

func InsertName(input InsertNameParameters) []NamesTrie {
	name := input.Name
	namesTrie := input.NamesTrie
	stateID := input.StateID

	if len(name) == 0 {
		return namesTrie
	}
	if len(namesTrie) == 0 {
		namePart := name[0]

		namesTrie = append(namesTrie, NamesTrie{NamePartTree: map[string]int{namePart: 1}, stateID: stateID})
		namesTrie = append(namesTrie, NamesTrie{stateID: -1})

		return namesTrie
	}
	namesTracker := 0

	for i := 0; i < len(name); i++ {
		namePart := name[i]
		nextNameID, ok := namesTrie[namesTracker].NamePartTree[namePart]

		if !ok {

			namesTrie = append(namesTrie, NamesTrie{stateID: -1})
			nextNameID = len(namesTrie) - 1

			if namesTrie[namesTracker].NamePartTree == nil {
				namesTrie[namesTracker] = NamesTrie{NamePartTree: map[string]int{namePart: nextNameID}, stateID: namesTrie[namesTracker].stateID}

			} else {
				namesTrie[namesTracker].NamePartTree[namePart] = nextNameID
			}
		}
		namesTracker = nextNameID
	}

	// if item is new
	if namesTrie[namesTracker].stateID == -1 {
		namesTrie[namesTracker].stateID = stateID
	}
	return namesTrie

}
