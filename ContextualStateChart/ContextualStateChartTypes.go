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

MapValue("testKey", MapValueString("testKey2", "testValue2"))

{

	0: {
		id: 0
		MapValues:{"testKey": 1}
	},
	1: {
		id: 1
		MapValues:{"testKey2": 2}
	},
	2: {
		id: 2
		StringValue:"testValue2"
	},

}

ArrayValues("test1", "test2", "test3")

{

	0: {
		id: 0
		ArrayValues: [1, 2, 3]
	},
	1: {
		id: 1
		StringValue:"test1"
	},
	2: {
		id: 2
		StringValue:"test2"
	},
	3: {
		id: 3
		StringValue:"test3"
	},

}

ArrayValues(ArrayValues("test1", "test2"))
{

	0: {
		id: 0
		ArrayValues: [1]

	},
	1: {
		id: 1
		ArrayValues: [2, 3]
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
		ArrayValues: [1, 4]

	},
	1: {
		id: 1
		ArrayValues: [2, 3]
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
		ArrayValues: [5]
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
		ArrayValues: [2, 3]
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
func MapValueString(key string, value string) map[int]State {
	return nil
}
