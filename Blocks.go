package main

import (
	ContextualStateChartTypes "Contextual-State-Chart-Editor-Go/ContextualStateChart"
	"time"
)

type Parent struct {
	Link             Link `json:"Link,omitempty"`
	PositionInParent int  `json:"PositionInParent,omitempty"`
}

type Change struct {
	Value ContextualStateChartTypes.Atom `json:"Value,omitempty"`
	Type  string                         `json:"Type,omitempty"`
}

type Variable struct {
	Value   ContextualStateChartTypes.Atom `json:"Value,omitempty"`
	History []Change                       `json:"History,omitempty"`
}

type Link struct {
	Ids                       []string  `json:"Ids,omitempty"`
	UsageCount                int       `json:"UsageCount,omitempty"`
	TimeLastUsed              time.Time `json:"TimeLastUsed,omitempty"`
	ActiveConnectionLastIndex int       `json:"ActiveConnectionLastIndex,omitempty"`
}
type Block struct {
	Id          string                                `json:"Id"`
	Function    func(map[string]Block, []string) bool `json:"Function,omitempty"`
	Parents     map[string]Parent                     `json:"Parents,omitempty"`
	Sequence    []Link                                `json:"Link,omitempty"`
	Variables   map[string]Variable                   `json:"Variables,omitempty"`
	NextContext map[string]string                     `json:"NextContext,omitempty"`
}
