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
  "autoscaling:v1alpha1 prediction:v1alpha1 ensurance:v1alpha1" \
  --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
  --go-header-file "${SCRIPT_ROOT}/hack/boilerplate/boilerplate.go.txt"

KUBE_OPENAPI_PKG=`go list -mod=readonly -m -f '{{.Dir}}' k8s.io/kube-openapi`

go run ${KUBE_OPENAPI_PKG}/cmd/openapi-gen/openapi-gen.go --logtostderr \
	    -i k8s.io/metrics/pkg/apis/custom_metrics,k8s.io/metrics/pkg/apis/custom_metrics/v1beta1,k8s.io/metrics/pkg/apis/custom_metrics/v1beta2,k8s.io/metrics/pkg/apis/external_metrics,k8s.io/metrics/pkg/apis/external_metrics/v1beta1,k8s.io/metrics/pkg/apis/metrics,k8s.io/metrics/pkg/apis/metrics/v1beta1,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/version,k8s.io/api/core/v1 \
	    --build-tag autogenerated \
	    -h ./hack/boilerplate/boilerplate.go.txt \
	    -p ./pkg/generated/openapi \
	    -O zz_generated.openapi \
	    -o ./ \
	    -r /dev/null