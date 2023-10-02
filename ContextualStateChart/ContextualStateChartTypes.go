package ContextualStateChartTypes

import (
	// "fmt"
	// "fmt"
	// "reflect"
	// "fmt"
	// "fmt"
	// "fmt"
	"fmt"
	"sort"
	"strconv"
)

// Parents: NDParentStateName -> Id
// MapValues: 1D string -> Id
/*
local variable
1D name -> primitive value
1D name -> array of Id's
1D name -> map of string keys -> Id's

database variable
ND name -> primitive value
ND name -> array of Id's
ND name -> map of string keys -> Id's
*/

type Atom struct {
	Id           int            `json:"Id"`
	BoolValue    bool           `json:"BoolValue,omitempty"`
	IntValue     int            `json:"IntValue,omitempty"`
	StringValue  string         `json:"StringValue,omitempty"`
	MapValues    map[string]int `json:"MapValues,omitempty"`
	TypeValueSet string         `json:"TypeValueSet"`
	Parent       int            `json:"Parent,omitempty"`
}

func (a *Atom) Value() any {
	if a.TypeValueSet == "BoolValue" {
		return a.BoolValue
	} else if a.TypeValueSet == "IntValue" {
		return a.IntValue
	} else if a.TypeValueSet == "StringValue" {
		return a.StringValue
	}
	return nil
}

func (a *Atom) CloneWithOffset(j, childOffset, parentOffset int) (atomClone Atom) {
	// a.Parent + parentOffset
	// parent and child atoms are being cloned in a loop from caller and added to a map of type
	// map[int]Atom by adding new batches of entries by id as appending
	if a.TypeValueSet == "MapValues" {
		newMapValues := make(map[string]int)
		for key2, value2 := range a.MapValues {
			newMapValues[key2] = value2 + childOffset
		}
		return Atom{
			Id:           j,
			MapValues:    newMapValues,
			TypeValueSet: "MapValues",
			Parent:       a.Parent + parentOffset,
		}
	} else if a.TypeValueSet == "BoolValue" {
		return Atom{
			Id:           j,
			BoolValue:    a.BoolValue,
			TypeValueSet: "BoolValue",
			Parent:       a.Parent + parentOffset,
		}
	} else if a.TypeValueSet == "IntValue" {
		return Atom{
			Id:           j,
			IntValue:     a.IntValue,
			TypeValueSet: "IntValue",
			Parent:       a.Parent + parentOffset,
		}
	} else if a.TypeValueSet == "StringValue" {
		return Atom{
			Id:           j,
			StringValue:  a.StringValue,
			TypeValueSet: "StringValue",
			Parent:       a.Parent + parentOffset,
		}
	}
	return Atom{}
}

type AtomChan struct {
	Id           int            `json:"Id"`
	BoolValue    bool           `json:"BoolValue,omitempty"`
	IntValue     int            `json:"IntValue,omitempty"`
	StringValue  string         `json:"StringValue,omitempty"`
	MapValues    map[string]int `json:"MapValues,omitempty"`
	Channel      chan Atom      `json:"Channel,omitempty"`
	ChannelWrite chan<- Atom    `json:"ChannelWrite,omitempty"`
	ChannelRead  <-chan Atom    `json:"ChannelRead,omitempty"`

	TypeValueSet string `json:"TypeValueSet"`
	Parent       int    `json:"Parent,omitempty"`
}

func SaveString(s map[int]Atom, key int, newString string) {
	if entry, ok := s[key]; ok {
		entry.StringValue = newString
		entry.TypeValueSet = "StringValue"
		s[key] = entry
	}
}
func addAtoms(atoms, newAtoms map[int]Atom, newIndex int) map[int]Atom {

	// visiting keys in ascending order for offset formula to work
	firstNewIndex := newIndex
	// first newAtom is parent
	// parent's first parent is 0
	// child's parent is firstNewIndex
	value := newAtoms[0]
	// caller is adding 1 new parent
	atoms[newIndex] = value.CloneWithOffset(newIndex, firstNewIndex, 1)
	newIndex++

	for key := 1; key < len(newAtoms); key++ {
		value := newAtoms[key]
		atoms[newIndex] = value.CloneWithOffset(newIndex, firstNewIndex, firstNewIndex)
		newIndex++
	}
	return atoms
}

func arrayGetNewIndex(i int, elements ...any) (stringIndex string) {
	return strconv.Itoa(i)
}
func mapGetNewIndex(i int, elements ...any) (stringIndex string) {
	myString, _ := elements[i].(string)
	return myString

}

func arrayGetValueIndex(i int) (j int) {
	return i
}
func mapGetValueIndex(i int) (j int) {
	return i + 1
}

func addEntries(
	step int,
	getNewIndex func(int, ...any) string,
	valueIndex func(int) int,
	elements ...any) (Atoms map[int]Atom) {

	mapValues := make(map[string]int)
	atoms := make(map[int]Atom)

	j := 1
	for i := 0; i < len(elements); i += step {
		myString := getNewIndex(i, elements...)
		mapValues[myString] = j

		myElement := elements[valueIndex(i)]
		myBoolValue, okBoolValue := myElement.(bool)
		myIntValue, okIntValue := myElement.(int)
		myStringValue, okStringValue := myElement.(string)
		myAtomsValue, okAtomsValue := myElement.(map[int]Atom)

		if okBoolValue {
			atoms[j] = Atom{
				Id:           j,
				BoolValue:    myBoolValue,
				TypeValueSet: "BoolValue",
				Parent:       0}
		} else if okIntValue {
			atoms[j] = Atom{
				Id:           j,
				IntValue:     myIntValue,
				TypeValueSet: "IntValue",
				Parent:       0}
		} else if okStringValue {
			atoms[j] = Atom{
				Id:           j,
				StringValue:  myStringValue,
				TypeValueSet: "StringValue",
				Parent:       0}
		} else if okAtomsValue {
			atoms = addAtoms(atoms, myAtomsValue, j)
		}
		var offset int
		if okBoolValue || okIntValue || okStringValue {
			offset = 1
		} else if okAtomsValue {
			offset = len(myAtomsValue)
		}

		j += offset

	}
	atoms[0] = Atom{Id: 0, MapValues: mapValues, TypeValueSet: "MapValues", Parent: -1}
	return atoms
}
func ArrayValue(elements ...any) (Atoms map[int]Atom) {

	return addEntries(
		1,
		arrayGetNewIndex,
		arrayGetValueIndex,
		elements...)
}

func CollectMaps(elements ...any) (Atoms map[int]Atom) {
	// 0, 2, 4 element ids are strings
	// 1, 3, 5 element ids are values (bool, int, string, atom)
	if len(elements)%2 != 0 {
		return nil
	}
	return addEntries(
		2,
		mapGetNewIndex,
		mapGetValueIndex,
		elements...)
}

func makeString(states map[int]Atom, currentState int, indents, currentString string) (strings []string) {

	myState := states[currentState]
	typeName := myState.TypeValueSet
	myArray := make([]string, 0, 1)

	if typeName == "BoolValue" || typeName == "IntValue" || typeName == "StringValue" {
		if typeName == "BoolValue" {
			myArray = append(myArray, fmt.Sprintf("|%s%t|", indents, myState.BoolValue))
		} else if typeName == "IntValue" {
			myArray = append(myArray, fmt.Sprintf("|%s%d|", indents, myState.IntValue))
		} else if typeName == "StringValue" {
			myArray = append(myArray, "|"+indents+myState.StringValue+"|")
		}
		return myArray
	} else if typeName == "MapValues" {
		keys := make([]string, 0, len(myState.MapValues))
		for key1 := range myState.MapValues {
			keys = append(keys, key1)
		}
		sort.Strings(keys)
		for _, key := range keys {
			value := myState.MapValues[key]
			myArray = append(myArray, "|"+indents+key+": |")
			myArray = append(myArray, makeString(states, value, indents+"    ", "")...)
		}
	}
	return myArray
}
func convertToTree(states map[int]Atom) (strings []string) {
	return makeString(states, 0, "", "")
}

type Graph struct {
	Atoms map[int]Atom
}

const (
	DATA_STRUCTURE_IDS = "DataStructureIds"
)

func (g *Graph) GetState(stateName []string) *Graph {
	return &Graph{Atoms: g.Atoms}
}

func (s *Graph) GetVariable(variableName string) Atom {
	return s.Atoms[0]
}

func (g *Graph) InitGraph() {
	g.AddState(
		CollectMaps(DATA_STRUCTURE_IDS, CollectMaps()))
}
func (g *Graph) AddStateHelper(state map[int]Atom, newIndex int) (stateId int) {
	g.Atoms = addAtoms(g.Atoms, state, newIndex)
	return len(g.Atoms) - len(state)

}

func (g *Graph) AddState(state map[int]Atom) (stateId int) {
	// state keys are[0, len(states))
	if len(g.Atoms) == 0 {
		return g.AddStateHelper(state, 0)
	}
	return g.AddStateHelper(state, len(g.Atoms))
}

const (
	INPUT_ERROR = 0
	NOT_FOUND   = 1
	FOUND       = 2
)

func (g *Graph) GetAtom2(startAtom int, path []string) (idsFound []int) {

	tracker := startAtom
	for i := 0; i < len(path); i++ {
		currentBranch := path[i]
		nextEdge, ok := g.Atoms[tracker].MapValues[currentBranch]
		if !ok {
			return idsFound
		}
		idsFound = append(idsFound, nextEdge)
		tracker = nextEdge
	}
	return idsFound
}
func (g *Graph) GetAtom(startAtom int, path []string) (atomId int, currentPath []string, returnKind int) {

	// no clear way to know if it was unable to find item
	if len(path) == 0 {
		return -1, []string{}, INPUT_ERROR
	}

	tracker := startAtom
	pathFound := []string{}
	for i := 0; i < len(path); i++ {
		currentBranch := path[i]
		nextEdge, ok := g.Atoms[tracker].MapValues[currentBranch]
		if !ok {
			return tracker, pathFound, NOT_FOUND
		}
		pathFound = append(pathFound, currentBranch)

		tracker = nextEdge
	}
	return tracker, []string{}, FOUND
}
func (g *Graph) UpdateAtomMapValues(Id int, replacements map[string]int) {

	for item := range replacements {
		g.Atoms[Id].MapValues[item] = replacements[item]
	}
}
func (g *Graph) InitMapValues(startIndex int) {
	g.Atoms[startIndex] = Atom{
		Id:           startIndex,
		MapValues:    map[string]int{},
		TypeValueSet: "MapValues"}
}
func (g *Graph) TrieTreeInit() {
	pathToDataStructureIds := []string{DATA_STRUCTURE_IDS}
	dataStructureIdsId, _, returnKind1 := g.GetAtom(0, pathToDataStructureIds)
	length := len(g.Atoms)
	var trieTreeStartIndex int

	if returnKind1 == NOT_FOUND {
		// trieTreeStartIndex = length + 3
		// g.AddState(
		// 	CollectMaps(DATA_STRUCTURE_IDS,
		// 		CollectMaps("trie tree", trieTreeStartIndex)))
		// g.InitMapValues(trieTreeStartIndex)
		return
	}
	pathToTrieTreeId := []string{"trie tree"}
	_, _, returnKind2 := g.GetAtom(dataStructureIdsId, pathToTrieTreeId)

	if returnKind2 == NOT_FOUND {
		trieTreeStartIndex = length + 1
		// g.UpdateAtomMapValues(trieTreeId, map[string]int{"trie tree": length})
		// g.Atoms[length] = Atom{
		// 	Id:           length,
		// 	IntValue:     trieTreeStartIndex,
		// 	TypeValueSet: "IntValue"}

		g.InitMapValues(trieTreeStartIndex)
		return
	}

}

func (g *Graph) DoubleLinkListKeysAdd(path []string, startId int) (lastAtomNodeId int) {

	idsFound := g.GetAtom2(startId, path)
	foundPathLength := len(idsFound)
	pathLength := len(path)
	if foundPathLength < pathLength || len(idsFound) == 0 {
		return startId
	}

	length := len(g.Atoms)
	remainingPathLength := pathLength - foundPathLength

	id := idsFound[foundPathLength-1]
	// first atom added using CollectMaps is expected to have id = len(g.Atoms)
	g.Atoms[id].MapValues[path[foundPathLength]] = length
	lastString := path[pathLength-1]
	secondToLastString := path[pathLength-2]

	if remainingPathLength == 1 {
		g.Atoms[length] = Atom{
			Id:           length,
			MapValues:    map[string]int{lastString: length + 1},
			TypeValueSet: "MapValues",
			Parent:       id,
		}
		return length
	}

	mapList := CollectMaps(secondToLastString, lastString)

	for j := remainingPathLength; j >= foundPathLength; j-- {
		mapList = CollectMaps(path[j], mapList)
	}
	g.AddState(mapList)
	return length + remainingPathLength
}

func (g *Graph) TrieTreeAdd(strings []string, trieTreeId int) (newTrieTreeNodeId int) {

	// trieTreeId, _, _ := g.GetAtom(0, []string{DATA_STRUCTURE_IDS, "trie tree"})
	Id, path, returnKind := g.GetAtom(trieTreeId, strings)

	if returnKind == NOT_FOUND {
		pathLength := len(path)
		newIds := []int{}
		length := len(g.Atoms)
		remainingPath := []string{}
		remainingPathLength := len(strings) - pathLength
		for i := 0; i < remainingPathLength; i++ {
			newIds = append(newIds, length+i)
			remainingPath = append(remainingPath, strings[pathLength+i])
		}
		newIds = append(newIds,
			length+remainingPathLength,
			length+remainingPathLength+1)

		remainingPath = append(remainingPath,
			"Id",
			"id")
		g.Atoms[Id].MapValues[strings[pathLength]] = newIds[0]

		for j := 0; j < len(newIds); j++ {
			if remainingPath[j] != "id" {
				g.Atoms[newIds[j]] = Atom{
					Id:           newIds[j],
					MapValues:    map[string]int{remainingPath[j]: newIds[j+1]},
					TypeValueSet: "MapValues",
				}
			} else {
				g.Atoms[newIds[j]] = Atom{
					Id:           newIds[j],
					IntValue:     len(g.Atoms), // Add state right after TrieTreeAdd is done
					TypeValueSet: "IntValue",
				}
			}

		}

	}
	if returnKind == FOUND {
		// if no id attribute
		// 		add id attribute
	}
	return 0
}
