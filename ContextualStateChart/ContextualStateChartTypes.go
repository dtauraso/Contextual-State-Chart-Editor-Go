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

func MapValueString(key, value string) map[int]State {
	states := make(map[int]State)
	states[0] = State{ID: 0, MapValues: map[string]int{key: 1}, TypeValueSet: "MapValues"}
	states[1] = State{ID: 1, StringValue: value, TypeValueSet: "StringValue"}
	return states
}

func MapValueInt(key string, value int) map[int]State {
	states := make(map[int]State)
	states[0] = State{ID: 0, MapValues: map[string]int{key: 1}, TypeValueSet: "MapValues"}
	states[1] = State{ID: 1, IntValue: value, TypeValueSet: "IntValue"}
	return states
}

func MapValueBool(key string, value bool) map[int]State {
	states := make(map[int]State)
	states[0] = State{ID: 0, MapValues: map[string]int{key: 1}, TypeValueSet: "MapValues"}
	states[1] = State{ID: 1, BoolValue: value, TypeValueSet: "BoolValue"}
	return states
}

func MapValue(key string, value map[int]State) map[int]State {
	states := make(map[int]State)
	states[0] = State{ID: 0, MapValues: map[string]int{key: 1}, TypeValueSet: "MapValues"}

	states, _ = addStates(states, value, 1)

	return states
}

func ArrayValueStrings(strings ...string) map[int]State {
	states := make(map[int]State)
	arrayMapValues := make(map[string]int)

	for i := 0; i < len(strings); i++ {
		arrayMapValues[strconv.Itoa(i)] = i + 1
	}
	states[0] = State{ID: 0, MapValues: arrayMapValues, TypeValueSet: "MapValues"}
	for i, myString := range strings {
		states[i+1] = State{ID: i + 1, StringValue: myString, TypeValueSet: "StringValue"}
	}
	return states
}

func ArrayValueInts(ints ...int) map[int]State {
	states := make(map[int]State)
	arrayMapValues := make(map[string]int)

	for i := 0; i < len(ints); i++ {
		arrayMapValues[strconv.Itoa(i)] = i + 1
	}
	states[0] = State{ID: 0, MapValues: arrayMapValues, TypeValueSet: "MapValues"}
	for i, myInt := range ints {
		states[i+1] = State{ID: i + 1, IntValue: myInt, TypeValueSet: "IntValue"}
	}
	return states
}
func ArrayValueBools(bools ...bool) map[int]State {
	states := make(map[int]State)
	arrayMapValues := make(map[string]int)

	for i := 0; i < len(bools); i++ {
		arrayMapValues[strconv.Itoa(i)] = i + 1
	}
	states[0] = State{ID: 0, MapValues: arrayMapValues, TypeValueSet: "MapValues"}
	for i, myBool := range bools {
		states[i+1] = State{ID: i + 1, BoolValue: myBool, TypeValueSet: "BoolValue"}
	}
	return states
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
		_, okString := prevElement.(string)
		_, okInt := prevElement.(int)

		_, okBool := prevElement.(bool)
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
func CollectMaps(elements ...any) map[int]State {

	states := make(map[int]State)
	// each element[0] can only have 1 key
	firstElement, _ := elements[0].(map[int]State)
	firstKey1 := getFirstKey(firstElement[0].MapValues)

	mapValues := map[string]int{firstKey1: 1}
	states = AddNewEntry(mapValues, states, mapTest1, 1, elements...)
	return states

}
func CollectMaps2(elements ...any) map[int]State {
	// 0, 2, 4 element ids are strings
	// 1, 3, 5 element ids are values (bool, int, string, states)
	return nil
}
