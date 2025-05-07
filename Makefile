MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables

.PHONY: install
install:
	go install -mod=vendor -v .

generate:
	protoc \
	-I . \
	--go_out="paths=source_relative:." \
	authz/*.proto

example:
	DEBUG_PG_AUTHZ=true protoc example.proto --proto_path=. --proto_path=examplepb --go_out="paths=source_relative:examplepb" --authz_out="lang=go,paths=source_relative:examplepb"

.PHONY: adddep
adddep:
	go mod tidy -v
	go mod vendor

.PHONY: updatedeps
updatedeps:
	go get -d -u ./...
	go mod tidy -v
	go mod vendor
