
# Go parameters
GO=go
BUILDDIR = build
WASM = $(wildcard wasm/*)
GOROOT = $(shell go env GOROOT)
CC ?= go

# Build flags
BUILD_MODULE = $(shell cat go.mod | head -1 | cut -d ' ' -f 2)
BUILD_LD_FLAGS += -X $(BUILD_MODULE)/pkg/version.GitSource=${BUILD_MODULE}
BUILD_LD_FLAGS += -X $(BUILD_MODULE)/pkg/version.GitTag=$(shell git describe --tags --always)
BUILD_LD_FLAGS += -X $(BUILD_MODULE)/pkg/version.GitBranch=$(shell git name-rev HEAD --name-only --always)
BUILD_LD_FLAGS += -X $(BUILD_MODULE)/pkg/version.GitHash=$(shell git rev-parse HEAD)
BUILD_LD_FLAGS += -X $(BUILD_MODULE)/pkg/version.GoBuildTime=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
BUILD_FLAGS = -ldflags "-s -w ${BUILD_LD_FLAGS}"

# WASM build flags (uses main package for version info)
WASM_LD_FLAGS = -X main.GitSource=${BUILD_MODULE} -X main.GitTag=$(shell git describe --tags --always) -X main.GitBranch=$(shell git name-rev HEAD --name-only --always) -X main.GitHash=$(shell git rev-parse HEAD) -X main.GoBuildTime=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

# All targets
all: wasmbuild $(WASM)

# Rules for building
.PHONY: $(WASM)
$(WASM): mkdir
	@echo 'Building $@ with ${CC}'
	@$(BUILDDIR)/wasmbuild build --go=${CC} --go-flags='-ldflags "${WASM_LD_FLAGS}"' -o ${BUILDDIR}/$(shell basename $@).wasm ./$@

.PHONY: wasmbuild
wasmbuild: mkdir
	@echo 'Building wasmbuild'
	@${GO} build ${BUILD_FLAGS} -o ${BUILDDIR}/wasmbuild ./cmd/wasmbuild

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
