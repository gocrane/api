#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# This script holds common bash variables and utility functions.

# This function installs a Go tools by 'go get' command.
# Parameters:
#  - $1: package name, such as "sigs.k8s.io/controllers-tools/cmd/controllers-gen"
#  - $2: package version, such as "v0.4.1"
# Note:
#   Since 'go get' command will resolve and add dependencies to current module, that may update 'go.mod' and 'go.sum' file.
#   So we use a temporary directory to install the tools.
function util::install_tools() {
	local package="$1"
	local version="$2"

	temp_path=$(mktemp -d)
	pushd "${temp_path}" >/dev/null
	GO111MODULE=on go install "${package}"@"${version}"
	GOPATH=$(go env GOPATH | awk -F ':' '{print $1}')
	export PATH=$PATH:$GOPATH/bin
	popd >/dev/null
	rm -rf "${temp_path}"
}