#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

CONTROLLER_GEN_PKG="sigs.k8s.io/controller-tools/cmd/controller-gen"
CONTROLLER_GEN_VER="v0.4.1"

source hack/util.sh

echo "Generating with controller-gen"
util::install_tools ${CONTROLLER_GEN_PKG} ${CONTROLLER_GEN_VER} >/dev/null 2>&1
controller-gen crd paths=./prediction/... output:crd:dir=./artifacts/deploy
controller-gen crd paths=./autoscaling/... output:crd:dir=./artifacts/deploy
controller-gen crd paths=./ensurance/... output:crd:dir=./artifacts/deploy

controller-gen webhook paths=./prediction/... output:webhook:dir=./artifacts/deploy
controller-gen webhook paths=./autoscaling/... output:webhook:dir=./artifacts/deploy
controller-gen webhook paths=./ensurance/... output:webhook:dir=./artifacts/deploy

