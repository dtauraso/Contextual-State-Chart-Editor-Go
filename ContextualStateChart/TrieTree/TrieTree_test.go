package TrieTree

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type GotWant struct {
	got  InsertNameParameters
	want string
}

func TestTrieTree(t *testing.T) {

	testParameters := []GotWant{
		{got: InsertNameParameters{Name: []string{"test"}, StateID: 0}, want: "[{\"NextNamePartTree\":{\"test\":1},\"StateID\":0},{\"NextNamePartTree\":null,\"StateID\":-1}]"},
		{got: InsertNameParameters{Name: []string{"test", "test2"}, StateID: 1}, want: "[{\"NextNamePartTree\":{\"test\":1},\"StateID\":0},{\"NextNamePartTree\":{\"test2\":2},\"StateID\":-1},{\"NextNamePartTree\":null,\"StateID\":1}]"},
		{got: InsertNameParameters{Name: []string{"test", "test2", "test3"}, StateID: 2}, want: "[{\"NextNamePartTree\":{\"test\":1},\"StateID\":0},{\"NextNamePartTree\":{\"test2\":2},\"StateID\":-1},{\"NextNamePartTree\":{\"test3\":3},\"StateID\":1},{\"NextNamePartTree\":null,\"StateID\":2}]"},
		{got: InsertNameParameters{Name: []string{"testx", "test2", "test3"}, StateID: 3}, want: "[{\"NextNamePartTree\":{\"test\":1,\"testx\":4},\"StateID\":0},{\"NextNamePartTree\":{\"test2\":2},\"StateID\":-1},{\"NextNamePartTree\":{\"test3\":3},\"StateID\":1},{\"NextNamePartTree\":null,\"StateID\":2},{\"NextNamePartTree\":{\"test2\":5},\"StateID\":-1},{\"NextNamePartTree\":{\"test3\":6},\"StateID\":-1},{\"NextNamePartTree\":null,\"StateID\":3}]"},
	}
	namesTrie := TrieTree{}

	for i := 0; i < len(testParameters); i++ {

		name := testParameters[i].got.Name
		stateID := testParameters[i].got.StateID

		want := testParameters[i].want

		t.Run(fmt.Sprintf("insert [%s] -> %d", strings.Join(name, " "), stateID), func(t *testing.T) {

			namesTrie = namesTrie.InsertName(InsertNameParameters{Name: name, StateID: stateID})
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
