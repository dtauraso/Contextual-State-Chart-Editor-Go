package StarbucksTree

import (
    csc "github.com/dtauraso/Contextual-State-Chart-Editor-Go/ContextualStateChart"
)

type IStateNamePart struct {
	NamePart string
	State IState
}

type IState struct {
    FunctionCode func (csc.Graph) bool
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