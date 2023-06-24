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
}

func SaveString(s map[int]Atom, key int, newString string) {
	if entry, ok := s[key]; ok {
		entry.StringValue = newString
		entry.TypeValueSet = "StringValue"
		s[key] = entry
	}
}
func addStates(states, newStates map[int]Atom, newIndex int) (atoms map[int]Atom) {

	// visiting keys in ascending order for offset formula to work
	for key := 0; key < len(newStates); key++ {
		value := newStates[key]
		if value.TypeValueSet == "MapValues" {
			newMapValues := make(map[string]int)
			offset := newIndex - key
			for key2, value2 := range value.MapValues {
				newMapValues[key2] = value2 + offset
			}
			states[newIndex] = Atom{Id: newIndex, MapValues: newMapValues, TypeValueSet: "MapValues"}

		} else if value.TypeValueSet == "BoolValue" {
			states[newIndex] = Atom{Id: newIndex, BoolValue: value.BoolValue, TypeValueSet: "BoolValue"}
		} else if value.TypeValueSet == "IntValue" {
			states[newIndex] = Atom{Id: newIndex, IntValue: value.IntValue, TypeValueSet: "IntValue"}
		} else if value.TypeValueSet == "StringValue" {
			states[newIndex] = Atom{Id: newIndex, StringValue: value.StringValue, TypeValueSet: "StringValue"}
		}

		newIndex++
	}
	return states
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
	states := make(map[int]Atom)

	j := 1
	for i := 0; i < len(elements); i += step {
		myString := getNewIndex(i, elements...)
		mapValues[myString] = j

		myElement := elements[valueIndex(i)]
		myBoolValue, okBoolValue := myElement.(bool)
		myIntValue, okIntValue := myElement.(int)
		myStringValue, okStringValue := myElement.(string)
		myStatesValue, okStatesValue := myElement.(map[int]Atom)

		if okBoolValue {
			states[j] = Atom{Id: j, BoolValue: myBoolValue, TypeValueSet: "BoolValue"}
		} else if okIntValue {
			states[j] = Atom{Id: j, IntValue: myIntValue, TypeValueSet: "IntValue"}
		} else if okStringValue {
			states[j] = Atom{Id: j, StringValue: myStringValue, TypeValueSet: "StringValue"}
		} else if okStatesValue {
			states = addStates(states, myStatesValue, j)

		}
		var offset int
		if okBoolValue || okIntValue || okStringValue {
			offset = 1
		} else if okStatesValue {
			offset = len(myStatesValue)
		}

		j += offset

	}
	states[0] = Atom{Id: 0, MapValues: mapValues, TypeValueSet: "MapValues"}
	return states
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
	// 1, 3, 5 element ids are values (bool, int, string, states)
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
	States map[int]Atom
}

const (
	DATA_STRUCTURE_IDS = "DataStructureIds"
)

func (g *Graph) InitGraph() {
	g.AddState(
		CollectMaps(DATA_STRUCTURE_IDS, CollectMaps("StateIds", 0)))
}
func (g *Graph) AddStateHelper(state map[int]Atom, newIndex int) (stateId int) {
	g.States = addStates(g.States, state, newIndex)
	return len(g.States) - len(state)

}

func (g *Graph) AddState(state map[int]Atom) (stateId int) {
	// state keys are[0, len(states))
	if len(g.States) == 0 {
		return g.AddStateHelper(state, 0)
	}
	return g.AddStateHelper(state, len(g.States))
}

const (
	INPUT_ERROR = 0
	NOT_FOUND   = 1
	FOUND       = 2
)

func (g *Graph) GetAtom(startAtom int, path []string) (atomId int, currentPath []string, returnKind int) {

	// no clear way to know if it was unable to find item
	if len(path) == 0 {
		return -1, []string{}, INPUT_ERROR
	}

	tracker := startAtom
	pathFound := []string{}
	for i := 0; i < len(path); i++ {
		currentBranch := path[i]
		nextEdge, ok := g.States[tracker].MapValues[currentBranch]
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
		g.States[Id].MapValues[item] = replacements[item]
	}
}
func (g *Graph) InitMapValues(startIndex int) {
	g.States[startIndex] = Atom{
		Id:           startIndex,
		MapValues:    map[string]int{},
		TypeValueSet: "MapValues"}
}
func (g *Graph) TrieTreeInit() {
	pathToDataStructureIds := []string{DATA_STRUCTURE_IDS}
	dataStructureIdsId, _, returnKind1 := g.GetAtom(0, pathToDataStructureIds)
	length := len(g.States)
	var trieTreeStartIndex int

	if returnKind1 == NOT_FOUND {
		trieTreeStartIndex = length + 3
		g.AddState(
			CollectMaps(DATA_STRUCTURE_IDS,
				CollectMaps("trie tree", trieTreeStartIndex)))
		g.InitMapValues(trieTreeStartIndex)
		return
	}
	pathToTrieTreeId := []string{"trie tree"}
	trieTreeId, _, returnKind2 := g.GetAtom(dataStructureIdsId, pathToTrieTreeId)

	if returnKind2 == NOT_FOUND {
		trieTreeStartIndex = length + 1
		g.UpdateAtomMapValues(trieTreeId, map[string]int{"trie tree": length})
		g.States[length] = Atom{
			Id:           length,
			IntValue:     trieTreeStartIndex,
			TypeValueSet: "IntValue"}

		g.InitMapValues(trieTreeStartIndex)
		return
	}

}
func (g *Graph) TrieTreeAdd(strings []string) (newTrieTreeNodeId int) {

	trieTreeId, _, _ := g.GetAtom(0, []string{DATA_STRUCTURE_IDS, "trie tree"})
	Id, path, returnKind := g.GetAtom(trieTreeId, strings)

	if returnKind == NOT_FOUND {
		pathLength := len(path)
		newIds := []int{}
		length := len(g.States)
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
		g.States[Id].MapValues[strings[pathLength]] = newIds[0]

		for j := 0; j < len(newIds); j++ {
			if remainingPath[j] != "id" {
				g.States[newIds[j]] = Atom{
					Id:           newIds[j],
					MapValues:    map[string]int{remainingPath[j]: newIds[j+1]},
					TypeValueSet: "MapValues",
				}
			} else {
				g.States[newIds[j]] = Atom{
					Id:           newIds[j],
					IntValue:     len(g.States), // add in separate state id tracker
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
