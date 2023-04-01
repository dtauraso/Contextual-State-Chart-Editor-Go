package StarbucksTree

import (
	csc "github.com/dtauraso/Contextual-State-Chart-Editor-Go/ContextualStateChart"
	u "github.com/dtauraso/Contextual-State-Chart-Editor-Go/Utility"
)

type IStateNamePartTree struct {
	NPT   map[string]IStateNamePartTree
	State IState
}

type IState struct {
	FunctionCode        func(csc.Graph) bool
	EdgeKinds           map[string]IEdges
	Children            map[string]IStateNamePartTree
	HaveStartChildren   bool
	Variables           map[string]any
	LockedByStates      map[string]bool
	LockedByStatesCount int
	Array               []any
	Map                 map[string]any
}

type IEdges struct {
	Edges       [][]string
	AreParallel bool
}

var Customer = IStateNamePartTree{
	NPT: map[string]IStateNamePartTree{
		"Cashier": {
			State: IState{
				FunctionCode: u.ReturnTrue,
				EdgeKinds: map[string]IEdges{
					"StartChildren": {
						Edges:       [][]string{{"Place order"}},
						AreParallel: false,
					},
				},
				HaveStartChildren: true,
				Children: map[string]IStateNamePartTree{
					"Place order": {
						State: IState{
							FunctionCode: u.ReturnTrue,
							EdgeKinds: map[string]IEdges{
								"Next": {
									Edges: [][]string{
										{"Dig up money"},
										{"Sip coffee"},
									},
									AreParallel: true,
								},
							},
							HaveStartChildren: false,
						},
					},
					"Dig up money": {
						State: IState{
							FunctionCode: u.ReturnTrue,
							EdgeKinds: map[string]IEdges{
								"Next": {
									Edges: [][]string{
										{"Put away change"},
									},
									AreParallel: true,
								},
							},
							HaveStartChildren:   false,
							LockedByStates:      map[string]bool{"Compute Price": true},
							LockedByStatesCount: 1,
						},
					},
					"Put away change": {
						State: IState{
							FunctionCode:        u.ReturnTrue,
							LockedByStates:      map[string]bool{"Compute change": true},
							LockedByStatesCount: 1,
						},
					},
					"Sip coffee": {
						State: IState{
							FunctionCode:        u.ReturnTrue,
							LockedByStates:      map[string]bool{"Output buffer": true},
							LockedByStatesCount: 1,
						},
					},
				},
				Variables: map[string]any{"drink": "frap choco"},
			}},
		"Barista": {
			State: IState{
				FunctionCode: u.ReturnTrue,
			},
		},
	},
}

var Cashier = IStateNamePartTree{
	State: IState{
		FunctionCode: u.ReturnTrue,
		EdgeKinds: map[string]IEdges{
			"StartChildren": {
				Edges: [][]string{
					{"Take order", "from customer"},
				},
				AreParallel: true,
			},
		},
		HaveStartChildren: true,
		Children: map[string]IStateNamePartTree{
			"Take order": {
				NPT: map[string]IStateNamePartTree{
					"from customer": {
						State: IState{
							FunctionCode: u.ReturnTrue,
							EdgeKinds: map[string]IEdges{
								"Next": {
									Edges: [][]string{
										{"Compute Price"},
									},
									AreParallel: false,
								},
							},
							HaveStartChildren:   false,
							LockedByStates:      map[string]bool{"Place order": true},
							LockedByStatesCount: 1,
						},
					}},
			},
			"Compute Price": {
				State: IState{
					FunctionCode: u.ReturnTrue,
					EdgeKinds: map[string]IEdges{
						"Next": {
							Edges: [][]string{
								{"Compute change"},
							},
							AreParallel: true,
						},
					},
					HaveStartChildren: false,
				},
			},
		},
	},
}
var StateTree = Customer

func Test(input string) string {
	return input
}
func SayHello() string {
	return "Hi from package dir1"
}
