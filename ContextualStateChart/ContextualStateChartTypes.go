package ContextualStateChartTypes

type Graph struct {
}

type Value struct {
	BoolValue   bool   `json:"BoolValue,omitempty"`
	IntValue    int    `json:"IntValue,omitempty"`
	StringValue string `json:"StringValue,omitempty"`
	TypeName    string `json:"TypeName"`
}

type Edges struct {
	Edges       []int
	AreParallel bool
}

// Parents: NDParentStateName -> ID
// MapValues: 1D string -> ID
/*
local variable
1D name -> primitive value
1D name -> array of ID's
1D name -> map of string keys -> ID's

database variable
ND name -> primitive value
ND name -> array of ID's
ND name -> map of string keys -> ID's
*/
type State struct {
	ID                  int              `json:"ID"`
	Name                []string         `json:"Name"`
	FunctionCode        string           `json:"FunctionCode,omitempty"`
	EdgeKinds           map[string]Edges `json:"EdgeKinds,omitempty"`
	Parents             map[string]int   `json:"Parents,omitempty"`
	Value               Value            `json:"Value,omitempty"`
	ArrayValues         []int            `json:"ArrayValues,omitempty"`
	MapValues           map[string]any   `json:"MapValues,omitempty"`
	LockedByStates      map[string]bool  `json:"LockedByStates,omitempty"`
	LockedByStatesCount int              `json:"LockedByStatesCount,omitempty"`
}
type State2 struct {
	ID           int                    `json:"ID"`
	Name         []string               `json:"Name"`
	FunctionCode func(StateID int) bool `json:"FunctionCode,omitempty"`
	BoolValue    bool                   `json:"BoolValue,omitempty"`
	IntValue     int                    `json:"IntValue,omitempty"`
	StringValue  string                 `json:"StringValue,omitempty"`
	ArrayValues  []int                  `json:"ArrayValues,omitempty"`
	MapValues    map[string]int         `json:"MapValues,omitempty"`
}
