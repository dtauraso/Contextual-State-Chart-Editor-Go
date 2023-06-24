package TrieTree

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

type InsertGotWant struct {
	got  InputParameters
	want string
}

type SearchGotWant struct {
	got  []string
	want string
}

func TestTrieTree(t *testing.T) {

	fmt.Println("Test Insert")
	insertTestParameters := []InsertGotWant{
		{got: InputParameters{Name: []string{"test"}, StateId: 0}, want: "[{\"NextNamePartTree\":{\"test\":1},\"StateId\":-1},{\"NextNamePartTree\":null,\"StateId\":0}]"},
		{got: InputParameters{Name: []string{"test", "test2"}, StateId: 1}, want: "[{\"NextNamePartTree\":{\"test\":1},\"StateId\":-1},{\"NextNamePartTree\":{\"test2\":2},\"StateId\":0},{\"NextNamePartTree\":null,\"StateId\":1}]"},
		{got: InputParameters{Name: []string{"test", "test2", "test3"}, StateId: 2}, want: "[{\"NextNamePartTree\":{\"test\":1},\"StateId\":-1},{\"NextNamePartTree\":{\"test2\":2},\"StateId\":0},{\"NextNamePartTree\":{\"test3\":3},\"StateId\":1},{\"NextNamePartTree\":null,\"StateId\":2}]"},
		{got: InputParameters{Name: []string{"testx", "test2", "test3"}, StateId: 3}, want: "[{\"NextNamePartTree\":{\"test\":1,\"testx\":4},\"StateId\":-1},{\"NextNamePartTree\":{\"test2\":2},\"StateId\":0},{\"NextNamePartTree\":{\"test3\":3},\"StateId\":1},{\"NextNamePartTree\":null,\"StateId\":2},{\"NextNamePartTree\":{\"test2\":5},\"StateId\":-1},{\"NextNamePartTree\":{\"test3\":6},\"StateId\":-1},{\"NextNamePartTree\":null,\"StateId\":3}]"},
	}
	namesTrie := TrieTree{}

	for i := 0; i < len(insertTestParameters); i++ {

		name := insertTestParameters[i].got.Name
		stateId := insertTestParameters[i].got.StateId

		want := insertTestParameters[i].want

		t.Run(fmt.Sprintf("Insert [%s] -> %d", strings.Join(name, " "), stateId), func(t *testing.T) {

			namesTrie = namesTrie.Insert(InputParameters{Name: name, StateId: stateId})
			got, _ := json.Marshal(namesTrie)

			assertCorrectMessage(t, string(got), want)

		})
	}

	fmt.Println("Test Search")
	searchTestParameters := []SearchGotWant{
		{got: []string{"test"}, want: "0"},
		{got: []string{"test", "test2"}, want: "1"},
		{got: []string{"test", "test2", "test3"}, want: "2"},
		{got: []string{"testx", "test2", "test3"}, want: "3"},
	}

	for i := 0; i < len(searchTestParameters); i++ {

		got := searchTestParameters[i].got

		want := searchTestParameters[i].want

		t.Run(fmt.Sprintf("Search [%s] -> %s", strings.Join(got, " "), want), func(t *testing.T) {

			stateId := namesTrie.Search(got)
			gotMarshaled, _ := json.Marshal(stateId)
			assertCorrectMessage(t, string(gotMarshaled), want)

		})
	}

	fmt.Println("Test Search Input")
	searchInputTestParameters := []SearchGotWant{
		{got: []string{"test"}, want: "true"},
		{got: []string{"test", "test"}, want: "false"},
		{got: []string{"test", "test2"}, want: "true"},
		{got: []string{"test", "test2", "test3"}, want: "true"},
		{got: []string{"testx", "test2", "test3"}, want: "true"},
		{got: []string{"testx", "test1", "test3"}, want: "false"},
	}

	for i := 0; i < len(searchInputTestParameters); i++ {

		got := searchInputTestParameters[i].got

		want := searchInputTestParameters[i].want

		t.Run(fmt.Sprintf("Search [%s] -> %s", strings.Join(got, " "), want), func(t *testing.T) {

			isThere := namesTrie.SearchInput(got)
			gotMarshaled, _ := json.Marshal(isThere)
			assertCorrectMessage(t, string(gotMarshaled), want)

		})
	}
	err1 := os.WriteFile("output.txt", []byte("test"), 0644)
	if err1 != nil {
		panic(err1)
	}

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
