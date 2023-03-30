// dir1.go
package dir1


var X = "test string"
type state struct {
	NamePart string
	State *state
}


var StateTree = state{
	"name1", &state{"name2", nil},
}

func SayHello() string {
    return "Hi from package dir1"
}