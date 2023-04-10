package TrieTree

type NamesTrie struct {
	NamePartTree map[string]int `json:"NextNamePartTree"`
	StateID      int            `json:"StateID",omitempty`
}
type InputParameters struct {
	Name    []string
	StateID int
}

type TrieTree []NamesTrie

func (trieTree TrieTree) Insert(input InputParameters) TrieTree {
	name := input.Name
	stateID := input.StateID

	if len(name) == 0 {
		return trieTree
	}
	if len(trieTree) == 0 {
		namePart := name[0]

		trieTree = append(trieTree, NamesTrie{NamePartTree: map[string]int{namePart: 1}, StateID: -1})
		trieTree = append(trieTree, NamesTrie{StateID: stateID})

		return trieTree
	}

	namesTracker := 0

	for i := 0; i < len(name); i++ {
		namePart := name[i]
		nextNameID, ok := trieTree[namesTracker].NamePartTree[namePart]

		if !ok {
			trieTree = append(trieTree, NamesTrie{StateID: -1})
			nextNameID = len(trieTree) - 1

			if trieTree[namesTracker].NamePartTree == nil {
				trieTree[namesTracker].NamePartTree = map[string]int{namePart: nextNameID}

			} else {
				trieTree[namesTracker].NamePartTree[namePart] = nextNameID
			}
		}
		namesTracker = nextNameID
	}

	// if item is new
	if trieTree[namesTracker].StateID == -1 {
		trieTree[namesTracker].StateID = stateID
	}

	return trieTree

}

func (trieTree TrieTree) Search(input []string) int {

	if len(input) == 0 {
		return -1
	}
	namesTracker := 0
	for i := 0; i < len(input); i++ {
		namePart := input[i]
		nextNameID, ok := trieTree[namesTracker].NamePartTree[namePart]
		if !ok {
			return -2
		}
		namesTracker = nextNameID
	}
	return trieTree[namesTracker].StateID
}

func (trieTree TrieTree) SearchInput(input []string) bool {
	return trieTree.Search(input) > -1
}
