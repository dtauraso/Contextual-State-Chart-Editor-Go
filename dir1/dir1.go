// dir1.go
package dir1


var X = "test string"
type TestFunction func (string) string
type State struct {
	NamePart string
	State *State
    FunctionCode TestFunction
}


var StateTree = State{
	"name1", &State{"name2", nil, Test},
    Test,
}

func Test(input string) string {
    return input
}
func SayHello() string {
    return "Hi from package dir1"
}