
# Go parameters
GO=$(shell which go)
BUILDDIR=build
WASM=$(wildcard wasm/*)
GOROOT=$(shell go env GOROOT)
GOCC ?= go

# Build flags
BUILD_MODULE = $(shell cat go.mod | head -1 | cut -d ' ' -f 2)
LD_FLAGS = -X $(BUILD_MODULE)/pkg/version.GitSource=${BUILD_MODULE} -X $(BUILD_MODULE)/pkg/version.GitTag=$(shell git describe --tags --always) -X $(BUILD_MODULE)/pkg/version.GitBranch=$(shell git name-rev HEAD --name-only --always) -X $(BUILD_MODULE)/pkg/version.GitHash=$(shell git rev-parse HEAD) -X $(BUILD_MODULE)/pkg/version.GoBuildTime=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

# All targets
all: wasmbuild $(WASM)

# Rules for building
.PHONY: $(WASM)
$(WASM): wasmbuild
	@echo 'Building $@ with ${GOCC}'
	@$(BUILDDIR)/wasmbuild build --go=${GOCC} --go-flags='-ldflags "$(LD_FLAGS)"' -o ${BUILDDIR}/$(shell basename $@).wasm ./$@

.PHONY: wasmbuild
wasmbuild: mkdir
	@echo 'Building wasmbuild'
	@${GO} build -ldflags "$(LD_FLAGS)" -o ${BUILDDIR}/wasmbuild ./cmd/wasmbuild

.PHONY: test
test: tidy
	@$(GO) test -v ./pkg/js
	@$(GO) test -v ./pkg/dom
	@$(GO) test -v ./pkg/mvc
#	@$(GO) test -v ./pkg/bootstrap

.PHONY: jstest
jstest: tidy
	@$(GO) install github.com/agnivade/wasmbrowsertest@latest
	@GOOS=js GOARCH=wasm $(GO) test -v -exec="wasmbrowsertest" ./pkg/js
	@GOOS=js GOARCH=wasm $(GO) test -v -exec="wasmbrowsertest" ./pkg/dom
	@GOOS=js GOARCH=wasm $(GO) test -v -exec="wasmbrowsertest" ./pkg/mvc
#	@GOOS=js GOARCH=wasm $(GO) test -v -exec="wasmbrowsertest" ./pkg/bootstrap

.PHONY: mkdir
mkdir:
	@install -d $(BUILDDIR)

.PHONY: tidy
tidy: 
	$(GO) mod tidy

.PHONY: clean
clean: tidy
	@rm -fr $(BUILDDIR)
	$(GO) clean
