#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$PWD/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo "${GOPATH}"/src/k8s.io/code-generator)}

vendor/k8s.io/code-generator/generate-groups.sh all \
  github.com/kubernetes-incubator/external-storage/snapshot/pkg/client github.com/kubernetes-incubator/external-storage/snapshot/pkg/apis \
  volumesnapshot.external-storage.k8s.io:v1
