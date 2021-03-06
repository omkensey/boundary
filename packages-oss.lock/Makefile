# ***
# WARNING: Do not EDIT or MERGE this file, it is generated by packagespec.
# ***
include $(shell git rev-parse --show-toplevel)/packages*.lock/config.mk

.PHONY: packages commands build package write-builder-cache-keys \
	write-all-package-cache-keys build-all

GOOS ?= $(shell go env GOOS 2>/dev/null || echo linux)
GOARCH ?= $(shell go env GOARCH 2>/dev/null || echo amd64)

DEFAULT_PACKAGE_YQ := [ .packages[] | select(.inputs.GOOS=="$(GOOS)" and .inputs.GOARCH=="$(GOARCH)") ][0]
QUERY_DEFAULT_PACKAGESPEC = $(call QUERY_LOCK,$(DEFAULT_PACKAGE_YQ) | $(1))

# MK is shorthand for changing to repo root and selecting a make target
# from a file in this directory. All *.mk files in this directory assume
# the current working directory is the repo root. This Makefile exists
# to invoke those mk files correctly.
MK := $(MAKE) -C $(REPO_ROOT) -f $(LOCKDIR)/

# configure load-builder-cache target to load the most specific builder layer cache
# available as an archive in the build cache.

ifneq ($(PACKAGE_SPEC_ID),)

PACKAGE_CACHE_KEY_FILE := $(shell $(call QUERY_PACKAGESPEC,.meta.builtin.PACKAGE_CACHE_KEY_FILE))

# Loading the best available archive for a specific package build.

BUILD_LAYER_ARCHIVES := $(shell $(call QUERY_PACKAGESPEC,.meta.builtin.BUILD_LAYERS[].archive))
BEST_BUILD_LAYER_ARCHIVE := $(shell cd $(REPO_ROOT) && for F in $(BUILD_LAYER_ARCHIVES); do \
	if [ -f $$F ]; then echo $$F; exit 0; fi; done)
ifeq ($(BEST_BUILD_LAYER_ARCHIVE),)
load-builder-cache:
	@echo "No build layer archives found in build cache. Looked for: $(BUILD_LAYER_ARCHIVES)"
else
BEST_BUILD_LAYER_NAME := $(shell $(call QUERY_PACKAGESPEC,.meta.builtin.BUILD_LAYERS[] \
	| select(.archive=="$(BEST_BUILD_LAYER_ARCHIVE)") | .name))
BEST_BUILD_LAYER_LOAD_TARGET := $(BEST_BUILD_LAYER_NAME)-load

load-builder-cache:
	@$(MK)layer.mk $(BEST_BUILD_LAYER_LOAD_TARGET)

endif

else ifneq ($(LAYER_SPEC_ID),)

# Loading the best avilable archive for a specific layer build.

BUILD_LAYER_ARCHIVES := $(shell $(call QUERY_LOCK,.layers[] | select(.name=="$(LAYER_SPEC_ID)") \
	| .meta.builtin.LAYER_LIST[].archive))
BEST_BUILD_LAYER_ARCHIVE := $(shell cd $(REPO_ROOT) && for F in $(BUILD_LAYER_ARCHIVES); do \
	if [ -f $$F ]; then echo $$F; exit 0; fi; done)
ifeq ($(BEST_BUILD_LAYER_ARCHIVE),)

load-builder-cache:
	@echo "No build layer archives found in build cache. Looked for: $(BUILD_LAYER_ARCHIVES)"

else
BEST_BUILD_LAYER_NAME := $(shell $(call QUERY_LOCK,.layers[] | select(.name=="$(LAYER_SPEC_ID)") \
	| .meta.builtin.LAYER_LIST[] | select(.archive=="$(BEST_BUILD_LAYER_ARCHIVE)") | .name))
BEST_BUILD_LAYER_LOAD_TARGET := $(BEST_BUILD_LAYER_NAME)-load

load-builder-cache:
	@$(MK)layer.mk $(BEST_BUILD_LAYER_LOAD_TARGET)

endif

else

load-builder-cache:
	@echo "You must set PACKAGE_SPEC_ID or LAYER_SPEC_ID so we know which caches to look for."

endif

commands:
	@$(MAKE) -f packages.mk commands

ifeq ($(DIRTY_FILES),)
DIRTY_SOURCE_WARNING :=
else
DIRTY_SOURCE_WARNING = echo "==> SOURCE TREE IS DIRTY; $(1)"
endif

# build is a convenience target for local builds, do not use in CI.
# Instead, use `make package` specifying PACKAGE_SPEC_ID.
build:
	@$(call DIRTY_SOURCE_WARNING,PERFORMING DIRTY BUILD)
	@echo "==> Building default package for GOOS=$(GOOS) GOARCH=$(GOARCH)"
	@ALIASES=$$($(call QUERY_DEFAULT_PACKAGESPEC,.aliases[] | "alias type:\(.type) path:\(.path)") | column -t); \
		echo "$$ALIASES"
	@PACKAGE_SPEC_ID="$$($(call QUERY_DEFAULT_PACKAGESPEC,.packagespecid) | head -n1)"; \
		COMMAND="PACKAGE_SOURCE_ID=$$PACKAGE_SOURCE_ID PACKAGE_SPEC_ID=$$PACKAGE_SPEC_ID $(MK)build.mk package"; \
		echo "$$COMMAND"; \
		$(SHELL) "$$COMMAND"

# package-contents is a convenience target for local builds, do not use in CI.
package-contents:
	@$(call DIRTY_SOURCE_WARNING,GETTING CONTENTS OF DIRTY BUILD)
	@echo "==> Getting contents of default package for GOOS=$(GOOS) GOARCH=$(GOARCH)"
	@ALIASES=$$($(call QUERY_DEFAULT_PACKAGESPEC,.aliases[] | "alias type:\(.type) path:\(.path)") | column -t); \
		echo "$$ALIASES"
	@PACKAGE_SPEC_ID="$$($(call QUERY_DEFAULT_PACKAGESPEC,.packagespecid) | head -n1)"; \
		COMMAND="PACKAGE_SOURCE_ID=$$PACKAGE_SOURCE_ID PACKAGE_SPEC_ID=$$PACKAGE_SPEC_ID $(MK)build.mk package-contents"; \
		echo "$$COMMAND"; \
		$(SHELL) "$$COMMAND"

# copy-package-contents is a convenience target for local builds, do not use in CI.
copy-package-contents:
	@$(call DIRTY_SOURCE_WARNING,COPYING CONTENTS OF DIRTY BUILD)
	@echo "==> Getting contents of default package for GOOS=$(GOOS) GOARCH=$(GOARCH)"
	@ALIASES=$$($(call QUERY_DEFAULT_PACKAGESPEC,.aliases[] | "alias type:\(.type) path:\(.path)") | column -t); \
		echo "$$ALIASES"
	@PACKAGE_SPEC_ID="$$($(call QUERY_DEFAULT_PACKAGESPEC,.packagespecid) | head -n1)"; \
		COMMAND="PACKAGE_SOURCE_ID=$$PACKAGE_SOURCE_ID PACKAGE_SPEC_ID=$$PACKAGE_SPEC_ID $(MK)build.mk copy-package-contents"; \
		echo "$$COMMAND"; \
		$(SHELL) "$$COMMAND"

# meta is a convenience target for local builds, do not use in CI.
# Instead, use `make package-meta` specifying PACKAGE_SPEC_ID.
meta:
	@$(call DIRTY_SOURCE_WARNING,WRITING METADATA FOR DIRTY BUILD)
	@echo "==> Writing metacdata for default package (GOOS=$(GOOS) GOARCH=$(GOARCH))"
	@ALIASES=$$($(call QUERY_DEFAULT_PACKAGESPEC,.aliases[] | "alias type:\(.type) path:\(.path)") | column -t); \
		echo "$$ALIASES"
	@PACKAGE_SPEC_ID="$$($(call QUERY_DEFAULT_PACKAGESPEC,.packagespecid) | head -n1)"; \
		COMMAND="PACKAGE_SOURCE_ID=$$PACKAGE_SOURCE_ID PACKAGE_SPEC_ID=$$PACKAGE_SPEC_ID $(MK)build.mk package-meta"; \
		echo "$$COMMAND"; \
		$(SHELL) "$$COMMAND"

# build-all is a convenience target to sequentially build each package.
# It is mostly useful in the tutorial, do not use this in CI as it is much slower
# than building packages in parallel.
build-all:
	@PACKAGE_SPEC_IDS="$$($(call QUERY_LOCK,.packages[] | .packagespecid))"; \
		COUNT=$$(echo $$PACKAGE_SPEC_IDS | wc -w | xargs); \
		echo "==> Building all $$COUNT packages sequentially."; \
		for PACKAGE_SPEC_ID in $$PACKAGE_SPEC_IDS; do \
			COMMAND="PACKAGE_SOURCE_ID=$$PACKAGE_SOURCE_ID PACKAGE_SPEC_ID=$$PACKAGE_SPEC_ID $(MK)build.mk package"; \
			echo "$$COMMAND"; \
			$(SHELL) "$$COMMAND"; \
		done

# package expects PACKAGE_SPEC_ID to already be set, use this in CI.
package:
	@$(call DIRTY_SOURCE_WARNING,BUILDING DIRTY PACKAGE)
	@echo "==> Building package spec $(PACKAGE_SPEC_ID)"
	@ALIASES=$$($(call QUERY_PACKAGESPEC,.aliases[] | "alias type:\(.type) path:\(.path)") | column -t); \
		echo "$$ALIASES"
	@COMMAND="PACKAGE_SOURCE_ID=$$PACKAGE_SOURCE_ID PACKAGE_SPEC_ID=$$PACKAGE_SPEC_ID $(MK)build.mk package"; \
		echo "$$COMMAND"; \
		$(SHELL) "$$COMMAND"

# package-meta expects PACKAGE_SPEC_ID to already be set, use this in CI.
package-meta:
	@$(call DIRTY_SOURCE_WARNING,WRITING DIRTY METADATA FOR DIRTY PACKAGE)
	@echo "==> Writing metadata for package $(PACKAGE_SPEC_ID)"
	@ALIASES=$$($(call QUERY_PACKAGESPEC,.aliases[] | "alias type:\(.type) path:\(.path)") | column -t); \
		echo "$$ALIASES"
	@COMMAND="PACKAGE_SOURCE_ID=$$PACKAGE_SOURCE_ID PACKAGE_SPEC_ID=$$PACKAGE_SPEC_ID $(MK)build.mk package-meta"; \
		echo "$$COMMAND"; \
		$(SHELL) "$$COMMAND"

# package-meta expects PACKAGE_SPEC_ID to already be set, use this in CI.
package-meta-all:
	@$(call DIRTY_SOURCE_WARNING,WRITING DIRTY METADATA FOR DIRTY PACKAGES)
	@PACKAGE_SPEC_IDS="$$($(call QUERY_LOCK,.packages[] | .packagespecid))"; \
		COUNT=$$(echo $$PACKAGE_SPEC_IDS | wc -w | xargs); \
		echo "==> Writing $$COUNT packages' metadata..."; \
		for PACKAGE_SPEC_ID in $$PACKAGE_SPEC_IDS; do \
			export PACKAGE_SPEC_ID; \
			FILE="$(PACKAGE_SOURCE_ID)-$${PACKAGE_SPEC_ID}.zip.meta.json"; \
			OUT="$(PACKAGE_STORE)/$$FILE"; \
			COMMAND="$(call QUERY_PACKAGESPEC_BY_ID,env.PACKAGE_SPEC_ID,.) > $$OUT"; \
			echo "$$COMMAND"; \
			$(SHELL) "$$COMMAND"; \
		done

# aliases writes all alias symlinks for packages in the package store that
# match the current LOCKFILE and PRODUCT_REVISION. It does not cause a new build.
# If the package store contains no matchin binaries, then this does nothing.
aliases:
	@echo "==> Writing alias symlinks for existing packages in the store."; \
	cd $(REPO_ROOT); \
	PACKAGE_SPEC_IDS="$$($(call QUERY_LOCK,.packages[] | .packagespecid))"; \
		for PACKAGE_SPEC_ID in $$PACKAGE_SPEC_IDS; do \
			PACKAGE_FILE="$$PACKAGE_SOURCE_ID-$$PACKAGE_SPEC_ID.zip"; \
			PACKAGE="$(CACHE_ROOT)/packages/store/$$PACKAGE_FILE"; \
			[ -f $$PACKAGE ] || continue; \
			ALIASES=$$($(call QUERY_PACKAGESPEC_BY_ID,'$$PACKAGE_SPEC_ID',.aliases[] | "$(CACHE_ROOT)/packages/by-alias/\(.type)/\(.path)")); \
			for A in $$ALIASES; do \
				mkdir -p $$(dirname $$A); \
				$(LN) -rfs $$PACKAGE $$A; \
				echo "==> Alias written: $$A -> $$PACKAGE"; \
			done; \
		done


write-builder-cache-keys:
	@echo "==> Writing build layer cache keys"
	@$(MK)layer.mk write-cache-keys

write-package-cache-key:
	@if [ -z "$(PACKAGE_CACHE_KEY_FILE)" ]; then echo "Must set PACKAGE_SPEC_ID"; exit 1; fi
	@$(WRITE_PACKAGE_CACHE_KEY)
	@echo "==> Package cache key written: $(PACKAGE_CACHE_KEY_FILE)"

# WRITE_PACKAGE_CACHE_KEY writes the package cache key for PACKAGE_SPEC_ID.
# We reference this as an environment variable, so you can override it in a
# recipe rather than relying on the global setting.
define WRITE_PACKAGE_CACHE_KEY
( \
	cd $(REPO_ROOT); \
	KEY="PACKAGE_SPEC_ID=$$PACKAGE_SPEC_ID"$$'\n'"PACKAGE_SOURCE_ID=$(PACKAGE_SOURCE_ID)"; \
	FILE=$$(yq -r ".packages[] | select(.packagespecid==\"$$PACKAGE_SPEC_ID\") \
		| .meta.builtin.PACKAGE_CACHE_KEY_FILE" < $(LOCK)); \
	echo "$$FILE"; \
	echo "$$KEY"; \
	mkdir -p $$(dirname $$FILE); \
	echo "$$KEY" > "$$FILE";\
)
endef

write-all-package-cache-keys:
	@IDS="$$($(call QUERY_LOCK,.packages[].packagespecid))"; \
	for PACKAGE_SPEC_ID in $$IDS; do \
		$(WRITE_PACKAGE_CACHE_KEY); \
	done; \
	echo "==> All package cache keys written"

clean-builder-images:
	@IMAGES=$$(docker images --format '{{.Repository}}:{{.Tag}}' | grep '^$(BUILDER_IMAGE_PREFIX)' || true); \
		if [ -z "$$IMAGES" ]; then exit 0; fi; \
		docker rmi -f $$IMAGES

clean:
	@cd $(REPO_ROOT); rm -rf $(CACHE_ROOT)

clean-all: clean clean-builder-images

clean-all-prune: clean-all
	docker container prune
	docker image prune

RELEASER_DIR   := $(REPO_ROOT)/.packagespec/release

# REQUIRE_EXPORT requires a set of make variables to be nonempty,
# exits 1 if any are not, and exports each one otherwise.
# To be used in recipe bodies.
define REQUIRE_EXPORT
$(foreach VAR,$(1),[ -n "$($(VAR))" ] || { echo "Must set $(VAR)"; exit 1; }; export $(VAR)='$($(VAR))';)
endef

# EXPORT exports each named variable, if it exists.
define EXPORT
$(foreach VAR,$(1),export $(VAR)='$($(VAR))';)
endef

# INVOKE_RELEASER_TARGET invokes the named target (first arg) in the releaser
# repository, first calling REQUIRE_EXPORT on all the named variables (second arg).
define INVOKE_RELEASER_TARGET
	$(call REQUIRE_EXPORT,\
		PRODUCT_REPO_LOCAL    PRODUCT_REPO          PRODUCT_PATH \
		PRODUCT_CIRCLECI_SLUG PRODUCT_CIRCLECI_HOST RELEASE_SYSTEM_BRANCH \
		PRODUCT_RELEASE_REPO  SPEC                  LOCKDIR \
	) \
	( cd $(REPO_ROOT) && packagespec load -asset=PREP_TIME -lockdir "$(LOCKDIR)"; ); \
	( cd $(REPO_ROOT) && packagespec load -asset=WORK_DIR  -lockdir "$(LOCKDIR)"; ); \
	$(MAKE) -C $(RELEASER_DIR) $(1)
endef

# RELEASE_TARGETS are targets in the release repo we pass control to
# to perform release actions.
# Note: The release repo is only available to HashiCorp employees.
RELEASE_TARGETS := build-ci stage-config stage custom-build custom-build-config orchestrator stop-orchestrator bundle

# We always rev-parse the PRODUCT_REVISION to obtain the full SHA. This is required
# for downstream processes which use it to determine part of the package name.
$(RELEASE_TARGETS): PRODUCT_REVISION := $(shell git rev-parse $${PRODUCT_REVISION:-HEAD})
$(RELEASE_TARGETS): PRODUCT_VERSION ?= 0.0.0-$(USER)-snapshot
$(RELEASE_TARGETS): RELEASE_SYSTEM_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
custom-build: PRODUCT_VERSION := $(PRODUCT_VERSION)-$(PRODUCT_REVISION)
bundle:            PRODUCT_VERSION := $(shell $(call QUERY_LOCK,.packages[0].inputs.PRODUCT_VERSION))
orchestrator:      PRODUCT_VERSION := $(shell $(call QUERY_LOCK,.packages[0].inputs.PRODUCT_VERSION))
stop-orchestrator: PRODUCT_VERSION := $(shell $(call QUERY_LOCK,.packages[0].inputs.PRODUCT_VERSION))
$(RELEASE_TARGETS):
	@\
		echo $(PRODUCT_VERSION) \
		$(call REQUIRE_EXPORT,PRODUCT_REVISION PRODUCT_VERSION) \
		$(call INVOKE_RELEASER_TARGET,$@)

# QUERY_TARGETS are targets in the release repo that perform queries, and are therefore
# not necessarily bound to a specific PRODUCT_VERSION or PRODUCT_REVISION.
# We still export PRODUCT_VERSION and PRODUCT_REVISION because they can be used as query
# parameters.
QUERY_TARGETS := list-staged-builds list-promoted-builds list-custom-builds watch-ci

$(QUERY_TARGETS): RELEASE_SYSTEM_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
$(QUERY_TARGETS):
	@\
		$(call EXPORT,PRODUCT_REVISION PRODUCT_VERSION) \
		$(call INVOKE_RELEASER_TARGET,$@)

# BUNDLE_TARGETS are targets acting on specific staged bundles, identified by
# their BUNDLE_ID.
BUNDLE_TARGETS := publish-config publish inspect-staged-build workflow

$(BUNDLE_TARGETS): RELEASE_SYSTEM_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
$(BUNDLE_TARGETS):
	@\
		$(call REQUIRE_EXPORT,BUNDLE_ID) \
		$(call INVOKE_RELEASER_TARGET,$@)
