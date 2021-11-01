#!/usr/bin/env bash


# This script generates `*/api.pb.go` from the protobuf file `*/api.proto`.
# Example:
#   kube::protoc::generate_proto "${PREDICTION_V1}"

set -o errexit
set -o nounset
set -o pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../" && pwd -P)"
PREDICTION_V1="${ROOT}/prediction/v1alpha1"
ENSURANCE_V1="${ROOT}/ensurance/v1alpha1"

source "${ROOT}/hack/protoc.sh"
kube::protoc::generate_proto "${PREDICTION_V1}"
kube::protoc::generate_proto "${ENSURANCE_V1}"