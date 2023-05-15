package ContextualStateChartTypes

import (
	"reflect"
	"testing"
)

func TestMapValueString(t *testing.T) {

	want := map[int]State{
		0: {
			ID:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			StringValue:  "testValue",
			TypeValueSet: "StringValue",
		},
	}

	got := MapValueString("testKey", "testValue")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestMapValueInt(t *testing.T) {

	want := map[int]State{
		0: {
			ID:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			IntValue:     0,
			TypeValueSet: "IntValue",
		},
	}

	got := MapValueInt("testKey", 0)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestMapValueBool(t *testing.T) {

	want := map[int]State{
		0: {
			ID:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			BoolValue:    false,
			TypeValueSet: "BoolValue",
		},
	}

	got := MapValueBool("testKey", false)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue(t *testing.T) {

	want := map[int]State{

		0: {
			ID:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			MapValues:    map[string]int{"testKey2": 2},
			TypeValueSet: "MapValues",
		},
		2: {
			ID:           2,
			StringValue:  "testValue2",
			TypeValueSet: "StringValue",
		},
	}

	got := MapValue("testKey", MapValueString("testKey2", "testValue2"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue2(t *testing.T) {

	want := map[int]State{

		0: {
			ID:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			MapValues:    map[string]int{"0": 2, "1": 3, "2": 4},
			TypeValueSet: "MapValues",
		},
		2: {
			ID:           2,
			StringValue:  "test1",
			TypeValueSet: "StringValue",
		},
		3: {
			ID:           3,
			IntValue:     0,
			TypeValueSet: "IntValue",
		},
		4: {
			ID:           4,
			BoolValue:    false,
			TypeValueSet: "BoolValue",
		},
	}

	got := MapValue("testKey", ArrayValue("test1", 0, false))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestMapValue3(t *testing.T) {

	want := map[int]State{

		0: {
			ID:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			MapValues:    map[string]int{"testKey2": 2},
			TypeValueSet: "MapValues",
		},
		2: {
			ID:           2,
			StringValue:  "testValue2",
			TypeValueSet: "StringValue",
		},
	}

	got := MapValue("testKey", MapValueString("testKey2", "testValue2"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue4(t *testing.T) {

	want := map[int]State{

		0: {
			ID:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			MapValues:    map[string]int{"testKey2": 2},
			TypeValueSet: "MapValues",
		},
		2: {
			ID:           2,
			IntValue:     0,
			TypeValueSet: "IntValue",
		},
	}

	got := MapValue("testKey", MapValueInt("testKey2", 0))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue5(t *testing.T) {

	want := map[int]State{

		0: {
			ID:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			MapValues:    map[string]int{"testKey2": 2},
			TypeValueSet: "MapValues",
		},
		2: {
			ID:           2,
			BoolValue:    false,
			TypeValueSet: "BoolValue",
		},
	}

	got := MapValue("testKey", MapValueBool("testKey2", false))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValues(t *testing.T) {

	want := map[int]State{

		0: {
			ID: 0,
			MapValues: map[string]int{
				"0": 1,
				"1": 2,
				"2": 3},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			StringValue:  "test1",
			TypeValueSet: "StringValue",
		},
		2: {
			ID:           2,
			StringValue:  "test2",
			TypeValueSet: "StringValue",
		},
		3: {
			ID:           3,
			StringValue:  "test3",
			TypeValueSet: "StringValue",
		},
	}

	got := ArrayValueStrings("test1", "test2", "test3")

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValuesInts(t *testing.T) {

	want := map[int]State{

		0: {
			ID: 0,
			MapValues: map[string]int{
				"0": 1,
				"1": 2,
				"2": 3},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			IntValue:     0,
			TypeValueSet: "IntValue",
		},
		2: {
			ID:           2,
			IntValue:     1,
			TypeValueSet: "IntValue",
		},
		3: {
			ID:           3,
			IntValue:     2,
			TypeValueSet: "IntValue",
		},
	}

	got := ArrayValueInts(0, 1, 2)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValuesBools(t *testing.T) {

	want := map[int]State{

		0: {
			ID: 0,
			MapValues: map[string]int{
				"0": 1,
				"1": 2,
				"2": 3},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			BoolValue:    false,
			TypeValueSet: "BoolValue",
		},
		2: {
			ID:           2,
			BoolValue:    true,
			TypeValueSet: "BoolValue",
		},
		3: {
			ID:           3,
			BoolValue:    false,
			TypeValueSet: "BoolValue",
		},
	}

	got := ArrayValueBools(false, true, false)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestArrayValues2(t *testing.T) {

	want := map[int]State{

		0: {
			ID: 0,
			MapValues: map[string]int{
				"0": 1},
			TypeValueSet: "MapValues",
		},
		1: {
			ID: 1,
			MapValues: map[string]int{
				"0": 2,
				"1": 3,
				"2": 4},
			TypeValueSet: "MapValues",
		},
		2: {
			ID:           2,
			StringValue:  "test1",
			TypeValueSet: "StringValue",
		},
		3: {
			ID:           3,
			StringValue:  "test2",
			TypeValueSet: "StringValue",
		},
		4: {
			ID:           4,
			StringValue:  "test3",
			TypeValueSet: "StringValue",
		},
	}

	got := ArrayValue(ArrayValueStrings("test1", "test2", "test3"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValues2Ints(t *testing.T) {

	want := map[int]State{

		0: {
			ID: 0,
			MapValues: map[string]int{
				"0": 1},
			TypeValueSet: "MapValues",
		},
		1: {
			ID: 1,
			MapValues: map[string]int{
				"0": 2,
				"1": 3,
				"2": 4},
			TypeValueSet: "MapValues",
		},
		2: {
			ID:           2,
			IntValue:     0,
			TypeValueSet: "IntValue",
		},
		3: {
			ID:           3,
			IntValue:     1,
			TypeValueSet: "IntValue",
		},
		4: {
			ID:           4,
			IntValue:     2,
			TypeValueSet: "IntValue",
		},
	}

	got := ArrayValue(ArrayValueInts(0, 1, 2))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestArrayValues3(t *testing.T) {

	want := map[int]State{

		0: {
			ID: 0,
			MapValues: map[string]int{
				"0": 1,
				"1": 5},
			TypeValueSet: "MapValues",
		},
		1: {
			ID: 1,
			MapValues: map[string]int{
				"0": 2,
				"1": 3,
				"2": 4},
			TypeValueSet: "MapValues",
		},
		2: {
			ID:           2,
			StringValue:  "test1",
			TypeValueSet: "StringValue",
		},
		3: {
			ID:           3,
			StringValue:  "test2",
			TypeValueSet: "StringValue",
		},
		4: {
			ID:           4,
			StringValue:  "test3",
			TypeValueSet: "StringValue",
		},
		5: {
			ID:           5,
			StringValue:  "test4",
			TypeValueSet: "StringValue",
		},
	}

	got := ArrayValue(ArrayValueStrings("test1", "test2", "test3"), "test4")

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestArrayValues4(t *testing.T) {

	want := map[int]State{

		0: {
			ID: 0,
			MapValues: map[string]int{
				"0": 1,
				"1": 2},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			StringValue:  "test4",
			TypeValueSet: "StringValue",
		},
		2: {
			ID: 2,
			MapValues: map[string]int{
				"0": 3,
				"1": 4,
				"2": 5},
			TypeValueSet: "MapValues",
		},
		3: {
			ID:           3,
			StringValue:  "test1",
			TypeValueSet: "StringValue",
		},
		4: {
			ID:           4,
			StringValue:  "test2",
			TypeValueSet: "StringValue",
		},
		5: {
			ID:           5,
			StringValue:  "test3",
			TypeValueSet: "StringValue",
		},
	}

	got := ArrayValue("test4", ArrayValueStrings("test1", "test2", "test3"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCollectMaps(t *testing.T) {
	want := map[int]State{
		0: {
			ID: 0,
			MapValues: map[string]int{
				"Name":          1,
				"FunctionCode":  5,
				"FunctionCode2": 7},
			TypeValueSet: "MapValues",
		},
		1: {
			ID:           1,
			MapValues:    map[string]int{"Name": 2},
			TypeValueSet: "MapValues",
		},
		2: {
			ID: 2,
			MapValues: map[string]int{
				"0": 3,
				"1": 4,
			},
			TypeValueSet: "MapValues",
		},
		3: {
			ID:           3,
			StringValue:  "I am a test",
			TypeValueSet: "StringValue",
		},
		4: {
			ID:           4,
			StringValue:  "StarbucksMachine",
			TypeValueSet: "StringValue",
		},
		5: {
			ID:           5,
			MapValues:    map[string]int{"FunctionCode": 6},
			TypeValueSet: "MapValues",
		},
		6: {
			ID:           6,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
		},
		7: {
			ID:           7,
			MapValues:    map[string]int{"FunctionCode2": 8},
			TypeValueSet: "MapValues",
		},
		8: {
			ID:           8,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
		},
	}

	got :=
		CollectMaps(
			MapValue("Name", ArrayValueStrings("I am a test", "StarbucksMachine")),
			MapValueString("FunctionCode", "ReturnTrue"),
			MapValueString("FunctionCode2", "ReturnTrue"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCollectMaps2(t *testing.T) {
	want := map[int]State{
		0: {ID: 0,
			MapValues:    map[string]int{"test": 1, "test2": 11},
			TypeValueSet: "MapValues"},
		1: {ID: 1,
			MapValues:    map[string]int{"test": 2},
			TypeValueSet: "MapValues"},
		2: {
			ID: 2,
			MapValues: map[string]int{
				"Name":          3,
				"FunctionCode":  7,
				"FunctionCode2": 9},
			TypeValueSet: "MapValues",
		},
		3: {
			ID:           3,
			MapValues:    map[string]int{"Name": 4},
			TypeValueSet: "MapValues",
		},
		4: {
			ID: 4,
			MapValues: map[string]int{
				"0": 5,
				"1": 6,
			},
			TypeValueSet: "MapValues",
		},
		5: {
			ID:           5,
			StringValue:  "I am a test",
			TypeValueSet: "StringValue",
		},
		6: {
			ID:           6,
			StringValue:  "StarbucksMachine",
			TypeValueSet: "StringValue",
		},
		7: {
			ID:           7,
			MapValues:    map[string]int{"FunctionCode": 8},
			TypeValueSet: "MapValues",
		},
		8: {
			ID:           8,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
		},
		9: {
			ID:           9,
			MapValues:    map[string]int{"FunctionCode2": 10},
			TypeValueSet: "MapValues",
		},
		10: {
			ID:           10,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
		},

		11: {
			ID:           11,
			MapValues:    map[string]int{"test2": 12},
			TypeValueSet: "MapValues",
		},
		12: {
			ID: 12,
			MapValues: map[string]int{
				"Name":          13,
				"FunctionCode":  17,
				"FunctionCode2": 19},
			TypeValueSet: "MapValues",
		},
		13: {
			ID:           13,
			MapValues:    map[string]int{"Name": 14},
			TypeValueSet: "MapValues",
		},
		14: {
			ID: 14,
			MapValues: map[string]int{
				"0": 15,
				"1": 16,
			},
			TypeValueSet: "MapValues",
		},
		15: {
			ID:           15,
			StringValue:  "I am a test",
			TypeValueSet: "StringValue",
		},
		16: {
			ID:           16,
			StringValue:  "StarbucksMachine",
			TypeValueSet: "StringValue",
		},
		17: {
			ID:           17,
			MapValues:    map[string]int{"FunctionCode": 18},
			TypeValueSet: "MapValues",
		},
		18: {
			ID:           18,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
		},
		19: {
			ID:           19,
			MapValues:    map[string]int{"FunctionCode2": 20},
			TypeValueSet: "MapValues",
		},
		20: {
			ID:           20,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
		},
	}

	got :=
		CollectMaps(
			MapValue("test",
				CollectMaps(
					MapValue("Name", ArrayValueStrings("I am a test", "StarbucksMachine")),
					MapValueString("FunctionCode", "ReturnTrue"),
					MapValueString("FunctionCode2", "ReturnTrue"))),
			MapValue("test2",
				CollectMaps(
					MapValue("Name", ArrayValueStrings("I am a test", "StarbucksMachine")),
					MapValueString("FunctionCode", "ReturnTrue"),
					MapValueString("FunctionCode2", "ReturnTrue"))))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestStateExistance(t *testing.T) {
	want := map[int]State{

		// CollectMaps(
		0: {ID: 0,
			MapValues: map[string]int{
				"FunctionCode":        8,
				"LockedByStates":      42,
				"LockedByStatesCount": 46,
				"Name":                4,
				"Next":                22,
				"StartChildren":       10,
				"Values":              34,
				"parents":             1,
			},
			TypeValueSet: "MapValues"},

		// MapValue("parents", MapValueString("0", "-1"))
		1: {ID: 1, MapValues: map[string]int{"parents": 2},
			TypeValueSet: "MapValues"},
		2: {ID: 2, MapValues: map[string]int{"0": 3},
			TypeValueSet: "MapValues"},
		3: {ID: 3, StringValue: "-1",
			TypeValueSet: "StringValue"},

		// MapValue("Name", ArrayValueStrings("I am a test", "StarbucksMachine"))
		4: {ID: 4, MapValues: map[string]int{"Name": 5},
			TypeValueSet: "MapValues"},
		5: {ID: 5, MapValues: map[string]int{"0": 6, "1": 7},
			TypeValueSet: "MapValues"},
		6: {ID: 6, StringValue: "I am a test",
			TypeValueSet: "StringValue"},
		7: {ID: 7, StringValue: "StarbucksMachine",
			TypeValueSet: "StringValue"},

		// MapValueString("FunctionCode", "ReturnTrue"
		8: {ID: 8, MapValues: map[string]int{"FunctionCode": 9},
			TypeValueSet: "MapValues"},
		9: {ID: 9, StringValue: "ReturnTrue",
			TypeValueSet: "StringValue"},

		// MapValue("StartChildren",
		10: {ID: 10, MapValues: map[string]int{"StartChildren": 11},
			TypeValueSet: "MapValues"},

		// CollectMaps(
		11: {ID: 11, MapValues: map[string]int{"AreParallel": 20, "Edges": 12},
			TypeValueSet: "MapValues"},

		// MapValue("Edges",
		12: {ID: 12, MapValues: map[string]int{"Edges": 13},
			TypeValueSet: "MapValues"},

		// ArrayValue(
		13: {ID: 13, MapValues: map[string]int{"0": 14, "1": 17},
			TypeValueSet: "MapValues"},

		// ArrayValue("state1 name1", "state1 name2")
		14: {ID: 14, MapValues: map[string]int{"0": 15, "1": 16},
			TypeValueSet: "MapValues"},
		15: {ID: 15, StringValue: "state1 name1",
			TypeValueSet: "StringValue"},
		16: {ID: 16, StringValue: "state1 name2",
			TypeValueSet: "StringValue"},

		// ArrayValue("state2 name1", "state2 name2"))
		17: {ID: 17, MapValues: map[string]int{"0": 18, "1": 19},
			TypeValueSet: "MapValues"},
		18: {ID: 18, StringValue: "state2 name1",
			TypeValueSet: "StringValue"},
		19: {ID: 19, StringValue: "state2 name2",
			TypeValueSet: "StringValue"},

		// MapValueString("AreParallel", "true"))
		20: {ID: 20, MapValues: map[string]int{"AreParallel": 21},
			TypeValueSet: "MapValues"},
		21: {ID: 21, StringValue: "true",
			TypeValueSet: "StringValue"},

		/////////

		/////////
		// MapValue("Next",
		22: {ID: 22, MapValues: map[string]int{"Next": 23},
			TypeValueSet: "MapValues"},

		// CollectMaps(
		23: {ID: 23, MapValues: map[string]int{"AreParallel": 32, "Edges": 24},
			TypeValueSet: "MapValues"},

		// MapValue("Edges",
		24: {ID: 24, MapValues: map[string]int{"Edges": 25},
			TypeValueSet: "MapValues"},

		// ArrayValue(
		25: {ID: 25, MapValues: map[string]int{"0": 26, "1": 29},
			TypeValueSet: "MapValues"},

		// ArrayValue("state1 name1", "state1 name2")
		26: {ID: 26, MapValues: map[string]int{"0": 27, "1": 28},
			TypeValueSet: "MapValues"},
		27: {ID: 27, StringValue: "state1 name1",
			TypeValueSet: "StringValue"},
		28: {ID: 28, StringValue: "state1 name2",
			TypeValueSet: "StringValue"},

		// ArrayValue("state2 name1", "state2 name2"))
		29: {ID: 29, MapValues: map[string]int{"0": 30, "1": 31},
			TypeValueSet: "MapValues"},
		30: {ID: 30, StringValue: "state2 name1",
			TypeValueSet: "StringValue"},
		31: {ID: 31, StringValue: "state2 name2",
			TypeValueSet: "StringValue"},

		// MapValueString("AreParallel", "true"))
		32: {ID: 32, MapValues: map[string]int{"AreParallel": 33},
			TypeValueSet: "MapValues"},
		33: {ID: 33, StringValue: "true",
			TypeValueSet: "StringValue"},

		// MapValue("Values",
		34: {ID: 34, MapValues: map[string]int{"Values": 35},
			TypeValueSet: "MapValues"},

		// CollectMaps(
		35: {ID: 35, MapValues: map[string]int{
			"drinkOrder":   36,
			"orderQueue":   38,
			"outputBuffer": 40},
			TypeValueSet: "MapValues"},

		// MapValueString("drinkOrder", "")
		36: {ID: 36, MapValues: map[string]int{"drinkOrder": 37},
			TypeValueSet: "MapValues"},
		37: {ID: 37, StringValue: "",
			TypeValueSet: "StringValue"},

		// MapValueString("orderQueue", "")
		38: {ID: 38, MapValues: map[string]int{"orderQueue": 39},
			TypeValueSet: "MapValues"},
		39: {ID: 39, StringValue: "",
			TypeValueSet: "StringValue"},

		// MapValueString("outputBuffer", "")
		40: {ID: 40, MapValues: map[string]int{"outputBuffer": 41},
			TypeValueSet: "MapValues"},
		41: {ID: 41, StringValue: "",
			TypeValueSet: "StringValue"},

		// MapValue("LockedByStates",
		42: {ID: 42, MapValues: map[string]int{"LockedByStates": 43},
			TypeValueSet: "MapValues"},

		// CollectMaps(
		43: {ID: 43, MapValues: map[string]int{"11": 44},
			TypeValueSet: "MapValues"},

		// MapValueString("11", "true"))),
		44: {ID: 44, MapValues: map[string]int{"11": 45},
			TypeValueSet: "MapValues"},
		45: {ID: 45, StringValue: "true",
			TypeValueSet: "StringValue"},

		// MapValueString("LockedByStatesCount", "1")
		46: {ID: 46, MapValues: map[string]int{"LockedByStatesCount": 47},
			TypeValueSet: "MapValues"},
		47: {ID: 47, StringValue: "1",
			TypeValueSet: "StringValue"},
	}

	got :=
		CollectMaps( /* recorded */
			MapValue("parents", MapValueString("0", "-1")),                         /* recorded */
			MapValue("Name", ArrayValueStrings("I am a test", "StarbucksMachine")), /* recorded */
			MapValueString("FunctionCode", "ReturnTrue"),                           /* recorded */
			MapValue("StartChildren", /* recorded */
				CollectMaps( /* recorded */
					MapValue("Edges", /* recorded */
						ArrayValue( /* recorded */
							ArrayValue("state1 name1", "state1 name2"),   /* recorded */
							ArrayValue("state2 name1", "state2 name2"))), /* recorded */
					MapValueString("AreParallel", "true")), /* recorded */
			),
			MapValue("Next", /* recorded */
				CollectMaps( /* recorded */
					MapValue("Edges", /* recorded */
						ArrayValue( /* recorded */
							ArrayValue("state1 name1", "state1 name2"),   /* recorded */
							ArrayValue("state2 name1", "state2 name2"))), /* recorded */
					MapValueString("AreParallel", "true"))), /* recorded */
			MapValue("Values", /* recorded */
				CollectMaps( /* recorded */
					MapValueString("drinkOrder", ""),     /* recorded */
					MapValueString("orderQueue", ""),     /* recorded */ // failed
					MapValueString("outputBuffer", ""))), /* recorded */ // failed
			MapValue("LockedByStates", /* recorded */
				CollectMaps( /* recorded */
					MapValueString("11", "true"))), /* recorded */
			MapValueString("LockedByStatesCount", "1"), /* recorded */
		)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
