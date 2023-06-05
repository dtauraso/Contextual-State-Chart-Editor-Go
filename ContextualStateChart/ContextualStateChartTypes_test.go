package ContextualStateChartTypes

import (
	// "fmt"
	"errors"
	"reflect"
	"testing"
)

func TestMapValueString(t *testing.T) {

	want := map[int]Atom{
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

	got := CollectMaps("testKey", "testValue")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestMapValueInt(t *testing.T) {

	want := map[int]Atom{
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

	got := CollectMaps("testKey", 0)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestMapValueBool(t *testing.T) {

	want := map[int]Atom{
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

	got := CollectMaps("testKey", false)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue(t *testing.T) {

	want := map[int]Atom{

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

	got := CollectMaps("testKey", CollectMaps("testKey2", "testValue2"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue2(t *testing.T) {

	want := map[int]Atom{

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

	got := CollectMaps("testKey", ArrayValue("test1", 0, false))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestMapValue3(t *testing.T) {

	want := map[int]Atom{

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

	got := CollectMaps("testKey", CollectMaps("testKey2", "testValue2"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue4(t *testing.T) {

	want := map[int]Atom{

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

	got := CollectMaps("testKey", CollectMaps("testKey2", 0))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestMapValue5(t *testing.T) {

	want := map[int]Atom{

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

	got := CollectMaps("testKey", CollectMaps("testKey2", false))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValues(t *testing.T) {

	want := map[int]Atom{

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

	got := ArrayValue("test1", "test2", "test3")

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValuesInts(t *testing.T) {

	want := map[int]Atom{

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

	got := ArrayValue(0, 1, 2)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValuesBools(t *testing.T) {

	want := map[int]Atom{

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

	got := ArrayValue(false, true, false)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestArrayValues2(t *testing.T) {

	want := map[int]Atom{

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

	got := ArrayValue(ArrayValue("test1", "test2", "test3"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func TestArrayValues2Ints(t *testing.T) {

	want := map[int]Atom{

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

	got := ArrayValue(ArrayValue(0, 1, 2))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestArrayValues3(t *testing.T) {

	want := map[int]Atom{

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

	got := ArrayValue(ArrayValue("test1", "test2", "test3"), "test4")

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestArrayValues4(t *testing.T) {

	want := map[int]Atom{

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

	got := ArrayValue("test4", ArrayValue("test1", "test2", "test3"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCollectMaps(t *testing.T) {
	want := map[int]Atom{
		0: {
			ID: 0,
			MapValues: map[string]int{
				"Name":          1,
				"FunctionCode":  4,
				"FunctionCode2": 5},
			TypeValueSet: "MapValues",
		},
		1: {
			ID: 1,
			MapValues: map[string]int{
				"0": 2,
				"1": 3,
			},
			TypeValueSet: "MapValues",
		},
		2: {
			ID:           2,
			StringValue:  "I am a test",
			TypeValueSet: "StringValue",
		},
		3: {
			ID:           3,
			StringValue:  "StarbucksMachine",
			TypeValueSet: "StringValue",
		},
		4: {
			ID:           4,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
		},
		5: {
			ID:           5,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
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
		0: {ID: 0,
			MapValues:    map[string]int{"test": 1, "test2": 7},
			TypeValueSet: "MapValues"},
		1: {
			ID: 1,
			MapValues: map[string]int{
				"Name":          2,
				"FunctionCode":  5,
				"FunctionCode2": 6},
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
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
		},
		6: {
			ID:           6,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
		},
		7: {
			ID: 7,
			MapValues: map[string]int{
				"Name":          8,
				"FunctionCode":  11,
				"FunctionCode2": 12},
			TypeValueSet: "MapValues",
		},
		8: {
			ID: 8,
			MapValues: map[string]int{
				"0": 9,
				"1": 10,
			},
			TypeValueSet: "MapValues",
		},
		9: {
			ID:           9,
			StringValue:  "I am a test",
			TypeValueSet: "StringValue",
		},
		10: {
			ID:           10,
			StringValue:  "StarbucksMachine",
			TypeValueSet: "StringValue",
		},
		11: {
			ID:           11,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
		},
		12: {
			ID:           12,
			StringValue:  "ReturnTrue",
			TypeValueSet: "StringValue",
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
		0: {ID: 0,
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
			TypeValueSet: "MapValues"},

		// "parents", CollectMaps("0", "-1")
		1: {ID: 1, MapValues: map[string]int{"0": 2},
			TypeValueSet: "MapValues"},
		2: {ID: 2, IntValue: -1,
			TypeValueSet: "IntValue"},

		// "Name", ArrayValueStrings("I am a test", "StarbucksMachine")
		3: {ID: 3, MapValues: map[string]int{"0": 4, "1": 5},
			TypeValueSet: "MapValues"},
		4: {ID: 4, StringValue: "I am a test",
			TypeValueSet: "StringValue"},
		5: {ID: 5, StringValue: "StarbucksMachine",
			TypeValueSet: "StringValue"},

		// "FunctionCode", "ReturnTrue"
		6: {ID: 6, StringValue: "ReturnTrue",
			TypeValueSet: "StringValue"},

		// "StartChildren",
		// CollectMaps(
		7: {ID: 7, MapValues: map[string]int{"AreParallel": 15, "Edges": 8},
			TypeValueSet: "MapValues"},

		// "Edges",
		// ArrayValue(
		8: {ID: 8, MapValues: map[string]int{"0": 9, "1": 12},
			TypeValueSet: "MapValues"},

		// ArrayValue("state1 name1", "state1 name2")
		9: {ID: 9, MapValues: map[string]int{"0": 10, "1": 11},
			TypeValueSet: "MapValues"},
		10: {ID: 10, StringValue: "state1 name1",
			TypeValueSet: "StringValue"},
		11: {ID: 11, StringValue: "state1 name2",
			TypeValueSet: "StringValue"},

		// ArrayValue("state2 name1", "state2 name2"))
		12: {ID: 12, MapValues: map[string]int{"0": 13, "1": 14},
			TypeValueSet: "MapValues"},
		13: {ID: 13, StringValue: "state2 name1",
			TypeValueSet: "StringValue"},
		14: {ID: 14, StringValue: "state2 name2",
			TypeValueSet: "StringValue"},

		// "AreParallel", "true")
		15: {ID: 15, BoolValue: true,
			TypeValueSet: "BoolValue"},

		/////////

		/////////
		// "Next",
		// CollectMaps(
		16: {ID: 16, MapValues: map[string]int{"AreParallel": 24, "Edges": 17},
			TypeValueSet: "MapValues"},

		// "Edges",
		// ArrayValue(
		17: {ID: 17, MapValues: map[string]int{"0": 18, "1": 21},
			TypeValueSet: "MapValues"},

		// ArrayValue("state1 name1", "state1 name2")
		18: {ID: 18, MapValues: map[string]int{"0": 19, "1": 20},
			TypeValueSet: "MapValues"},
		19: {ID: 19, StringValue: "state1 name1",
			TypeValueSet: "StringValue"},
		20: {ID: 20, StringValue: "state1 name2",
			TypeValueSet: "StringValue"},

		// ArrayValue("state2 name1", "state2 name2"))
		21: {ID: 21, MapValues: map[string]int{"0": 22, "1": 23},
			TypeValueSet: "MapValues"},
		22: {ID: 22, StringValue: "state2 name1",
			TypeValueSet: "StringValue"},
		23: {ID: 23, StringValue: "state2 name2",
			TypeValueSet: "StringValue"},

		// "AreParallel", "true")
		24: {ID: 24, BoolValue: true,
			TypeValueSet: "BoolValue"},

		// "Values",
		// CollectMaps(
		25: {ID: 25, MapValues: map[string]int{
			"drinkOrder":   26,
			"orderQueue":   27,
			"outputBuffer": 28},
			TypeValueSet: "MapValues"},

		// "drinkOrder", ArrayValue()
		26: {ID: 26, MapValues: map[string]int{},
			TypeValueSet: "MapValues"},

		// "orderQueue", ArrayValue()
		27: {ID: 27, MapValues: map[string]int{},
			TypeValueSet: "MapValues"},

		// "outputBuffer", ArrayValue()
		28: {ID: 28, MapValues: map[string]int{},
			TypeValueSet: "MapValues"},

		// "LockedByStates",
		// CollectMaps(
		29: {ID: 29, MapValues: map[string]int{"11": 30},
			TypeValueSet: "MapValues"},

		// "11", "true"),
		30: {ID: 30, BoolValue: true,
			TypeValueSet: "BoolValue"},

		// "LockedByStatesCount", "1"
		31: {ID: 31, IntValue: 1,
			TypeValueSet: "IntValue"},
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
		myGraph := Graph{States: map[int]Atom{}}
		firstIDWant := 0

		firstIDGot := myGraph.AddState(ArrayValue("I am a test", "StarbucksMachine"))

		if firstIDWant != firstIDGot {
			t.Fatalf("wanted %v, got %v", firstIDWant, firstIDGot)
		}
	})

	t.Run("3 prior states before adding", func(t *testing.T) {
		myGraph := Graph{States: map[int]Atom{}}
		firstIDWant := 4

		firstGraph := ArrayValue("I am a test", "StarbucksMachine", "test")
		myGraph.AddState(firstGraph)
		firstIDGot := myGraph.AddState(ArrayValue("I am a test", "StarbucksMachine"))

		if firstIDWant != firstIDGot {
			t.Fatalf("wanted %v, got %v", firstIDWant, firstIDGot)
		}
	})
}

func TestGetAtom(t *testing.T) {
	myGraph := Graph{States: CollectMaps(
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

		id1, _ := myGraph.getAtom(0, []string{})

		want := 0
		got := id1
		if want != got {
			t.Fatalf("wanted %v, got %v", want, got)
		}
	})

	t.Run("path does not exist", func(t *testing.T) {

		id1, idError := myGraph.getAtom(0, []string{"not there"})

		want, errorWant := -1, errors.New("|not there|")
		got, errorGot := id1, idError

		if want != got || errorWant.Error() != errorGot.Error() {
			t.Fatalf("wanted %v%v, got %v%v", want, errorWant, got, errorGot)
		}
	})
	t.Run("path exists and has length 1", func(t *testing.T) {

		id1, idError := myGraph.getAtom(0, []string{"test"})

		want := 1
		got, errorGot := id1, idError

		if want != got || nil != errorGot {
			t.Fatalf("wanted %vnil, got %v%v", want, got, errorGot)
		}
	})
}
