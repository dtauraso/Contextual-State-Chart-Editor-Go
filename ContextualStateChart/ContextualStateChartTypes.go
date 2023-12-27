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
	"sync"
	"time"
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
	Id                 int            `json:"Id"`
	BoolValue          bool           `json:"BoolValue,omitempty"`
	IntValue           int            `json:"IntValue,omitempty"`
	StringValue        string         `json:"StringValue,omitempty"`
	MapValues          map[string]int `json:"MapValues,omitempty"`
	TypeValueSet       string         `json:"TypeValueSet"`
	AtomParent         int            `json:"AtomParent,omitempty"`
	AtomParentChildKey string         `json:"AtomParentChildKey,omitempty"`
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

func (a *Atom) CloneWithOffset(j, childOffset, parentOffset int, childKey string) (atomClone Atom) {
	// a.AtomParent + parentOffset
	// parent and child atoms are being cloned in a loop from caller and added to a map of type
	// map[int]Atom by adding new batches of entries by id as appending
	// a.AtomParentChildKey is "" if a is a new parent from last round in loop from caller
	newChildKey := a.AtomParentChildKey
	if len(newChildKey) == 0 {
		newChildKey = childKey
	}
	if a.TypeValueSet == "MapValues" {
		newMapValues := make(map[string]int)
		for key2, value2 := range a.MapValues {
			newMapValues[key2] = value2 + childOffset
		}
		return Atom{
			Id:                 j,
			MapValues:          newMapValues,
			TypeValueSet:       "MapValues",
			AtomParent:         a.AtomParent + parentOffset,
			AtomParentChildKey: newChildKey,
		}
	} else if a.TypeValueSet == "BoolValue" {
		return Atom{
			Id:                 j,
			BoolValue:          a.BoolValue,
			TypeValueSet:       "BoolValue",
			AtomParent:         a.AtomParent + parentOffset,
			AtomParentChildKey: newChildKey,
		}
	} else if a.TypeValueSet == "IntValue" {
		return Atom{
			Id:                 j,
			IntValue:           a.IntValue,
			TypeValueSet:       "IntValue",
			AtomParent:         a.AtomParent + parentOffset,
			AtomParentChildKey: newChildKey,
		}
	} else if a.TypeValueSet == "StringValue" {
		return Atom{
			Id:                 j,
			StringValue:        a.StringValue,
			TypeValueSet:       "StringValue",
			AtomParent:         a.AtomParent + parentOffset,
			AtomParentChildKey: newChildKey,
		}
	}
	return Atom{}
}

func (a Atom) IsLeaf() bool {
	return a.TypeValueSet == "BoolValue" ||
		a.TypeValueSet == "IntValue" ||
		a.TypeValueSet == "StringValue"
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
	AtomParent   int    `json:"AtomParent,omitempty"`
}

func SaveString(s map[int]Atom, key int, newString string) {
	if entry, ok := s[key]; ok {
		entry.StringValue = newString
		entry.TypeValueSet = "StringValue"
		s[key] = entry
	}
}
func addAtoms(atoms, newAtoms map[int]Atom, newIndex int, childKey string) map[int]Atom {

	// assumes addEntries is the caller
	// visiting keys in ascending order for offset formula to work
	firstNewIndex := newIndex
	// first newAtom is parent
	// parent's first parent is 0
	// child's parent is firstNewIndex
	value := newAtoms[0]
	// caller is adding 1 new parent
	atoms[newIndex] = value.CloneWithOffset(newIndex, firstNewIndex, 1, childKey)
	newIndex++

	for key := 1; key < len(newAtoms); key++ {
		value := newAtoms[key]
		atoms[newIndex] = value.CloneWithOffset(newIndex, firstNewIndex, firstNewIndex, "")
		newIndex++
	}
	for childKey := range atoms[firstNewIndex].MapValues {
		childId := atoms[firstNewIndex].MapValues[childKey]
		childAtom := atoms[childId]
		childAtom.AtomParentChildKey = childKey
		atoms[childId] = childAtom
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

	// designed to be used only with nested CollectMaps and ArrayValue
	mapValues := make(map[string]int)
	atoms := make(map[int]Atom)

	j := 1
	for i := 0; i < len(elements); i += step {
		myString := getNewIndex(i, elements...)
		mapValues[myString] = j

		myElement := elements[valueIndex(i)]
		// assumes value is primitive unless value is of type map[int]Atom
		myBoolValue, okBoolValue := myElement.(bool)
		myIntValue, okIntValue := myElement.(int)
		myStringValue, okStringValue := myElement.(string)
		myAtomsValue, okAtomsValue := myElement.(map[int]Atom)

		if okBoolValue {
			atoms[j] = Atom{
				Id:                 j,
				BoolValue:          myBoolValue,
				TypeValueSet:       "BoolValue",
				AtomParent:         0,
				AtomParentChildKey: myString}
		} else if okIntValue {
			atoms[j] = Atom{
				Id:                 j,
				IntValue:           myIntValue,
				TypeValueSet:       "IntValue",
				AtomParent:         0,
				AtomParentChildKey: myString}
		} else if okStringValue {
			atoms[j] = Atom{
				Id:                 j,
				StringValue:        myStringValue,
				TypeValueSet:       "StringValue",
				AtomParent:         0,
				AtomParentChildKey: myString}
		} else if okAtomsValue {
			atoms = addAtoms(atoms, myAtomsValue, j, myString)
		}
		var offset int
		if okBoolValue || okIntValue || okStringValue {
			offset = 1
		} else if okAtomsValue {
			offset = len(myAtomsValue)
		}

		j += offset

	}
	atoms[0] = Atom{Id: 0, MapValues: mapValues, TypeValueSet: "MapValues", AtomParent: -1, AtomParentChildKey: ""}
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
	g.AddAtoms(
		CollectMaps(DATA_STRUCTURE_IDS, CollectMaps()))
}

func (g *Graph) AddAtomsHelper(atoms map[int]Atom, newIndex int) (stateId int) {
	g.Atoms = addAtoms(g.Atoms, atoms, newIndex, "")
	return len(g.Atoms) - len(atoms)

}

func (g *Graph) AddAtoms(atoms map[int]Atom) (stateId int) {
	// state keys are[0, len(states))
	if len(g.Atoms) == 0 {
		return g.AddAtomsHelper(atoms, 0)
	}
	return g.AddAtomsHelper(atoms, len(g.Atoms))
}

const (
	INPUT_ERROR = 0
	NOT_FOUND   = 1
	FOUND       = 2
)

func (g *Graph) GetValues(startAtom int, keys []string) (idsFound []int, ok bool) {

	tracker := startAtom
	for i := 0; i < len(keys); i++ {
		currentBranch := keys[i]
		nextEdge, ok := g.Atoms[tracker].MapValues[currentBranch]
		if !ok {
			return idsFound, false
		}
		idsFound = append(idsFound, nextEdge)
		tracker = nextEdge
	}
	return idsFound, true
}
func (g *Graph) GetAtom2(startAtom int, keys []string) (atom Atom) {

	idsFound, ok := g.GetValues(startAtom, keys)
	if ok {
		return g.Atoms[idsFound[len(idsFound)-1]]
	}
	return Atom{}
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

func validatePath(path ...any) bool {

	for i := 0; i < len(path)-1; i++ {
		_, ok := path[i].(string)
		if !ok {
			return false
		}
	}
	lastItem := path[len(path)-1]
	_, okBool := lastItem.(bool)
	_, okInt := lastItem.(int)
	_, okString := lastItem.(string)

	return okBool || okInt || okString

}

func arePrimitivesEqual(value any, a Atom) bool {

	valueBool, okBool := value.(bool)
	valueInt, okInt := value.(int)
	valueString, okString := value.(string)

	typeName := a.TypeValueSet

	if okBool && typeName == "BoolValue" {
		return valueBool == a.BoolValue
	} else if okInt && typeName == "IntValue" {
		return valueInt == a.IntValue
	} else if okString && typeName == "StringValue" {
		return valueString == a.StringValue
	}
	return false

}

func (g *Graph) DoubleLinkTreeKeysMatch(startId int, path ...any) bool {
	return false
}
func appendAtoms(atoms, newAtoms map[int]Atom, newIndex int) map[int]Atom {

	for key := 0; key < len(newAtoms); key++ {
		value := newAtoms[key]
		atoms[newIndex] = value.CloneWithOffset(newIndex, newIndex, newIndex, "")
		newIndex++
	}
	return atoms
}
func (g *Graph) DoubleLinkTreeKeysValueAdd(startId int, path ...any) (lastAtomNodeId int) {

	if len(path) == 0 {
		return startId
	}
	// value can be bool, int, or string
	if !validatePath(path...) {
		return startId
	}
	if len(path) < 2 {

		return startId
	}
	keys := []string{}
	valueLocation := len(path) - 1
	for i := 0; i < valueLocation; i++ {
		keys = append(keys, path[i].(string))
	}

	idsFound, ok := g.GetValues(startId, keys)

	idsFoundLength := len(idsFound)

	if idsFoundLength > 0 {
		arePrimitivesEqual := arePrimitivesEqual(
			path[valueLocation],
			g.Atoms[idsFound[idsFoundLength-1]])
		areKeysEqual := ok
		// check keys and value for match
		// keys match
		if areKeysEqual {
			// value matches
			if arePrimitivesEqual {
				return startId
			}
		}
	}

	lastParentId := 0

	if idsFoundLength == 1 {
		if !g.Atoms[idsFound[0]].IsLeaf() {
			lastParentId = idsFound[0]
		}
	} else if idsFoundLength >= 2 {
		lastIdFound := idsFound[idsFoundLength-1]
		if g.Atoms[lastIdFound].IsLeaf() {
			lastParentId = idsFound[idsFoundLength-2]
		} else {
			lastParentId = lastIdFound
		}
	}

	// unfound path

	g.Atoms[lastParentId].MapValues[path[idsFoundLength].(string)] = len(g.Atoms)

	valueString := path[len(path)-1]
	lastKeyString := path[len(path)-2]

	mapList := CollectMaps(lastKeyString, valueString)

	for i := len(path) - 1; i >= idsFoundLength; i-- {
		mapList = CollectMaps(path[i], mapList)
	}
	fmt.Println(mapList)
	g.Atoms = appendAtoms(g.Atoms, mapList, len(g.Atoms))
	return len(g.Atoms) - 1
}

/*
Delete(nodes, starting id, canContinueF, deleteF)

	tracker = starting id
	while canContinue(nodes, tracker)
		bump up trcker to parent of current tracker
		deletef(nodes, pref of tracker)

canContiueF0(nodes, tracker)

	visits whaever attributes or nodes needed to find out if we are done

deleteF0(nodes, tracker)

	deletes whatever nodes need to be deleted

Delete(nodes, starting id, canContinueF0, deleteF0)
Delete(nodes, starting id, canContinueF1, deleteF1)
*/
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

type TimeStep struct {
	id    int
	value Atom
}

func HierarchicalTimelines() {
	const (
		barista   = "barista"
		customer  = "customer"
		movement  = "movement"
		makeDrink = "makeDrink"
	)

	myGraph := Graph{Atoms: CollectMaps(
		"timelineTree", CollectMaps(
			"timeline", CollectMaps(
				"0", "a", "2", "c", "4", "b", "6", "a"),
			"children", ArrayValue(
				CollectMaps(
					"timeline", CollectMaps(
						"0", "a", "2", "c", "4", "b", "6", "a")),
				CollectMaps(
					"timeline", CollectMaps(
						"1", "4")),
				CollectMaps(
					"timeline", CollectMaps(
						"0", "1", "1", "2", "4", "1")),
			)),
	)}

	atom := myGraph.GetAtom2(0, []string{"Timelines"})
	/*
		{map[
			0:{0 false 0  map[Timelines:1] MapValues -1 }
			1:{1 false 0  map[0:2 1:6 2:8] MapValues 0 Timelines}
			2:{2 false 0  map[0:3 1:4 4:5] MapValues 1 0}
			3:{3 false 0  map[] IntValue 2 0}
			4:{4 false 2  map[] IntValue 2 1}
			5:{5 false 1  map[] IntValue 2 4}
			6:{6 false 0  map[1:7] MapValues 1 1}
			7:{7 false 0 makeDrink map[] StringValue 6 1}
			8:{8 false 0  map[0:9 1:10 4:11] MapValues 1 2}
			9:{9 false 0  map[] IntValue 8 0}
			10:{10 false 2  map[] IntValue 8 1}
			11:{11 false 1  map[] IntValue 8 4}]}
	*/
	// fmt.Printf("%v\n%v", atom, myGraph)

	timelines := make(map[int][]int)
	timestepChannel := make(chan TimeStep)
	for timestep := 0; timestep < 10; timestep++ {
		for key, i := range atom.MapValues {
			go func(key string, i int) {
				timelineAtomId := myGraph.Atoms[i].MapValues[fmt.Sprint(timestep)]
				timestepChannel <- TimeStep{i, myGraph.Atoms[timelineAtomId]}
			}(key, i)
		}
		for i := 0; i < len(atom.MapValues); i++ {
			r := <-timestepChannel
			_, ok := timelines[r.id]
			if !ok {
				timelines[r.id] = []int{}
			}
			timelines[r.id] = append(timelines[r.id], r.value.IntValue)
		}
	}
	fmt.Println(timelines)
	/*
		case 1: nothing is there to match
		case 2: there is 1 match but there is nothing to predict
		case 3: there is 1 match and there is at least 1 item to predict
	*/
	// Constructing a sample tree
	// root := &Node{
	// 	ID: 1,
	// 	Children: []*Node{
	// 		{ID: 2, Children: []*Node{
	// 			{ID: 4, Children: []*Node{}},
	// 			{ID: 5, Children: []*Node{}},
	// 		}},
	// 		{ID: 3, Children: []*Node{
	// 			{ID: 6, Children: []*Node{}},
	// 			{ID: 7, Children: []*Node{}},
	// 		}},
	// 	},
	// }

	// Use a WaitGroup to wait for all goroutines to finish
	// var wg sync.WaitGroup
	// wg.Add(1)

	// Start the process with the root node
	// go runGoroutines(root, &wg)

	// Wait for the root goroutine to finish
	// wg.Wait()
	// Example usage
	sets1 := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, 5, 6},
		{3, 4, 5},
	}

	result1 := intersectSets(sets1)

	fmt.Println("Intersection of sets:", result1)

	sets2 := [][]int{
		{1, 2, 3, 4, 5, 6},
		{2, 3, 5, 6},
		{3, 4, 5},
	}

	result2 := intersectSets(sets2)

	fmt.Println("Intersection of sets:", result2)

}

func intersectSets(sets [][]int) []int {
	// Create a map to store the frequency of each element
	elementFrequency := make(map[int]int)

	// Count the frequency of each element in the sets
	for _, set := range sets {
		uniqueSet := make(map[int]bool)
		for _, elem := range set {
			uniqueSet[elem] = true
		}

		// Increment the frequency of each element in the map
		for elem := range uniqueSet {
			elementFrequency[elem]++
		}
	}

	// Sort the elements based on their frequency in descending order
	sortedElements := make([]int, 0, len(elementFrequency))
	for elem, freq := range elementFrequency {
		sortedElements = append(sortedElements, elem*freq)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedElements)))
	return []int{}
	// Extract the original elements from the sorted list
	// result := make([]int, 0, len(sortedElements))
	// for _, elem := range sortedElements {

	// result = append(result, elem/elementFrequency[elem])
}

// return result

// Node represents a node in the tree data structure
type Node struct {
	ID       int
	Children []*Node
}

// Function that simulates some work for a node
func simulateWork(nodeID int) {
	fmt.Printf("Node %d is doing some work\n", nodeID)
	time.Sleep(time.Second)
	fmt.Printf("Node %d finished its work\n", nodeID)
}

// Function to execute the goroutines for a node and its children
func runGoroutines(node *Node, wg *sync.WaitGroup) {
	defer wg.Done()

	// Run goroutines for children first
	childWg := &sync.WaitGroup{}
	for _, child := range node.Children {
		childWg.Add(1)
		go runGoroutines(child, childWg)
	}

	// Wait for all child goroutines to finish
	childWg.Wait()

	// Now, execute the goroutine for the current node
	simulateWork(node.ID)
}

func Pattern() {

	// sequence of blocks for different directions
	// detect repeating
	// 1 small spiral using different colors
	// 1 large spiral using different colors
	// detect repeating for small spiral
	// detect parts of small spiral as part of large spiral
	// update small spiral to have parts of large spiral
	// 1 small large spiral
	// detect part of small large spiral using current spiral template
	// 1 wierd spiral (70% or less match with spiral detector)
	// detect spiral parts and generate spiral using the spiral parts it has detected

}
