
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
all: wasmbuild npm generate $(WASM)

NPM_CARBON_OUTFILE=$(if $(NPM_CARBON_DIST_DIR),$(abspath $(NPM_CARBON_DIST_DIR))/bundle.js,bundle.js)
NPM_AUTH_OUTFILE=$(if $(NPM_AUTH_DIST_DIR),$(abspath $(NPM_AUTH_DIST_DIR))/auth.js,auth.js)

# Generate icon names from the npm bundle.
wasm/carbon-app/content/icon_names.go: npm/carbon/bundle.js
	@echo 'Generating icon names'
	@cd wasm/carbon-app/content && $(GO) generate

.PHONY: generate
generate: wasm/carbon-app/content/icon_names.go

# Rules for building
.PHONY: $(WASM)
$(WASM): wasmbuild generate
	@echo 'Building $@ with ${GOCC}'
	@$(BUILDDIR)/wasmbuild build --go=${GOCC} --go-flags='-ldflags "$(LD_FLAGS)"' -o ${BUILDDIR}/$(shell basename $@).wasm ./$@

.PHONY: npm
npm: npm/carbon npm/auth

.PHONY: npm/carbon
npm/carbon: npm/carbon/index.js npm/carbon/package.json
	@echo 'Building npm/carbon bundle'
	@cd npm/carbon && install -d "$(dir $(NPM_CARBON_OUTFILE))" && npm install && OUTFILE='$(NPM_CARBON_OUTFILE)' npm run build && if [ -n "$(NPM_CARBON_DIST_DIR)" ]; then rm -f bundle.js icons-generated.js assets/themes.css assets/grid.css; rm -rf assets/icons; fi

.PHONY: npm/auth
npm/auth: npm/auth/auth.ts npm/auth/package.json
	@echo 'Building npm/auth bundle'
	@cd npm/auth && install -d "$(dir $(NPM_AUTH_OUTFILE))" && npm install && OUTFILE='$(NPM_AUTH_OUTFILE)' npm run build && if [ -n "$(NPM_AUTH_DIST_DIR)" ]; then rm -f auth.js; fi

npm/carbon/bundle.js: npm/carbon/index.js npm/carbon/package.json
	@echo 'Building npm/carbon bundle'
	@cd npm/carbon && npm install && npm run build
	@if [ -n "$(NPM_CARBON_DIST_DIR)" ]; then \
		echo 'Copying npm/carbon assets to $(NPM_CARBON_DIST_DIR)'; \
		install -d "$(NPM_CARBON_DIST_DIR)"; \
		cp npm/carbon/bundle.js "$(NPM_CARBON_DIST_DIR)/bundle.js"; \
		rm -rf "$(NPM_CARBON_DIST_DIR)/assets"; \
		cp -R npm/carbon/assets "$(NPM_CARBON_DIST_DIR)/assets"; \
	fi

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
	@rm -f npm/carbon/bundle.js
	@rm -f npm/auth/auth.js
	@rm -f wasm/carbon-app/content/icon_names.go
	$(GO) clean
