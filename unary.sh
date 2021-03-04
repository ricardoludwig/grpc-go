#!/bin/bash

protoc unary/protos/unary.proto --go_out=plugins=grpc:.

if [ $? -ne 0 ]; then
  echo "Error - Something wrong happens"
  exit 1
fi

echo "Success"
