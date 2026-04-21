
# Go parameters
GO=$(shell which go)
BUILDDIR=build
NPM_CARBON_DIST_DIR ?= npm/carbon
NPM_CARBON_BUNDLE := $(abspath $(NPM_CARBON_DIST_DIR))/bundle.js
WASM=$(wildcard wasm/*)
GOROOT=$(shell go env GOROOT)
GOCC ?= go

# Build flags
BUILD_MODULE = $(shell cat go.mod | head -1 | cut -d ' ' -f 2)
LD_FLAGS = -X $(BUILD_MODULE)/pkg/version.GitSource=${BUILD_MODULE} -X $(BUILD_MODULE)/pkg/version.GitTag=$(shell git describe --tags --always) -X $(BUILD_MODULE)/pkg/version.GitBranch=$(shell git name-rev HEAD --name-only --always) -X $(BUILD_MODULE)/pkg/version.GitHash=$(shell git rev-parse HEAD) -X $(BUILD_MODULE)/pkg/version.GoBuildTime=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

# All targets
all: wasmbuild npm generate $(WASM)

NPM_CARBON_OUTFILE=$(if $(NPM_CARBON_DIST_DIR),$(abspath $(NPM_CARBON_DIST_DIR))/bundle.js,bundle.js)
NPM_AUTH_OUTDIR=$(if $(NPM_AUTH_DIST_DIR),$(abspath $(NPM_AUTH_DIST_DIR)),.)

# Generate icon names from the npm bundle.
wasm/carbon-app/content/icon_names.go: $(NPM_CARBON_BUNDLE)
	@echo 'Generating icon names'
	@cd wasm/carbon-app/content && $(GO) generate

.PHONY: generate
generate: wasm/carbon-app/content/icon_names.go

# Rules for building
.PHONY: $(WASM)
$(WASM): wasmbuild generate
	@echo 'Building $@ with ${GOCC}'
	@$(BUILDDIR)/wasmbuild build --go=${GOCC} --go-flags='-ldflags "$(LD_FLAGS)"' -o ${BUILDDIR}/$(shell basename $@).wasm ./$@

.PHONY: npm npm/carbon npm/auth
npm: npm/carbon npm/auth

npm/carbon: $(NPM_CARBON_BUNDLE)

$(NPM_CARBON_BUNDLE): npm/carbon/index.js npm/carbon/package.json npm/carbon/gen-icons.mjs
	@echo 'Building npm/carbon bundle'
	@cd npm/carbon && npm install && CARBON_DIST_DIR='$(abspath $(NPM_CARBON_DIST_DIR))' npm run build

npm/auth: npm/auth/auth.ts npm/auth/token.ts npm/auth/package.json
	@echo 'Building npm/auth bundle'
	@cd npm/auth && install -d "$(NPM_AUTH_OUTDIR)" && npm install && OUTDIR='$(NPM_AUTH_OUTDIR)' npm run build

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
	@if [ '$(abspath $(NPM_CARBON_DIST_DIR))' = '$(abspath npm/carbon)' ]; then \
		rm -f npm/carbon/bundle.js; \
		rm -f npm/carbon/assets/themes.css; \
		rm -f npm/carbon/assets/grid.css; \
		rm -fr npm/carbon/assets/icons; \
	else \
		rm -fr '$(abspath $(NPM_CARBON_DIST_DIR))'; \
	fi
	@rm -f npm/carbon/icons-generated.js
	@rm -f wasm/carbon-app/content/icon_names.go
	$(GO) clean
