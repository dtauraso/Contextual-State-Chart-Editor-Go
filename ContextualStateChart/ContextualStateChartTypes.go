package ContextualStateChartTypes

import (
	// "fmt"
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
	i := 1
	for key := 0; key < len(value); key++ {
		value := value[key]
		if !reflect.ValueOf(value.MapValues).IsZero() {
			newMapValues := make(map[string]int)
			for key2, value2 := range value.MapValues {
				newMapValues[key2] = value2 + (i - key)
			}
			states[i] = State{ID: i, MapValues: newMapValues}
		} else if !reflect.ValueOf(value.StringValue).IsZero() {
			states[i] = State{ID: i, StringValue: value.StringValue}
		}

		i++
	}

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

/*

ArrayValues(ArrayValues("test1", "test2"))
{

	0: {
		id: 0
		MapValues: {"0": 1}

	},
	1: {
		id: 1
		MapValues: {"0": 2, "1": 3}
	},
	2: {
		id: 2
		StringValue:"test1"
	},
	3: {
		id: 3
		StringValue:"test2"
	},

}
ArrayValues(ArrayValues("test1", "test2"), ArrayValues("test3"))

{

	0: {
		id: 0
		MapValues: {"0": 1, "1": 4}

	},
	1: {
		id: 1
		MapValues: {"0": 2, "1": 3}
	},
	2: {
		id: 2
		StringValue:"test1"
	},
	3: {
		id: 3
		StringValue:"test2"
	},
	4: {
		id: 4
		MapValues: {"0": 5}
	},
	5: {
		id: 5
		StringValue:"test3"
	},

}

MapValue("testKey", ArrayValues("test1", "test2"))

{

	0: {
		id: 0
		MapValues:{"testKey": 1}
	},
	1: {
		id: 1
		MapValues: {"0": 2, "1": 3}
	},
	2: {
		id: 2
		StringValue:"test1"
	},
	3: {
		id: 3
		StringValue:"test2"
	},
	}
*/
