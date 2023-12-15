package main

import ContextualStateChartTypes "Contextual-State-Chart-Editor-Go/ContextualStateChart"

type Parent struct {
	Link             Link `json:"Link,omitempty"`
	PositionInParent int  `json:"PositionInParent,omitempty"`
}

type Variable struct {
	Value   ContextualStateChartTypes.Atom   `json:"Value,omitempty"`
	History []ContextualStateChartTypes.Atom `json:"History,omitempty"`
}

type Link struct {
	Ids        []string `json:"Ids,omitempty"`
	UsageCount int      `json:"UsageCount,omitempty"`
}
type Block struct {
	Id          string                                `json:"Id"`
	Function    func(map[string]Block, []string) bool `json:"Function,omitempty"`
	Parents     []Parent                              `json:"Parents,omitempty"`
	Sequence    []Link                                `json:"Link,omitempty"`
	Variables   map[string]Variable                   `json:"Variables,omitempty"`
	NextContext map[string]string                     `json:"NextContext,omitempty"`
}
