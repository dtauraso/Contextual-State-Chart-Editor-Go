// dir1.go
package dir1


var X = "test string"
type TestFunction func (string) string
type state struct {
	NamePart string
	State *state
    FunctionCode TestFunction
}


var StateTree = state{
	"name1", &state{"name2", nil, Test},
    Test,
}

func Test(input string) string {
    return input
}
func SayHello() string {
    return "Hi from package dir1"
}