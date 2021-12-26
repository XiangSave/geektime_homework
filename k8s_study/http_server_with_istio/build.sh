#!/bin/bash

BUILD_SERVERS=(http-server01)
VERSION="1.0"


function changeName() {
  pkgName=$1
  sed -i "s#{{ VAR_SERVER_NAME }}#${pkgName}#g" Dockerfile
}

function recoverName(){
  pkgName=$1
  sed -i "s#${pkgName}#{{ VAR_SERVER_NAME }}#g" Dockerfile
}

for name in ${BUILD_SERVERS[*]};do
  changeName "$name"
  docker build -t "${name}":${VERSION} .
  recoverName "$name"
done