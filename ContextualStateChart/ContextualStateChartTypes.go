package ContextualStateChartTypes

import (
	"fmt"
	"reflect"
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
	ID          int            `json:"ID"`
	BoolValue   bool           `json:"BoolValue,omitempty"`
	IntValue    int            `json:"IntValue,omitempty"`
	StringValue string         `json:"StringValue,omitempty"`
	MapValues   map[string]int `json:"MapValues,omitempty"`
}

func SaveString(s map[int]State, key int, newString string) {
	if entry, ok := s[key]; ok {
		entry.StringValue = newString
		s[key] = entry
	}
}

func MapValueString(key, value string) map[int]State {
	states := make(map[int]State)
	states[0] = State{ID: 0, MapValues: map[string]int{key: 1}}
	states[1] = State{ID: 1, StringValue: value}
	return states
}

func MapValue(key string, value map[int]State) map[int]State {
	states := make(map[int]State)
	states[0] = State{ID: 0, MapValues: map[string]int{key: 1}}

	states, _ = addStates(states, value, 1)

	return states
}

func ArrayValueStrings(strings ...string) map[int]State {
	states := make(map[int]State)
	arrayMapValues := make(map[string]int)

	for i := 0; i < len(strings); i++ {
		arrayMapValues[strconv.Itoa(i)] = i + 1
	}
	states[0] = State{ID: 0, MapValues: arrayMapValues}
	for i, myString := range strings {
		states[i+1] = State{ID: i + 1, StringValue: myString}
	}
	return states
}
func addStates(states, newStates map[int]State, newIndex int) (map[int]State, int) {

	// visiting keys in ascending order for offset formula to work
	for key := 0; key < len(newStates); key++ {
		value := newStates[key]
		if !reflect.ValueOf(value.MapValues).IsZero() {
			newMapValues := make(map[string]int)
			for key2, value2 := range value.MapValues {
				newMapValues[key2] = value2 + (newIndex - key)
			}
			states[newIndex] = State{ID: newIndex, MapValues: newMapValues}
		} else if !reflect.ValueOf(value.StringValue).IsZero() {
			states[newIndex] = State{ID: newIndex, StringValue: value.StringValue}
		}

		newIndex++
	}
	return states, newIndex
}
func ArrayValue(elements ...any) map[int]State {
	states := make(map[int]State)
	arrayMapValues := map[string]int{"0": 1}

	for i := 1; i < len(elements); i++ {

		newIndex := strconv.Itoa(i)

		prevElement := elements[i-1]
		_, okString := prevElement.(string)
		if okString {
			arrayMapValues[newIndex] = i + 1
		}
		myStates, okStates := prevElement.(map[int]State)
		if okStates {
			arrayMapValues[newIndex] = i + len(myStates)
		}
	}
	states[0] = State{ID: 0, MapValues: arrayMapValues}

	newIndex := 1
	for i := 0; i < len(elements); i++ {

		element := elements[i]
		myString, okString := element.(string)

		if okString {
			states[newIndex] = State{ID: newIndex, StringValue: myString}
			newIndex++
		}
		myStates, okStates := element.(map[int]State)
		if okStates {
			states, newIndex = addStates(states, myStates, newIndex)
		}

	}
	return states

}

func getFirstKey(mapValues map[string]int) string {

	keys := make([]string, 0, len(mapValues))
	for key := range mapValues {
		keys = append(keys, key)
	}
	return keys[0]
}

func convertIntToString(myInt int) string {
	return strconv.Itoa(myInt)
}

func CollectMaps(elements ...any) map[int]State {
	states := make(map[int]State)
	// each element[0] can only have 1 key
	arrayMapValues := map[string]int{"0": 1}
	var newIndex string
	for i := 1; i < len(elements); i++ {

		element, _ := elements[i].(map[int]State)
		firstKey := getFirstKey(element[0].MapValues)
		if newIndex == firstKey {
			fmt.Println("map keys must be unique")
			return nil
		}
		newIndex = firstKey

		prevElement := elements[i-1]
		_, okString := prevElement.(string)
		if okString {
			arrayMapValues[newIndex] = i + 1
		}
		myStates, okStates := prevElement.(map[int]State)
		if okStates {
			arrayMapValues[newIndex] = i + len(myStates)
		}
	}
	states[0] = State{ID: 0, MapValues: arrayMapValues}

	newIndex := 1
	for i := 0; i < len(elements); i++ {

		element := elements[i]
		myString, okString := element.(string)

		if okString {
			states[newIndex] = State{ID: newIndex, StringValue: myString}
			newIndex++
		}
		myStates, okStates := element.(map[int]State)
		if okStates {
			states, newIndex = addStates(states, myStates, newIndex)
		}

	}
	return states

}
