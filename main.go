package main

import (
	"github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/pquerna/protoc-gen-authz/internal/pgaz"
)

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG_PG_AUTHZ")).
		RegisterModule(pgaz.New()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
