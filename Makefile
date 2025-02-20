.PHONY: buf-go
buf-go:
	go run github.com/bufbuild/buf/cmd/buf@v1.50 generate \
		--template proto/buf.gen.yaml

.PHONY: buf-python
buf-python:
	go run github.com/bufbuild/buf/cmd/buf@v1.50 generate \
		--template proto/buf.gen.python.yaml
