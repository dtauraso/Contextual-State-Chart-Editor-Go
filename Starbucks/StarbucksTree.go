package StarbucksTree

import (
    csc "github.com/dtauraso/Contextual-State-Chart-Editor-Go/ContextualStateChart"
	u "github.com/dtauraso/Contextual-State-Chart-Editor-Go/Utility"

)

type IStateNamePart struct {
	NamePart string
    NextNamePart *IStateNamePart
	State IState
}

type IState struct {
    FunctionCode func (csc.Graph) bool
    EdgeKinds map[string]IEdges
    Children map[string]IStateNamePart
    HaveStartChildren bool
    Variables map[string]any
    LockedByStates map[string] bool
    LockedByStatesCount int
}

type IEdges struct {
    Edges [][]string
    AreParallel bool
}

var Customer = IStateNamePart{
    NamePart: "Cashier",
        State: IState{
            FunctionCode: u.ReturnTrue,
            EdgeKinds: map[string]IEdges{
                "StartChildren": IEdges{
                    Edges: [][]string{[]string{"Place order"}},
                    AreParallel: false,
                },
            },
            HaveStartChildren: true,
            Children: map[string]IStateNamePart{
                "Place order": IStateNamePart{
                    State: IState{
                        FunctionCode: u.ReturnTrue,
                        EdgeKinds: map[string]IEdges{
                            "Next": IEdges{
                                Edges: [][]string{
                                        []string{"Dig up money"},
                                        []string{"Sip coffee"},
                                        },
                                AreParallel: true,
                            },
                        },
                        HaveStartChildren: false,
                    },
                },
                "Dig up money": IStateNamePart{
                    State: IState{
                        FunctionCode: u.ReturnTrue,
                        EdgeKinds: map[string]IEdges{
                            "Next": IEdges{
                                Edges: [][]string{
                                        []string{"Put away change"},
                                },
                                AreParallel: true,
                            },
                        },
                        HaveStartChildren: false,
                        LockedByStates: map[string]bool{"Compute Price": true},
                        LockedByStatesCount: 1,
                    },
                },
            },
            // Variables: map[string]any{"test": 1}
        },
}

var StateTree = Customer

func Test(input string) string {
    return input
}
func SayHello() string {
    return "Hi from package dir1"
}