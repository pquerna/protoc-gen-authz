package pgaz

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type Module struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

var _ pgs.Module = (*Module)(nil)

func New() pgs.Module {
	return &Module{ModuleBase: &pgs.ModuleBase{}}
}
