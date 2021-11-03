#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# For all commands, the working directory is the parent directory(repo root).
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "${REPO_ROOT}"

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(
  cd "${SCRIPT_ROOT}"
  go mod vendor
  ls -d -1 ./vendor/k8s.io/code-generator
)}

bash "${CODEGEN_PKG}/generate-groups.sh" all \
  github.com/gocrane-io/api/pkg/generated \
  github.com/gocrane-io/api \
  "autoscaling:v1alpha1 prediction:v1alpha1" \
  --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
  --go-header-file "${SCRIPT_ROOT}/hack/boilerplate/boilerplate.go.txt"