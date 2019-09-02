.PHONY: gen
gen:
	prototool generate
	protoc -I. --go-grpc-error_out=../../.. example/proto/*.proto

.PHONY: install
install:
	go install ./protoc-gen-go-grpc-error

.PHONY: clean
clean:
	rm -rf ./gen
	rm -rf ./example/gen
