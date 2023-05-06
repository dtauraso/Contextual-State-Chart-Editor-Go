package ContextualStateChartTypes

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
	ArrayValues []int          `json:"ArrayValues,omitempty"`
	MapValues   map[string]int `json:"MapValues,omitempty"`
}

func SaveString(s map[int]State, key int, newString string) {
	if entry, ok := s[key]; ok {
		entry.StringValue = newString
		s[key] = entry
	}
}

/*
MapValueString("testKey", "testValue")

{

	0: {
		id: 0
		MapValues:{"testKey": 1}
	},
	1: {
		id: 1
		StringValue:"testValue"
	},

}
*/
func MapValueString(key string, value string) map[int]State {
	return nil
}
