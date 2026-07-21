package util

import (
	"sync"
)

//go:generate go run ../_tools emb-structs -i ../_tools/html5entities.json -o ./html5entities.gen.go

var _html5entitiesOnce sync.Once
var _html5entitiesMap map[string]*HTML5Entity

func buildHTML5Entities() { _ = "STUB: not implemented"; return }

type HTML5Entity struct {
	Name       string
	Characters []byte
}

func LookUpHTML5EntityByName(name string) (*HTML5Entity, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
