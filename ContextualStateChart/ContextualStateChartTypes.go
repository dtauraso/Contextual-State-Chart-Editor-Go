package ContextualStateChartTypes

type Graph struct {
}

type Value struct {
	IntValue    int    `json:"intValue",omitvalue`
	StringValue string `json:"stringValue",omitvalue`
}
type State struct {
	ID                  int
	name                []string
	FunctionCode        func(Graph) bool
	EdgeKinds           map[string]Edges
	Parents             map[string]int
	StartChildren       []int
	HaveStartChildren   bool
	Value               Value
	arrayValues         []int          `json:"arrayValues",omitvalue`
	mapValues           map[string]int `json:"arrayValues",omitvalue`
	LockedByStates      map[string]bool
	LockedByStatesCount int
	// Database            map[string]IStateNamePartTree
}
type Edges struct {
	Edges       [][]string
	AreParallel bool
}
type IDatabase struct {
	Array []any
	Map   map[string]any
}
