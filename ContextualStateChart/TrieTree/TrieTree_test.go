package TrieTree

import (
	"fmt"
	s "strconv"
	"strings"
	"testing"
)

func TestTrieTree(t *testing.T) {

	testParameters := []InsertNameParameters{
		InsertNameParameters{Name: []string{"test"}, StateID: 0},
		InsertNameParameters{Name: []string{"test", "test2"}, StateID: 1},
		InsertNameParameters{Name: []string{"test", "test2", "test3"}, StateID: 2},
		InsertNameParameters{Name: []string{"testx", "test2", "test3"}, StateID: 3},
	}
	for i := 0; i < len(testParameters); i++ {

		name := testParameters[i].Name
		stateID := testParameters[i].StateID
		t.Run(fmt.Sprintf("insert [%s] -> %d", strings.Join(name, " "), stateID), func(t *testing.T) {
			namesTrie := TrieTree{}
			got := s.Itoa(namesTrie.InsertName(InsertNameParameters{Name: name, StateID: stateID}))
			want := s.Itoa(stateID)
			assertCorrectMessage(t, got, want)

		})
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
