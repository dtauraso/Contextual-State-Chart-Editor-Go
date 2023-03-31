package StarbucksTree


var X = "test string"
type TestFunction func (string) string
type TestFunction2 func (any) bool
type IStateNamePart struct {
	NamePart string
	State *IState
}

type IState struct {
    FunctionCode TestFunction2
}

type IEdgeGroups struct {

}
type IEdges struct {
    Edges []IEdge
    AreParallel bool
}

type IEdge struct {
    NextStateName []string
}

var Customer = IStateNamePart{
    "Cashier", &IState{FunctionCode: nil },
}

// var StateTree = State{
// 	"name1", &State{"name2", nil},
// }

func Test(input string) string {
    return input
}
func SayHello() string {
    return "Hi from package dir1"
}