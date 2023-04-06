package TrieTree

type NamesTrie struct {
	NamePartTree map[string]int `json:"NextNamePartTree",omitempty`
	stateID      int            `json:"StateID",omitempty`
}
type InsertNameParameters struct {
	namesTrie []NamesTrie
	name      []string
	stateID   int
}

func InsertName(input InsertNameParameters) {
	name := input.name
	namesTrie := input.namesTrie
	stateID := input.stateID

	namesTracker := 0
	for i := 0; i < len(name); i++ {
		namePart := name[i]
		nextNameID, ok := namesTrie[namesTracker].NamePartTree[namePart]
		if !ok {
			namesTrie = append(namesTrie, NamesTrie{NamePartTree: map[string]int{namePart: -1}})
			nextNameID = len(namesTrie)
			namesTrie[namesTracker].NamePartTree[namePart] = nextNameID
		}
		namesTracker = nextNameID
	}
	namesTrie[namesTracker].stateID = stateID

}
