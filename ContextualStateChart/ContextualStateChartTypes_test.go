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

func TestState(t *testing.T) {
	want := map[int]State{
		0: {
			ID: 0,
			MapValues: map[string]int{
				"Name":         1,
				"FunctionCode": 5},
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
	}

	got :=
		CollectMaps(
			MapValue("Name", ArrayValueStrings("I am a test", "StarbucksMachine")),
			MapValueString("FunctionCode", "ReturnTrue"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

// func TestState2(t *testing.T) {
// 	want := nil

// 	got :=
// 		CollectMaps(
// 			MapValue("test",
// 				CollectMaps(
// 					MapValue("Name", ArrayValueStrings("I am a test", "StarbucksMachine")),
// 					MapValueString("FunctionCode", "ReturnTrue"),
// 					MapValue("Name2", ArrayValueStrings("I am a test", "StarbucksMachine")),
// 					MapValueString("FunctionCode2", "ReturnTrue"))),
// 			MapValue("test",
// 				CollectMaps(
// 					MapValue("Name", ArrayValueStrings("I am a test", "StarbucksMachine")),
// 					MapValueString("FunctionCode", "ReturnTrue"),
// 					MapValue("Name2", ArrayValueStrings("I am a test", "StarbucksMachine")),
// 					MapValueString("FunctionCode2", "ReturnTrue"))),
// 		)

// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("wanted %v, got %v", want, got)
// 	}
// }
