#!/usr/bin/env bash

# Generic Go app image builder
# Gets its appname from the directory name

set -o errexit

if [[ -f .env ]]
then
  source .env
else
  echo "Could not find a .env file!"
  exit 1
fi

if [[ -z ${APPNAME} ]]
then 
  echo "APPNAME variable not found in .env file"
  exit 1
fi

if [[ -z ${GOPATH_SRC_DIR} ]]
then
  echo "GOPATH_SRC_DIR variable not found in .env file"
  exit 1
fi

appname="${APPNAME}"
srcdir="${GOPATH_SRC_DIR}"
appdir="/go/src/${srcdir}/${appname}"

builder="${appname}_builder"
builder_image="${appname}:builder"
builder_dockerfile="Dockerfile-builder"

app_image="${appname}:latest"
app_dockerfile="Dockerfile"

function cleanup() {
  echo "Cleaning up..."
  if $(docker ps -a | grep $builder)
  then
    docker rm $builder
  fi
}

# Always cleanup on exit
trap cleanup EXIT

function build_binary() {
  docker build --tag ${builder_image} \
    --file ${builder_dockerfile} .
  
  docker create --name ${builder} ${builder_image}
  
  mkdir -p ./pkg
  
  docker cp ${builder}:${bindir}/${appname} ./pkg
}

function build_image() {
  docker build --tag ${app_image} \
    --file ${app_dockerfile} .
}

populate_dockerfile_builder() {
  local builder_file="Dockerfile-builder"
  sed -i "s/^FROM.*$/FROM ${GOLANG_IMAGE}/" $builder_file
  sed -i "s/^LABEL maintainer=.*$/LABEL maintainer=\"${MAINTAINER}\"/" $builder_file
  sed -i "s/^ENV appname=.*$/ENV appname=${appname}/" $builder_file
  sed -i "s|^ENV appdir=.*$|ENV appdir=${appdir}|" $builder_file

  echo "Building Go binary with:"
  cat $builder_file
}

populate_dockerfile() {
  local docker_file="Dockerfile" 
  sed -i "s/^LABEL maintainer=.*$/LABEL maintainer=\"${MAINTAINER}\"/" $docker_file
  sed -i "s|^CMD.*$|CMD [ \"/${appname}\" ]|" $docker_file

  echo "Building image with:"
  cat $docker_file
}

main() {

  populate_dockerfile_builder
  build_binary

  populate_dockerfile
  build_image
}

main "$@"
