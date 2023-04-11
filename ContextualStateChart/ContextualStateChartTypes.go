package ContextualStateChartTypes

type Graph struct {
}

type Value struct {
	IntValue    int    `json:"intValue,omitempty"`
	StringValue string `json:"stringValue,omitempty"`
}
type State struct {
	ID                  int              `json:"ID"`
	Name                []string         `json:"name"`
	FunctionCode        func(Graph) bool `json:"FunctionCode,omitempty"`
	EdgeKinds           map[string]Edges `json:"EdgeKinds,omitempty"`
	Parents             map[string]int   `json:"Parents,omitempty"`
	StartChildren       []int            `json:"StartChildren,omitempty"`
	HaveStartChildren   bool             `json:"HaveStartChildren,omitempty"`
	Value               Value            `json:"Value,omitempty"`
	ArrayValues         []int            `json:"arrayValues,omitempty"`
	MapValues           map[string]int   `json:"mapValues,omitempty"`
	LockedByStates      map[string]bool  `json:"LockedByStates,omitempty"`
	LockedByStatesCount int              `json:"LockedByStatesCount,omitempty"`
}
type Edges struct {
	Edges       [][]string
	AreParallel bool
}
type IDatabase struct {
	Array []any
	Map   map[string]any
}
