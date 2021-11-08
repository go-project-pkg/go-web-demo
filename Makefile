# Usage components ============================================================

define USAGE_OPTIONS

Options:

   BINS        The binaries to build. Default is all in cmd/.
               This option is available for: make build/build.multiarch
               Example: make build BINS="apiserver otherbin"
endef
export USAGE_OPTIONS

# Includes ====================================================================

include scripts/makefiles/share.makefile
include scripts/makefiles/go.makefile
include scripts/makefiles/tools.makefile

# Targets =====================================================================

# Build all by default.
.DEFAULT_GOAL := all

.PHONY: all
all: tidy build

.PHONY: tidy
tidy:
	@${GO} mod tidy

##  build: Compile packages and dependencies to generate bin file for current platform.
.PHONY: build
build:
	@${MAKE} go.build

##  build.multiarch: Build for multiple platforms. See option PLATFORMS.
.PHONY: build.multiarch
build.multiarch:
	@${MAKE} go.build.multiarch

##  lint: Check syntax and style of Go source code.
.PHONY: lint
lint:
	@${MAKE} go.lint

##  test: Run unit test.
.PHONY: test
test:
	@${MAKE} go.test

##  cover: Run unit test and get test coverage.
.PHONY: cover 
cover:
	@${MAKE} go.test.cover

##  clean: Remove all files that are created by building.
.PHONY: clean
clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf ${OUTPUT_DIR}

##  help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make [TARGETS] [OPTIONS] \n\nTargets:\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"