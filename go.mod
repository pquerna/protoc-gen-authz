module github.com/pquerna/protoc-gen-authz

go 1.24

require (
	github.com/dave/jennifer v1.7.1
	github.com/golang/protobuf v1.5.4
	github.com/lyft/protoc-gen-star/v2 v2.0.4
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/spf13/afero v1.14.0 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	golang.org/x/tools v0.32.0 // indirect
)

// https://github.com/lyft/protoc-gen-star/pull/132
replace github.com/lyft/protoc-gen-star/v2 => github.com/pquerna/protoc-gen-star/v2 v2.0.0-20250415201647-653a078eb414
