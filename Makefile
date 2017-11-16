DIRS     ?= $(shell find . -name '*.go' | grep --invert-match 'vendor' | xargs -n 1 dirname | sort --unique)
PKGS     ?= $(shell go list ./... | grep --invert-match --regexp 'vendor')
PKG_NAME ?= lob

BFLAGS ?=
LFLAGS ?=
TFLAGS ?=

TS_METHOD ?=
TS_SUITE  ?=

LINT_CONF   ?= ./.lint.json

default: build

build:
	@echo "---> Building"
	go build -v -o ./bin/$(PKG_NAME) $(BFLAGS)

build_all:
	@echo "---> Building: (darwin, amd64)"
	GOOS=darwin GOARCH=amd64 go build -v -o ./bin/$(PKG_NAME)_osx $(BFLAGS)
	@echo "---> Building: (windows, amd64)"
	GOOS=windows GOARCH=amd64 go build -v -o ./bin/$(PKG_NAME)_win.exe $(BFLAGS)

lint:
	@echo "---> Linting... this might take a minute"
	gometalinter --config=$(LINT_CONF) --vendor --tests --deadline=2m $(LFLAGS) $(DIRS)

test:
	@echo "---> Running unit tests"
	go test $(PKGS) -cover -short $(TFLAGS) | tee ./coverage.out

test_integration:
	@echo "---> Running full test suite"
	go test $(PKGS) -cover $(TFLAGS) | tee ./coverage.out

clean:
	@echo "---> Cleaning"
	@rm -rf ./bin

install_tools:
	@echo "--> Installing tools"
	go get -u -v github.com/alecthomas/gometalinter
	gometalinter --install

uninstall_tools:
	@echo "--> Uninstalling tools"
	go clean -i github.com/alecthomas/gometalinter

.PHONY: build build_all lint test cover clean install_tools uninstall_tools
