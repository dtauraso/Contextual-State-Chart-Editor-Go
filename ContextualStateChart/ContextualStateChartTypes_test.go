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
		0: {ID: 0,
			MapValues: map[string]int{
				"parents":             1,
				"Name":                20,
				"FunctionCode":        30,
				"StartChildren":       40,
				"Next":                50,
				"Values":              60,
				"LockedByStates":      70,
				"LockedByStatesCount": 80}}}

	got :=
		CollectMaps(
			MapValue("parents", MapValueString("0", "-1")),
			MapValue("Name", ArrayValueStrings("I am a test", "StarbucksMachine")),
			MapValueString("FunctionCode", "ReturnTrue"),
			MapValue("StartChildren",
				CollectMaps(
					MapValue("Edges",
						ArrayValue(
							ArrayValue("state1 name1", "state1 name2"),
							ArrayValue("state2 name1", "state2 name2"))),
					MapValueString("AreParallel", "true")),
			),
			MapValue("Next",
				CollectMaps(
					MapValue("Edges",
						ArrayValue(
							ArrayValue("state1 name1", "state1 name2"),
							ArrayValue("state2 name1", "state2 name2"))),
					MapValueString("AreParallel", "true"))),
			MapValue("Values",
				CollectMaps(
					MapValueString("drinkOrder", ""),
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
