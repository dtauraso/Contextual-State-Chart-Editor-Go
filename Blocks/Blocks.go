package block

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

/*func(map[string]Block, []string) bool*/
type Block struct {
	Id          string              `json:"Id"`
	Parents     []Parent            `json:"Parents,omitempty"`
	Sequence    []Link              `json:"Link,omitempty"`
	Variables   map[string]Variable `json:"Variables,omitempty"`
	NestedBlock map[string]string   `json:"NestedBlock,omitempty"`
}

type Blocks struct {
	Blocks map[string]Block `json:"Blocks,omitempty"`
	MaxInt int              `json:"MaxInt,omitempty"`
}
