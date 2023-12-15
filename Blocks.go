package main

import ContextualStateChartTypes "Contextual-State-Chart-Editor-Go/ContextualStateChart"

type Parent struct {
	IdList           []int `json:"IdList,omitempty"`
	PositionInParent int   `json:"PositionInParent,omitempty"`
}

type Variable struct {
	Value   ContextualStateChartTypes.Atom   `json:"Value,omitempty"`
	History []ContextualStateChartTypes.Atom `json:"History,omitempty"`
}

type Block struct {
	Id          int                 `json:"Id"`
	Name        string              `json:"Name,omitempty"`
	Parents     map[int]Parent      `json:"Parents,omitempty"`
	Sequence    [][]int             `json:"Sequence,omitempty"`
	Variables   map[string]Variable `json:"Variables,omitempty"`
	NextContext map[string]int      `json:"NextContext,omitempty"`
}
