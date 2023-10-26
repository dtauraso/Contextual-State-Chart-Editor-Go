package ContextualStateChartTypes

import (
	// "fmt"
	// "errors"
	"fmt"
	"reflect"
	"testing"
)

func TestMapValueString(t *testing.T) {

	want := map[int]Atom{
		0: {
			Id:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			StringValue:  "testValue",
			TypeValueSet: "StringValue",
			AtomParent:   0,
		},
	}

	got := CollectMaps("testKey", "testValue")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestMapValueInt(t *testing.T) {

	want := map[int]Atom{
		0: {
			Id:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			IntValue:     0,
			TypeValueSet: "IntValue",
			AtomParent:   0,
		},
	}

	got := CollectMaps("testKey", 0)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestMapValueBool(t *testing.T) {

	want := map[int]Atom{
		0: {
			Id:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			BoolValue:    false,
			TypeValueSet: "BoolValue",
			AtomParent:   0,
		},
	}

	got := CollectMaps("testKey", false)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue1(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			MapValues:    map[string]int{"testKey2": 2},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			StringValue:  "testValue2",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
	}

	got := CollectMaps("testKey", CollectMaps("testKey2", "testValue2"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue2(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			MapValues:    map[string]int{"0": 2, "1": 3, "2": 4},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			StringValue:  "test1",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
		3: {
			Id:           3,
			IntValue:     0,
			TypeValueSet: "IntValue",
			AtomParent:   1,
		},
		4: {
			Id:           4,
			BoolValue:    false,
			TypeValueSet: "BoolValue",
			AtomParent:   1,
		},
	}

	got := CollectMaps("testKey", ArrayValue("test1", 0, false))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestMapValue3(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			MapValues:    map[string]int{"testKey2": 2},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			StringValue:  "testValue2",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
	}

	got := CollectMaps("testKey", CollectMaps("testKey2", "testValue2"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue4(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			MapValues:    map[string]int{"testKey2": 2},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			IntValue:     0,
			TypeValueSet: "IntValue",
			AtomParent:   1,
		},
	}

	got := CollectMaps("testKey", CollectMaps("testKey2", 0))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue5(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id:           0,
			MapValues:    map[string]int{"testKey": 1},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			MapValues:    map[string]int{"testKey2": 2},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			BoolValue:    false,
			TypeValueSet: "BoolValue",
			AtomParent:   1,
		},
	}

	got := CollectMaps("testKey", CollectMaps("testKey2", false))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValues1(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id: 0,
			MapValues: map[string]int{
				"0": 1,
				"1": 2,
				"2": 3},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			StringValue:  "test1",
			TypeValueSet: "StringValue",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			StringValue:  "test2",
			TypeValueSet: "StringValue",
			AtomParent:   0,
		},
		3: {
			Id:           3,
			StringValue:  "test3",
			TypeValueSet: "StringValue",
			AtomParent:   0,
		},
	}

	got := ArrayValue("test1", "test2", "test3")

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValuesInts(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id: 0,
			MapValues: map[string]int{
				"0": 1,
				"1": 2,
				"2": 3},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			IntValue:     0,
			TypeValueSet: "IntValue",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			IntValue:     1,
			TypeValueSet: "IntValue",
			AtomParent:   0,
		},
		3: {
			Id:           3,
			IntValue:     2,
			TypeValueSet: "IntValue",
			AtomParent:   0,
		},
	}

	got := ArrayValue(0, 1, 2)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValuesBools(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id: 0,
			MapValues: map[string]int{
				"0": 1,
				"1": 2,
				"2": 3},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			BoolValue:    false,
			TypeValueSet: "BoolValue",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			BoolValue:    true,
			TypeValueSet: "BoolValue",
			AtomParent:   0,
		},
		3: {
			Id:           3,
			BoolValue:    false,
			TypeValueSet: "BoolValue",
			AtomParent:   0,
		},
	}

	got := ArrayValue(false, true, false)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestArrayValues21(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id: 0,
			MapValues: map[string]int{
				"0": 1},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id: 1,
			MapValues: map[string]int{
				"0": 2,
				"1": 3,
				"2": 4},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			StringValue:  "test1",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
		3: {
			Id:           3,
			StringValue:  "test2",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
		4: {
			Id:           4,
			StringValue:  "test3",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
	}

	got := ArrayValue(ArrayValue("test1", "test2", "test3"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValues2Ints(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id: 0,
			MapValues: map[string]int{
				"0": 1},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id: 1,
			MapValues: map[string]int{
				"0": 2,
				"1": 3,
				"2": 4},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			IntValue:     0,
			TypeValueSet: "IntValue",
			AtomParent:   1,
		},
		3: {
			Id:           3,
			IntValue:     1,
			TypeValueSet: "IntValue",
			AtomParent:   1,
		},
		4: {
			Id:           4,
			IntValue:     2,
			TypeValueSet: "IntValue",
			AtomParent:   1,
		},
	}

	got := ArrayValue(ArrayValue(0, 1, 2))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestArrayValues3(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id: 0,
			MapValues: map[string]int{
				"0": 1,
				"1": 5},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id: 1,
			MapValues: map[string]int{
				"0": 2,
				"1": 3,
				"2": 4},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			StringValue:  "test1",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
		3: {
			Id:           3,
			StringValue:  "test2",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
		4: {
			Id:           4,
			StringValue:  "test3",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
		5: {
			Id:           5,
			StringValue:  "test4",
			TypeValueSet: "StringValue",
			AtomParent:   0,
		},
	}

	got := ArrayValue(ArrayValue("test1", "test2", "test3"), "test4")

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestArrayValues4(t *testing.T) {

	want := map[int]Atom{

		0: {
			Id: 0,
			MapValues: map[string]int{
				"0": 1,
				"1": 2},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id:           1,
			StringValue:  "test4",
			TypeValueSet: "StringValue",
			AtomParent:   0,
		},
		2: {
			Id: 2,
			MapValues: map[string]int{
				"0": 3,
				"1": 4,
				"2": 5},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		3: {
			Id:           3,
			StringValue:  "test1",
			TypeValueSet: "StringValue",
			AtomParent:   2,
		},
		4: {
			Id:           4,
			StringValue:  "test2",
			TypeValueSet: "StringValue",
			AtomParent:   2,
		},
		5: {
			Id:           5,
			StringValue:  "test3",
			TypeValueSet: "StringValue",
			AtomParent:   2,
		},
	}

	got := ArrayValue("test4", ArrayValue("test1", "test2", "test3"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCollectMaps1(t *testing.T) {
	want := map[int]Atom{
		0: {
			Id: 0,
			MapValues: map[string]int{
				"Name":          1,
				"FunctionCode":  4,
				"FunctionCode2": 5},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id: 1,
			MapValues: map[string]int{
				"0": 2,
				"1": 3,
			},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		2: {
			Id:           2,
			StringValue:  "I am a test",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
		3: {
			Id:           3,
			StringValue:  "StarbucksMachine",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
		4: {
			Id:           4,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
			AtomParent:   0,
		},
		5: {
			Id:           5,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
			AtomParent:   0,
		},
	}

	got :=
		CollectMaps(
			"Name", ArrayValue("I am a test", "StarbucksMachine"),
			"FunctionCode", "ReturnTrue",
			"FunctionCode2", "ReturnTrue")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCollectMaps2(t *testing.T) {
	want := map[int]Atom{
		0: {Id: 0,
			MapValues:    map[string]int{"test": 1, "test2": 7},
			TypeValueSet: "MapValues",
			AtomParent:   -1,
		},
		1: {
			Id: 1,
			MapValues: map[string]int{
				"Name":          2,
				"FunctionCode":  5,
				"FunctionCode2": 6},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		2: {
			Id: 2,
			MapValues: map[string]int{
				"0": 3,
				"1": 4,
			},
			TypeValueSet: "MapValues",
			AtomParent:   1,
		},
		3: {
			Id:           3,
			StringValue:  "I am a test",
			TypeValueSet: "StringValue",
			AtomParent:   2,
		},
		4: {
			Id:           4,
			StringValue:  "StarbucksMachine",
			TypeValueSet: "StringValue",
			AtomParent:   2,
		},
		5: {
			Id:           5,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
		6: {
			Id:           6,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
			AtomParent:   1,
		},
		7: {
			Id: 7,
			MapValues: map[string]int{
				"Name":          8,
				"FunctionCode":  11,
				"FunctionCode2": 12},
			TypeValueSet: "MapValues",
			AtomParent:   0,
		},
		8: {
			Id: 8,
			MapValues: map[string]int{
				"0": 9,
				"1": 10,
			},
			TypeValueSet: "MapValues",
			AtomParent:   7,
		},
		9: {
			Id:           9,
			StringValue:  "I am a test",
			TypeValueSet: "StringValue",
			AtomParent:   8,
		},
		10: {
			Id:           10,
			StringValue:  "StarbucksMachine",
			TypeValueSet: "StringValue",
			AtomParent:   8,
		},
		11: {
			Id:           11,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
			AtomParent:   7,
		},
		12: {
			Id:           12,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
			AtomParent:   7,
		},
	}

	got :=
		CollectMaps(
			"test",
			CollectMaps(
				"Name", ArrayValue("I am a test", "StarbucksMachine"),
				"FunctionCode", "ReturnTrue",
				"FunctionCode2", "ReturnTrue"),
			"test2",
			CollectMaps(
				"Name", ArrayValue("I am a test", "StarbucksMachine"),
				"FunctionCode", "ReturnTrue",
				"FunctionCode2", "ReturnTrue"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestStateExistance(t *testing.T) {
	want := map[int]Atom{

		// CollectMaps(
		0: {Id: 0,
			MapValues: map[string]int{
				"FunctionCode":        6,
				"LockedByStates":      29,
				"LockedByStatesCount": 31,
				"Name":                3,
				"Next":                16,
				"StartChildren":       7,
				"Values":              25,
				"parents":             1,
			},
			TypeValueSet: "MapValues",
			AtomParent:   -1},

		// "parents", CollectMaps("0", "-1")
		1: {Id: 1, MapValues: map[string]int{"0": 2},
			TypeValueSet: "MapValues",
			AtomParent:   0},
		2: {Id: 2, IntValue: -1,
			TypeValueSet: "IntValue",
			AtomParent:   1},

		// "Name", ArrayValueStrings("I am a test", "StarbucksMachine")
		3: {Id: 3, MapValues: map[string]int{"0": 4, "1": 5},
			TypeValueSet: "MapValues",
			AtomParent:   0},
		4: {Id: 4, StringValue: "I am a test",
			TypeValueSet: "StringValue",
			AtomParent:   3},
		5: {Id: 5, StringValue: "StarbucksMachine",
			TypeValueSet: "StringValue",
			AtomParent:   3},

		// "FunctionCode", "ReturnTrue"
		6: {Id: 6, StringValue: "ReturnTrue",
			TypeValueSet: "StringValue",
			AtomParent:   0},

		// "StartChildren",
		// CollectMaps(
		7: {Id: 7, MapValues: map[string]int{"AreParallel": 15, "Edges": 8},
			TypeValueSet: "MapValues",
			AtomParent:   0},

		// "Edges",
		// ArrayValue(
		8: {Id: 8, MapValues: map[string]int{"0": 9, "1": 12},
			TypeValueSet: "MapValues",
			AtomParent:   7},

		// ArrayValue("state1 name1", "state1 name2")
		9: {Id: 9, MapValues: map[string]int{"0": 10, "1": 11},
			TypeValueSet: "MapValues",
			AtomParent:   8},
		10: {Id: 10, StringValue: "state1 name1",
			TypeValueSet: "StringValue",
			AtomParent:   9},
		11: {Id: 11, StringValue: "state1 name2",
			TypeValueSet: "StringValue",
			AtomParent:   9},

		// ArrayValue("state2 name1", "state2 name2"))
		12: {Id: 12, MapValues: map[string]int{"0": 13, "1": 14},
			TypeValueSet: "MapValues",
			AtomParent:   8},
		13: {Id: 13, StringValue: "state2 name1",
			TypeValueSet: "StringValue",
			AtomParent:   12},
		14: {Id: 14, StringValue: "state2 name2",
			TypeValueSet: "StringValue",
			AtomParent:   12},

		// "AreParallel", "true")
		15: {Id: 15, BoolValue: true,
			TypeValueSet: "BoolValue",
			AtomParent:   7},

		/////////

		/////////
		// "Next",
		// CollectMaps(
		16: {Id: 16, MapValues: map[string]int{"AreParallel": 24, "Edges": 17},
			TypeValueSet: "MapValues",
			AtomParent:   0},

		// "Edges",
		// ArrayValue(
		17: {Id: 17, MapValues: map[string]int{"0": 18, "1": 21},
			TypeValueSet: "MapValues",
			AtomParent:   16},

		// ArrayValue("state1 name1", "state1 name2")
		18: {Id: 18, MapValues: map[string]int{"0": 19, "1": 20},
			TypeValueSet: "MapValues",
			AtomParent:   17},
		19: {Id: 19, StringValue: "state1 name1",
			TypeValueSet: "StringValue",
			AtomParent:   18},
		20: {Id: 20, StringValue: "state1 name2",
			TypeValueSet: "StringValue",
			AtomParent:   18},

		// ArrayValue("state2 name1", "state2 name2"))
		21: {Id: 21, MapValues: map[string]int{"0": 22, "1": 23},
			TypeValueSet: "MapValues",
			AtomParent:   17},
		22: {Id: 22, StringValue: "state2 name1",
			TypeValueSet: "StringValue",
			AtomParent:   21},
		23: {Id: 23, StringValue: "state2 name2",
			TypeValueSet: "StringValue",
			AtomParent:   21},

		// "AreParallel", "true")
		24: {Id: 24, BoolValue: true,
			TypeValueSet: "BoolValue",
			AtomParent:   16},

		// "Values",
		// CollectMaps(
		25: {Id: 25, MapValues: map[string]int{
			"drinkOrder":   26,
			"orderQueue":   27,
			"outputBuffer": 28},
			TypeValueSet: "MapValues",
			AtomParent:   0},

		// "drinkOrder", ArrayValue()
		26: {Id: 26, MapValues: map[string]int{},
			TypeValueSet: "MapValues",
			AtomParent:   25},

		// "orderQueue", ArrayValue()
		27: {Id: 27, MapValues: map[string]int{},
			TypeValueSet: "MapValues",
			AtomParent:   25},

		// "outputBuffer", ArrayValue()
		28: {Id: 28, MapValues: map[string]int{},
			TypeValueSet: "MapValues",
			AtomParent:   25},

		// "LockedByStates",
		// CollectMaps(
		29: {Id: 29, MapValues: map[string]int{"11": 30},
			TypeValueSet: "MapValues",
			AtomParent:   0},

		// "11", "true"),
		30: {Id: 30, BoolValue: true,
			TypeValueSet: "BoolValue",
			AtomParent:   29},

		// "LockedByStatesCount", "1"
		31: {Id: 31, IntValue: 1,
			TypeValueSet: "IntValue",
			AtomParent:   0},
	}

	got :=
		CollectMaps(
			"parents", CollectMaps("0", -1),
			"Name", ArrayValue("I am a test", "StarbucksMachine"),
			"FunctionCode", "ReturnTrue",
			"StartChildren", CollectMaps(
				"Edges", ArrayValue(
					ArrayValue("state1 name1", "state1 name2"),
					ArrayValue("state2 name1", "state2 name2")),
				"AreParallel", true),
			"Next", CollectMaps(
				"Edges", ArrayValue(
					ArrayValue("state1 name1", "state1 name2"),
					ArrayValue("state2 name1", "state2 name2")),
				"AreParallel", true),
			"Values", CollectMaps(
				"drinkOrder", ArrayValue(),
				"orderQueue", ArrayValue(),
				"outputBuffer", ArrayValue()),
			"LockedByStates", CollectMaps(
				"11", true),
			"LockedByStatesCount", 1,
		)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestStateConnections(t *testing.T) {
	want := []string{
		"|FunctionCode: |",
		"|    ReturnTrue|",
		"|LockedByStates: |",
		"|    11: |",
		"|        true|",
		"|LockedByStatesCount: |",
		"|    1|",
		"|Name: |",
		"|    0: |",
		"|        I am a test|",
		"|    1: |",
		"|        StarbucksMachine|",
		"|Next: |",
		"|    AreParallel: |",
		"|        true|",
		"|    Edges: |",
		"|        0: |",
		"|            0: |",
		"|                state1 name1|",
		"|            1: |",
		"|                state1 name2|",
		"|        1: |",
		"|            0: |",
		"|                state2 name1|",
		"|            1: |",
		"|                state2 name2|",
		"|StartChildren: |",
		"|    AreParallel: |",
		"|        true|",
		"|    Edges: |",
		"|        0: |",
		"|            0: |",
		"|                state1 name1|",
		"|            1: |",
		"|                state1 name2|",
		"|        1: |",
		"|            0: |",
		"|                state2 name1|",
		"|            1: |",
		"|                state2 name2|",
		"|Values: |",
		"|    drinkOrder: |",
		"|    orderQueue: |",
		"|    outputBuffer: |",
		"|parents: |",
		"|    0: |",
		"|        -1|",
	}

	got :=
		convertToTree(CollectMaps(
			"parents", CollectMaps("0", -1),
			"Name", ArrayValue("I am a test", "StarbucksMachine"),
			"FunctionCode", "ReturnTrue",
			"StartChildren", CollectMaps(
				"Edges", ArrayValue(
					ArrayValue("state1 name1", "state1 name2"),
					ArrayValue("state2 name1", "state2 name2")),
				"AreParallel", true),
			"Next", CollectMaps(
				"Edges", ArrayValue(
					ArrayValue("state1 name1", "state1 name2"),
					ArrayValue("state2 name1", "state2 name2")),
				"AreParallel", true),
			"Values", CollectMaps(
				"drinkOrder", ArrayValue(),
				"orderQueue", ArrayValue(),
				"outputBuffer", ArrayValue()),
			"LockedByStates", CollectMaps(
				"11", true),
			"LockedByStatesCount", 1,
		))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestAddStates(t *testing.T) {
	t.Run("0 prior states before adding", func(t *testing.T) {
		myGraph := Graph{Atoms: map[int]Atom{}}
		firstIdWant := 0

		firstIdGot := myGraph.AddAtoms(ArrayValue("I am a test", "StarbucksMachine"))

		if firstIdWant != firstIdGot {
			t.Fatalf("wanted %v, got %v", firstIdWant, firstIdGot)
		}
	})

	t.Run("3 prior states before adding", func(t *testing.T) {
		myGraph := Graph{Atoms: map[int]Atom{}}
		firstIdWant := 4

		firstGraph := ArrayValue("I am a test", "StarbucksMachine", "test")
		myGraph.AddAtoms(firstGraph)
		firstIdGot := myGraph.AddAtoms(ArrayValue("I am a test", "StarbucksMachine"))

		if firstIdWant != firstIdGot {
			t.Fatalf("wanted %v, got %v", firstIdWant, firstIdGot)
		}
	})
}

func TestGetAtom(t *testing.T) {
	myGraph := Graph{Atoms: CollectMaps(
		"test",
		CollectMaps(
			"Name", ArrayValue("I am a test", "StarbucksMachine"),
			"FunctionCode", "ReturnTrue",
			"FunctionCode2", "ReturnTrue"),
		"test2",
		CollectMaps(
			"Name", ArrayValue("I am a test", "StarbucksMachine"),
			"FunctionCode", "ReturnTrue",
			"FunctionCode2", "ReturnTrue"))}
	t.Run("path has length == 0", func(t *testing.T) {

		idsFound, _ := myGraph.GetValues(0, []string{})

		wantIdsFound := []int{}
		gotIdsFound := idsFound
		if !(len(wantIdsFound) == len(gotIdsFound) && len(wantIdsFound) == 0) {
			t.Fatalf("wanted |%v|, got |%v|", wantIdsFound, gotIdsFound)
		}
	})

	t.Run("path does not exist 1", func(t *testing.T) {

		idsFound, _ := myGraph.GetValues(0, []string{"not there"})

		wantIdsFound := []int{}
		gotIdsFound := idsFound
		if !(len(idsFound) == 0 && len(wantIdsFound) == len(idsFound)) {
			t.Fatalf("wanted |%v|, got |%v|", wantIdsFound, gotIdsFound)
		}
	})
	t.Run("path does not exist 2", func(t *testing.T) {

		idsFound, _ := myGraph.GetValues(0, []string{"test", "not there"})

		wantPath := []int{1}
		gotPath := idsFound
		if len(wantPath) != len(gotPath) {
			t.Fatalf("wanted |%v|, got |%v|", wantPath, gotPath)
		}
	})
	t.Run("path exists and has length 1", func(t *testing.T) {

		idsFound, _ := myGraph.GetValues(0, []string{"test"})

		wantPath := []int{1}
		gotPath := idsFound
		if !reflect.DeepEqual(wantPath, gotPath) {
			t.Fatalf("wanted |%v|, got |%v|", wantPath, gotPath)
		}
	})

	t.Run("path exists and has length 2", func(t *testing.T) {

		idsFound, _ := myGraph.GetValues(0, []string{"test", "Name"})

		wantPath := []int{1, 2}
		gotPath := idsFound
		if !reflect.DeepEqual(wantPath, gotPath) {
			t.Fatalf("wanted |%v|, got |%v|", wantPath, gotPath)
		}
	})

	t.Run("path exists and has length 3", func(t *testing.T) {

		idsFound, _ := myGraph.GetValues(0, []string{"test", "Name", "0"})

		wantPath := []int{1, 2, 3}
		gotPath := idsFound
		if !reflect.DeepEqual(wantPath, gotPath) {
			t.Fatalf("wanted |%v|, got |%v|", wantPath, gotPath)
		}
	})
}

func TestUpdateAtom(t *testing.T) {
	want := Graph{Atoms: map[int]Atom{0: {Id: 0,
		MapValues: map[string]int{"a": 2, "b": 3, "c": 4}}}}

	got := Graph{Atoms: map[int]Atom{0: {Id: 0,
		MapValues: map[string]int{"a": 1, "b": 3}}}}
	got.UpdateAtomMapValues(0, map[string]int{"a": 2, "c": 4})

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestDoubleLinkListKeysAdd(t *testing.T) {
	presetGraph1 := Graph{
		Atoms: map[int]Atom{
			0: {MapValues: map[string]int{},
				TypeValueSet: "MapValues",
				AtomParent:   -1},
		},
	}

	t.Run("Add 0 keys", func(t *testing.T) {
		want := 0

		// search for key
		// if key is not there add it
		got := presetGraph1.DoubleLinkTreeKeysValueAdd(0)

		if want != got {
			t.Fatalf("wanted %v, got %v", want, got)
		}

	})
	t.Run("Add 1 (key, value) pair", func(t *testing.T) {
		want := 2

		// search for key
		// if key is not there add it
		got := presetGraph1.DoubleLinkTreeKeysValueAdd(0, "key1", "value1")
		fmt.Println(presetGraph1)
		/*
			generated
			map[
				0:{0 false 0  map[key1:1] MapValues -1}
				1:{1 false 0  map[value1:2] MapValues 0}
				2:{2 false 0  map[key1:3] MapValues 1}
				3:{3 false 0 value1 map[] StringValue 2}]



			append
				   {map[
					0:{0 false 0  map[key1:1] MapValues -1}
					1:{1 false 0  map[key1:2] MapValues 0}
					2:{2 false 0  map[value1:4] MapValues 2}
					3:{3 false 0  map[key1:6] MapValues 4}
					4:{4 false 0 value1 map[] StringValue 6}]}
		*/
		if want != got {
			t.Fatalf("wanted %v, got %v", want, got)
		}

	})
	// t.Run("Add 2 keys", func(t *testing.T) {
	// 	want := 4

	// 	got := myGraph.DoubleLinkListKeysValueAdd(0, "test1", "test2", "test4", "test5")
	// 	/*
	// 		0:{0 false 0  map[test4:1] MapValues -1}
	// 		1:{1 false 0  map[test4:2] MapValues 0}
	// 		2:{2 false 0 test5 map[] StringValue 1}

	// 		0:{0 false 0  map[test1:1]  -1}
	// 		1:{1 false 0  map[test2:2 test4:4]  0}
	// 		2:{2 false 0 test3 map[]  1}
	// 		3:{3 false 0  map[test4:4] MapValues 1}
	// 		4:{4 false 0  map[test4:5] MapValues 0}
	// 		5:{5 false 0  map[test4:6] MapValues 4}
	// 		6:{6 false 0 test5 map[] StringValue 5}
	// 	*/
	// 	// fmt.Println(myGraph)
	// 	if want != got {
	// 		t.Fatalf("wanted %v, got %v", want, got)
	// 	}

	// })
}

// func TestTrieTreeInit(t *testing.T) {
// 	myGraph := Graph{Atoms: map[int]Atom{
// 		0: {
// 			Id: 0,
// 			MapValues: map[string]int{
// 				"data structure Id's": 1,
// 			},
// 			TypeValueSet: "MapValues",
// 		},
// 		1: {
// 			Id: 1,
// 			MapValues: map[string]int{
// 				"trie tree": 2,
// 			},
// 			TypeValueSet: "MapValues",
// 		},
// 		2: {
// 			Id:           2,
// 			IntValue:     3,
// 			TypeValueSet: "IntValue",
// 		},
// 		3: {
// 			Id:           3,
// 			MapValues:    map[string]int{},
// 			TypeValueSet: "MapValues",
// 		},
// 	}}
// 	t.Run("no 'data structure Id's'", func(t *testing.T) {
// 		want := myGraph

// 		got := Graph{Atoms: map[int]Atom{}}
// 		got.TrieTreeInit()

// 		if !reflect.DeepEqual(want, got) {
// 			t.Fatalf("wanted %v, got %v", want, got)
// 		}
// 	})
// 	t.Run("no 'trie tree'", func(t *testing.T) {
// 		want := myGraph

// 		got := Graph{Atoms: map[int]Atom{
// 			0: {
// 				Id: 0,
// 				MapValues: map[string]int{
// 					"data structure Id's": 1,
// 				},
// 				TypeValueSet: "MapValues",
// 			},
// 			1: {
// 				Id:           1,
// 				MapValues:    map[string]int{},
// 				TypeValueSet: "MapValues",
// 			}}}
// 		got.TrieTreeInit()

// 		if !reflect.DeepEqual(want, got) {
// 			t.Fatalf("wanted %v, got %v", want, got)
// 		}
// 	})

// }
