#!/usr/bin/env bash
set -e

OLM_REPO=$1
OLM_BUNDLE_REPO=$2
TAG=$3
VERSION=$4

echo "Running with $1 $2 $3"
OLM_REPO_E=$(echo $OLM_REPO | sed -e 's/[\/&]/\\&/g')

echo "Calculator Versions"
VERSIONS=$(git tag | xargs)
# VERSIONS=$(ls deploy/olm-catalog/redhat-marketplace-operator | grep -E '^[[:digit:]]+\.[[:digit:]]+\.[[:digit:]]+$' | grep -Ev "^$VERSION$")
# VERSIONS=$(echo "$VERSIONS,$OLM_REPO:$TAG")

echo "Making version list"
VERSIONS_LIST=$(echo $VERSIONS | xargs | sed -e "s/ / $OLM_REPO_E:/g" | sed -e "s/^/$OLM_REPO_E:/g" | sed -e 's/ /,/g')

echo $VERSIONS_LIST

echo "Using tag ${OLM_BUNDLE_REPO}:${TAG}"
echo "Building index with $VERSIONS_LIST"
echo ""
./testbin/opm index add -u docker --generate --bundles "$VERSIONS_LIST" --tag "${OLM_BUNDLE_REPO}:${TAG}" --permissive
docker build -f custom-index.Dockerfile -t "${OLM_BUNDLE_REPO}:${TAG}" .
docker push "${OLM_BUNDLE_REPO}:${TAG}"
