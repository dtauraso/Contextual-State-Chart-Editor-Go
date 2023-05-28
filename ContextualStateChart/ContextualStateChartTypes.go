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

type State struct {
	ID           int            `json:"ID"`
	BoolValue    bool           `json:"BoolValue,omitempty"`
	IntValue     int            `json:"IntValue,omitempty"`
	StringValue  string         `json:"StringValue,omitempty"`
	MapValues    map[string]int `json:"MapValues,omitempty"`
	Channel      chan State     `json:"Channel,omitempty"`
	ChannelWrite chan<- State   `json:"ChannelWrite,omitempty"`
	ChannelRead  <-chan State   `json:"ChannelRead,omitempty"`
	TypeValueSet string         `json:"TypeValueSet"`
}
type Graph struct {
	states     []State
	deletedIDs []int
}

func SaveString(s map[int]State, key int, newString string) {
	if entry, ok := s[key]; ok {
		entry.StringValue = newString
		entry.TypeValueSet = "StringValue"
		s[key] = entry
	}
}
func addStates(states, newStates map[int]State, newIndex int) map[int]State {

	// visiting keys in ascending order for offset formula to work
	for key := 0; key < len(newStates); key++ {
		value := newStates[key]
		if value.TypeValueSet == "MapValues" {
			newMapValues := make(map[string]int)
			offset := newIndex - key
			for key2, value2 := range value.MapValues {
				newMapValues[key2] = value2 + offset
			}
			states[newIndex] = State{ID: newIndex, MapValues: newMapValues, TypeValueSet: "MapValues"}

		} else if value.TypeValueSet == "BoolValue" {
			states[newIndex] = State{ID: newIndex, BoolValue: value.BoolValue, TypeValueSet: "BoolValue"}
		} else if value.TypeValueSet == "IntValue" {
			states[newIndex] = State{ID: newIndex, IntValue: value.IntValue, TypeValueSet: "IntValue"}
		} else if value.TypeValueSet == "StringValue" {
			states[newIndex] = State{ID: newIndex, StringValue: value.StringValue, TypeValueSet: "StringValue"}
		}

		newIndex++
	}
	return states
}

func arrayGetNewIndex(i int, elements ...any) string {
	return strconv.Itoa(i)
}
func mapGetNewIndex(i int, elements ...any) string {
	myString, _ := elements[i].(string)
	return myString

}

func arrayGetValueIndex(i int) int {
	return i
}
func mapGetValueIndex(i int) int {
	return i + 1
}

func addEntries(
	step int,
	getNewIndex func(int, ...any) string,
	valueIndex func(int) int,
	elements ...any) map[int]State {

	mapValues := make(map[string]int)
	states := make(map[int]State)

	j := 1
	for i := 0; i < len(elements); i += step {
		myString := getNewIndex(i, elements...)
		mapValues[myString] = j

		myElement := elements[valueIndex(i)]
		myBoolValue, okBoolValue := myElement.(bool)
		myIntValue, okIntValue := myElement.(int)
		myStringValue, okStringValue := myElement.(string)
		myStatesValue, okStatesValue := myElement.(map[int]State)

		if okBoolValue {
			states[j] = State{ID: j, BoolValue: myBoolValue, TypeValueSet: "BoolValue"}
		} else if okIntValue {
			states[j] = State{ID: j, IntValue: myIntValue, TypeValueSet: "IntValue"}
		} else if okStringValue {
			states[j] = State{ID: j, StringValue: myStringValue, TypeValueSet: "StringValue"}
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
	states[0] = State{ID: 0, MapValues: mapValues, TypeValueSet: "MapValues"}
	return states
}
func ArrayValue(elements ...any) map[int]State {

	return addEntries(
		1,
		arrayGetNewIndex,
		arrayGetValueIndex,
		elements...)
}

func CollectMaps(elements ...any) map[int]State {
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
func makeString(states map[int]State, currentState int, indents, currentString string) []string {

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
func convertToTree(states map[int]State) []string {
	return makeString(states, 0, "", "")
}
