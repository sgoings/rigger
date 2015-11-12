ifndef GOPATH
$(error No GOPATH set)
endif

MAIN_GO := rigger.go
BIN_NAME := bin/rigger

export GO15VENDOREXPERIMENT=1

build: $(MAIN_GO)
	go build -o $(BIN_NAME) -ldflags "-X main.version=${VERSION}" $<

resolve:
	glide up

test:
	./rigger stubbs:test --module rigger

test-%:
	./rigger stubbs:test --module rigger --plan $(subst test-,,$@)

.PHONY: test build resolve
