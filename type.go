package skiplist

import "math/rand"

type elementNode struct {
	next []*Element
}

type Element struct {
	elementNode
	key   uint64
	value interface{}
}

// Key allows retrieval of the key for a given Element
func (e *Element) Key() uint64 {
	return e.key
}

// Value allows retrieval of the value for a given Element
func (e *Element) Value() interface{} {
	return e.value
}

type SkipList struct {
	elementNode
	maxLevel       int
	length         int
	randSource     rand.Source
	probability    float64
	probTable      []float64
	prevNodesCache []*elementNode
}
