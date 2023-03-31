package StarbucksTree


type IStateNamePart struct {
	NamePart string
	State *IState
}

type IState struct {
    FunctionCode func (*string) bool
    EdgeKinds map[string]IEdges
}

type IEdges struct {
    Edges [][]string
    AreParallel bool
}


var Customer = IStateNamePart{
    "Cashier", &IState{FunctionCode: nil },
}

var StateTree = Customer

func Test(input string) string {
    return input
}
func SayHello() string {
    return "Hi from package dir1"
}