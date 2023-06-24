package TrieTree

type NamesTrie struct {
	NamePartTree map[string]int `json:"NextNamePartTree"`
	StateId      int            `json:"StateId",omitempty`
}
type InputParameters struct {
	Name    []string
	StateId int
}

type TrieTree []NamesTrie

func (trieTree TrieTree) Insert(input InputParameters) TrieTree {
	name := input.Name
	stateId := input.StateId

	if len(name) == 0 {
		return trieTree
	}
	if len(trieTree) == 0 {
		namePart := name[0]

		trieTree = append(trieTree, NamesTrie{NamePartTree: map[string]int{namePart: 1}, StateId: -1})
		trieTree = append(trieTree, NamesTrie{StateId: stateId})

		return trieTree
	}

	namesTracker := 0

	for i := 0; i < len(name); i++ {
		namePart := name[i]
		nextNameId, ok := trieTree[namesTracker].NamePartTree[namePart]

		if !ok {
			trieTree = append(trieTree, NamesTrie{StateId: -1})
			nextNameId = len(trieTree) - 1

			if trieTree[namesTracker].NamePartTree == nil {
				trieTree[namesTracker].NamePartTree = map[string]int{namePart: nextNameId}

			} else {
				trieTree[namesTracker].NamePartTree[namePart] = nextNameId
			}
		}
		namesTracker = nextNameId
	}

	// if item is new
	if trieTree[namesTracker].StateId == -1 {
		trieTree[namesTracker].StateId = stateId
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
		nextNameId, ok := trieTree[namesTracker].NamePartTree[namePart]
		if !ok {
			return -2
		}
		namesTracker = nextNameId
	}
	return trieTree[namesTracker].StateId
}

func (trieTree TrieTree) SearchInput(input []string) bool {
	return trieTree.Search(input) > -1
}
