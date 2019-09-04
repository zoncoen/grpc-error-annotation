.PHONY: gen
gen:
	prototool generate

.PHONY: install
install:
	go install ./protoc-gen-go-grpc-error

.PHONY: example
example:
	protoc -I. --go-grpc-error_out=../../.. example/proto/*.proto

.PHONY: clean
clean:
	rm -rf ./gen
	rm -rf ./example/gen
