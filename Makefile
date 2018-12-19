# Lit all back service on the project
SERVICES=port

install-proto-validate:
	go get -d github.com/lyft/protoc-gen-validate

generate-pb:
	for service in $(SERVICES) ; do \
		make -C app/services/$$service gen-proto; \
	done