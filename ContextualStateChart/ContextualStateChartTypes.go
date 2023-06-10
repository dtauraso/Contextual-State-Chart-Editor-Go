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

// Parents: NDParentStateName -> ID
// MapValues: 1D string -> ID
/*
local variable
1D name -> primitive value
1D name -> array of ID's
1D name -> map of string keys -> ID's

database variable
ND name -> primitive value
ND name -> array of ID's
ND name -> map of string keys -> ID's
*/

type Atom struct {
	ID           int            `json:"ID"`
	BoolValue    bool           `json:"BoolValue,omitempty"`
	IntValue     int            `json:"IntValue,omitempty"`
	StringValue  string         `json:"StringValue,omitempty"`
	MapValues    map[string]int `json:"MapValues,omitempty"`
	Channel      chan Atom      `json:"Channel,omitempty"`
	ChannelWrite chan<- Atom    `json:"ChannelWrite,omitempty"`
	ChannelRead  <-chan Atom    `json:"ChannelRead,omitempty"`
	TypeValueSet string         `json:"TypeValueSet"`
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
			states[newIndex] = Atom{ID: newIndex, MapValues: newMapValues, TypeValueSet: "MapValues"}

		} else if value.TypeValueSet == "BoolValue" {
			states[newIndex] = Atom{ID: newIndex, BoolValue: value.BoolValue, TypeValueSet: "BoolValue"}
		} else if value.TypeValueSet == "IntValue" {
			states[newIndex] = Atom{ID: newIndex, IntValue: value.IntValue, TypeValueSet: "IntValue"}
		} else if value.TypeValueSet == "StringValue" {
			states[newIndex] = Atom{ID: newIndex, StringValue: value.StringValue, TypeValueSet: "StringValue"}
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
			states[j] = Atom{ID: j, BoolValue: myBoolValue, TypeValueSet: "BoolValue"}
		} else if okIntValue {
			states[j] = Atom{ID: j, IntValue: myIntValue, TypeValueSet: "IntValue"}
		} else if okStringValue {
			states[j] = Atom{ID: j, StringValue: myStringValue, TypeValueSet: "StringValue"}
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
	states[0] = Atom{ID: 0, MapValues: mapValues, TypeValueSet: "MapValues"}
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

func (g *Graph) AddStateHelper(state map[int]Atom, newIndex int) (stateID int) {
	g.States = addStates(g.States, state, newIndex)
	return len(g.States) - len(state)

}

func (g *Graph) AddState(state map[int]Atom) (stateID int) {
	if len(g.States) == 0 {
		return g.AddStateHelper(state, 0)
	}
	return g.AddStateHelper(state, len(g.States))
}

func (g *Graph) GetAtom(startAtom int, path []string) (atomID int, currentPath []string) {

	if len(path) == 0 {
		return -1, []string{}
	}

	tracker := startAtom
	pathFound := []string{}
	for i := 0; i < len(path); i++ {
		currentBranch := path[i]
		nextEdge, ok := g.States[tracker].MapValues[currentBranch]
		if !ok {
			return tracker, pathFound
		}
		pathFound = append(pathFound, currentBranch)

		tracker = nextEdge
	}
	return tracker, []string{}
}
func (g *Graph) InitMapValues(startIndex int) {
	g.States[startIndex] = Atom{
		ID:           startIndex,
		MapValues:    map[string]int{},
		TypeValueSet: "MapValues"}
}
func (g *Graph) TrieTreeInit() {
	pathToDataStructureIDs := []string{"data structure ID's"}
	dataStructureIDsID, path1 := g.GetAtom(0, pathToDataStructureIDs)
	length := len(g.States)
	var trieTreeStartIndex int

	if len(path1) == 0 {
		trieTreeStartIndex = length + 3
		g.AddState(
			CollectMaps("data structure ID's",
				CollectMaps("trie tree", trieTreeStartIndex)))
		g.InitMapValues(trieTreeStartIndex)
		return
	}
	pathToTrieTreeID := []string{"trie tree"}
	_, path2 := g.GetAtom(dataStructureIDsID, pathToTrieTreeID)
	if len(path2) == 0 {
		trieTreeStartIndex = length + 2
		g.AddState(
			CollectMaps("trie tree", trieTreeStartIndex))
		g.InitMapValues(trieTreeStartIndex)
		return
	}

}
func (g *Graph) TrieTreeAdd(strings []string) (newTrieTreeNodeID int) {

	trieTreeID, _ := g.GetAtom(0, []string{"data structure ID's", "trie tree"})
	ID, path := g.GetAtom(trieTreeID, strings)
	// no match and first item

	if ID == 0 {
		if len(path) == 0 {

		}
	}
	return 0
}
