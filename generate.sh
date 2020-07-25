#!/bin/bash
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
if [ $? -ne 0 ]; then
  echo "Falha na execução"
  exit 1
fi
echo "Executado com sucesso"
