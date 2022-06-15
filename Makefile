PROTO_FILES = $(shell find api -name *.proto)

.PHONY: api
api:
	protoc --go_out=. $(PROTO_FILES)

.PHONY: clean
clean:
	rm -rf api/*.go
