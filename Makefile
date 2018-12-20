# Lit all back service on the project
SERVICES=port

build:
	go build -o port-service entry/entry.go

install-proto-validate:
	go get -d github.com/lyft/protoc-gen-validate

install-linter:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.12.3

run-linter:
	golangci-lint run -v

generate-pb:
	for service in $(SERVICES) ; do \
		make -C app/services/$$service gen-proto; \
	done