#!/bin/bash

protoc solar-system/protos/solar-system.proto --go_out=plugins=grpc:.

if [ $? -ne 0 ]; then
  echo "Error - Something wrong happens"
  exit 1
fi

echo "Success"
