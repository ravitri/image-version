#!/bin/bash
## Script to fetch version from a Openshift Release Image
## Usage: ./image-version.sh ${IMAGE_DIGEST}
## Example: ./image-version.sh quay.io/openshift-release-dev/ocp-release@sha256:dd75546170e65d7d17130de10a6ffeb425f960399640632cbc8426b9da338458

IMAGE_DIGEST=$1

if [[ -z $IMAGE_DIGEST ]]
then
	echo "Usage $0 IMAGE_DIGEST"
	exit 1
fi

SHA=$(echo $IMAGE_DIGEST | awk -F":" '{ print $2 }')

CONFIG_DIGEST=$(curl -s -H"Accept: application/vnd.docker.distribution.manifest.v2+json" "https://quay.io/v2/openshift-release-dev/ocp-release/manifests/sha256:$SHA" | jq -r .config.digest)

VERSION=$(curl -sL "https://quay.io/v2/openshift-release-dev/ocp-release/blobs/${CONFIG_DIGEST}" | jq '.config.Labels["io.openshift.release"]')

echo "Version is $VERSION"
