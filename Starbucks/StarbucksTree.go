package StarbucksTree


type IStateNamePart struct {
	NamePart string
	State IState
}

type IState struct {
    FunctionCode func (any) bool
    EdgeKinds map[string]IEdges
    Children []IStateNamePart
    HaveStartChildren bool
    Variables map[string]any
}

type IEdges struct {
    Edges [][]string
    AreParallel bool
}

var Customer = IStateNamePart{
    "Cashier", IState{FunctionCode: nil, Variables: map[string]any{"test": 1} },
}

var StateTree = Customer

func Test(input string) string {
    return input
}
func SayHello() string {
    return "Hi from package dir1"
}