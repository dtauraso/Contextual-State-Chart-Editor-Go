package main

import ContextualStateChartTypes "Contextual-State-Chart-Editor-Go/ContextualStateChart"

type Parent struct {
	IdList           []string `json:"IdList,omitempty"`
	PositionInParent int      `json:"PositionInParent,omitempty"`
}

type Variable struct {
	Value   ContextualStateChartTypes.Atom   `json:"Value,omitempty"`
	History []ContextualStateChartTypes.Atom `json:"History,omitempty"`
}

type Block struct {
	Id          string              `json:"Id"`
	Parents     []Parent            `json:"Parents,omitempty"`
	Sequence    [][]string          `json:"Sequence,omitempty"`
	Variables   map[string]Variable `json:"Variables,omitempty"`
	NextContext map[string]int      `json:"NextContext,omitempty"`
}
