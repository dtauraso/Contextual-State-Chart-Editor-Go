package TrieTree

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type GotWant struct {
	got  InputParameters
	want string
}

func TestTrieTree(t *testing.T) {

	fmt.Println("Test Insert")
	insertTestParameters := []GotWant{
		{got: InputParameters{Name: []string{"test"}, StateID: 0}, want: "[{\"NextNamePartTree\":{\"test\":1},\"StateID\":0},{\"NextNamePartTree\":null,\"StateID\":-1}]"},
		{got: InputParameters{Name: []string{"test", "test2"}, StateID: 1}, want: "[{\"NextNamePartTree\":{\"test\":1},\"StateID\":0},{\"NextNamePartTree\":{\"test2\":2},\"StateID\":-1},{\"NextNamePartTree\":null,\"StateID\":1}]"},
		{got: InputParameters{Name: []string{"test", "test2", "test3"}, StateID: 2}, want: "[{\"NextNamePartTree\":{\"test\":1},\"StateID\":0},{\"NextNamePartTree\":{\"test2\":2},\"StateID\":-1},{\"NextNamePartTree\":{\"test3\":3},\"StateID\":1},{\"NextNamePartTree\":null,\"StateID\":2}]"},
		{got: InputParameters{Name: []string{"testx", "test2", "test3"}, StateID: 3}, want: "[{\"NextNamePartTree\":{\"test\":1,\"testx\":4},\"StateID\":0},{\"NextNamePartTree\":{\"test2\":2},\"StateID\":-1},{\"NextNamePartTree\":{\"test3\":3},\"StateID\":1},{\"NextNamePartTree\":null,\"StateID\":2},{\"NextNamePartTree\":{\"test2\":5},\"StateID\":-1},{\"NextNamePartTree\":{\"test3\":6},\"StateID\":-1},{\"NextNamePartTree\":null,\"StateID\":3}]"},
	}
	namesTrie := TrieTree{}

	for i := 0; i < len(insertTestParameters); i++ {

		name := insertTestParameters[i].got.Name
		stateID := insertTestParameters[i].got.StateID

		want := insertTestParameters[i].want

		t.Run(fmt.Sprintf("Insert [%s] -> %d", strings.Join(name, " "), stateID), func(t *testing.T) {

			namesTrie = namesTrie.Insert(InputParameters{Name: name, StateID: stateID})
			got, _ := json.Marshal(namesTrie)

			assertCorrectMessage(t, string(got), want)

		})
	}

	fmt.Println("Test Search")
	searchTestParameters := []GotWant{
		{got: InputParameters{Name: []string{"test"}, StateID: 0}, want: "0"},
		{got: InputParameters{Name: []string{"test", "test2"}, StateID: 1}, want: "1"},
		{got: InputParameters{Name: []string{"test", "test2", "test3"}, StateID: 2}, want: "2"},
		{got: InputParameters{Name: []string{"testx", "test2", "test3"}, StateID: 3}, want: "3"},
	}

	for i := 0; i < len(searchTestParameters); i++ {

		name := searchTestParameters[i].got.Name
		stateID := searchTestParameters[i].got.StateID

		want := searchTestParameters[i].want

		t.Run(fmt.Sprintf("Search [%s] -> %d", strings.Join(name, " "), stateID), func(t *testing.T) {

			namesTrie = namesTrie.Search(InputParameters{Name: name, StateID: stateID})
			got, _ := json.Marshal(namesTrie)

			assertCorrectMessage(t, string(got), want)

		})
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
