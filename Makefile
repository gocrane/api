GOOS ?= $(shell go env GOOS)
SOURCES := $(shell find . -type f  -name '*.go')

# Git information
GIT_VERSION ?= $(shell git describe --tags --dirty)
GIT_COMMIT_HASH ?= $(shell git rev-parse HEAD)
GIT_TREESTATE = "clean"
GIT_DIFF = $(shell git diff --quiet >/dev/null 2>&1; if [ $$? -eq 1 ]; then echo "1"; fi)
ifeq ($(GIT_DIFF), 1)
    GIT_TREESTATE = "dirty"
endif
BUILDDATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS := "-X github.com/gocrane/api/pkg/version.gitVersion=$(GIT_VERSION) \
                      -X github.com/gocrane/api/pkg/version.gitCommit=$(GIT_COMMIT_HASH) \
                      -X github.com/gocrane/api/pkg/version.gitTreeState=$(GIT_TREESTATE) \
                      -X github.com/gocrane/api/pkg/version.buildDate=$(BUILDDATE)"

# Images management
REGISTRY?="ccr.ccs.tencentyun.com/kube-orm"
REGISTRY_USER_NAME?=""
REGISTRY_PASSWORD?=""
REGISTRY_SERVER_ADDRESS?="ccr.ccs.tencentyun.com"

# Set your version by env or using latest tags from git
VERSION?=""
ifeq ($(VERSION), "")
    LATEST_TAG=$(shell git describe --tags)
    ifeq ($(LATEST_TAG),)
        # Forked repo may not sync tags from upstream, so give it a default tag to make CI happy.
        VERSION="unknown"
    else
        VERSION=$(LATEST_TAG)
    endif
endif


clean:
	rm -rf output

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./prediction/... ./ensurance/... ./recommendation/...

# Run go vet against code
vet:
	go vet ./pkg/... ./prediction/... ./ensurance/... ./recommendation/...

test:
	go test --race --v ./pkg/...


.PHONY: update
update: fmt vet
	hack/update-all.sh

.PHONY: verify
verify:
	hack/verify-all.sh

.PHONY: staticcheck
staticcheck:
	hack/verify-staticcheck.sh

.PHONY: all
all:
	hack/update-all.sh && hack/verify-all.sh
