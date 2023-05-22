package ContextualStateChartTypes

import (
	// "fmt"
	// "fmt"
	// "reflect"
	// "fmt"
	// "fmt"
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
	TypeValueSet string         `json:"TypeValueSet"`
}

func SaveString(s map[int]State, key int, newString string) {
	if entry, ok := s[key]; ok {
		entry.StringValue = newString
		entry.TypeValueSet = "StringValue"
		s[key] = entry
	}
}
func addStates(states, newStates map[int]State, newIndex int) (map[int]State, int) {

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
	return states, newIndex
}
func AddNewEntry(
	values map[string]int,
	states map[int]State,
	getNewIndex func(int, ...any) string,
	elementStartingIndex int,
	elements ...any) map[int]State {

	// add new entry
	var newIndex string
	j := 1
	for i := 1; i < len(elements); i++ {

		newIndex = getNewIndex(i, elements...)

		prevElement := elements[i-1]
		_, okBool := prevElement.(bool)
		_, okInt := prevElement.(int)
		_, okString := prevElement.(string)
		myStates, okStates := prevElement.(map[int]State)

		var offset int
		if okString || okInt || okBool {
			offset = 1
		} else if okStates {
			offset = len(myStates)
		}
		values[newIndex] = j + offset
		j += offset

	}
	states[0] = State{ID: 0, MapValues: values, TypeValueSet: "MapValues"}

	// copy over elements to states
	newIndex2 := 1
	for i := 0; i < len(elements); i++ {

		element := elements[i]
		myBool, okBool := element.(bool)
		myInt, okInt := element.(int)

		myString, okString := element.(string)
		myStates, okStates := element.(map[int]State)

		if okBool {
			states[newIndex2] = State{ID: newIndex2, BoolValue: myBool, TypeValueSet: "BoolValue"}
			newIndex2++
		} else if okInt {
			states[newIndex2] = State{ID: newIndex2, IntValue: myInt, TypeValueSet: "IntValue"}
			newIndex2++
		} else if okString {
			states[newIndex2] = State{ID: newIndex2, StringValue: myString, TypeValueSet: "StringValue"}
			newIndex2++
		} else if okStates {
			states, newIndex2 = addStates(states, myStates, newIndex2)
		}

	}
	return states
}

func ArrayValue(elements ...any) map[int]State {
	states := make(map[int]State)
	if len(elements) == 0 {
		states[0] = State{ID: 0, MapValues: map[string]int{}, TypeValueSet: "MapValues"}
		return states
	}
	arrayMapValues := map[string]int{"0": 1}
	states = AddNewEntry(arrayMapValues, states, arrayTest1, 0, elements...)
	return states

}

func getFirstKey(mapValues map[string]int) string {

	keys := make([]string, 0, len(mapValues))
	for key := range mapValues {
		keys = append(keys, key)
	}
	return keys[0]
}

func arrayTest1(i int, elements ...any) string {
	return strconv.Itoa(i)
}
func mapTest1(i int, elements ...any) string {
	element, _ := elements[i].(map[int]State)
	return getFirstKey(element[0].MapValues)

}
func ArrayValue2(elements ...any) map[int]State {
	return nil
}

func CollectMaps(elements ...any) map[int]State {
	// 0, 2, 4 element ids are strings
	// 1, 3, 5 element ids are values (bool, int, string, states)
	if len(elements)%2 != 0 {
		return nil
	}
	mapValues := make(map[string]int)
	states := make(map[int]State)

	j := 1
	for i := 0; i < len(elements); i += 2 {
		myString, _ := elements[i].(string)
		mapValues[myString] = j

		myElement := elements[i+1]
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
			states, _ = addStates(states, myStatesValue, j)

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
