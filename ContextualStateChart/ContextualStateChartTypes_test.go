package ContextualStateChartTypes

import (
	"reflect"
	"testing"
)

func TestMapValueString(t *testing.T) {

	want := map[int]State{
		0: {
			ID:        0,
			MapValues: map[string]int{"testKey": 1},
		},
		1: {
			ID:          1,
			StringValue: "testValue",
		},
	}

	got := MapValueString("testKey", "testValue")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestMapValue(t *testing.T) {

	want := map[int]State{

		0: {
			ID:        0,
			MapValues: map[string]int{"testKey": 1},
		},
		1: {
			ID:        1,
			MapValues: map[string]int{"testKey2": 2},
		},
		2: {
			ID:          2,
			StringValue: "testValue2",
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
			ID:        0,
			MapValues: map[string]int{"testKey": 1},
		},
		1: {
			ID:        1,
			MapValues: map[string]int{"0": 2, "1": 3},
		},
		2: {
			ID:          2,
			StringValue: "test1",
		},
		3: {
			ID:          3,
			StringValue: "test2",
		},
	}

	got := MapValue("testKey", ArrayValue("test1", "test2"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestMapValue3(t *testing.T) {

	want := map[int]State{

		0: {
			ID:        0,
			MapValues: map[string]int{"testKey": 1},
		},
		1: {
			ID:        1,
			MapValues: map[string]int{"testKey2": 2},
		},
		2: {
			ID:          2,
			StringValue: "testValue2",
		},
	}

	got := MapValue("testKey", MapValueString("testKey2", "testValue2"))

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
		},
		1: {
			ID:          1,
			StringValue: "test1",
		},
		2: {
			ID:          2,
			StringValue: "test2",
		},
		3: {
			ID:          3,
			StringValue: "test3",
		},
	}

	got := ArrayValueStrings("test1", "test2", "test3")

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
		},
		1: {
			ID: 1,
			MapValues: map[string]int{
				"0": 2,
				"1": 3,
				"2": 4},
		},
		2: {
			ID:          2,
			StringValue: "test1",
		},
		3: {
			ID:          3,
			StringValue: "test2",
		},
		4: {
			ID:          4,
			StringValue: "test3",
		},
	}

	got := ArrayValue(ArrayValueStrings("test1", "test2", "test3"))

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
		},
		1: {
			ID: 1,
			MapValues: map[string]int{
				"0": 2,
				"1": 3,
				"2": 4},
		},
		2: {
			ID:          2,
			StringValue: "test1",
		},
		3: {
			ID:          3,
			StringValue: "test2",
		},
		4: {
			ID:          4,
			StringValue: "test3",
		},
		5: {
			ID:          5,
			StringValue: "test4",
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
		},
		1: {
			ID:          1,
			StringValue: "test4",
		},
		2: {
			ID: 2,
			MapValues: map[string]int{
				"0": 3,
				"1": 4,
				"2": 5},
		},
		3: {
			ID:          3,
			StringValue: "test1",
		},
		4: {
			ID:          4,
			StringValue: "test2",
		},
		5: {
			ID:          5,
			StringValue: "test3",
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
		},
		1: {
			ID:        1,
			MapValues: map[string]int{"Name": 2},
		},
		2: {
			ID: 2,
			MapValues: map[string]int{
				"0": 3,
				"1": 4,
			},
		},
		3: {
			ID:          3,
			StringValue: "I am a test",
		},
		4: {
			ID:          4,
			StringValue: "StarbucksMachine",
		},
		5: {
			ID:        5,
			MapValues: map[string]int{"FunctionCode": 6},
		},
		6: {
			ID:          6,
			StringValue: "ReturnTrue",
		},
		7: {
			ID:        7,
			MapValues: map[string]int{"FunctionCode2": 8},
		},
		8: {
			ID:          8,
			StringValue: "ReturnTrue",
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
		0: {ID: 0, MapValues: map[string]int{"test": 1, "test2": 11}},
		1: {ID: 1, MapValues: map[string]int{"test": 2}},
		2: {
			ID: 2,
			MapValues: map[string]int{
				"Name":          3,
				"FunctionCode":  7,
				"FunctionCode2": 9},
		},
		3: {
			ID:        3,
			MapValues: map[string]int{"Name": 4},
		},
		4: {
			ID: 4,
			MapValues: map[string]int{
				"0": 5,
				"1": 6,
			},
		},
		5: {
			ID:          5,
			StringValue: "I am a test",
		},
		6: {
			ID:          6,
			StringValue: "StarbucksMachine",
		},
		7: {
			ID:        7,
			MapValues: map[string]int{"FunctionCode": 8},
		},
		8: {
			ID:          8,
			StringValue: "ReturnTrue",
		},
		9: {
			ID:        9,
			MapValues: map[string]int{"FunctionCode2": 10},
		},
		10: {
			ID:          10,
			StringValue: "ReturnTrue",
		},

		11: {ID: 11, MapValues: map[string]int{"test2": 12}},
		12: {
			ID: 12,
			MapValues: map[string]int{
				"Name":          13,
				"FunctionCode":  17,
				"FunctionCode2": 19},
		},
		13: {
			ID:        13,
			MapValues: map[string]int{"Name": 14},
		},
		14: {
			ID: 14,
			MapValues: map[string]int{
				"0": 15,
				"1": 16,
			},
		},
		15: {
			ID:          15,
			StringValue: "I am a test",
		},
		16: {
			ID:          16,
			StringValue: "StarbucksMachine",
		},
		17: {
			ID:        17,
			MapValues: map[string]int{"FunctionCode": 18},
		},
		18: {
			ID:          18,
			StringValue: "ReturnTrue",
		},
		19: {
			ID:        19,
			MapValues: map[string]int{"FunctionCode2": 20},
		},
		20: {
			ID:          20,
			StringValue: "ReturnTrue",
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

func TestState(t *testing.T) {
	want := map[int]State{

		// CollectMaps(
		0: {ID: 0,
			MapValues: map[string]int{
				"parents":             1,
				"Name":                4,
				"FunctionCode":        8,
				"StartChildren":       10,
				"Next":                23,
				"Values":              -1,
				"LockedByStates":      -1,
				"LockedByStatesCount": -1}},

		// MapValue("parents", MapValueString("0", "-1"))
		1: {ID: 1, MapValues: map[string]int{"parents": 2}},
		2: {ID: 2, MapValues: map[string]int{"0": 3}},
		3: {ID: 3, StringValue: "-1"},

		// MapValue("Name", ArrayValueStrings("I am a test", "StarbucksMachine"))
		4: {ID: 4, MapValues: map[string]int{"Name": 5}},
		5: {ID: 5, MapValues: map[string]int{"0": 6, "1": 7}},
		6: {ID: 6, StringValue: "I am a test"},
		7: {ID: 7, StringValue: "StarbucksMachine"},

		// MapValueString("FunctionCode", "ReturnTrue"
		8: {ID: 8, MapValues: map[string]int{"FunctionCode": 9}},
		9: {ID: 9, StringValue: "ReturnTrue"},

		// MapValue("StartChildren",
		10: {ID: 10, MapValues: map[string]int{"StartChildren": 11}},

		// CollectMaps(
		11: {ID: 11, MapValues: map[string]int{"Edges": 12, "AreParallel": 21}},

		// MapValue("Edges",
		12: {ID: 12, MapValues: map[string]int{"Edges": 13}},

		// ArrayValue(
		13: {ID: 13, MapValues: map[string]int{"0": 14}},
		14: {ID: 14, MapValues: map[string]int{"0": 15, "1": 18}},

		// ArrayValue("state1 name1", "state1 name2")
		15: {ID: 15, MapValues: map[string]int{"0": 16, "1": 17}},
		16: {ID: 16, StringValue: "state1 name1"},
		17: {ID: 17, StringValue: "state1 name2"},

		// ArrayValue("state2 name1", "state2 name2"))
		18: {ID: 18, MapValues: map[string]int{"0": 19, "1": 20}},
		19: {ID: 19, StringValue: "state1 name1"},
		20: {ID: 20, StringValue: "state1 name2"},

		// MapValueString("AreParallel", "true"))
		21: {ID: 21, MapValues: map[string]int{"AreParallel": 22}},
		22: {ID: 22, StringValue: "true"},

		/////////

		/////////
		// MapValue("Next",
		23: {ID: 23, MapValues: map[string]int{"Next": 24}},

		// CollectMaps(
		24: {ID: 24, MapValues: map[string]int{"Edges": 25, "AreParallel": 34}},

		// MapValue("Edges",
		25: {ID: 25, MapValues: map[string]int{"Edges": 26}},

		// ArrayValue(
		26: {ID: 26, MapValues: map[string]int{"0": 27}},
		27: {ID: 27, MapValues: map[string]int{"0": 28, "1": 31}},

		// ArrayValue("state1 name1", "state1 name2")
		28: {ID: 28, MapValues: map[string]int{"0": 29, "1": 30}},
		29: {ID: 29, StringValue: "state1 name1"},
		30: {ID: 30, StringValue: "state1 name2"},

		// ArrayValue("state2 name1", "state2 name2"))
		31: {ID: 31, MapValues: map[string]int{"0": 32, "1": 33}},
		32: {ID: 32, StringValue: "state1 name1"},
		33: {ID: 33, StringValue: "state1 name2"},

		// MapValueString("AreParallel", "true"))
		34: {ID: 34, MapValues: map[string]int{"AreParallel": 35}},
		35: {ID: 35, StringValue: "true"},

		// MapValue("Values",
		36: {ID: 36, MapValues: map[string]int{"Values": 37}},

		// CollectMaps(
		37: {ID: 37, MapValues: map[string]int{
			"drinkOrder":   38,
			"orderQueue":   40,
			"outputBuffer": 42}},

		// MapValueString("drinkOrder", "")
		38: {ID: 38, MapValues: map[string]int{"drinkOrder": 39}},
		39: {ID: 39, StringValue: ""},

		// MapValueString("orderQueue", "")
		40: {ID: 40, MapValues: map[string]int{"orderQueue": 41}},
		41: {ID: 41, StringValue: ""},

		// MapValueString("outputBuffer", "")

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
					MapValueString("drinkOrder", ""), /* recorded */
					MapValueString("orderQueue", ""),
					MapValueString("outputBuffer", ""))),
			MapValue("LockedByStates",
				CollectMaps(
					MapValueString("11", "true"))),
			MapValueString("LockedByStatesCount", "1"),
		)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
