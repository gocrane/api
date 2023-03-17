#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

CONTROLLER_GEN_PKG="sigs.k8s.io/controller-tools/cmd/controller-gen"
CONTROLLER_GEN_VER="v0.7.0"

source hack/util.sh

echo "Generating with controller-gen"
util::install_tools ${CONTROLLER_GEN_PKG} ${CONTROLLER_GEN_VER} >/dev/null 2>&1
controller-gen crd paths=./prediction/... output:crd:dir=./artifacts/deploy
controller-gen crd paths=./autoscaling/... output:crd:dir=./artifacts/deploy
controller-gen crd paths=./ensurance/... output:crd:dir=./artifacts/deploy
controller-gen crd paths=./analysis/... output:crd:dir=./artifacts/deploy
controller-gen crd paths=./topology/... output:crd:dir=./artifacts/deploy
controller-gen crd paths=./co2e/... output:crd:dir=./artifacts/deploy

controller-gen webhook paths=./prediction/... output:webhook:dir=./artifacts/deploy
controller-gen webhook paths=./autoscaling/... output:webhook:dir=./artifacts/deploy
controller-gen webhook paths=./ensurance/... output:webhook:dir=./artifacts/deploy
controller-gen webhook paths=./analysis/... output:webhook:dir=./artifacts/deploy
controller-gen webhook paths=./topology/... output:webhook:dir=./artifacts/deploy
controller-gen webhook paths=./co2e/... output:webhook:dir=./artifacts/deploy

