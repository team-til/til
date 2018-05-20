#!/bin/bash

PROJ_ROOT="$(dirname "$(dirname "$(readlink "$0")")")"

mkdir -p ${PROJ_ROOT}/server/_proto
mkdir -p ${PROJ_ROOT}/frontend/_proto
mkdir -p ${PROJ_ROOT}/grpcgateway/_proto

protoc \
  --plugin=protoc-gen-ts=${PROJ_ROOT}/frontend/node_modules/.bin/protoc-gen-ts \
  -I ${PROJ_ROOT}/proto \
  -I $GOPATH/src \
  -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --js_out=import_style=commonjs,binary:${PROJ_ROOT}/frontend/_proto \
  --go_out=plugins=grpc:${PROJ_ROOT}/server/_proto \
  --grpc-gateway_out=logtostderr=true:${PROJ_ROOT}/server/_proto \
  --ts_out=service=true:${PROJ_ROOT}/frontend/_proto \
${PROJ_ROOT}/proto/til_service.proto
 
 protoc \
  -I ${PROJ_ROOT}/proto \
  -I $GOPATH/src \
  -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:${PROJ_ROOT}/grpcgateway/_proto \
  --grpc-gateway_out=logtostderr=true:${PROJ_ROOT}/grpcgateway/_proto \
${PROJ_ROOT}/proto/til_service.proto