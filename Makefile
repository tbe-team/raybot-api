.PHONY: tool-install
tool-install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: clean
clean:
	git clean -fxd

.PHONY: buf-go
buf-go:
	go run github.com/bufbuild/buf/cmd/buf@v1.54.0 generate

.PHONY: buf-lint
buf-lint:
	go run github.com/bufbuild/buf/cmd/buf@v1.54.0 lint

.PHONY: buf-format
buf-format:
	go run github.com/bufbuild/buf/cmd/buf@v1.54.0 format -w

.PHONY: golangci-lint
golangci-lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8 \
		run --config .golangci.yml

.PHONY: lint
lint: buf-lint buf-format golangci-lint
